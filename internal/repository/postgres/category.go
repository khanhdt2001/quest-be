package postgres

import (
	"context"

	"github.com/quest-be/internal/repository/model"
)

type ICategories interface {
	InsertCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	UpdateCategory(ctx context.Context, category *model.Category) (*model.Category, error)
	FindCategoryById(ctx context.Context, id uint64) (*model.Category, error)
	FindCategoryByName(ctx context.Context, name string) (*model.Category, error)
	DeleteCategory(ctx context.Context, category *model.Category) error
}

func (r *Database) InsertCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	if err := r.Gorm.Create(&category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *Database) UpdateCategory(ctx context.Context, category *model.Category) (*model.Category, error) {
	if err := r.Gorm.Updates(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *Database) FindCategoryById(ctx context.Context, id uint64) (*model.Category, error) {
	category := &model.Category{}
	if err := r.Gorm.Where("id = ?", id).First(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *Database) FindCategoryByName(ctx context.Context, name string) (*model.Category, error) {
	category := &model.Category{}
	if err := r.Gorm.Where("name = ?", name).First(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *Database) DeleteCategory(ctx context.Context, category *model.Category) error {
	if err := r.Gorm.Delete(category).Error; err != nil {
		return err
	}
	return nil
}
