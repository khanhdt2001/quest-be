package postgres

import (
	"context"

	"github.com/quest-be/internal/repository/model"
)

type IQuestion interface {
	InsertQuestion(ctx context.Context, question *model.Question) (*model.Question, error)
	UpdateQuestion(ctx context.Context, question *model.Question) (*model.Question, error)
	FindQuestionById(ctx context.Context, id uint64) (*model.Question, error)
	FindQuestionByLevel(ctx context.Context, level *model.Level) ([]model.Question, error)
	DeleteQuestion(ctx context.Context, question *model.Question) error
}

func (r *Database) InsertQuestion(ctx context.Context, question *model.Question) (*model.Question, error) {
	if err := r.Gorm.Create(&question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (r *Database) UpdateQuestion(ctx context.Context, question *model.Question) (*model.Question, error) {
	if err := r.Gorm.Updates(question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (r *Database) FindQuestionById(ctx context.Context, id uint64) (*model.Question, error) {
	question := &model.Question{}
	if err := r.Gorm.Where("id = ?", id).First(question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (r *Database) FindQuestionByLevel(ctx context.Context, level *model.Level) ([]model.Question, error) {
	var questions []model.Question
	if err := r.Gorm.Where("level_id = ?", level.ID).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *Database) DeleteQuestion(ctx context.Context, question *model.Question) error {
	if err := r.Gorm.Delete(question).Error; err != nil {
		return err
	}
	return nil
}
