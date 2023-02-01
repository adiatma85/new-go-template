package entity

type PaginationParam struct {
	Limit int `form:"limit"`
	Page  int `form:"page"`
}
