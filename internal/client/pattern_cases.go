package client

import (
	"comet/utils"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	pb "protos/maps"
)

func (m *MapsConsumerGroup) HandleAddPatternRequest(ctx context.Context, message *kafka.Message) {
	var request pb.Pattern
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	r, err := m.srv.AddPattern(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] mcg.srv.AddPattern", "error", err)
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
			Partition: utils.AddPatternResponsePartition},
		Key:   []byte{utils.AddPatternResponse},
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

func (m *MapsConsumerGroup) HandlePatternRequest(ctx context.Context, message *kafka.Message) {
	var request pb.Pattern
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.Pattern(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.Pattern", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.PatternResponsePartition},
		Key:   []byte{utils.PatternResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.layerWorker] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) HandleDeletePatternRequest(ctx context.Context, message *kafka.Message) {
	var request pb.Pattern
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.DeletePattern(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.DeletePattern", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.DeletePatternResponsePartition},
		Key:   []byte{utils.DeletePatternResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.layerWorker] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) HandlePatternsRequest(ctx context.Context, message *kafka.Message) {
	var id = message.Headers[0].Key

	model, err := m.srv.Patterns(ctx)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.Patterns", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.PatternsResponsePartition},
		Key:   []byte{utils.PatternsResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.layerWorker] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) PatternSwitcher(ctx context.Context, message *kafka.Message) bool {
	switch message.Key[0] {
	case utils.PatternRequest:
		m.HandlePatternRequest(ctx, message)
		return true
	case utils.PatternsRequest:
		m.HandlePatternsRequest(ctx, message)
		return true
	case utils.AddPatternRequest:
		m.HandleAddPatternRequest(ctx, message)
		return true
	case utils.DeletePatternRequest:
		m.HandleDeletePatternRequest(ctx, message)
		return true
	default:
		return false
	}
}
