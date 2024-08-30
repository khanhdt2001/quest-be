package postgres

import (
	"context"

	"github.com/quest-be/internal/repository/model"
)

type IOtp interface {
	InsertOtp(ctx context.Context, otp *model.Otp) (*model.Otp, error)
	UpdateOtp(ctx context.Context, otp *model.Otp) (*model.Otp, error)
	FindOtpByUserId(ctx context.Context, userId uint64) (*model.Otp, error)
	DeleteOtp(ctx context.Context, otp *model.Otp) error
}

func (r *Database) InsertOtp(ctx context.Context, otp *model.Otp) (*model.Otp, error) {
	if err := r.Gorm.Create(&otp).Error; err != nil {
		return nil, err
	}
	return otp, nil
}

func (r *Database) UpdateOtp(ctx context.Context, otp *model.Otp) (*model.Otp, error) {
	if err := r.Gorm.Updates(otp).Error; err != nil {
		return nil, err
	}
	return otp, nil
}

func (r *Database) FindOtpByUserId(ctx context.Context, userId uint64) (*model.Otp, error) {
	otp := &model.Otp{}
	if err := r.Gorm.Where("user_id = ?", userId).First(otp).Error; err != nil {
		return nil, err
	}
	return otp, nil
}

func (r *Database) DeleteOtp(ctx context.Context, otp *model.Otp) error {
	if err := r.Gorm.Delete(otp).Error; err != nil {
		return err
	}
	return nil
}
