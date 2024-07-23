package server

import (
	"comet/utils"
	"context"
	"github.com/hashicorp/go-hclog"
	"maps-service/internal/models"
	pb "protos/maps"
)

func (m *MapsService) Pattern(ctx context.Context, rr *pb.Pattern) (*pb.PatternMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "Pattern")
	defer span.End()

	r, err := m.db.Pattern(rr.Id)
	if err != nil {
		log.Error("[server.Pattern] r.db.Pattern", "error", err)
		return &pb.PatternMessage{Code: int32(utils.CodeInternal), Pattern: &pb.Pattern{}}, models.InternalError
	}

	return &pb.PatternMessage{
		Code: int32(utils.CodeOK),
		Pattern: &pb.Pattern{
			Id:           r.ID,
			Name:         r.Name,
			X:            r.X,
			Y:            r.Y,
			Img:          r.Img,
			CreateUserId: r.CreateUserID,
			CreateUserIp: r.CreateUserIP,
			CreatedAt:    r.CreatedAt.String(),
			UpdateUserId: r.UpdateUserID,
			UpdateUserIp: r.UpdateUserIP,
			UpdatedAt:    r.UpdatedAt.String(),
		},
	}, nil
}

func (m *MapsService) Patterns(ctx context.Context) (*pb.PatternsMessage, error) {
	log := hclog.Default()
	var response []*pb.Pattern
	tr := m.trace
	ctx, span := tr.Start(ctx, "Patterns")
	defer span.End()

	r, err := m.db.Patterns()
	if err != nil {
		log.Error("[server.Patterns] r.db.Patterns", "error", err)
		return &pb.PatternsMessage{Code: int32(utils.CodeInternal), Patterns: response}, models.InternalError
	}

	for _, v := range *r {
		response = append(response, &pb.Pattern{
			Id:           v.ID,
			Name:         v.Name,
			Img:          v.Img,
			X:            v.X,
			Y:            v.Y,
			CreateUserId: v.CreateUserID,
			CreateUserIp: v.CreateUserIP,
			UpdateUserId: v.UpdateUserID,
			UpdateUserIp: v.UpdateUserIP,
			CreatedAt:    v.CreatedAt.String(),
			UpdatedAt:    v.UpdatedAt.String(),
		})
	}

	return &pb.PatternsMessage{
		Code:     int32(utils.CodeOK),
		Patterns: response,
	}, nil
}

func (m *MapsService) AddPattern(ctx context.Context, rr *pb.Pattern) (*pb.PatternMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddPattern")
	defer span.End()

	r, err := m.db.AddPattern(
		rr.Name,
		rr.Img,
		rr.X,
		rr.Y,
		rr.CreateUserIp,
		rr.CreateUserId,
	)
	if err != nil {
		log.Error("[server.AddPattern] r.db.AddPattern", "error", err)
		return &pb.PatternMessage{Code: int32(utils.CodeInternal), Pattern: &pb.Pattern{}}, models.InternalError
	}

	return &pb.PatternMessage{
		Code: int32(utils.CodeOK),
		Pattern: &pb.Pattern{
			Id:           r.ID,
			Name:         r.Name,
			Img:          r.Img,
			X:            r.X,
			Y:            r.Y,
			CreateUserId: r.CreateUserID,
			CreateUserIp: r.CreateUserIP,
			UpdateUserId: r.UpdateUserID,
			UpdateUserIp: r.UpdateUserIP,
			CreatedAt:    r.CreatedAt.String(),
			UpdatedAt:    r.UpdatedAt.String(),
		},
	}, nil
}

func (m *MapsService) DeletePattern(ctx context.Context, rr *pb.Pattern) (*pb.PatternMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeletePattern")
	defer span.End()

	err := m.db.DeletePattern(rr.Id)
	if err != nil {
		log.Error("[server.DeletePattern] m.db.DeletePattern", "error", err)
		return &pb.PatternMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.PatternMessage{
		Code: int32(utils.CodeOK),
	}, nil
}
