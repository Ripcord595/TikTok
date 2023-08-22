package model

import (
	"time"
)

// Comment 模型对应数据库中的 comment 表
type Comment struct {
	ID            int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId        int64     `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	VideoId       int64     `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	CommentText   string    `gorm:"column:comment_text" db:"comment_text" json:"comment_text" form:"comment_text"`
	CancelComment int64     `gorm:"column:cancel_comment" db:"cancel_comment" json:"cancel_comment" form:"cancel_comment"` //  0表示已评论，1表示取消评论
	CreateTime    time.Time `gorm:"column:create_time" db:"create_time" json:"create_time" form:"create_time"`
}
