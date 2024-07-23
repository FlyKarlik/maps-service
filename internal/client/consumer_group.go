package client

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"maps-service/config"
	"maps-service/internal/server"
	"maps-service/internal/server/repository"
	"sync"
)

type MapsConsumerGroup struct {
	Brokers []string
	GroupID string
	cfg     *config.Config
	log     hclog.Logger
	mapsUC  *repository.Repository
	Client  *KafkaClient
	srv     *server.MapsService
}

func NewMapsConsumerGroup(
	brokers []string,
	groupID string,
	cfg *config.Config,
	repo *repository.Repository,
	client *KafkaClient,
	log hclog.Logger,
	srv *server.MapsService,
) *MapsConsumerGroup {
	return &MapsConsumerGroup{
		Brokers: brokers,
		GroupID: groupID,
		cfg:     cfg,
		mapsUC:  repo,
		Client:  client,
		log:     log,
		srv:     srv,
	}
}

func (m *MapsConsumerGroup) consumerLayer(
	ctx context.Context,
	cancel context.CancelFunc,
	workersNum int) {

	log := hclog.Default()
	defer cancel()

	log.Info("Starting consumer group...")

	wg := &sync.WaitGroup{}
	for i := 0; i <= workersNum; i++ {
		wg.Add(1)
		go m.layerWorker(ctx, cancel, wg, i)
	}
}

func (m *MapsConsumerGroup) RunConsumers(ctx context.Context, cancel context.CancelFunc) {
	go m.consumerLayer(ctx, cancel, 1)
}
