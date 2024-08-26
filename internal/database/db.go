package database

import (
	"fmt"

	"github.com/quest-be/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Gorm *gorm.DB
}

func New(log bool) (*Database, error) {
	var logMode logger.LogLevel
	if log {
		logMode = logger.Info
	}

	cfg := &gorm.Config{
		Logger:                 logger.Default.LogMode(logMode),
		SkipDefaultTransaction: true,
	}
	dbUrl := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		util.Default.PostgresUser,
		util.Default.PostgresPassword,
		util.Default.PostgresHost,
		util.Default.PostgresPort,
		util.Default.PostgresDB,
	)
	db, err := gorm.Open(postgres.Open(dbUrl), cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to open postgres conn: %w", err)
	}

	return &Database{Gorm: db}, nil
}

func Setup(db *Database) error {
	return nil
}
