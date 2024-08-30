package handler

import (
	"context"
	"time"

	"github.com/quest-be/constant"
	"github.com/quest-be/internal/repository/model"
	"github.com/quest-be/internal/repository/postgres"
	"github.com/quest-be/util"
)

type IOtpHandler interface {
	CreateOtp(ctx context.Context, userId uint64, otp string) error
	VerifyOtp(ctx context.Context, userId uint64, otp string) error
	ResendOtp(ctx context.Context, userId uint64) error
}

type OtpHandler struct {
	db *postgres.Database
}

func NewOtpHandler(db *postgres.Database) IOtpHandler {
	return &OtpHandler{db: db}
}

// CreateOtp implements IOtpHandler.
func (o *OtpHandler) CreateOtp(ctx context.Context, userId uint64, otp string) error {
	_, err := o.db.InsertOtp(ctx, &model.Otp{
		UserId:    userId,
		OTP:       otp,
		ExpiredAt: time.Now().Add(constant.OTP_EXP_TIME),
	})
	return err
}

// VerifyOtp implements IOtpHandler.
func (o *OtpHandler) VerifyOtp(ctx context.Context, userId uint64, code string) error {
	otp, err := o.db.FindOtpByUserId(ctx, userId)
	if err != nil {
		return err
	}

	if otp.OTP != code {
		return constant.ErrInvalidOTP
	}
	if time.Now().After(otp.ExpiredAt) {
		return constant.ErrExpiredOTP
	}
	err = o.db.DeleteOtp(ctx, otp)
	if err != nil {
		return err
	}
	return nil
}

func (o *OtpHandler) ResendOtp(ctx context.Context, userId uint64) error {
	otp, err := o.db.FindOtpByUserId(ctx, userId)
	if err != nil {
		return err
	}
	if time.Now().Before(otp.ExpiredAt) {
		return nil
	}
	otp.OTP = util.RandomString(6)
	otp.ExpiredAt = time.Now().Add(constant.OTP_EXP_TIME)
	_, err = o.db.UpdateOtp(ctx, otp)
	if err != nil {
		return err
	}
	return nil
}
