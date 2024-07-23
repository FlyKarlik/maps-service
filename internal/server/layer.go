package server

import (
	"comet/utils"
	"context"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"maps-service/internal/models"
	pb "protos/maps"
)

// AddLayer layer add, grpc server method
func (m *MapsService) AddLayer(ctx context.Context, rr *pb.MLayer) (*pb.LayerMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddLayer")
	defer span.End()

	layer, err := m.db.AddLayer(
		rr.Name,
		rr.LayerType,
		rr.TableId,
		rr.CreateUserIp,
		rr.CreateUserId,
	)

	if err != nil {
		log.Error("[server.AddLayer] m.db.AddLayer", "error", err)
		return &pb.LayerMessage{Code: int32(utils.CodeInternal), Layer: &pb.MLayer{}}, models.InternalError
	}

	return &pb.LayerMessage{
		Code: int32(utils.CodeOK),
		Layer: &pb.MLayer{
			Id:           layer.ID,
			Name:         layer.Name,
			LayerType:    layer.LayerType,
			TableId:      layer.TableID,
			CreateUserId: layer.CreateUserID,
			CreateUserIp: layer.CreateUserIP,
			UpdateUserId: layer.UpdateUserID,
			UpdateUserIp: layer.UpdateUserIP,
			CreatedAt:    layer.CreatedAt.String(),
			UpdatedAt:    layer.UpdatedAt.String(),
		},
	}, nil
}

// Layer get layer by id, grpc server method
func (m *MapsService) Layer(ctx context.Context, rr *pb.MLayer) (*pb.LayerMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "Layer")
	defer span.End()

	layer, err := m.db.Layer(rr.Id)
	if err != nil {
		log.Error("[server.Layer] m.db.Layer", "error", err)
		return &pb.LayerMessage{Code: int32(utils.CodeNotFound), Layer: &pb.MLayer{}}, models.InternalError
	}

	return &pb.LayerMessage{
		Code: int32(utils.CodeOK),
		Layer: &pb.MLayer{
			Id:           layer.ID,
			Name:         layer.Name,
			LayerType:    layer.LayerType,
			TableId:      layer.TableID,
			CreateUserId: layer.CreateUserID,
			CreateUserIp: layer.CreateUserIP,
			UpdateUserId: layer.UpdateUserID,
			UpdateUserIp: layer.UpdateUserIP,
			CreatedAt:    layer.CreatedAt.String(),
			UpdatedAt:    layer.UpdatedAt.String(),
		},
	}, nil
}

// Layers get all layers, grpc server method
func (m *MapsService) Layers(ctx context.Context, _ *emptypb.Empty) (*pb.LayersMessage, error) {
	var layer []*pb.MLayer
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "Layers")
	defer span.End()

	layers, err := m.db.Layers()
	if err != nil {
		log.Error("[server.Layers] m.db.Layers", "error", err)
		return &pb.LayersMessage{Code: int32(utils.CodeInternal), Layers: layer}, models.InternalError
	}

	for _, v := range *layers {
		layer = append(layer, &pb.MLayer{
			Id:           v.ID,
			Name:         v.Name,
			LayerType:    v.LayerType,
			TableId:      v.TableID,
			CreateUserId: v.CreateUserID,
			CreateUserIp: v.CreateUserIP,
			UpdateUserId: v.UpdateUserID,
			UpdateUserIp: v.UpdateUserIP,
			CreatedAt:    v.CreatedAt.String(),
			UpdatedAt:    v.UpdatedAt.String(),
		})
	}

	return &pb.LayersMessage{
		Code:   int32(utils.CodeOK),
		Layers: layer,
	}, nil
}

// EditLayer edit layer, grpc server method
func (m *MapsService) EditLayer(ctx context.Context, rr *pb.MLayer) (*pb.LayerMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "EditLayer")
	defer span.End()

	l, err := m.db.EditLayer(
		rr.Id,
		rr.Name,
		rr.LayerType,
		rr.TableId,
		rr.UpdateUserIp,
		rr.UpdateUserId,
	)

	if err != nil {
		log.Error("[server.EditLayer] m.db.EditLayer", "error", err)
		return &pb.LayerMessage{Code: int32(utils.CodeInternal), Layer: &pb.MLayer{}}, models.InternalError
	}

	return &pb.LayerMessage{
		Code: int32(utils.CodeOK),
		Layer: &pb.MLayer{
			Id:           l.ID,
			Name:         l.Name,
			LayerType:    l.LayerType,
			TableId:      l.TableID,
			CreateUserId: l.CreateUserID,
			CreateUserIp: l.CreateUserIP,
			UpdateUserId: l.UpdateUserID,
			UpdateUserIp: l.UpdateUserIP,
			CreatedAt:    l.CreatedAt.String(),
			UpdatedAt:    l.UpdatedAt.String(),
		},
	}, nil
}

// DeleteLayer delete layer by id, grpc server method
func (m *MapsService) DeleteLayer(ctx context.Context, rr *pb.MLayer) (*pb.LayerMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeleteLayer")
	defer span.End()

	err := m.db.DeleteLayer(rr.Id)
	if err != nil {
		log.Error("[server.DeleteLayer] m.db.DeleteLayer", "error", err)
		return &pb.LayerMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.LayerMessage{Code: int32(utils.CodeOK)}, nil
}
