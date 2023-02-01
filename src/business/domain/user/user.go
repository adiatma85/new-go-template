package user

import (
	"context"

	"github.com/adiatma85/url-shortener/src/business/entity"
	"github.com/adiatma85/url-shortener/src/utils/querybuilder"
	"gorm.io/gorm"
)

// Function goes here
type Interface interface {
	Create(ctx context.Context, newUser entity.User) (entity.User, error)
	GetList(ctx context.Context, userParam entity.UserParam) ([]entity.User, error)
	Get(ctx context.Context, userParam entity.UserParam) (entity.User, error)
	Update(ctx context.Context, userParam entity.UserParam, updateParam entity.User) (entity.User, error)
	Delete(ctx context.Context, userParam entity.UserParam) error
}

type user struct {
	db *gorm.DB
	qb querybuilder.Interface
	// Log
	// Redis
	// JSON Parser
}

func Init(db *gorm.DB, qb querybuilder.Interface) Interface {
	u := &user{
		db: db,
		qb: qb,
	}

	return u
}

// Must return new user instance and error
func (u *user) Create(ctx context.Context, newUser entity.User) (entity.User, error) {
	result := u.db.Create(&newUser)
	return newUser, result.Error
}

// Must return a user and an error if occured
func (u *user) GetList(ctx context.Context, userParam entity.UserParam) ([]entity.User, error) {
	users := []entity.User{}
	queryBuilder := u.qb.ProcessPagination(ctx, userParam.PaginationParam)

	result := queryBuilder.Model(&entity.User{}).Where(userParam).Find(&users)

	if result.Error != nil {
		return []entity.User{}, result.Error
	}

	return users, nil
}

func (u *user) Get(ctx context.Context, userParam entity.UserParam) (entity.User, error) {
	user := entity.User{}
	result := u.db.Model(&user).Where(userParam).First(&user)

	if result.Error != nil {
		return entity.User{}, result.Error
	}

	return user, nil
}

func (u *user) Update(ctx context.Context, userParam entity.UserParam, updateParam entity.User) (entity.User, error) {
	u.db.Model(&updateParam).Where(userParam).Updates(updateParam)
	return u.Get(ctx, userParam)
}

func (u *user) Delete(ctx context.Context, userParam entity.UserParam) error {
	return u.db.Model(&entity.User{}).Delete(userParam).Error
}
