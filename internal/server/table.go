package server

import (
	"comet/utils"
	"context"
	"github.com/hashicorp/go-hclog"
	"google.golang.org/genproto/googleapis/rpc/code"
	"maps-service/internal/models"
	pb "protos/maps"
)

// AddTable ...
func (m *MapsService) AddTable(ctx context.Context, rr *pb.Table) (*pb.TableMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "AddTable")
	defer span.End()

	r, err := m.db.AddTable(
		rr.Name,
		rr.Alias,
		rr.GeometryType,
		rr.TableType,
		rr.IsRelated,
		rr.IsVersioned,
		rr.IsArchived,
		rr.IsGeometryNullable,
		rr.Srid,
	)

	if err != nil {
		log.Error("[server.AddTable] m.db.AddTable", "error", err)
		return &pb.TableMessage{Code: int32(utils.CodeInternal), Table: &pb.Table{}}, models.InternalError
	}

	return &pb.TableMessage{
		Code: int32(utils.CodeOK),
		Table: &pb.Table{
			Id:                 r.ID,
			Name:               r.Name,
			Alias:              r.Alias,
			GeometryType:       r.GeometryType,
			TableType:          r.TableType,
			IsRelated:          r.IsRelated,
			IsVersioned:        r.IsVersioned,
			IsArchived:         r.IsArchived,
			IsGeometryNullable: r.IsGeometryNullable,
			Srid:               r.SRID,
		},
	}, nil
}

// Table ...
func (m *MapsService) Table(ctx context.Context, rr *pb.Table) (*pb.TableMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "Table")
	defer span.End()

	r, err := m.db.Table(rr.Id)
	if err != nil {
		log.Error("[server.Table] m.db.Table", "error", err)
		return &pb.TableMessage{Code: int32(utils.CodeInternal), Table: &pb.Table{}}, models.InternalError
	}

	return &pb.TableMessage{
		Code: int32(utils.CodeOK),
		Table: &pb.Table{
			Id:                 r.ID,
			Name:               r.Name,
			Alias:              r.Alias,
			GeometryType:       r.GeometryType,
			TableType:          r.TableType,
			IsRelated:          r.IsRelated,
			IsVersioned:        r.IsVersioned,
			IsArchived:         r.IsArchived,
			IsGeometryNullable: r.IsGeometryNullable,
			Srid:               r.SRID,
		},
	}, nil
}

// Tables ...
func (m *MapsService) Tables(ctx context.Context) (*pb.TablesMessage, error) {
	log := hclog.Default()
	var response []*pb.Table
	tr := m.trace
	ctx, span := tr.Start(ctx, "Tables")
	defer span.End()

	r, err := m.db.Tables()
	if err != nil {
		log.Error("[server.Tables] m.db.Tables", "error", err)
		return &pb.TablesMessage{Code: int32(utils.CodeInternal), Tables: response}, models.InternalError
	}

	for _, v := range *r {
		response = append(response, &pb.Table{
			Id:                 v.ID,
			Name:               v.Name,
			Alias:              v.Alias,
			GeometryType:       v.GeometryType,
			TableType:          v.TableType,
			IsRelated:          v.IsRelated,
			IsVersioned:        v.IsVersioned,
			IsArchived:         v.IsArchived,
			IsGeometryNullable: v.IsGeometryNullable,
			Srid:               v.SRID,
		})
	}

	return &pb.TablesMessage{
		Code:   int32(utils.CodeOK),
		Tables: response,
	}, nil
}

// DeleteTable ...
func (m *MapsService) DeleteTable(ctx context.Context, rr *pb.Table) (*pb.TableMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "DeleteTable")
	defer span.End()

	err := m.db.DeleteTable(rr.Id)
	if err != nil {
		log.Error("[server.DeleteTable] m.db.DeleteTable", "error", err)
		return &pb.TableMessage{Code: int32(utils.CodeInternal)}, models.InternalError
	}

	return &pb.TableMessage{Code: int32(utils.CodeOK)}, nil
}

// EditTable ...
func (m *MapsService) EditTable(ctx context.Context, rr *pb.Table) (*pb.TableMessage, error) {
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "EditTable")
	defer span.End()

	r, err := m.db.EditTable(
		rr.Id,
		rr.Name,
		rr.Alias,
		rr.GeometryType,
		rr.TableType,
		rr.IsRelated,
		rr.IsVersioned,
		rr.IsArchived,
		rr.IsGeometryNullable,
		rr.Srid,
	)

	if err != nil {
		log.Error("[server.EditTable] m.db.EditTable", "error", err)
		return &pb.TableMessage{Code: int32(utils.CodeInternal), Table: &pb.Table{}}, models.InternalError
	}

	return &pb.TableMessage{
		Code: int32(utils.CodeOK),
		Table: &pb.Table{
			Id:                 r.ID,
			Name:               r.Name,
			Alias:              r.Alias,
			GeometryType:       r.GeometryType,
			TableType:          r.TableType,
			IsRelated:          r.IsRelated,
			IsVersioned:        r.IsVersioned,
			IsArchived:         r.IsArchived,
			IsGeometryNullable: r.IsGeometryNullable,
			Srid:               r.SRID,
		},
	}, nil
}

func (m *MapsService) TableColumns(ctx context.Context, rr *pb.Table) (*pb.ColumnsMessage, error) {
	var response []*pb.Column
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "TableColumns")
	defer span.End()

	r, err := m.db.TableColumns(rr.Name)
	if err != nil {
		log.Error("[server.TableColumns] m.db.TableColumns", "error", err)
		return &pb.ColumnsMessage{Code: int32(utils.CodeInternal), Columns: response}, err
	}

	for _, v := range r {
		response = append(response, &pb.Column{
			Id:           v.ID,
			TableId:      v.TableID,
			DomainId:     &v.DomainID,
			Name:         v.Name,
			Alias:        v.Alias,
			DataType:     v.DataType,
			Nullable:     v.Nullable,
			Length:       int32(v.Length),
			DefaultValue: v.DefaultValue,
		})
	}

	return &pb.ColumnsMessage{
		Code:    int32(utils.CodeOK),
		Columns: response,
	}, nil
}

func (m *MapsService) TableColumnUniqueValues(ctx context.Context, rr *pb.ColumnUnique) (*pb.ColumnUniqueMessage, error) {
	var response []string
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "TableColumnUniqueValues")
	defer span.End()

	r, dataType, err := m.db.TableColumnUniqueValues(rr.TableName, rr.ColumnName)
	if err != nil {
		log.Error("[server.TableColumnUniqueValues] m.db.TableColumnUniqueValues", "error", err)
		return &pb.ColumnUniqueMessage{Code: int32(utils.CodeInternal), Unique: &pb.ColumnUnique{}}, err
	}

	for _, v := range *r {
		response = append(response, v.Value)
	}

	return &pb.ColumnUniqueMessage{
		Code: int32(code.Code_OK),
		Unique: &pb.ColumnUnique{
			Unique:     response,
			DataType:   dataType,
			ColumnName: rr.ColumnName,
		},
	}, nil
}

func (m *MapsService) TableFeatures(ctx context.Context, rr *pb.TableFeaturesRequest) (*pb.TableFeatureMessage, error) {
	var response []*pb.TableFeature
	log := hclog.Default()

	tr := m.trace
	ctx, span := tr.Start(ctx, "TableFeatures")
	defer span.End()

	r, err := m.db.TableFeatures(rr.Layers, rr.Xmin, rr.Ymin, rr.Xmax, rr.Ymax)
	if err != nil {
		log.Error("[server.TableFeatures] m.db.TableFeatures", "error", err)
		return &pb.TableFeatureMessage{Code: int32(utils.CodeInternal), Features: response}, err
	}

	for _, v := range *r {
		response = append(response, &pb.TableFeature{
			Feature:      v.Feature,
			Geometry:     v.Geometry,
			Layer:        v.Layer,
			Alias:        v.Alias,
			LayerFeature: v.LayerFeature,
		})
	}

	return &pb.TableFeatureMessage{
		Code:     int32(utils.CodeOK),
		Features: response,
	}, nil
}
