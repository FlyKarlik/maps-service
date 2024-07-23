package client

import (
	"comet/utils"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	pb "protos/maps"
)

func (m *MapsConsumerGroup) HandleGroupRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGroup
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.Group(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.Group", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.GroupResponsePartition},
		Key:   []byte{utils.GroupResponse},
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

func (m *MapsConsumerGroup) HandleGroupsRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGroup
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.Groups(ctx, nil)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.Groups", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.GroupsResponsePartition},
		Key:   []byte{utils.GroupsResponse},
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

func (m *MapsConsumerGroup) HandleEditGroupRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGroup
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.EditGroup(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.EditGroup", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.EditGroupResponsePartition},
		Key:   []byte{utils.EditGroupResponse},
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

func (m *MapsConsumerGroup) HandleDeleteGroupRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGroup
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.DeleteGroup(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.DeleteGroup", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.DeleteGroupResponsePartition},
		Key:   []byte{utils.DeleteGroupResponse},
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

func (m *MapsConsumerGroup) HandleAddGroupRequest(ctx context.Context, message *kafka.Message) {
	var request pb.MGroup
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.AddGroup(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.AddGroup", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.AddGroupResponsePartition},
		Key:   []byte{utils.AddGroupResponse},
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
