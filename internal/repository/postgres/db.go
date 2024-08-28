package postgres

import (
	"fmt"

	"github.com/quest-be/internal/repository/model"
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
	conn := db.Gorm.Begin()

	if err := conn.AutoMigrate(
		&model.User{},
	); err != nil {
		conn.Rollback()
		return fmt.Errorf("failed to migrate user table: %w", err)
	}
	if err := conn.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit migration: %w", err)
	}
	return nil
}
