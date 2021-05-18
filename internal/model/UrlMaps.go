package models

import (
	"time"
)

type UrlMap struct {
	ID        int64 `gorm:"primaryKey"`
	LongUrl   string
	ShortUrl  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoCreateTime"`
}

func (UrlMap) TableName() string {
	return "url_maps"
}
