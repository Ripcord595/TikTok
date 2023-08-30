package services

import (
	"errors"
	"tiktok/internal/repository/cache"
	"tiktok/internal/repository/models"
)

type VideoList struct {
	Videos []*models.Video `json:"video_list,omitempty"`
}

func QueryVideoListByUserId(userId int64) (*VideoList, error) {
	return NewQueryVideoListByUserIdFlow(userId).Do()
}

func NewQueryVideoListByUserIdFlow(userId int64) *QueryVideoListByUserIdFlow {
	return &QueryVideoListByUserIdFlow{userId: userId}
}

type QueryVideoListByUserIdFlow struct {
	userId int64
	videos []*models.Video

	videoList *VideoList
}

func (q *QueryVideoListByUserIdFlow) Do() (*VideoList, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.videoList, nil
}

func (q *QueryVideoListByUserIdFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return errors.New("用户不存在")
	}

	return nil
}

func (q *QueryVideoListByUserIdFlow) packData() error {
	err := models.NewVideoDAO().QueryVideoListByUserId(q.userId, &q.videos)
	if err != nil {
		return err
	}
	var userInfo models.UserInfo
	err = models.NewUserInfoDAO().QueryUserInfoById(q.userId, &userInfo)
	p := cache.NewProxyIndexMap()
	if err != nil {
		return err
	}
	for i := range q.videos {
		q.videos[i].Author = userInfo
		q.videos[i].IsFavorite = p.GetVideoFavorState(q.userId, q.videos[i].Id)
	}

	q.videoList = &VideoList{Videos: q.videos}

	return nil
}
