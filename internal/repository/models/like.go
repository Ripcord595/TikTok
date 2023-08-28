package models

type Like struct {
	ID         int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	UserId     int64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	VideoId    int64 `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	CancleLike int64 `gorm:"column:cancle_like" db:"cancle_like" json:"cancle_like" form:"cancle_like"`
}
