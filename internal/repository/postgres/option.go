package postgres

import (
	"context"

	"github.com/quest-be/internal/repository/model"
)

type IOption interface {
	InsertOption(ctx context.Context, option *model.Option) (*model.Option, error)
	UpdateOption(ctx context.Context, option *model.Option) (*model.Option, error)
	FindOptionById(ctx context.Context, id uint64) (*model.Option, error)
	FindOptionByQuestion(ctx context.Context, question *model.Question) ([]model.Option, error)
	DeleteOption(ctx context.Context, option *model.Option) error
}

func (r *Database) InsertOption(ctx context.Context, option *model.Option) (*model.Option, error) {
	if err := r.Gorm.Create(&option).Error; err != nil {
		return nil, err
	}
	return option, nil
}

func (r *Database) UpdateOption(ctx context.Context, option *model.Option) (*model.Option, error) {
	if err := r.Gorm.Updates(option).Error; err != nil {
		return nil, err
	}
	return option, nil
}

func (r *Database) FindOptionById(ctx context.Context, id uint64) (*model.Option, error) {
	option := &model.Option{}
	if err := r.Gorm.Where("id = ?", id).First(option).Error; err != nil {
		return nil, err
	}
	return option, nil
}

func (r *Database) FindOptionByQuestion(ctx context.Context, question *model.Question) ([]model.Option, error) {
	var options []model.Option
	if err := r.Gorm.Where("question_id = ?", question.ID).Find(&options).Error; err != nil {
		return nil, err
	}
	return options, nil
}

func (r *Database) DeleteOption(ctx context.Context, option *model.Option) error {
	if err := r.Gorm.Delete(option).Error; err != nil {
		return err
	}
	return nil
}
