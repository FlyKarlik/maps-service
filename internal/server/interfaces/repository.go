package interfaces

import (
	"maps-service/config"
	"maps-service/internal/models"
	pb "protos/maps"
)

type Repository interface {
	Layer(id string) (*models.Layer, error)
	Layers() (*[]models.Layer, error)
	AddLayer(name, layerType, tableID, createUserIp, createUserID string) (*models.Layer, error)
	EditLayer(id, name, layerType, tableID, updateUserIp, updateUserID string) (*models.Layer, error)
	DeleteLayer(id string) error

	AddGroup(name, createUserIp, createUserID string) (*models.Group, error)
	EditGroup(id, name, updateUserIp, updateUserID string) (*models.Group, error)
	DeleteGroup(id string) error
	Group(id string) (*models.Group, error)
	Groups() (*[]models.Group, error)

	Map(id string) (*models.Map, error)
	Maps() (*[]models.Map, error)
	AddMap(name, picture, describe, createUserIp, createUserID string, active bool) (*models.Map, error)
	EditMap(id, name, picture, describe, updateUserIp, updateUserID string, active bool) (*models.Map, error)
	DeleteMap(id string) error

	AddGroupLayerRelation(groupID, layerID string) (*models.GroupLayerRelation, error)
	DeleteGroupLayerRelation(groupID, layerID string) error
	GroupLayerRelations() (*[]models.GroupLayerRelation, error)
	LayerRelationGroups(layerID string) (*[]models.Group, error)
	GroupRelationLayers(groupID string) (*[]models.Layer, error)
	GroupLayerOrderDown(id string) (*models.GroupLayerRelation, error)
	GroupLayerOrderUp(id string) (*models.GroupLayerRelation, error)

	AddMapGroupRelation(mapID, groupID string) (*models.MapGroupRelation, error)
	DeleteMapGroupRelation(mapID, groupID string) error
	MapGroupRelations() (*[]models.MapGroupRelation, error)
	MapRelationGroups(mapID string) (*[]models.Group, error)
	GroupRelationMaps(groupID string) (*[]models.Map, error)
	MapGroupOrderDown(id string) (*models.MapGroupRelation, error)
	MapGroupOrderUp(id string) (*models.MapGroupRelation, error)

	AddStyle(s models.Style) (*models.Style, error)
	EditStyle(s models.Style) (*models.Style, error)
	DeleteStyle(id string) error
	Style(id string) (*models.Style, error)
	Styles() (*[]models.Style, error)
	StylesPagination(page, pageSize int) (*[]models.Style, error)

	AddLayerStyleRelation(layerID, styleID string) (*models.LayerStyleRelation, error)
	DeleteLayerStyleRelation(id string) error
	LayerStyleRelations() (*[]models.LayerStyleRelation, error)
	LayerRelationStyles(layerID string) (*models.Style, error)
	StyleRelationLayers(styleID string) (*[]models.Layer, error)

	StyledMap(mapID string, cfg *config.Config) (*pb.StyledMap, error)

	Pattern(patternID string) (*models.Pattern, error)
	Patterns() (*[]models.Pattern, error)
	AddPattern(name, img string, x, y int32, createUserID, createUserIP string) (*models.Pattern, error)
	DeletePattern(patternID string) error

	AddTable(
		name, alias, geometryType, tableType string,
		isRelated, isVersioned, IsArchived, isGeometryNullable bool,
		SRID int32,
	) (*models.Table, error)
	Table(tableID string) (*models.Table, error)
	Tables() (*[]models.Table, error)
	EditTable(
		id, name, alias, geometryType, tableType string,
		isRelated, isVersioned, IsArchived, isGeometryNullable bool,
		SRID int32,
	) (*models.Table, error)
	TableColumns(tableID string) ([]*models.Column, error)
	DeleteTable(tableID string) error
	TableColumnUniqueValues(tableName, columnName string) (*[]models.UniqueValue, string, error)
	TableFeatures(layers string, xMin, yMin, xMax, yMax float32) (*[]models.TableFeature, error)
}
