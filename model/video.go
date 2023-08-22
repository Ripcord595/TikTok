package model

import (
	"time"
)

// Video 表示视频数据模型

type Video struct {
	ID          int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	AuthorId    int64     `gorm:"column:author_id" db:"author_id" json:"author_id" form:"author_id"`
	PlayUrl     string    `gorm:"column:play_url" db:"play_url" json:"play_url" form:"play_url"`
	CoverUrl    string    `gorm:"column:cover_url" db:"cover_url" json:"cover_url" form:"cover_url"`
	PublishTime time.Time `gorm:"column:publish_time" db:"publish_time" json:"publish_time" form:"publish_time"`
	Title       string    `gorm:"column:title" db:"title" json:"title" form:"title"`
}
