package client

import (
	"comet/utils"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	pb "protos/maps"
)

func (m *MapsConsumerGroup) HandleAddMapGroupRelationRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	r, err := m.srv.AddMapGroupRelation(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] mcg.srv.AddMapGroupRelation", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	data, err := proto.Marshal(r)

	if err != nil {
		m.log.Error("[client.Worker] proto.Marshal", "error", err)
	}

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.AddMapGroupRelationResponsePartition},
		Key:   []byte{utils.AddMapGroupRelationResponse},
		Value: data,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.Worker] m.Client.P.Produce", "error", err)
	}

	return
}

func (m *MapsConsumerGroup) HandleDeleteMapGroupRelationRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.DeleteMapGroupRelation(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.DeleteMapGroupRelation", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.DeleteMapGroupRelationResponsePartition},
		Key:   []byte{utils.DeleteMapGroupRelationResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.Worker] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) HandleMapGroupRelationsRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.MapGroupRelations(ctx, nil)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.MapGroupRelations", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.MapGroupRelationsResponsePartition},
		Key:   []byte{utils.MapGroupRelationsResponse},
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

func (m *MapsConsumerGroup) HandleMapRelationGroupsRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MMap
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.MapRelationGroups(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.MapRelationGroups", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.MapRelationGroupsResponsePartition},
		Key:   []byte{utils.MapRelationGroupsResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.Worker] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) HandleGroupRelationMapsRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGroup
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.GroupRelationMaps(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.GroupRelationMaps", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.GroupRelationMapsResponsePartition},
		Key:   []byte{utils.GroupRelationMapsResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.Worker] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) HandleMapGroupOrderDownRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.m.HandleMapGroupOrderDownRequest] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.MapGroupOrderDown(ctx, &request)
	if err != nil {
		m.log.Error("[client.m.HandleMapGroupOrderDownRequest] m.srv.MapGroupOrderDown", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.m.HandleMapGroupOrderDownRequest] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.MapGroupOrderDownResponsePartition},
		Key:   []byte{utils.MapGroupOrderDownResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.m.HandleMapGroupOrderDownRequest] m.Client.P.Produce", "error", err)
	}

	return
}

func (m *MapsConsumerGroup) HandleMapGroupOrderUpRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.m.HandleMapGroupOrderUpRequest] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.MapGroupOrderUp(ctx, &request)
	if err != nil {
		m.log.Error("[client.m.HandleMapGroupOrderUpRequest] m.srv.MapGroupOrderUp", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.m.HandleMapGroupOrderUpRequest] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.MapGroupOrderUpResponsePartition},
		Key:   []byte{utils.MapGroupOrderUpResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.m.HandleMapGroupOrderUpRequest] m.Client.P.Produce", "error", err)
	}

	return
}
