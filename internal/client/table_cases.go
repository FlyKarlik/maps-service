package client

import (
	"comet/utils"
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/protobuf/proto"
	pb "protos/maps"
)

func (m *MapsConsumerGroup) HandleAddTableRequest(ctx context.Context, message *kafka.Message) {
	var request pb.Table
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	r, err := m.srv.AddTable(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] mcg.srv.AddTable", "error", err)
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
			Partition: utils.AddTableResponsePartition},
		Key:   []byte{utils.AddTableResponse},
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

func (m *MapsConsumerGroup) HandleTableRequest(ctx context.Context, message *kafka.Message) {
	var request pb.Table
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.Table(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.Table", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.TableResponsePartition},
		Key:   []byte{utils.TableResponse},
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

func (m *MapsConsumerGroup) HandleEditTableRequest(ctx context.Context, message *kafka.Message) {
	var request pb.Table
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.EditTable(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.EditTable", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.EditTableResponsePartition},
		Key:   []byte{utils.EditTableResponse},
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

func (m *MapsConsumerGroup) HandleDeleteTableRequest(ctx context.Context, message *kafka.Message) {
	var request pb.Table
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.DeleteTable(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.DeleteTable", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.DeleteTableResponsePartition},
		Key:   []byte{utils.DeleteTableResponse},
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

func (m *MapsConsumerGroup) HandleTablesRequest(ctx context.Context, message *kafka.Message) {
	var id = message.Headers[0].Key

	model, err := m.srv.Tables(ctx)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.Tables", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.TablesResponsePartition},
		Key:   []byte{utils.TablesResponse},
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

func (m *MapsConsumerGroup) HandleTableColumnsRequest(ctx context.Context, message *kafka.Message) {
	var request pb.Table
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.TableColumns(ctx, &request)
	if err != nil {
		m.log.Error("[client.layerWorker] m.srv.TableColumns", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.layerWorker] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.TableColumnsResponsePartition},
		Key:   []byte{utils.TableColumnsResponse},
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

func (m *MapsConsumerGroup) HandleTableColumnUniqueValuesRequest(ctx context.Context, message *kafka.Message) {
	var request pb.ColumnUnique

	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.HandleTableColumnUniqueValuesRequest] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.TableColumnUniqueValues(ctx, &request)
	if err != nil {
		m.log.Error("[client.HandleTableColumnUniqueValuesRequest] m.srv.TableColumnUniqueValues", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.HandleTableColumnUniqueValuesRequest] m.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)
	if err != nil {
		m.log.Error("[client.HandleTableColumnUniqueValuesRequest] proto.Marshal", "error", err)
	}

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.TableColumnUniqueValuesResponsePartition},
		Key:   []byte{utils.TableColumnUniqueValuesResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.HandleTableColumnUniqueValuesRequest] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) HandleTableFeaturesRequest(ctx context.Context, message *kafka.Message) {
	var request pb.TableFeaturesRequest
	var id = message.Headers[0].Key

	err := proto.Unmarshal(message.Value, &request)
	if err != nil {
		m.log.Error("[client.HandleTableFeaturesRequest] proto.Unmarshal", "error", err)
	}

	model, err := m.srv.TableFeatures(ctx, &request)
	if err != nil {
		m.log.Error("[client.HandleTableFeaturesRequest] m.srv.TableColumns", "error", err)
	}

	if _, err = m.Client.C.CommitMessage(message); err != nil {
		m.log.Error("[client.HandleTableFeaturesRequest] mcg.Client.C.CommitMessage", "error", err)
	}

	dataByte, err := proto.Marshal(model)

	err = m.Client.P.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &m.cfg.KafkaResponseTopic,
			Partition: utils.TableFeaturesResponsePartition},
		Key:   []byte{utils.TableFeaturesResponse},
		Value: dataByte,
		Headers: []kafka.Header{
			{Key: id},
		},
	}, nil)

	if err != nil {
		m.log.Error("[client.HandleTableFeaturesRequest] m.Client.P.Produce", "error", err)

	}

	return
}

func (m *MapsConsumerGroup) TableSwitcher(ctx context.Context, message *kafka.Message) bool {
	switch message.Key[0] {
	case utils.AddTableRequest:
		m.HandleAddTableRequest(ctx, message)
		return true
	case utils.TableRequest:
		m.HandleTableRequest(ctx, message)
		return true
	case utils.TablesRequest:
		m.HandleTablesRequest(ctx, message)
		return true
	case utils.TableColumnsRequest:
		m.HandleTableColumnsRequest(ctx, message)
		return true
	case utils.DeleteTableRequest:
		m.HandleDeleteTableRequest(ctx, message)
		return true
	case utils.EditTableRequest:
		m.HandleEditTableRequest(ctx, message)
		return true
	case utils.TableColumnUniqueValuesRequest:
		m.HandleTableColumnUniqueValuesRequest(ctx, message)
		return true
	case utils.TableFeaturesRequest:
		m.HandleTableFeaturesRequest(ctx, message)
		return true
	default:
		return false
	}
}
