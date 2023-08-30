package services

import (
	"errors"
	"tiktok/internal/repository/models"
)

const (
	FOLLOW = 1
	CANCEL = 2
)

var (
	ErrIvdAct    = errors.New("未定义操作")
	ErrIvdFolUsr = errors.New("关注用户不存在")
)

type PostFollowActionFlow struct {
	userId     int64
	userToId   int64
	actionType int
}

func NewPostFollowActionFlow(userId int64, userToId int64, actionType int) *PostFollowActionFlow {
	return &PostFollowActionFlow{userId: userId, userToId: userToId, actionType: actionType}
}

func (p *PostFollowActionFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(p.userToId) {
		return ErrIvdFolUsr
	}
	if p.actionType != FOLLOW && p.actionType != CANCEL {
		return ErrIvdAct
	}
	if p.userId == p.userToId {
		return ErrIvdAct
	}
	return nil
}
