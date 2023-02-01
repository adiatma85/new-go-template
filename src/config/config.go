package config

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type AppConfig struct {
	APP_NAME string `mapstructure:"APP_NAME"`
	APP_PORT string `mapstructure:"APP_PORT"`
	APP_KEY  string `mapstructure:"DB_DATABASE"`

	// Database Configuration
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_DATABSE  string `mapstructure:"DB_DATABASE"`
	DB_USERNAME string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`

	// Redis Configuration
}

func Setup() (*gorm.DB, error) {
	config := AppConfig{}
	viper.SetConfigFile("./env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return setupMySQL(config)
}

// Init for MySQL
func setupMySQL(cfg AppConfig) (*gorm.DB, error) {
	dsn := generateDsn(cfg)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

// Generating Dsn for MySql Database
func generateDsn(cfg AppConfig) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_DATABSE)
	return dsn
}
