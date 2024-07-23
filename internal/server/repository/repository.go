package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
	"maps-service/internal/models"
	"time"
)

type Repository struct {
	DB   *gorm.DB
	DbSP *sqlx.DB
}

// NewRepository create new Repository
func NewRepository(db *gorm.DB, dbSP *sqlx.DB) *Repository {
	return &Repository{
		DB:   db,
		DbSP: dbSP,
	}
}

// Layer get layer by id
func (r *Repository) Layer(id string) (*models.Layer, error) {
	var resultLayer models.Layer

	result := r.DB.Where(models.Layer{ID: id}).First(&resultLayer)

	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.First error: %w", result.Error)
	}

	return &resultLayer, nil
}

// Layers get all layers
func (r *Repository) Layers() (*[]models.Layer, error) {
	var resultLayers []models.Layer

	result := r.DB.Find(&resultLayers)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &resultLayers, nil
}

// AddLayer add new layer
func (r *Repository) AddLayer(name, layerType, tableID, createUserIp, createUserID string) (*models.Layer, error) {
	model := models.Layer{
		Name:         name,
		LayerType:    layerType,
		TableID:      tableID,
		CreateUserIP: createUserIp,
		CreateUserID: createUserID,
		CreatedAt:    time.Now(),
	}
	result := r.DB.Create(&model)

	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &model, nil
}

// EditLayer edit layer by id
func (r *Repository) EditLayer(id, name, layerType, tableID, updateUserIp, updateUserID string) (*models.Layer, error) {
	model := &models.Layer{
		ID: id,
	}
	result := r.DB.Model(model).Updates(models.Layer{
		Name:         name,
		LayerType:    layerType,
		TableID:      tableID,
		UpdateUserIP: updateUserIp,
		UpdateUserID: updateUserID,
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

// DeleteLayer delete layer by id
func (r *Repository) DeleteLayer(id string) error {
	result := r.DB.Delete(&models.Layer{ID: id})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete error: %w", result.Error)
	}

	return nil
}

// AddGroup add new group
func (r *Repository) AddGroup(name, createUserIp, createUserID string) (*models.Group, error) {
	model := models.Group{
		Name:         name,
		CreateUserID: createUserID,
		CreateUserIP: createUserIp,
		CreatedAt:    time.Now(),
	}

	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &model, nil
}

// EditGroup edit group by id
func (r *Repository) EditGroup(id, name, updateUserIp, updateUserID string) (*models.Group, error) {
	model := models.Group{
		ID:           id,
		Name:         name,
		UpdateUserIP: updateUserIp,
		UpdateUserID: updateUserID,
		UpdatedAt:    time.Now(),
	}

	result := r.DB.Model(&model).Updates(model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Updates error: %w", result.Error)
	}

	result = r.DB.Where("id = ?", id).First(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.First error: %w", result.Error)
	}

	return &model, nil
}

// DeleteGroup delete group by id
func (r *Repository) DeleteGroup(id string) error {
	result := r.DB.Delete(&models.Group{ID: id})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete error: %w", result.Error)
	}

	return nil
}

// Group get group by id
func (r *Repository) Group(id string) (*models.Group, error) {
	var resultGroup models.Group

	result := r.DB.Where(models.Group{ID: id}).First(&resultGroup)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.First error: %w", result.Error)
	}

	return &resultGroup, nil
}

// Groups get all groups
func (r *Repository) Groups() (*[]models.Group, error) {
	var resultGroups []models.Group

	result := r.DB.Find(&resultGroups)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &resultGroups, nil
}

// Map get map by id
func (r *Repository) Map(id string) (*models.Map, error) {
	var resultMap models.Map

	result := r.DB.Where("id = ?", id).First(&resultMap)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.First error: %w", result.Error)
	}

	return &resultMap, nil
}

// Maps get all maps
func (r *Repository) Maps() (*[]models.Map, error) {
	var resultMaps []models.Map

	result := r.DB.Find(&resultMaps)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &resultMaps, nil
}

// AddMap add map
func (r *Repository) AddMap(name, picture, describe, createUserIp, createUserID string, active bool) (*models.Map, error) {
	model := models.Map{
		Name:         name,
		Picture:      picture,
		Describe:     describe,
		CreateUserIP: createUserIp,
		CreateUserID: createUserID,
		Active:       active,
		CreatedAt:    time.Now(),
	}

	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &model, nil
}

// EditMap update map by id
func (r *Repository) EditMap(id, name, picture, describe, updateUserIp, updateUserID string, active bool) (*models.Map, error) {
	model := models.Map{
		ID:           id,
		Name:         name,
		Picture:      picture,
		Describe:     describe,
		Active:       active,
		UpdateUserIP: updateUserIp,
		UpdateUserID: updateUserID,
		UpdatedAt:    time.Now(),
	}

	result := r.DB.Model(&model).Updates(model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Updates error: %w", result.Error)
	}

	return &model, nil
}

// DeleteMap delete map by id
func (r *Repository) DeleteMap(id string) error {
	result := r.DB.Delete(&models.Map{ID: id})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete error: %w", result.Error)
	}

	return nil
}

// AddLayerStyleRelation add new relation
func (r *Repository) AddLayerStyleRelation(layerID, styleID string) (*models.LayerStyleRelation, error) {
	model := models.LayerStyleRelation{
		LayerID: layerID,
		StyleID: styleID,
	}

	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &model, nil
}

// DeleteLayerStyleRelation  delete relation by id
func (r *Repository) DeleteLayerStyleRelation(id string) error {
	result := r.DB.Delete(&models.LayerStyleRelation{ID: id})

	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete error: %w", result.Error)
	}

	return nil
}

// LayerStyleRelations get all relation
func (r *Repository) LayerStyleRelations() (*[]models.LayerStyleRelation, error) {
	var relations []models.LayerStyleRelation

	result := r.DB.Find(&relations)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &relations, nil
}

// LayerRelationStyles  get layer`s styles
func (r *Repository) LayerRelationStyles(layerID string) (*models.Style, error) {
	var style models.Style

	result := r.DB.Raw(layerRelationStylesJoinStyle, layerID).Scan(&style)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Raw.Scan")
	}

	return &style, nil
}

// StyleRelationLayers get style`s layers
func (r *Repository) StyleRelationLayers(styleID string) (*[]models.Layer, error) {
	var layers []models.Layer

	result := r.DB.Raw(styleRelationLayersJoinLayer, styleID).Scan(&layers)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Raw.Scan error: %w", result.Error)
	}

	return &layers, nil
}
