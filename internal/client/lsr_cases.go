package client

import (
	"comet/utils"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	pb "protos/maps"
)

func (m *MapsConsumerGroup) HandleAddLayerStyleRelationRequest(ctx context.Context, message *kafka.Message) {
	var request pb.LSRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	r, err := m.srv.AddLayerStyleRelation(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] mcg.srv.AddLayerStyleRelation", "error", err)
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
			Partition: utils.AddLayerStyleRelationResponsePartition},
		Key:   []byte{utils.AddLayerStyleRelationResponse},
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

func (m *MapsConsumerGroup) HandleDeleteLayerStyleRelationRequest(ctx context.Context, message *kafka.Message) {
	var request pb.LSRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.DeleteLayerStyleRelation(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.DeleteLayerStyleRelation", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.DeleteLayerStyleRelationResponsePartition},
		Key:   []byte{utils.DeleteLayerStyleRelationResponse},
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

func (m *MapsConsumerGroup) HandleLayerStyleRelationsRequest(ctx context.Context, message *kafka.Message) {
	var request pb.LSRelation
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.LayerStyleRelations(ctx, nil)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.LayerStyleRelations", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.LayerStyleRelationsResponsePartition},
		Key:   []byte{utils.LayerStyleRelationsResponse},
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

func (m *MapsConsumerGroup) HandleLayerRelationStylesRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MLayer
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.LayerRelationStyles(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.LayerRelationStyles", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.LayerRelationStylesResponsePartition},
		Key:   []byte{utils.LayerRelationStylesResponse},
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

func (m *MapsConsumerGroup) HandleStyleRelationLayersRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MStyle
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.Worker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.StyleRelationLayers(ctx, &request)
	if err != nil {
		m.log.Error("[client.Worker] m.srv.StyleRelationLayers", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.Worker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.StyleRelationLayersResponsePartition},
		Key:   []byte{utils.StyleRelationLayersResponse},
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

func (m *MapsConsumerGroup) LayerStyleRelationSwitcher(ctx context.Context, message *kafka.Message) bool {
	switch message.Key[0] {
	case utils.AddLayerStyleRelationRequest:
		m.HandleAddLayerStyleRelationRequest(ctx, message)
		return true
	case utils.DeleteLayerStyleRelationRequest:
		m.HandleDeleteLayerStyleRelationRequest(ctx, message)
		return true
	case utils.LayerStyleRelationsRequest:
		m.HandleLayerStyleRelationsRequest(ctx, message)
		return true
	case utils.LayerRelationStylesRequest:
		m.HandleLayerRelationStylesRequest(ctx, message)
		return true
	case utils.StyleRelationLayersRequest:
		m.HandleStyleRelationLayersRequest(ctx, message)
		return true
	default:
		return false
	}
}
