package client

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	"github.com/hashicorp/go-hclog"
	"maps-service/config"
)

type StatusMessage struct {
	Code    int
	Message *proto.Message
}

type KafkaClient struct {
	C *kafka.Consumer
	P *kafka.Producer
}

func NewKafkaClient(cfg *config.Config) (*KafkaClient, error) {
	log := hclog.Default()

	configConsumer := kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBrokers,
		"client.id":         cfg.ServiceName,
		"auto.offset.reset": "earliest",
		"group.id":          cfg.ServiceName,
	}

	configProducer := kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBrokers,
		"client.id":         cfg.ServiceName,
	}

	c, err := kafka.NewConsumer(&configConsumer)
	if err != nil {
		log.Error("[client.NewKafkaClient] kafka.NewConsumer", "error", err)
		return nil, err
	}

	err = c.Subscribe(cfg.KafkaRequestTopic, nil)
	if err != nil {
		log.Error("[client.NewKafkaClient] c.Subscribe", "error", err)
		return nil, err
	}

	p, err := kafka.NewProducer(&configProducer)
	if err != nil {
		log.Error("[client.NewKafkaClient] kafka.NewProducer", "error", err)
		return nil, err
	}

	return &KafkaClient{
		C: c,
		P: p,
	}, nil
}
