package client

import (
	"comet/utils"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	pb "protos/maps"
)

func (m *MapsConsumerGroup) HandleAddMapRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MMap
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	r, err := m.srv.AddMap(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] mcg.mapsUC.AddMap", "error", err)
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
			Partition: utils.AddMapResponsePartition},
		Key:   []byte{utils.AddMapResponse},
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

func (m *MapsConsumerGroup) HandleMapRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MMap
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.Map(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.Map", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.MapResponsePartition},
		Key:   []byte{utils.MapResponse},
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

func (m *MapsConsumerGroup) HandleEditMapRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MMap
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.EditMap(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.EditMap", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.EditMapResponsePartition},
		Key:   []byte{utils.EditMapResponse},
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

func (m *MapsConsumerGroup) HandleDeleteMapRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MMap
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.DeleteMap(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.DeleteMap", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.DeleteMapResponsePartition},
		Key:   []byte{utils.DeleteMapResponse},
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

func (m *MapsConsumerGroup) HandleMapsRequest(ctx context.Context, message *kafka.Message) {
	var id = message.Headers[0].Key

	model, err := m.srv.Maps(ctx, nil)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.Map", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.MapsResponsePartition},
		Key:   []byte{utils.MapsResponse},
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

func (m *MapsConsumerGroup) MapSwitcher(ctx context.Context, message *kafka.Message) bool {
	switch message.Key[0] {
	case utils.AddMapRequest:
		m.HandleAddMapRequest(ctx, message)
		return true
	case utils.MapRequest:
		m.HandleMapRequest(ctx, message)
		return true
	case utils.EditMapRequest:
		m.HandleEditMapRequest(ctx, message)
		return true
	case utils.DeleteMapRequest:
		m.HandleDeleteMapRequest(ctx, message)
		return true
	case utils.MapsRequest:
		m.HandleMapsRequest(ctx, message)
		return true
	default:
		return false
	}
}
