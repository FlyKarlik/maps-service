package repository

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"maps-service/internal/models"
	"strings"
)

func (r *Repository) AddTable(
	name, alias, geometryType, tableType string,
	isRelated, isVersioned, IsArchived, isGeometryNullable bool,
	SRID int32,
) (*models.Table, error) {
	model := models.Table{
		Name:               name,
		Alias:              alias,
		GeometryType:       geometryType,
		TableType:          tableType,
		IsRelated:          isRelated,
		IsVersioned:        isVersioned,
		IsArchived:         IsArchived,
		IsGeometryNullable: isGeometryNullable,
		SRID:               SRID,
	}

	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &model, nil
}

func (r *Repository) Table(tableID string) (*models.Table, error) {
	var model models.Table

	result := r.DB.Where(models.Table{ID: tableID}).First(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.First error: %w", result.Error)
	}

	return &model, nil
}

func (r *Repository) Tables() (*[]models.Table, error) {
	var model []models.Table

	result := r.DB.Find(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &model, nil
}

func (r *Repository) DeleteTable(tableID string) error {
	result := r.DB.Delete(&models.Table{ID: tableID})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete error: %w", result.Error)
	}

	return nil
}

func (r *Repository) EditTable(
	id, name, alias, geometryType, tableType string,
	isRelated, isVersioned, IsArchived, isGeometryNullable bool,
	SRID int32) (*models.Table, error) {

	model := &models.Table{ID: id}

	result := r.DB.Model(model).Updates(models.Table{
		Name:               name,
		Alias:              alias,
		GeometryType:       geometryType,
		TableType:          tableType,
		IsRelated:          isRelated,
		IsVersioned:        isVersioned,
		IsArchived:         IsArchived,
		IsGeometryNullable: isGeometryNullable,
		SRID:               SRID,
	})

	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Updates error: %w", result.Error)
	}

	result = r.DB.Where("id = ?", id).First(model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.First error: %w", result.Error)
	}

	return model, nil
}

func (r *Repository) TableColumns(tableID string) ([]*models.Column, error) {
	var columns []*models.Column
	var table models.Table

	result := r.DB.Where(models.Table{ID: tableID}).First(&table)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Where.First failed: %w", result.Error)
	}

	result = r.DB.Where(models.Column{TableID: table.ID}).Find(&columns)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Where.Find failed: %w", result.Error)
	}

	return columns, nil
}

func (r *Repository) TableColumnUniqueValues(tableName, columnName string) (*[]models.UniqueValue, string, error) {
	var values []models.UniqueValue
	var column models.Column
	var table models.Table

	result := r.DB.Where(models.Table{Name: tableName}).First(&table)
	if result.Error != nil {
		return nil, "", fmt.Errorf("r.DB.Where.First failed: %w", result.Error)
	}

	result = r.DB.Where(models.Column{TableID: table.ID, Name: columnName}).First(&column)
	if result.Error != nil {
		return nil, "", fmt.Errorf("r.DB.Where.Find failed: %w", result.Error)
	}

	err := r.DbSP.Ping()
	if err != nil {
		if err != nil {
			return nil, "", fmt.Errorf("r.DbSP.Ping failed: %w", result.Error)
		}
	}

	err = r.DbSP.Select(
		&values,
		fmt.Sprintf(GetColumnUniqueValuesQuery, columnName, "ippd."+table.Name),
	)
	if err != nil {
		return nil, "", fmt.Errorf("r.DbSP.Select failed: %w", result.Error)
	}

	return &values, column.DataType, nil
}

// TableFeatures ищет пересечения во всех таблицах представленных на карте и выдает пересекающиеся обекты в виде JSON.
func (r *Repository) TableFeatures(layers string, xMin, yMin, xMax, yMax float32) (*[]models.TableFeature, error) {
	var features []models.TableFeature

	log := hclog.Default()

	tables := strings.Split(layers, ",")

	for _, v := range tables {
		q, err := GetTableFeatureIntersectWithPolygon(v, xMin, yMin, xMax, yMax)
		if err != nil {
			log.Warn("[repository.TableFeatures] GetTableFeatureIntersectWithPolygon failed", "error", err)
			continue
		}

		var feature models.TableFeature
		err = r.DbSP.Get(&feature, q)
		if err != nil {
			log.Warn("[repository.TableFeatures] r.DbSP.Get failed", "error", err)
			continue
		}

		var layer models.Layer
		result := r.DB.Model(models.Layer{}).Where("table_id = ?", "ippd."+v).First(&layer)
		if result.Error != nil {
			log.Warn("[repository.TableFeatures] r.DB.Model.Where.First failed", "error", err)
			continue
		}

		var table models.Table
		result = r.DB.Model(models.Table{}).Where("name = ?", v).First(&table)
		if result.Error != nil {
			log.Warn("[repository.TableFeatures] r.DB.Model.Where.First table failed", "error", err)
			continue
		}

		var columns []models.Column
		result = r.DB.Where(models.Column{TableID: table.ID}).Find(&columns)
		if result.Error != nil {
			log.Warn("[repository.TableFeatures] r.DB.Find columns failed", "error", err)
			continue
		}

		Alias := make(map[string]string)
		for _, v := range columns {
			Alias[v.Name] = v.Alias
		}

		feature.Alias = Alias
		feature.Layer = layer.Name
		feature.LayerFeature = strings.Join([]string{"ippd", v}, ".")

		features = append(features, feature)
	}

	return &features, nil
}
