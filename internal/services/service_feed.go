package services

import (
	"tiktok/internal/repository/models"
	"tiktok/pkg/util"
	"time"
)

const (
	MaxVideoNum = 30
)

type FeedVideoList struct {
	Videos   []*models.Video `json:"video_list,omitempty"`
	NextTime int64           `json:"next_time,omitempty"`
}

func QueryFeedVideoList(userId int64, latestTime time.Time) (*FeedVideoList, error) {
	return NewQueryFeedVideoListFlow(userId, latestTime).Do()
}

type QueryFeedVideoListFlow struct {
	userId     int64
	latestTime time.Time

	videos   []*models.Video
	nextTime int64

	feedVideo *FeedVideoList
}

func NewQueryFeedVideoListFlow(userId int64, latestTime time.Time) *QueryFeedVideoListFlow {
	return &QueryFeedVideoListFlow{userId: userId, latestTime: latestTime}
}

func (q *QueryFeedVideoListFlow) Do() (*FeedVideoList, error) {
	q.checkNum()

	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.feedVideo, nil
}

func (q *QueryFeedVideoListFlow) checkNum() {
	if q.userId > 0 {
	}

	if q.latestTime.IsZero() {
		q.latestTime = time.Now()
	}
}

func (q *QueryFeedVideoListFlow) prepareData() error {
	err := models.NewVideoDAO().QueryVideoListByLimitAndTime(MaxVideoNum, q.latestTime, &q.videos)
	if err != nil {
		return err
	}
	latestTime, _ := util.FillVideoListFields(q.userId, &q.videos) //不是致命错误，不返回

	if latestTime != nil {
		q.nextTime = (*latestTime).UnixNano() / 1e6
		return nil
	}
	q.nextTime = time.Now().Unix() / 1e6
	return nil
}

func (q *QueryFeedVideoListFlow) packData() error {
	q.feedVideo = &FeedVideoList{
		Videos:   q.videos,
		NextTime: q.nextTime,
	}
	return nil
}
