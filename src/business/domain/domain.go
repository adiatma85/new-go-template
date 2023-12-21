package domain

import (
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/parser"
	"github.com/adiatma85/own-go-sdk/redis"
	"github.com/adiatma85/own-go-sdk/sql"
	"github.com/adiatma85/url-shortener/src/business/domain/user"
)

type Domain struct {
	User user.Interface
}

type InitParam struct {
	Log   log.Interface
	Db    sql.Interface
	Json  parser.JSONInterface
	Redis redis.Interface
}

func Init(param InitParam) *Domain {
	domain := &Domain{
		User: user.Init(user.InitParam{Log: param.Log, Db: param.Db, Json: param.Json, Redis: param.Redis}),
	}

	return domain
}
