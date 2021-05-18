package models

import (
	"time"
)

//User 用户模型
type UrlMap struct {
	ID        int64 `gorm:"primaryKey"`
	LongUrl   string
	ShortUrl  string
	CreatedAt time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
	UpdatedAt time.Time `gorm:"autoCreateTime"` // 使用秒级时间戳填充创建时间
}

//TableName 指定数据库名称
func (UrlMap) TableName() string {
	return "url_maps"
}
