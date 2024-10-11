package postgres

import (
	"context"

	"github.com/quest-be/internal/repository/model"
)

type ILevel interface {
	InsertLevel(ctx context.Context, level *model.Level) (*model.Level, error)
	UpdateLevel(ctx context.Context, level *model.Level) (*model.Level, error)
	FindLevelById(ctx context.Context, id uint64) (*model.Level, error)
	FindLevelByCategory(ctx context.Context, category *model.Category) ([]model.Level, error)
	DeleteLevel(ctx context.Context, level *model.Level) error
}

func (r *Database) InsertLevel(ctx context.Context, level *model.Level) (*model.Level, error) {
	if err := r.Gorm.Create(&level).Error; err != nil {
		return nil, err
	}
	return level, nil
}

func (r *Database) UpdateLevel(ctx context.Context, level *model.Level) (*model.Level, error) {
	if err := r.Gorm.Updates(level).Error; err != nil {
		return nil, err
	}
	return level, nil
}

func (r *Database) FindLevelById(ctx context.Context, id uint64) (*model.Level, error) {
	level := &model.Level{}
	if err := r.Gorm.Where("id = ?", id).First(level).Error; err != nil {
		return nil, err
	}
	return level, nil
}

func (r *Database) FindLevelByCategory(ctx context.Context, category *model.Category) ([]model.Level, error) {
	var levels []model.Level
	if err := r.Gorm.Where("category_id = ?", category.ID).Find(&levels).Error; err != nil {
		return nil, err
	}
	return levels, nil
}

func (r *Database) DeleteLevel(ctx context.Context, level *model.Level) error {
	if err := r.Gorm.Delete(level).Error; err != nil {
		return err
	}
	return nil
}
