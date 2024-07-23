package repository

import (
	"fmt"
	"maps-service/internal/models"
	"time"
)

// AddStyle add new style
func (r *Repository) AddStyle(s models.Style) (*models.Style, error) {
	s.CreatedAt = time.Now()

	result := r.DB.Create(&s)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &s, nil
}

// EditStyle update style by id
func (r *Repository) EditStyle(s models.Style) (*models.Style, error) {

	s.UpdatedAt = time.Now()

	result := r.DB.Model(&s).Updates(s)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Create error: %w", result.Error)
	}

	return &s, nil
}

// DeleteStyle delete style by id
func (r *Repository) DeleteStyle(id string) error {
	result := r.DB.Delete(&models.Style{ID: id})
	if result.Error != nil {
		return fmt.Errorf("r.DB.Delete error: %w", result.Error)
	}

	return nil
}

// Style get style by id
func (r *Repository) Style(id string) (*models.Style, error) {
	var style models.Style

	result := r.DB.Where("id = ?", id).First(&style)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.First error: %w", result.Error)
	}

	return &style, nil
}

// Styles get all styles
func (r *Repository) Styles() (*[]models.Style, error) {
	var styles []models.Style

	result := r.DB.Find(&styles)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Find error: %w", result.Error)
	}

	return &styles, nil
}

func (r *Repository) StylesPagination(page, pageSize int) (*[]models.Style, error) {
	var styles []models.Style

	fmt.Println("PageSize", pageSize, "page", page)

	result := r.DB.Offset(pageSize * (page - 1)).Limit(pageSize).Find(&styles)
	if result.Error != nil {
		return nil, fmt.Errorf("r.DB.Offset(pageSize * (page - 1)).Limit(pageSize).Find(&styles) error: %w", result.Error)
	}

	return &styles, nil
}
