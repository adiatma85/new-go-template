package mysql

import (
	"fmt"

	"github.com/adiatma85/url-shortener/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(cfg config.AppConfig) (*gorm.DB, error) {
	dsn := generateDsn(cfg)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func generateDsn(cfg config.AppConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_DATABSE)
	return dsn
}
