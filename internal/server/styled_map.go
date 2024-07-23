package server

import (
	"comet/utils"
	"context"
	"github.com/hashicorp/go-hclog"
	"maps-service/internal/models"
	pb "protos/maps"
)

func (m *MapsService) StyledMap(ctx context.Context, rr *pb.MMap) (*pb.StyledMapMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "StyledMap")
	defer span.End()

	r, err := m.rdb.StyledMapsRepository.Get(rr.Id)
	if err != nil {
		log.Error("[server.StyledMap] m.rdb.StyledMapsRepository.Get", "error", err)
	} else {
		return &pb.StyledMapMessage{
			Code: int32(utils.CodeOK),
			Map:  r,
		}, nil
	}

	r, err = m.db.StyledMap(rr.Id, m.cfg)
	if err != nil {
		log.Error("[server.StyledMap] m.db.StyledMap", "error", err)
		return &pb.StyledMapMessage{Code: int32(utils.CodeInternal), Map: &pb.StyledMap{}}, models.InternalError
	}

	if err := m.rdb.StyledMapsRepository.Set(r); err != nil {
		log.Error("[server.StyledMap] m.rdb.StyledMapsRepository.Set", "error", err)
	}

	return &pb.StyledMapMessage{
		Code: int32(utils.CodeOK),
		Map:  r,
	}, nil
}
