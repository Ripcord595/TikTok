package services

import (
	"errors"
	"tiktok/internal/repository/models"
)

var (
	ErrUserNotExist = errors.New("用户不存在或已注销")
)

type FollowList struct {
	UserList []*models.UserInfo `json:"user_list"`
}

type QueryFollowListFlow struct {
	userId int64

	userList []*models.UserInfo

	*FollowList
}

func NewQueryFollowListFlow(userId int64) *QueryFollowListFlow {
	return &QueryFollowListFlow{userId: userId}
}

func (q *QueryFollowListFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return ErrUserNotExist
	}
	return nil
}

func (q *QueryFollowListFlow) packData() error {
	q.FollowList = &FollowList{UserList: q.userList}

	return nil
}
