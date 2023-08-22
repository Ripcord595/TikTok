package model

// Follow 模型对应数据库中的 follow 表
type Follow struct {
	ID           int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId       int64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	FollowerId   int64 `gorm:"column:follower_id" db:"follower_id" json:"follower_id" form:"follower_id"`
	CancelFollow int64 `gorm:"column:cancel_follow" db:"cancel_follow" json:"cancel_follow" form:"cancel_follow"` //  0表示关注，1则表示未关注
}
