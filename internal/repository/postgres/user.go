package postgres

import (
	"context"

	"github.com/quest-be/internal/repository/model"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *model.User) (*model.User, error)
	FindUserById(ctx context.Context, id int64) (*model.User, error)
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
}

func (r *Database) InsertUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.Gorm.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Database) FindUserById(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	if err := r.Gorm.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Database) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	if err := r.Gorm.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *Database) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.Gorm.Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
