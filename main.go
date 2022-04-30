package main

import (
	"log"

	"github.com/adiatma85/golang-alter-url-shortener/config"
	"github.com/adiatma85/golang-alter-url-shortener/handler"
	"github.com/adiatma85/golang-alter-url-shortener/storage/redis"
	"github.com/valyala/fasthttp"
)

func main() {
	configuration, err := config.FromFile("./configuration.json")
	if err != nil {
		log.Fatal(err)
	}

	service, err := redis.New(configuration.Redis.Host, configuration.Redis.Port)
	if err != nil {
		log.Fatal(err)
	}

	defer service.Close()

	router := handler.New(configuration.Options.Schema, configuration.Options.Prefix, service)

	log.Fatal(fasthttp.ListenAndServe(":"+configuration.Server.Port, router.Handler))
}
