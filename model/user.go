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

// UserRequest 表示用户登录、注册时的请求数据模型
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserResponse 表示用户登录、注册时的响应数据模型
type UserResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserID     int64  `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}
