package handler

import (
	"context"
	"fmt"
	"strings"

	"github.com/quest-be/constant"
	"github.com/quest-be/internal/repository/model"
	"github.com/quest-be/internal/repository/postgres"
	"github.com/quest-be/internal/service/dto"
	"github.com/quest-be/util"
)

type IAuthHandler interface {
	LoginByPassword(ctx context.Context, req *dto.LoginByPasswordRequest) (string, error)
	SignUp(ctx context.Context, req *dto.SignUpRequest) error
	VerifyUser(ctx context.Context, req *dto.VerifyUserRequest) error
	SetUserHandler(userHandler IUserHandler)
	SetOtpHandler(otpHandler IOtpHandler)
}

type AuthHandler struct {
	db         *postgres.Database
	userHander IUserHandler
	otpHandler IOtpHandler
}

func NewAuthHandler(db *postgres.Database) IAuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) SetUserHandler(userHandler IUserHandler) {
	h.userHander = userHandler
}

func (h *AuthHandler) SetOtpHandler(otpHandler IOtpHandler) {
	h.otpHandler = otpHandler
}

func (h *AuthHandler) SignUp(ctx context.Context, req *dto.SignUpRequest) error {
	// create user entity, hash password, insert to db
	// check if this email is already registered
	// if not, create user
	// if yes, return error
	_, err := h.userHander.GetUserByEmail(ctx, req.Email)
	if err == nil {
		return constant.ErrUserAlreadyExist
	}
	if err != nil && !strings.Contains(err.Error(), "not found") {
		return err
	}
	otp := util.RandomString(6)

	user, err := h.userHander.CreateUser(ctx, dto.CreateUserRequest{
		Email:         req.Email,
		Password:      req.Password,
		LastLoginType: string(model.PASSWORD),
	})

	if err != nil {
		return err
	}

	err = h.otpHandler.CreateOtp(ctx, user.Id, otp)
	if err != nil {
		return err
	}

	go util.SendMail(
		req.Email,
		constant.WelcomeSubject,
		fmt.Sprintf(constant.WelcomeBody, otp))

	return err
}

func (h *AuthHandler) VerifyUser(ctx context.Context, req *dto.VerifyUserRequest) error {
	// get user by email
	// check if otp is correct
	// if correct, update user status to active
	// if not, return error
	user, err := h.userHander.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return err
	}
	if user.IsVerified {
		return constant.ErrUserAlreadyVerified
	}

	err = h.otpHandler.VerifyOtp(ctx, user.Id, req.OTP)
	if err != nil {
		return err
	}
	user.IsVerified = true
	_, err = h.userHander.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (h *AuthHandler) LoginByPassword(ctx context.Context, req *dto.LoginByPasswordRequest) (string, error) {
	// get user by email
	// check if password is correct
	// if correct, return jwt token
	// if not, return error
	user, err := h.userHander.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}
	if !user.IsVerified {
		return "", constant.ErrUserNotVerified
	}

	err = util.CompareHashAndPassword(user.PassWordHashed, req.Password)
	if err != nil {
		return "", err
	}

	token, err := util.CreateToken(user.Id, constant.JWT_EXP_TIME)
	if err != nil {
		return "", err
	}
	return token, nil
}
