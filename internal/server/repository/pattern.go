package repository

import (
	"fmt"
	"maps-service/internal/models"
)

func (r *Repository) Pattern(patternID string) (*models.Pattern, error) {
	var model models.Pattern

	result := r.DB.Where(models.Pattern{ID: patternID}).First(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.First error: %w", result.Error)
	}

	return &model, nil
}

func (r *Repository) Patterns() (*[]models.Pattern, error) {
	var model []models.Pattern

	result := r.DB.Find(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &model, nil
}

func (r *Repository) AddPattern(name, img string, x, y int32, createUserID, createUserIP string) (*models.Pattern, error) {
	model := models.Pattern{
		Name:         name,
		Img:          img,
		X:            x,
		Y:            y,
		CreateUserID: createUserID,
		CreateUserIP: createUserIP,
	}

	result := r.DB.Create(&model)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &model, nil
}

func (r *Repository) DeletePattern(patternID string) error {
	result := r.DB.Delete(&models.Pattern{ID: patternID})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete error: %w", result.Error)
	}

	return nil
}
