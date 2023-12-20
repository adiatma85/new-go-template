package user

import (
	"context"
	"fmt"
	"time"

	"github.com/adiatma85/own-go-sdk/codes"
	"github.com/adiatma85/own-go-sdk/errors"
	"github.com/adiatma85/own-go-sdk/log"
	"github.com/adiatma85/own-go-sdk/null"
	userDom "github.com/adiatma85/url-shortener/src/business/domain/user"
	"github.com/adiatma85/url-shortener/src/business/entity"
)

type Interface interface {
	Create(ctx context.Context, req entity.CreateUserParam) (entity.User, error)
	Get(ctx context.Context, params entity.UserParam) (entity.User, error)
	GetListAsAdmin(ctx context.Context, params entity.UserParam) ([]entity.User, *entity.Pagination, error)
	Update(ctx context.Context, updateParam entity.UpdateUserParam, selectParam entity.UserParam) error
	Delete(ctx context.Context, selectParam entity.UserParam) error

	// Improvement kedepannya
	// CheckPassword(ctx context.Context, params entity.UserCheckPasswordParam, userParam entity.UserParam) (entity.HTTPMessage, error)
	// ChangePassword(ctx context.Context, passwordChangeParam entity.UserChangePasswordParam, userParam entity.UserParam) (entity.HTTPMessage, error)
	// Activate(ctx context.Context, selectParam entity.UserParam) error
	// SignInWithPassword(ctx context.Context, req entity.UserLoginRequest) (entity.UserLoginResponse, error)
	// RefreshToken(ctx context.Context, param entity.UserRefreshTokenParam) (entity.RefreshTokenResponse, error)
}

type InitParam struct {
	Log  log.Interface
	User userDom.Interface
}

type user struct {
	log  log.Interface
	user userDom.Interface
}

var Now = time.Now

func Init(param InitParam) Interface {
	u := &user{
		log:  param.Log,
		user: param.User,
	}

	return u
}

func (u *user) Create(ctx context.Context, req entity.CreateUserParam) (entity.User, error) {
	var result entity.User

	result, err := u.validateUser(ctx, req)
	if err != nil {
		return result, err
	}

	req.CreatedBy = null.StringFrom(fmt.Sprintf("%v", entity.SystemUser))

	return u.user.Create(ctx, req)
}

func (u *user) validateUser(ctx context.Context, req entity.CreateUserParam) (entity.User, error) {
	var result entity.User

	if req.Password != req.ConfirmPassword {
		return result, errors.NewWithCode(codes.CodePasswordDoesNotMatch, "password does not match")
	}

	user, err := u.user.Get(ctx, entity.UserParam{
		Email: null.StringFrom(req.Email),
	})
	if err != nil && errors.GetCode(err) != codes.CodeSQLRecordDoesNotExist {
		return result, err
	}

	if user != result {
		return result, errors.NewWithCode(codes.CodeConflict, "email is exists")
	}

	return result, nil
}

func (u *user) Get(ctx context.Context, params entity.UserParam) (entity.User, error) {
	return u.user.Get(ctx, params)
}

func (u *user) GetListAsAdmin(ctx context.Context, params entity.UserParam) ([]entity.User, *entity.Pagination, error) {
	params.IncludePagination = true
	users, pg, err := u.user.GetList(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	return users, pg, nil
}

func (u *user) Update(ctx context.Context, updateParam entity.UpdateUserParam, selectParam entity.UserParam) error {
	return u.user.Update(ctx, updateParam, selectParam)
}

func (u *user) Delete(ctx context.Context, selectParam entity.UserParam) error {
	deleteParam := entity.UpdateUserParam{
		Status:    null.Int64From(-1),
		DeletedAt: null.TimeFrom(Now()),
		DeletedBy: null.StringFrom(fmt.Sprintf("%v", entity.SystemUser)),
	}

	return u.user.Update(ctx, deleteParam, selectParam)
}
