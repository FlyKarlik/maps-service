package server

import (
	"comet/utils"
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"maps-service/internal/models"
	pb "protos/maps"
)

// AddGroupLayerRelation add relation group and layer, grpc method
func (m *MapsService) AddGroupLayerRelation(ctx context.Context, rr *pb.GroupLayerRelation) (*pb.GroupLayerRelationMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddGroupLayerRelation")
	defer span.End()

	r, err := m.db.AddGroupLayerRelation(rr.GroupId, rr.LayerId)
	if err != nil {
		log.Error("[server.AddGroupLayerRelation] m.db.AddGroupLayerRelation", "error", err)
		return &pb.GroupLayerRelationMessage{Code: int32(utils.CodeInternal), Relation: &pb.GroupLayerRelation{}}, models.InternalError
	}

	return &pb.GroupLayerRelationMessage{
		Code: int32(utils.CodeOK),
		Relation: &pb.GroupLayerRelation{
			Id:         r.ID,
			LayerOrder: r.LayerOrder,
			GroupId:    r.GroupID,
			LayerId:    r.LayerID,
		},
	}, nil
}

// DeleteGroupLayerRelation delete relation by id, grpc server method
func (m *MapsService) DeleteGroupLayerRelation(ctx context.Context, rr *pb.GroupLayerRelation) (*pb.GroupLayerRelationMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeleteGroupLayerRelation")
	defer span.End()

	err := m.db.DeleteGroupLayerRelation(rr.GroupId, rr.LayerId)
	if err != nil {
		log.Error("[server.DeleteGroupLayerRelation] m.db.DeleteGroupLayerRelation", "error", err)
		return &pb.GroupLayerRelationMessage{Code: int32(utils.CodeNotFound)}, models.InternalError
	}

	return &pb.GroupLayerRelationMessage{
		Code: int32(utils.CodeOK),
	}, nil
}

// GroupLayerRelations get all relations, grpc method
func (m *MapsService) GroupLayerRelations(ctx context.Context, _ *emptypb.Empty) (*pb.GroupLayerRelationsMessage, error) {
	var response []*pb.GroupLayerRelation
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "GroupLayerRelations")
	defer span.End()

	relations, err := m.db.GroupLayerRelations()
	if err != nil {
		log.Error("[server.GroupLayerRelations] m.db.GroupLayerRelations", "error", err)
		return &pb.GroupLayerRelationsMessage{Code: int32(utils.CodeInternal), Relations: response}, models.InternalError
	}

	for _, r := range *relations {
		response = append(response, &pb.GroupLayerRelation{
			Id:         r.ID,
			LayerId:    r.LayerID,
			GroupId:    r.GroupID,
			LayerOrder: r.LayerOrder,
		})
	}

	return &pb.GroupLayerRelationsMessage{Code: int32(utils.CodeOK), Relations: response}, nil
}

// LayerRelationGroups get layer`s groups, grpc method
func (m *MapsService) LayerRelationGroups(ctx context.Context, rr *pb.MLayer) (*pb.LGMessage, error) {
	var response []*pb.MGroup
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "LayerRelationGroups")
	defer span.End()

	groups, err := m.db.LayerRelationGroups(rr.Id)
	if err != nil {
		log.Error("[server.LayerRelationGroups] m.db.LayerRelationGroups", "error", err)
		return &pb.LGMessage{Code: int32(utils.CodeInternal), Groups: response}, models.InternalError
	}

	if len(*groups) < 1 {
		log.Error("[server.LayerRelationGroups] len(*groups)", "error", fmt.Errorf("groups length less 1"))
		return &pb.LGMessage{Code: int32(utils.CodeNotFound), Groups: response}, models.NotFoundError
	}

	for _, g := range *groups {
		response = append(response, &pb.MGroup{
			Id:           g.ID,
			Name:         g.Name,
			CreateUserId: g.CreateUserID,
			UpdateUserId: g.UpdateUserID,
			CreateUserIp: g.CreateUserIP,
			UpdateUserIp: g.UpdateUserIP,
			CreatedAt:    g.CreatedAt.String(),
			UpdatedAt:    g.UpdatedAt.String(),
		})
	}

	return &pb.LGMessage{Code: int32(utils.CodeOK), Groups: response}, nil
}

// GroupRelationLayers get group`s layers, grpc server method
func (m *MapsService) GroupRelationLayers(ctx context.Context, rr *pb.MGroup) (*pb.GLMessage, error) {
	var response []*pb.MLayer
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "GroupRelationLayers")
	defer span.End()

	layers, err := m.db.GroupRelationLayers(rr.Id)
	if err != nil {
		log.Error("[server.GroupRelationLayers] m.db.GroupRelationLayers", "error", err)
		return &pb.GLMessage{Code: int32(utils.CodeInternal), Layers: response}, models.InternalError
	}

	if len(*layers) < 1 {
		log.Error("[server.GroupRelationLayers] len(*layers)", "error", fmt.Errorf("layers length less 1"))
		return &pb.GLMessage{Code: int32(utils.CodeNotFound), Layers: response}, models.NotFoundError
	}

	for _, layer := range *layers {
		response = append(response, &pb.MLayer{
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
		})
	}

	return &pb.GLMessage{
		Code:   int32(utils.CodeOK),
		Layers: response,
	}, nil
}

// GroupLayerOrderUp up the relation order in the group, grpc method
func (m *MapsService) GroupLayerOrderUp(ctx context.Context, rr *pb.GroupLayerRelation) (*pb.GroupLayerRelationMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "GroupLayerOrderUp")
	defer span.End()

	r, err := m.db.GroupLayerOrderUp(rr.Id)
	if err != nil {
		log.Error("[server.GroupLayerOrderUp] m.db.GroupLayerOrderUp", "error", err)
		return &pb.GroupLayerRelationMessage{Code: int32(utils.CodeInternal), Relation: &pb.GroupLayerRelation{}}, models.InternalError
	}

	return &pb.GroupLayerRelationMessage{
		Code: int32(utils.CodeOK),
		Relation: &pb.GroupLayerRelation{
			Id:         r.ID,
			LayerOrder: r.LayerOrder,
			GroupId:    r.GroupID,
			LayerId:    r.LayerID,
		},
	}, nil
}

// GroupLayerOrderDown down the relation order in the group, grpc method
func (m *MapsService) GroupLayerOrderDown(ctx context.Context, rr *pb.GroupLayerRelation) (*pb.GroupLayerRelationMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "GroupLayerOrderDown")
	defer span.End()

	r, err := m.db.GroupLayerOrderDown(rr.Id)
	if err != nil {
		log.Error("[server.GroupLayerOrderDown] m.db.GroupLayerOrderDown", "error", err)
		return &pb.GroupLayerRelationMessage{Code: int32(utils.CodeInternal), Relation: &pb.GroupLayerRelation{}}, models.InternalError
	}

	return &pb.GroupLayerRelationMessage{
		Code: int32(utils.CodeOK),
		Relation: &pb.GroupLayerRelation{
			Id:         r.ID,
			LayerOrder: r.LayerOrder,
			GroupId:    r.GroupID,
			LayerId:    r.LayerID,
		},
	}, nil
}
