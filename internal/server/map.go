package server

import (
	"comet/utils"
	"context"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"maps-service/internal/models"
	pb "protos/maps"
)

// Map get map by id, grpc server method
func (m *MapsService) Map(ctx context.Context, rr *pb.MMap) (*pb.MapMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "Map")
	defer span.End()

	m_, err := m.db.Map(rr.Id)
	if err != nil {
		log.Error("[server.Map] r.db.Map", "error", err)
		return &pb.MapMessage{Code: int32(utils.CodeInternal), Map: &pb.MMap{}}, models.InternalError
	}

	if len(m_.ID) < 1 {
		log.Error("[server.Map] len(m_.ID)", "error", "not found")
		return &pb.MapMessage{Code: int32(utils.CodeNotFound), Map: &pb.MMap{}}, models.InternalError
	}

	return &pb.MapMessage{
		Code: int32(utils.CodeOK),
		Map: &pb.MMap{
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
		},
	}, nil
}

// Maps get all maps, grpc server method
func (m *MapsService) Maps(ctx context.Context, _ *emptypb.Empty) (*pb.MapsMessage, error) {
	log := hclog.Default()
	var response []*pb.MMap
	tr := m.trace
	ctx, span := tr.Start(ctx, "Maps")
	defer span.End()

	maps, err := m.db.Maps()
	if err != nil {
		log.Error("[server.Maps] m.db.Error", "error", err)
		return &pb.MapsMessage{Code: int32(utils.CodeInternal), Maps: response}, models.InternalError
	}

	for _, m_ := range *maps {
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

// AddMap add new map, grpc server method
func (m *MapsService) AddMap(ctx context.Context, rr *pb.MMap) (*pb.MapMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddMap")
	defer span.End()

	m_, err := m.db.AddMap(
		rr.Name,
		rr.Picture,
		rr.Describe,
		rr.CreateUserIp,
		rr.CreateUserId,
		rr.Active,
	)
	if err != nil {
		log.Error("[server.AddMap] r.db.AddMap", "error", err)
		return &pb.MapMessage{Code: int32(utils.CodeInternal), Map: &pb.MMap{}}, models.InternalError
	}

	return &pb.MapMessage{
		Code: int32(utils.CodeOK),
		Map: &pb.MMap{
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
		},
	}, nil
}

// EditMap edit map by id, grpc server method
func (m *MapsService) EditMap(ctx context.Context, rr *pb.MMap) (*pb.MapMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "EditMap")
	defer span.End()

	m_, err := m.db.EditMap(
		rr.Id,
		rr.Name,
		rr.Picture,
		rr.Describe,
		rr.UpdateUserIp,
		rr.UpdateUserId,
		rr.Active,
	)
	if err != nil {
		log.Error("[server.EditMap] m.db.EditMap", "error", err)
		return &pb.MapMessage{Code: int32(utils.CodeInternal), Map: &pb.MMap{}}, models.InternalError
	}

	return &pb.MapMessage{
		Code: int32(utils.CodeOK),
		Map: &pb.MMap{
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
		},
	}, nil
}

// DeleteMap delete map by id, grpc server method
func (m *MapsService) DeleteMap(ctx context.Context, rr *pb.MMap) (*pb.MapMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeleteMap")
	defer span.End()

	err := m.db.DeleteMap(rr.Id)
	if err != nil {
		log.Error("[server.DeleteMap] m.db.DeleteMap", "error", err)
		return &pb.MapMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.MapMessage{
		Code: int32(utils.CodeOK),
	}, nil
}
