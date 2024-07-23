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

// AddMapGroupRelation add new relation
func (m *MapsService) AddMapGroupRelation(ctx context.Context, rr *pb.MGRelation) (*pb.MGRelationMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddMapGroupRelation")
	defer span.End()
	fmt.Println(rr)
	r, err := m.db.AddMapGroupRelation(rr.MapId, rr.GroupId)
	if err != nil {
		log.Error("[server.AddMapGroupRelation] m.db.AddMapGroupRelation", "error", err)
		return &pb.MGRelationMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.MGRelationMessage{
		Code: int32(utils.CodeOK),
		Relation: &pb.MGRelation{
			Id:         r.ID,
			MapId:      r.MapID,
			GroupOrder: r.GroupOrder,
			GroupId:    r.GroupID,
		},
	}, nil
}

func (m *MapsService) DeleteMapGroupRelation(ctx context.Context, rr *pb.MGRelation) (*pb.MGRelationMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeleteMapGroupRelation")
	defer span.End()

	err := m.db.DeleteMapGroupRelation(rr.MapId, rr.GroupId)
	if err != nil {
		log.Error("[server.DeleteMapGroupRelation] m.db.DeleteMapGroupRelation", "error", err)
		return &pb.MGRelationMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.MGRelationMessage{
		Code: int32(utils.CodeOK),
	}, nil
}

func (m *MapsService) MapGroupRelations(ctx context.Context, _ *emptypb.Empty) (*pb.MGRelationsMessage, error) {
	var response []*pb.MGRelation
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "MapGroupRelations")
	defer span.End()

	r, err := m.db.MapGroupRelations()
	if err != nil {
		log.Error("[server.MapGroupRelations] m.db.MapGroupRelations", "error", err)
		return &pb.MGRelationsMessage{
			Code:      int32(utils.CodeInternal),
			Relations: response,
		}, models.InternalError
	}

	for _, v := range *r {
		response = append(response, &pb.MGRelation{
			Id:         v.ID,
			MapId:      v.MapID,
			GroupId:    v.GroupID,
			GroupOrder: v.GroupOrder,
		})
	}

	return &pb.MGRelationsMessage{
		Code:      int32(utils.CodeOK),
		Relations: response,
	}, nil
}

func (m *MapsService) MapRelationGroups(ctx context.Context, rr *pb.MMap) (*pb.GroupsMessage, error) {
	var response []*pb.MGroup
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "MapRelationGroups")
	defer span.End()

	r, err := m.db.MapRelationGroups(rr.Id)
	if err != nil {
		log.Error("[server.MapRelationGroups] m.db.MapRelationGroups", "error", err)
		return &pb.GroupsMessage{
			Code:   int32(utils.CodeInternal),
			Groups: response,
		}, models.InternalError
	}

	if len(*r) < 1 {
		log.Error("[server.MapRelationGroups] len(*r)", "error", fmt.Errorf("r length less 1"))
		return &pb.GroupsMessage{
			Code:   int32(utils.CodeNotFound),
			Groups: response,
		}, models.NotFoundError
	}

	for _, g := range *r {
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

	return &pb.GroupsMessage{
		Code:   int32(utils.CodeOK),
		Groups: response,
	}, nil
}

func (m *MapsService) GroupRelationMaps(ctx context.Context, rr *pb.MGroup) (*pb.MapsMessage, error) {
	var response []*pb.MMap
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "GroupRelationMaps")
	defer span.End()

	r, err := m.db.GroupRelationMaps(rr.Id)
	if err != nil {
		log.Error("[server.GroupRelationMaps] m.db.GroupRelationMaps", "error", err)
		return &pb.MapsMessage{
			Code: int32(utils.CodeOK),
			Maps: response,
		}, models.InternalError
	}

	if len(*r) < 1 {
		log.Error("[server.GroupRelationMaps] len(*r)", "error", fmt.Errorf("r length less 1"))
		return &pb.MapsMessage{
			Code: int32(utils.CodeNotFound),
			Maps: response,
		}, models.NotFoundError
	}

	for _, m_ := range *r {
		response = append(response, &pb.MMap{
			Id:           m_.ID,
			Name:         m_.Name,
			Picture:      m_.Picture,
			Describe:     m_.Describe,
			Active:       m_.Active,
			CreateUserId: m_.CreateUserID,
			CreateUserIp: m_.CreateUserIP,
			CreatedAt:    m_.CreatedAt.String(),
			UpdateUserId: m_.UpdateUserID,
			UpdateUserIp: m_.UpdateUserIP,
			UpdatedAt:    m_.UpdatedAt.String(),
		})
	}

	return &pb.MapsMessage{
		Code: int32(utils.CodeOK),
		Maps: response,
	}, nil
}

func (m *MapsService) MapGroupOrderUp(ctx context.Context, rr *pb.MGRelation) (*pb.MGRelationMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "MapGroupOrderUp")
	defer span.End()

	r, err := m.db.MapGroupOrderUp(rr.Id)
	if err != nil {
		log.Error("[server.MapGroupOrderUp] m.db.MapGroupOrderUp", "error", err)
		return &pb.MGRelationMessage{Code: int32(utils.CodeInternal), Relation: &pb.MGRelation{}}, models.InternalError
	}

	return &pb.MGRelationMessage{
		Code: int32(utils.CodeOK),
		Relation: &pb.MGRelation{
			Id:         r.ID,
			GroupOrder: r.GroupOrder,
			GroupId:    r.GroupID,
			MapId:      r.MapID,
		},
	}, nil
}

func (m *MapsService) MapGroupOrderDown(ctx context.Context, rr *pb.MGRelation) (*pb.MGRelationMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "MapGroupOrderDown")
	defer span.End()

	r, err := m.db.MapGroupOrderDown(rr.Id)
	if err != nil {
		log.Error("[server.MapGroupOrderDown] m.db.MapGroupOrderDown", "error", err)
		return &pb.MGRelationMessage{Code: int32(utils.CodeInternal), Relation: &pb.MGRelation{}}, models.InternalError
	}

	return &pb.MGRelationMessage{
		Code: int32(utils.CodeOK),
		Relation: &pb.MGRelation{
			Id:         r.ID,
			GroupOrder: r.GroupOrder,
			GroupId:    r.GroupID,
			MapId:      r.MapID,
		},
	}, nil
}
