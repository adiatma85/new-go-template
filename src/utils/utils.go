package utils

import (
	"github.com/adiatma85/url-shortener/src/utils/querybuilder"
	"gorm.io/gorm"
)

// Utilitas untuk nanganin pagination
type Utils struct {
	QueryBuilder querybuilder.Interface
}

func Init(db *gorm.DB) *Utils {
	util := &Utils{
		QueryBuilder: querybuilder.Init(db),
	}

	return util
}
