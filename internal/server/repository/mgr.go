package repository

import (
	"fmt"
	"gorm.io/gorm"
	"maps-service/internal/models"
)

// AddMapGroupRelation add new relation
func (r *Repository) AddMapGroupRelation(mapID, groupID string) (*models.MapGroupRelation, error) {
	var relations []models.MapGroupRelation

	result := r.DB.Find(&relations, models.MapGroupRelation{MapID: mapID})
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	model := models.MapGroupRelation{
		MapID:      mapID,
		GroupID:    groupID,
		GroupOrder: int32(len(relations) + 1),
	}

	result = r.DB.Create(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &model, nil
}

// DeleteMapGroupRelation delete relation by id
func (r *Repository) DeleteMapGroupRelation(mapID, groupID string) error {
	var model models.MapGroupRelation
	var relations []models.MapGroupRelation

	result := r.DB.Where(models.MapGroupRelation{MapID: mapID, GroupID: groupID}).First(&model)
	if result.Error != nil {
		return fmt.Errorf("r.DB.Where(model) error: %w", result.Error)
	}

	result = r.DB.Find(&relations, models.MapGroupRelation{MapID: model.MapID})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Find(relations) error: %w", result.Error)
	}

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		for i := model.GroupOrder; i < int32(len(relations)); i++ {
			result = tx.Model(&relations[i]).Update("group_order", relations[i].GroupOrder-1)

			if result.Error != nil {
				return fmt.Errorf("tx.Model(&relations[i]).Update error: %w", result.Error)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("r.DB.Transaction error: %w", result.Error)
	}

	result = r.DB.Delete(&models.MapGroupRelation{ID: model.ID})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete.Where error: %w", result.Error)
	}

	return nil
}

// MapGroupRelations get all relations
func (r *Repository) MapGroupRelations() (*[]models.MapGroupRelation, error) {
	var relations []models.MapGroupRelation

	result := r.DB.Find(&relations)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &relations, nil
}

// MapRelationGroups Map`s relation groups
func (r *Repository) MapRelationGroups(mapID string) (*[]models.Group, error) {
	var groups []models.Group

	result := r.DB.Raw(mapRelationGroupsJoinGroup, mapID).Scan(&groups)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Raw.Scan error: %w", result.Error)
	}

	return &groups, nil
}

// GroupRelationMaps group`s relation maps
func (r *Repository) GroupRelationMaps(groupID string) (*[]models.Map, error) {
	var maps []models.Map

	result := r.DB.Raw(groupRelationMapsJoinMap, groupID).Scan(&maps)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Raw.Scan error: %w", result.Error)
	}

	return &maps, nil
}

// MapGroupOrderUp up model order, return error if model is last element in the group
func (r *Repository) MapGroupOrderUp(id string) (*models.MapGroupRelation, error) {
	var model models.MapGroupRelation
	var relations []models.MapGroupRelation
	var nextModel models.MapGroupRelation

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where(models.MapGroupRelation{ID: id}).First(&model).Error
		if err != nil {
			return fmt.Errorf("tx.Where(model) failed: %w", err)
		}

		err = tx.Find(&relations, models.MapGroupRelation{MapID: model.MapID}).Error
		if err != nil {
			return fmt.Errorf("tx.Find(&relations, models.MapGroupRelation) failed: %w", err)
		}

		if model.GroupOrder+1 > int32(len(relations)) {
			return fmt.Errorf("model is the last element")
		}

		err = tx.Where(models.MapGroupRelation{MapID: model.MapID, GroupOrder: model.GroupOrder + 1}).First(&nextModel).Error
		if err != nil {
			return fmt.Errorf("tx.Where().First(nextModel) failed: %w", err)
		}

		err = tx.Model(&model).Update("group_order", model.GroupOrder+1).Error
		if err != nil {
			return fmt.Errorf("tx.Model.Update(model) failed: %w", err)
		}

		err = tx.Model(&nextModel).Update("group_order", nextModel.GroupOrder-1).Error
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

// MapGroupOrderDown down model order, return error if model is first element in the group
func (r *Repository) MapGroupOrderDown(id string) (*models.MapGroupRelation, error) {
	var model models.MapGroupRelation
	var relations []models.MapGroupRelation
	var nextModel models.MapGroupRelation

	err := r.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Where(models.MapGroupRelation{ID: id}).First(&model).Error
		if err != nil {
			return fmt.Errorf("tx.Where(model) failed: %w", err)
		}

		err = tx.Find(&relations, models.MapGroupRelation{MapID: model.MapID}).Error
		if err != nil {
			return fmt.Errorf("tx.Find(&relations, models.MapGroupRelation) failed: %w", err)
		}

		if model.GroupOrder-1 < 1 {
			return fmt.Errorf("model is the first element")
		}

		err = tx.Where(models.MapGroupRelation{MapID: model.MapID, GroupOrder: model.GroupOrder - 1}).First(&nextModel).Error
		if err != nil {
			return fmt.Errorf("tx.Where().First(nextModel) failed: %w", err)
		}

		err = tx.Model(&model).Update("group_order", model.GroupOrder-1).Error
		if err != nil {
			return fmt.Errorf("tx.Model.Update(model) failed: %w", err)
		}

		err = tx.Model(&nextModel).Update("group_order", nextModel.GroupOrder+1).Error
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
