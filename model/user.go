package model

// User 表示用户数据模型
// User  用户表
type User struct {
	ID              int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Username        string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Password        string `gorm:"column:password" db:"password" json:"password" form:"password"`
	Token           string `gorm:"column:token" db:"token" json:"token" form:"token"`
	Avatar          string `gorm:"column:avatar" db:"avatar" json:"avatar" form:"avatar"`
	BackgroundImage string `gorm:"column:background_image" db:"background_image" json:"background_image" form:"background_image"`
	FavoriteCount   int64  `gorm:"column:favorite_count" db:"favorite_count" json:"favorite_count" form:"favorite_count"`
	FollowCount     int64  `gorm:"column:follow_count" db:"follow_count" json:"follow_count" form:"follow_count"`
	FollowerCount   int64  `gorm:"column:follower_count" db:"follower_count" json:"follower_count" form:"follower_count"`
	IsFollow        int64  `gorm:"column:is_follow" db:"is_follow" json:"is_follow" form:"is_follow"`
	Name            string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Signature       string `gorm:"column:signature" db:"signature" json:"signature" form:"signature"`
	TotalFavorited  int64  `gorm:"column:total_favorited" db:"total_favorited" json:"total_favorited" form:"total_favorited"`
	WorkCount       int64  `gorm:"column:work_count" db:"work_count" json:"work_count" form:"work_count"`
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

type UserInfoGet struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	User       *User   `json:"user"`        // 用户信息
}
