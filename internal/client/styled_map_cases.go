package client

import (
	"comet/utils"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	pb "protos/maps"
)

func (m *MapsConsumerGroup) HandleStyledMapRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MMap
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	r, err := m.srv.StyledMap(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] mcg.srv.StyledMap", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	data, err := proto.Marshal(r)

	if err != nil {
		m.log.Error("[client.layerWorker] proto.Marshal", "error", err)
	}

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.StyledMapResponsePartition},
		Key:   []byte{utils.StyledMapResponse},
		Value: data,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.layerWorker] m.Client.P.Produce", "error", err)
	}

	return
}

func (m *MapsConsumerGroup) StyledMapSwitcher(ctx context.Context, message *kafka.Message) bool {
	switch message.Key[0] {
	case utils.StyledMapRequest:
		m.HandleStyledMapRequest(ctx, message)
		return true
	default:
		return false
	}
}
