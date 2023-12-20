package usecase

import (
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/url-shortener/src/business/domain"
	"github.com/adiatma85/url-shortener/src/business/usecase/user"
)

type Usecase struct {
	User user.Interface
}

type InitParam struct {
	Log log.Interface
	Dom *domain.Domain
}

func Init(param InitParam) *Usecase {
	usecase := &Usecase{
		User: user.Init(user.InitParam{Log: param.Log, User: param.Dom.User}),
	}

	return usecase
}
