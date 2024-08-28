package repository

import "github.com/quest-be/internal/repository/postgres"

type Store interface {
	postgres.IUserRepository
}
