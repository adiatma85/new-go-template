package entity

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	OriginalUrl string `json:"originalUrl"`
	ShortenUrl  string `json:"shortenUrl"`
	Visit       int    `json:"visit"`
	UserID      uint   `json:"userID"`
	User        User   `json:"user"`
}
