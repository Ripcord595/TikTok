package model

import "time"

// User 表示用户数据模型
type User struct {
	ID        uint      // 用户ID
	Username  string    // 用户名
	Email     string    // 电子邮件
	Password  string    // 密码（可能需要加密存储）
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
}
