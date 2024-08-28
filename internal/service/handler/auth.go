package handler

import (
	"github.com/quest-be/internal/repository/postgres"
)

type AuthHandler struct {
	db *postgres.Database
}
