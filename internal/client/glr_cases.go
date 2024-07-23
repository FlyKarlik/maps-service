package client

import (
	"comet/utils"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	pb "protos/maps"
)

func (m *MapsConsumerGroup) HandleAddGroupLayerRelationRequest(ctx context.Context, message *kafka.Message) {
	var request pb.GroupLayerRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	r, err := m.srv.AddGroupLayerRelation(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] mcg.mapsUC.AddGroupLayerRelation", "error", err)
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
			Partition: utils.AddGroupLayerRelationResponsePartition},
		Key:   []byte{utils.AddGroupLayerRelationResponse},
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

func (m *MapsConsumerGroup) HandleDeleteGroupLayerRelationRequest(ctx context.Context, message *kafka.Message) {
	var request pb.GroupLayerRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.DeleteGroupLayerRelation(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.DeleteGroupLayerRelation", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.DeleteGroupLayerRelationResponsePartition},
		Key:   []byte{utils.DeleteGroupLayerRelationResponse},
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

func (m *MapsConsumerGroup) HandleGroupLayerRelationsRequest(ctx context.Context, message *kafka.Message) {
	var request pb.GroupLayerRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.GroupLayerRelations(ctx, nil)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.GroupLayerRelations", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.GroupLayerRelationsResponsePartition},
		Key:   []byte{utils.GroupLayerRelationsResponse},
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

func (m *MapsConsumerGroup) HandleGroupRelationLayersRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGroup
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.GroupRelationLayers(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.GroupRelationLayers", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.GroupRelationLayersResponsePartition},
		Key:   []byte{utils.GroupRelationLayersResponse},
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

func (m *MapsConsumerGroup) HandleLayerRelationGroupsRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MLayer
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.LayerRelationGroups(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.LayerRelationGroups", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.LayerRelationGroupsResponsePartition},
		Key:   []byte{utils.LayerRelationGroupsResponse},
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

func (m *MapsConsumerGroup) HandleGroupLayerOrderUpRequest(ctx context.Context, message *kafka.Message) {
	var request pb.GroupLayerRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.m.HandleGroupLayerOrderUpRequest] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.GroupLayerOrderUp(ctx, &request)
	if err != nil {
		m.log.Error("[client.m.HandleGroupLayerOrderUpRequest] m.srv.GroupLayerOrderUp", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.m.HandleGroupLayerOrderUpRequest] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.GroupLayerOrderUpResponsePartition},
		Key:   []byte{utils.GroupLayerOrderUpResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.m.HandleGroupLayerOrderUpRequest] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) HandleGroupLayerOrderDownRequest(ctx context.Context, message *kafka.Message) {
	var request pb.GroupLayerRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.m.HandleGroupLayerOrderDownRequest] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.GroupLayerOrderDown(ctx, &request)
	if err != nil {
		m.log.Error("[client.m.HandleGroupLayerOrderDownRequest] m.srv.GroupLayerOrderDown", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.m.HandleGroupLayerOrderDownRequest] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.GroupLayerOrderDownResponsePartition},
		Key:   []byte{utils.GroupLayerOrderDownResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.m.HandleGroupLayerOrderDownRequest] m.Client.P.Produce", "error", err)
	}

	return
}
