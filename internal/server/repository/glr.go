package repository

import (
	"fmt"
	"gorm.io/gorm"
	"maps-service/internal/models"
)

// AddGroupLayerRelation add new relation
func (r *Repository) AddGroupLayerRelation(groupID, layerID string) (*models.GroupLayerRelation, error) {
	var relations []models.GroupLayerRelation

	result := r.DB.Find(&relations, models.GroupLayerRelation{GroupID: groupID})
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	model := models.GroupLayerRelation{
		GroupID:    groupID,
		LayerID:    layerID,
		LayerOrder: int32(1 + len(relations)),
	}

	result = r.DB.Create(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &model, nil
}

// DeleteGroupLayerRelation delete relation by id
func (r *Repository) DeleteGroupLayerRelation(groupID, layerID string) error {
	var model models.GroupLayerRelation
	var relations []models.GroupLayerRelation

	result := r.DB.Where(models.GroupLayerRelation{GroupID: groupID, LayerID: layerID}).First(&model)
	if result.Error != nil {
		return fmt.Errorf("r.DB.Where(model) error: %w", result.Error)
	}

	result = r.DB.Find(&relations, models.GroupLayerRelation{GroupID: model.GroupID})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Find(relations) error: %w", result.Error)
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {

		for i := model.LayerOrder; i < int32(len(relations)); i++ {
			result = tx.Model(&relations[i]).Update("layer_order", relations[i].LayerOrder-1)

			if result.Error != nil {
				return fmt.Errorf("tx.Model(&relations[i]).Update error: %w", result.Error)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("r.DB.Transaction error: %w", result.Error)
	}

	result = r.DB.Delete(&models.GroupLayerRelation{ID: model.ID})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete error: %w", result.Error)
	}

	return nil
}

// GroupLayerRelations get add relations
func (r *Repository) GroupLayerRelations() (*[]models.GroupLayerRelation, error) {
	var relations []models.GroupLayerRelation

	result := r.DB.Find(&relations)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &relations, nil
}

// LayerRelationGroups get layer`s groups
func (r *Repository) LayerRelationGroups(layerID string) (*[]models.Group, error) {
	var resultGroups []models.Group

	result := r.DB.Raw(layerGroupRelationJoins, layerID).Scan(&resultGroups)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Raw error: %w", result.Error)
	}

	return &resultGroups, nil
}

// GroupRelationLayers get group`s layers
func (r *Repository) GroupRelationLayers(groupID string) (*[]models.Layer, error) {
	var resultLayers []models.Layer

	result := r.DB.Raw(groupRelationLayersJoinLayer, groupID).Scan(&resultLayers)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Raw error: %w", result.Error)
	}

	return &resultLayers, nil
}

// GroupLayerOrderUp up model order, return error if model is last element in the group
func (r *Repository) GroupLayerOrderUp(id string) (*models.GroupLayerRelation, error) {
	var model models.GroupLayerRelation
	var relations []models.GroupLayerRelation
	var nextModel models.GroupLayerRelation

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where(models.GroupLayerRelation{ID: id}).First(&model).Error
		if err != nil {
			return fmt.Errorf("tx.Where(model) failed: %w", err)
		}

		err = tx.Find(&relations, models.GroupLayerRelation{GroupID: model.GroupID}).Error
		if err != nil {
			return fmt.Errorf("tx.Find(&relations, models.GroupLayerRelation) failed: %w", err)
		}

		if model.LayerOrder+1 > int32(len(relations)) {
			return fmt.Errorf("model is the last element")
		}

		err = tx.Where(models.GroupLayerRelation{GroupID: model.GroupID, LayerOrder: model.LayerOrder + 1}).First(&nextModel).Error
		if err != nil {
			return fmt.Errorf("tx.Where().First(nextModel) failed: %w", err)
		}

		err = tx.Model(&model).Update("layer_order", model.LayerOrder+1).Error
		if err != nil {
			return fmt.Errorf("tx.Model.Update(model) failed: %w", err)
		}

		err = tx.Model(&nextModel).Update("layer_order", nextModel.LayerOrder-1).Error
		if err != nil {
			return fmt.Errorf("tx.Model.Update(netxModel) failed: %w", err)
		}

		return nil

	})
	if err != nil {
		return nil, fmt.Errorf("r.DB.Transaction error: %w", err)
	}

	return &model, nil
}

// GroupLayerOrderDown down model order, return error if model is first element in the group
func (r *Repository) GroupLayerOrderDown(id string) (*models.GroupLayerRelation, error) {
	var model models.GroupLayerRelation
	var relations []models.GroupLayerRelation
	var nextModel models.GroupLayerRelation

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where(models.GroupLayerRelation{ID: id}).First(&model).Error
		if err != nil {
			return fmt.Errorf("tx.Where(model) failed: %w", err)
		}

		err = tx.Find(&relations, models.GroupLayerRelation{GroupID: model.GroupID}).Error
		if err != nil {
			return fmt.Errorf("tx.Find(&relations, models.GroupLayerRelation) failed: %w", err)
		}

		if model.LayerOrder-1 < 1 {
			return fmt.Errorf("model is the first element")
		}

		err = tx.Where(models.GroupLayerRelation{GroupID: model.GroupID, LayerOrder: model.LayerOrder - 1}).First(&nextModel).Error
		if err != nil {
			return fmt.Errorf("tx.Where().First(nextModel) failed: %w", err)
		}

		err = tx.Model(&model).Update("layer_order", model.LayerOrder-1).Error
		if err != nil {
			return fmt.Errorf("tx.Model.Update(model) failed: %w", err)
		}

		err = tx.Model(&nextModel).Update("layer_order", nextModel.LayerOrder+1).Error
		if err != nil {
			return fmt.Errorf("tx.Model.Update(netxModel) failed: %w", err)
		}

		return nil

	})
	if err != nil {
		return nil, fmt.Errorf("r.DB.Transaction error: %w", err)
	}

	return &model, nil
}
