package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"-"`
}

// Password Hashing goes here
func (u *User) BeforeCreate() {

}

type UserParam struct {
	ID       uint
	Username string
	PaginationParam
}
