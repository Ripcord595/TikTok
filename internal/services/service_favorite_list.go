package services

import (
	"errors"
	"tiktok/internal/repository/models"
)

type FavorList struct {
	Videos []*models.Video `json:"video_list"`
}

func QueryFavorVideoList(userId int64) (*FavorList, error) {
	return NewQueryFavorVideoListFlow(userId).Do()
}

type QueryFavorVideoListFlow struct {
	userId int64

	videos []*models.Video

	videoList *FavorList
}

func NewQueryFavorVideoListFlow(userId int64) *QueryFavorVideoListFlow {
	return &QueryFavorVideoListFlow{userId: userId}
}

func (q *QueryFavorVideoListFlow) Do() (*FavorList, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.videoList, nil
}

func (q *QueryFavorVideoListFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return errors.New("用户状态异常")
	}
	return nil
}

func (q *QueryFavorVideoListFlow) prepareData() error {
	err := models.NewVideoDAO().QueryFavorVideoListByUserId(q.userId, &q.videos)
	if err != nil {
		return err
	}
	for i := range q.videos {
		var userInfo models.UserInfo
		err = models.NewUserInfoDAO().QueryUserInfoById(q.videos[i].UserInfoId, &userInfo)
		if err == nil {
			q.videos[i].Author = userInfo
		}
		q.videos[i].IsFavorite = true
	}
	return nil
}

func (q *QueryFavorVideoListFlow) packData() error {
	q.videoList = &FavorList{Videos: q.videos}
	return nil
}
