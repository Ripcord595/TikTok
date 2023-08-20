package model

import "time"

// Video 表示视频数据模型
type Video struct {
	ID          uint      // 视频ID
	UserID      uint      // 上传视频的用户ID
	Title       string    // 视频标题
	Description string    // 视频描述
	URL         string    // 视频存储的URL
	Likes       uint      // 点赞数
	CreatedAt   time.Time // 创建时间
	UpdatedAt   time.Time // 更新时间
}
