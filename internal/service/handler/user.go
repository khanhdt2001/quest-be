package handler

import (
	"context"

	"github.com/quest-be/internal/repository/model"
	"github.com/quest-be/internal/repository/postgres"
	"github.com/quest-be/internal/service/dto"
	"github.com/quest-be/util"
)

type IUserHandler interface {
	CreateUser(ctx context.Context, userDto dto.CreateUserRequest) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserById(ctx context.Context, id int64) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
}

type UserHandler struct {
	db *postgres.Database
}

func NewUserHandler(db *postgres.Database) IUserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) CreateUser(ctx context.Context, userDto dto.CreateUserRequest) (*model.User, error) {

	var (
		password string
		err      error
	)

	if model.LoginType(userDto.LastLoginType) == model.PASSWORD {
		password, err = util.GenerateFromPassword(userDto.Password)
		if err != nil {
			return nil, err
		}
	}

	user, err := h.db.InsertUser(ctx, &model.User{
		Email:          userDto.Email,
		PassWordHashed: password,
		LastLoginType:  model.LoginType(userDto.LastLoginType),
	})
	
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *UserHandler) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := h.db.FindUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *UserHandler) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	user, err := h.db.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *UserHandler) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	user, err := h.db.UpdateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
