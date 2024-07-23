package main

import (
	"comet/db"
	"comet/utils"
	"context"
	"flag"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"maps-service/config"
	"maps-service/internal/client"
	"maps-service/internal/models"
	"maps-service/internal/server"
	"maps-service/internal/server/repository"
	"maps-service/internal/server/repository/redis"
	"net"
	"os"
	protos "protos/maps"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()
	log := hclog.Default()

	cfg, err := config.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to load environment: %w", err)
	}

	tracer, err := utils.InitTracerProvider(cfg.JaegerHost, "maps-service")
	if err != nil {
		return fmt.Errorf("failed init Tracer: %w", err)
	}

	dns2, err := cfg.GetDbDnsSp()
	if err != nil {
		return fmt.Errorf("failed get dns2 string: %w", err)
	}

	databaseSP, err := sqlx.Connect("postgres", dns2)
	defer func(databaseSP *sqlx.DB) {
		err := databaseSP.Close()
		if err != nil {
			_ = fmt.Errorf("error while closing DB connection %v", err)
		}
	}(databaseSP)
	if err != nil {
		return fmt.Errorf("failed to connect to databaseSP: %w", err)
	}

	dns1, err := cfg.GetDbDns()
	if err != nil {
		return fmt.Errorf("failed get dns2 string: %w", err)
	}

	database, err := db.NewDatabase(dns1)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Info("Postgres", "dns1", dns1)
	log.Info("Postgres", "dns2", dns2)

	err = database.AutoMigrate(
		&models.Group{}, &models.MapGroupRelation{},
		&models.History{}, &models.Layer{},
		&models.GroupLayerRelation{}, &models.LayerStyleRelation{},
		&models.Map{}, &models.Pattern{},
		&models.Style{},
		&models.Table{}, &models.Column{},
		&models.Srids{},
	)
	if err != nil {
		return fmt.Errorf("failed AutoMigrate database: %w", err)
	}

	cert, err := credentials.NewServerTLSFromFile("cert/server.crt", "cert/server.pem")
	if err != nil {
		return fmt.Errorf("failed to setup TLS: %w", err)
	}

	// Create a new gRPC srv
	gs := grpc.NewServer(
		grpc.Creds(cert),
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
	)

	redisStyledMapsClient, err := redis.NewStyledMapsClient(ctx, cfg)
	if err != nil {
		return fmt.Errorf("redis.NewStyledMapsClient: %w", err)
	}

	redisRepos := redis.NewRepository(redisStyledMapsClient)
	repoMaps := repository.NewRepository(database, databaseSP)
	srv := server.NewMaps(repoMaps, redisRepos, tracer, cfg)
	protos.RegisterMapsServiceServer(gs, srv)
	reflection.Register(gs)

	c, err := client.NewKafkaClient(cfg)

	mapsCG := client.NewMapsConsumerGroup(
		[]string{cfg.KafkaBrokers},
		"maps",
		cfg,
		repoMaps,
		c,
		log,
		srv,
	)
	mapsCG.RunConsumers(ctx, cancel)

	l, err := net.Listen("tcp", cfg.ServerHost)
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	log.Info("service is running.")

	// listen for requests
	err = gs.Serve(l)
	if err != nil {
		return fmt.Errorf("failed serving: %w", err)
	}

	return nil
}
