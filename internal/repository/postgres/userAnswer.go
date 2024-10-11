package postgres

import (
	"context"

	"github.com/quest-be/internal/repository/model"
)

type IUserAnswer interface {
	InsertUserAnswer(ctx context.Context, userAnswer *model.UserAnswer) (*model.UserAnswer, error)
	UpdateUserAnswer(ctx context.Context, userAnswer *model.UserAnswer) (*model.UserAnswer, error)
	FindUserAnswerById(ctx context.Context, id uint64) (*model.UserAnswer, error)
	FindUserAnswerByUserId(ctx context.Context, userId uint64) ([]model.UserAnswer, error)
	DeleteUserAnswer(ctx context.Context, userAnswer *model.UserAnswer) error
}

func (r *Database) InsertUserAnswer(ctx context.Context, userAnswer *model.UserAnswer) (*model.UserAnswer, error) {
	if err := r.Gorm.Create(&userAnswer).Error; err != nil {
		return nil, err
	}
	return userAnswer, nil
}

func (r *Database) UpdateUserAnswer(ctx context.Context, userAnswer *model.UserAnswer) (*model.UserAnswer, error) {
	if err := r.Gorm.Updates(userAnswer).Error; err != nil {
		return nil, err
	}
	return userAnswer, nil
}

func (r *Database) FindUserAnswerById(ctx context.Context, id uint64) (*model.UserAnswer, error) {
	userAnswer := &model.UserAnswer{}
	if err := r.Gorm.Where("id = ?", id).First(userAnswer).Error; err != nil {
		return nil, err
	}
	return userAnswer, nil
}

func (r *Database) FindUserAnswerByUserId(ctx context.Context, userId uint64) ([]model.UserAnswer, error) {
	var userAnswers []model.UserAnswer
	if err := r.Gorm.Where("user_id = ?", userId).Find(&userAnswers).Error; err != nil {
		return nil, err
	}
	return userAnswers, nil
}

func (r *Database) DeleteUserAnswer(ctx context.Context, userAnswer *model.UserAnswer) error {
	if err := r.Gorm.Delete(userAnswer).Error; err != nil {
		return err
	}
	return nil
}
