package services

import (
	"errors"
	"tiktok/internal/repository/cache"
	"tiktok/internal/repository/models"
)

const (
	PLUS  = 1
	MINUS = 2
)

func PostFavorState(userId, videoId, actionType int64) error {
	return NewPostFavorStateFlow(userId, videoId, actionType).Do()
}

type PostFavorStateFlow struct {
	userId     int64
	videoId    int64
	actionType int64
}

func NewPostFavorStateFlow(userId, videoId, action int64) *PostFavorStateFlow {
	return &PostFavorStateFlow{
		userId:     userId,
		videoId:    videoId,
		actionType: action,
	}
}

func (p *PostFavorStateFlow) Do() error {
	var err error
	if err = p.checkNum(); err != nil {
		return err
	}

	switch p.actionType {
	case PLUS:
		err = p.PlusOperation()
	case MINUS:
		err = p.MinusOperation()
	default:
		return errors.New("未定义的操作")
	}
	return err
}

func (p *PostFavorStateFlow) PlusOperation() error {
	err := models.NewVideoDAO().PlusOneFavorByUserIdAndVideoId(p.userId, p.videoId)
	if err != nil {
		return errors.New("不要重复点赞")
	}
	cache.NewProxyIndexMap().UpdateVideoFavorState(p.userId, p.videoId, true)
	return nil
}

func (p *PostFavorStateFlow) MinusOperation() error {
	err := models.NewVideoDAO().MinusOneFavorByUserIdAndVideoId(p.userId, p.videoId)
	if err != nil {
		return errors.New("点赞数目已经为0")
	}
	cache.NewProxyIndexMap().UpdateVideoFavorState(p.userId, p.videoId, false)
	return nil
}

func (p *PostFavorStateFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(p.userId) {
		return errors.New("用户不存在")
	}
	if p.actionType != PLUS && p.actionType != MINUS {
		return errors.New("未定义的行为")
	}
	return nil
}
