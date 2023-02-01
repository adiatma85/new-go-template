package main

import (
	"fmt"

	"github.com/adiatma85/url-shortener/src/business/domain"
	"github.com/adiatma85/url-shortener/src/config"
	"github.com/adiatma85/url-shortener/src/utils"
)

type AppConfig struct {
	APP_NAME string `mapstructure:"APP_NAME"`
	APP_PORT string `mapstructure:"APP_PORT"`
	APP_KEY  string `mapstructure:"DB_DATABASE"`

	// Mysql Config

}

func main() {
	// fmt.Println("Teting jancok")
	// Do the configuration
	gormDB, err := config.Setup()
	if err != nil {
		// Log error
		fmt.Println("log error")
		return
	}

	// Utils init
	util := utils.Init(gormDB)

	// Domain init
	_ = domain.Init(gormDB, util)

	// Usecase init

	// Gin Init

}
