package server

import (
	"go.opentelemetry.io/otel/trace"
	"maps-service/config"
	"maps-service/internal/server/interfaces"
	"maps-service/internal/server/repository/redis"
	protos "protos/maps"
)

// MapsService grpc service declaration
type MapsService struct {
	protos.UnimplementedMapsServiceServer
	db    interfaces.Repository
	rdb   *redis.Repository
	trace trace.Tracer
	cfg   *config.Config
}

// NewMaps Creates a new Maps server
func NewMaps(db interfaces.Repository, rdb *redis.Repository, tracer trace.Tracer, cfg *config.Config) *MapsService {
	return &MapsService{
		db:    db,
		rdb:   rdb,
		trace: tracer,
		cfg:   cfg,
	}
}
