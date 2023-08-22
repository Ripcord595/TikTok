package model

// User  用户表\n
type User struct {
	ID       int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Password string `gorm:"column:password" db:"password" json:"password" form:"password"`
	Token    string `gorm:"column:token" db:"token" json:"token" form:"token"`
}
