package server

import (
	"comet/utils"
	"context"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"maps-service/internal/models"
	pb "protos/maps"
	"strconv"
	"time"
)

func (m *MapsService) AddStyle(ctx context.Context, rr *pb.MStyle) (*pb.StyleMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddStyle")
	defer span.End()

	rr.CreatedAt = time.Now().String()

	model := models.FromMStyle(rr)
	r, err := m.db.AddStyle(model)

	if err != nil {
		log.Error("[server.AddStyle] m.db.AddStyle", "error", err)
		return &pb.StyleMessage{Code: int32(utils.CodeInternal), Style: &pb.MStyle{}}, models.InternalError
	}

	return &pb.StyleMessage{Code: int32(utils.CodeOK),
		Style: models.Style2Protobuf(*r)}, nil
}

func (m *MapsService) EditStyle(ctx context.Context, rr *pb.MStyle) (*pb.StyleMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "EditStyle")
	defer span.End()

	model := models.FromMStyle(rr)

	r, err := m.db.EditStyle(model)
	if err != nil {
		log.Error("[server.EditStyle] m.db.EditStyle", "error", err)
		return &pb.StyleMessage{Code: int32(utils.CodeInternal), Style: &pb.MStyle{}}, models.InternalError
	}

	return &pb.StyleMessage{Code: int32(utils.CodeOK),
		Style: models.Style2Protobuf(*r)}, nil
}

func (m *MapsService) DeleteStyle(ctx context.Context, rr *pb.MStyle) (*pb.StyleMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeleteStyle")
	defer span.End()

	err := m.db.DeleteStyle(rr.Id)
	if err != nil {
		log.Error("[server.DeleteStyle] m.db.DeleteStyle", "error", err)
		return &pb.StyleMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.StyleMessage{Code: int32(utils.CodeOK)}, nil
}

func (m *MapsService) Style(ctx context.Context, rr *pb.MStyle) (*pb.StyleMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "Style")
	defer span.End()

	style, err := m.db.Style(rr.Id)
	if err != nil {
		log.Error("[server.Style] m.db.Style", "error", err)
		return &pb.StyleMessage{Code: int32(utils.CodeInternal), Style: &pb.MStyle{}}, models.InternalError
	}

	return &pb.StyleMessage{Code: int32(utils.CodeOK),
		Style: models.Style2Protobuf(*style)}, nil
}

func (m *MapsService) Styles(ctx context.Context, _ *emptypb.Empty) (*pb.StylesMessage, error) {
	log := hclog.Default()
	var addStyleRes []*pb.MStyle
	tr := m.trace
	ctx, span := tr.Start(ctx, "Styles")
	defer span.End()

	r, err := m.db.Styles()
	if err != nil {
		log.Error("[server.Style] m.db.Style", "error", err)
		return &pb.StylesMessage{Code: int32(utils.CodeInternal), Styles: addStyleRes}, models.InternalError
	}

	for _, v := range *r {
		addStyleRes = append(addStyleRes, models.Style2Protobuf(v))
	}

	return &pb.StylesMessage{
		Code:   int32(utils.CodeOK),
		Styles: addStyleRes,
	}, nil
}

func (m *MapsService) StylesPagination(ctx context.Context, rr *pb.StylesPagination) (*pb.StylesMessage, error) {
	var styles []*pb.MStyle

	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "StylesPagination")
	defer span.End()

	p, err := strconv.Atoi(rr.Page)
	if err != nil {
		log.Error("[server.Style] strconv.Atoi(page)", "error", err)
		return &pb.StylesMessage{Code: int32(utils.CodeInternal), Styles: styles}, models.InternalError
	}

	pS, err := strconv.Atoi(rr.PageSize)
	if err != nil {
		log.Error("[server.Style] strconv.Atoi(pageSize)", "error", err)
		return &pb.StylesMessage{Code: int32(utils.CodeInternal), Styles: styles}, models.InternalError
	}

	r, err := m.db.StylesPagination(p, pS)
	if err != nil {
		log.Error("[server.Style] m.db.Style", "error", err)
		return &pb.StylesMessage{Code: int32(utils.CodeInternal), Styles: styles}, models.InternalError
	}

	for _, v := range *r {
		styles = append(styles, models.Style2Protobuf(v))
	}

	return &pb.StylesMessage{
		Code:   int32(utils.CodeOK),
		Styles: styles,
	}, nil
}
