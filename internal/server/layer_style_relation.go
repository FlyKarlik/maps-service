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

func (m *MapsService) AddLayerStyleRelation(ctx context.Context, rr *pb.LSRelation) (*pb.LSRMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddLayerStyleRelation")
	defer span.End()

	r, err := m.db.AddLayerStyleRelation(rr.LayerId, rr.StyleId)
	if err != nil {
		log.Error("[m.AddLayerStyleRelation] m.db.AddLayerStyleRelation", "error", err)
		return &pb.LSRMessage{Code: int32(utils.CodeInternal), Relation: &pb.LSRelation{}}, models.InternalError
	}

	return &pb.LSRMessage{
		Code: int32(utils.CodeOK),
		Relation: &pb.LSRelation{
			Id:      r.ID,
			LayerId: r.LayerID,
			StyleId: r.StyleID,
		},
	}, nil
}

func (m *MapsService) DeleteLayerStyleRelation(ctx context.Context, rr *pb.LSRelation) (*pb.LSRMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeleteLayerStyleRelation")
	defer span.End()

	err := m.db.DeleteLayerStyleRelation(rr.Id)
	if err != nil {
		log.Error("[m.DeleteLayerStyleRelation] m.db.DeleteLayerStyleRelation", "error", err)
		return &pb.LSRMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.LSRMessage{
		Code: int32(utils.CodeOK),
	}, nil
}

func (m *MapsService) LayerStyleRelations(ctx context.Context, _ *emptypb.Empty) (*pb.LSRsMessage, error) {
	var response []*pb.LSRelation
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "LayerStyleRelations")
	defer span.End()

	r, err := m.db.LayerStyleRelations()
	if err != nil {
		log.Error("[m.LayerStyleRelations] m.db.LayerStyleRelations", "error", err)
		return &pb.LSRsMessage{Code: int32(utils.CodeInternal), Relations: response}, models.InternalError
	}

	if len(*r) < 1 {
		log.Error("[m.LayerStyleRelations] len(*r)", "error", fmt.Errorf("r length less 1"))
		return &pb.LSRsMessage{Code: int32(utils.CodeNotFound), Relations: response}, models.NotFoundError
	}

	for _, v := range *r {
		response = append(response, &pb.LSRelation{
			Id:      v.ID,
			LayerId: v.LayerID,
			StyleId: v.StyleID,
		})
	}

	return &pb.LSRsMessage{
		Code:      int32(utils.CodeOK),
		Relations: response,
	}, nil
}

func (m *MapsService) LayerRelationStyles(ctx context.Context, rr *pb.MLayer) (*pb.LRSMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "LayerRelationStyles")
	defer span.End()

	r, err := m.db.LayerRelationStyles(rr.Id)
	if err != nil {
		log.Error("[m.LayerRelationStyles] m.db.LayerRelationStyles", "error", err)
		return &pb.LRSMessage{Code: int32(utils.CodeInternal), Style: models.Style2Protobuf(*r)}, models.InternalError
	}

	return &pb.LRSMessage{
		Code:  int32(utils.CodeOK),
		Style: models.Style2Protobuf(*r),
	}, nil
}

func (m *MapsService) StyleRelationLayers(ctx context.Context, rr *pb.MStyle) (*pb.SRLMessage, error) {
	log := hclog.Default()
	var response []*pb.MLayer
	tr := m.trace
	ctx, span := tr.Start(ctx, "StyleRelationLayers")
	defer span.End()

	r, err := m.db.StyleRelationLayers(rr.Id)
	if err != nil {
		log.Error("[m.StyleRelationLayers] m.db.StyleRelationLayers", "error", err)
		return &pb.SRLMessage{Code: int32(utils.CodeOK), Layers: response}, models.InternalError
	}

	for _, m_ := range *r {
		response = append(response, &pb.MLayer{
			Id:           m_.ID,
			Name:         m_.Name,
			TableId:      m_.TableID,
			LayerType:    m_.LayerType,
			CreateUserId: m_.CreateUserID,
			CreateUserIp: m_.CreateUserIP,
			CreatedAt:    m_.CreatedAt.String(),
			UpdateUserId: m_.UpdateUserID,
			UpdateUserIp: m_.UpdateUserIP,
			UpdatedAt:    m_.UpdatedAt.String(),
		})
	}

	return &pb.SRLMessage{
		Code:   int32(utils.CodeOK),
		Layers: response,
	}, nil
}
