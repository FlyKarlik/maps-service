package server

import (
	"comet/utils"
	"context"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"maps-service/internal/models"
	pb "protos/maps"
)

// AddGroup add new group, grpc method
func (m *MapsService) AddGroup(ctx context.Context, rr *pb.MGroup) (*pb.GroupMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddGroup")
	defer span.End()

	g, err := m.db.AddGroup(rr.Name, rr.CreateUserIp, rr.CreateUserId)
	if err != nil {
		log.Error("[server.AddGroup] m.db.AddGroup", "error", err)
		return &pb.GroupMessage{Code: int32(utils.CodeInternal), Group: &pb.MGroup{}}, models.InternalError
	}

	return &pb.GroupMessage{
		Code: int32(utils.CodeOK),
		Group: &pb.MGroup{
			Id:           g.ID,
			Name:         g.Name,
			CreateUserId: g.CreateUserID,
			UpdateUserId: g.UpdateUserID,
			CreateUserIp: g.CreateUserIP,
			UpdateUserIp: g.UpdateUserIP,
			CreatedAt:    g.CreatedAt.String(),
			UpdatedAt:    g.UpdatedAt.String(),
		},
	}, nil
}

// EditGroup edit group, grpc method
func (m *MapsService) EditGroup(ctx context.Context, rr *pb.MGroup) (*pb.GroupMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "EditGroup")
	defer span.End()

	g, err := m.db.EditGroup(
		rr.Id,
		rr.Name,
		rr.UpdateUserIp,
		rr.UpdateUserId,
	)
	if err != nil {
		log.Error("[server.EditGroup] m.db.EditGroup", "error", err)
		return &pb.GroupMessage{Code: int32(utils.CodeInternal), Group: &pb.MGroup{}}, models.InternalError
	}

	return &pb.GroupMessage{
		Code: int32(utils.CodeOK),
		Group: &pb.MGroup{
			Id:           g.ID,
			Name:         g.Name,
			CreateUserId: g.CreateUserID,
			UpdateUserId: g.UpdateUserID,
			CreateUserIp: g.CreateUserIP,
			UpdateUserIp: g.UpdateUserIP,
			CreatedAt:    g.CreatedAt.String(),
			UpdatedAt:    g.UpdatedAt.String(),
		},
	}, nil
}

// DeleteGroup delete group by id, grpc method
func (m *MapsService) DeleteGroup(ctx context.Context, rr *pb.MGroup) (*pb.GroupMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeleteGroup")
	defer span.End()

	err := m.db.DeleteGroup(rr.Id)
	if err != nil {
		log.Error("[server.DeleteGroup] m.db.DeleteGroup", "error", err)
		return &pb.GroupMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.GroupMessage{
		Code: int32(utils.CodeOK),
	}, nil
}

// Group get group by id, grpc method
func (m *MapsService) Group(ctx context.Context, rr *pb.MGroup) (*pb.GroupMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "Group")
	defer span.End()

	g, err := m.db.Group(rr.Id)
	if err != nil {
		log.Error("[server.Group] m.db.Group", "error", err)
		return &pb.GroupMessage{Code: int32(utils.CodeInternal), Group: &pb.MGroup{}}, models.InternalError
	}

	return &pb.GroupMessage{
		Code: int32(utils.CodeOK),
		Group: &pb.MGroup{
			Id:           g.ID,
			Name:         g.Name,
			CreateUserId: g.CreateUserID,
			UpdateUserId: g.UpdateUserID,
			CreateUserIp: g.CreateUserIP,
			UpdateUserIp: g.UpdateUserIP,
			CreatedAt:    g.CreatedAt.String(),
			UpdatedAt:    g.UpdatedAt.String(),
		},
	}, nil
}

// Groups get all groups, grpc method
func (m *MapsService) Groups(ctx context.Context, _ *emptypb.Empty) (*pb.GroupsMessage, error) {
	var response []*pb.MGroup

	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "Groups")
	defer span.End()

	groups, err := m.db.Groups()
	if err != nil {
		log.Error("[server.Groups] m.db.Groups", "error", err)
		return &pb.GroupsMessage{Code: int32(utils.CodeInternal), Groups: response}, models.InternalError
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

	return &pb.GroupsMessage{
		Code:   int32(utils.CodeOK),
		Groups: response,
	}, nil
}
