package database

import (
	"fmt"
	"time"

	"github.com/brajes/go-rest-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDatabase(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s posrt=%s sslmode=%s TimeZone=UTC", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
