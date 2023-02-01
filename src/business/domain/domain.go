package domain

import (
	"github.com/adiatma85/url-shortener/src/business/domain/user"
	"github.com/adiatma85/url-shortener/src/utils"
	"gorm.io/gorm"
)

type Domain struct {
	User user.Interface
}

func Init(db *gorm.DB, util *utils.Utils) *Domain {
	domain := &Domain{
		User: user.Init(db, util.QueryBuilder),
		// Url dan sebagainya nanti di sini
	}

	return domain
}
