package querybuilder

import (
	"context"

	"github.com/adiatma85/url-shortener/src/business/entity"
	"gorm.io/gorm"
)

type Interface interface {
	ProcessPagination(ctx context.Context, params entity.PaginationParam) *gorm.DB
}

type queryBuilder struct {
	db *gorm.DB
}

func Init(db *gorm.DB) Interface {
	qb := &queryBuilder{
		db: db,
	}

	return qb
}

func (qb *queryBuilder) ProcessPagination(ctx context.Context, params entity.PaginationParam) *gorm.DB {
	_, limit, offset := qb.pagination(params)

	queryBuilder := qb.db.Limit(limit).Offset(offset).Order("ASC")

	return queryBuilder
}

// Process Pagination and limit
// Mengembalikan page, limit, dan offset
func (qb *queryBuilder) pagination(params entity.PaginationParam) (int, int, int) {
	if params.Page < 1 {
		params.Page = 1
	}

	if params.Limit < 1 {
		params.Limit = 10
	}

	offset := (params.Page - 1) * params.Limit

	return params.Page, params.Limit, offset
}
