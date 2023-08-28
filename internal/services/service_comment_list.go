package services

import (
	"errors"
	"fmt"
	"tiktok/internal/repository/models"
	"tiktok/pkg/util"
)

type CommentsList struct {
	Comments []*models.Comment `json:"comment_list"`
}

func QueryCommentList(userId, videoId int64) (*CommentsList, error) {
	return NewQueryCommentListFlow(userId, videoId).Do()
}

type QueryCommentListFlow struct {
	userId  int64
	videoId int64

	comments []*models.Comment

	commentList *CommentsList
}

func NewQueryCommentListFlow(userId, videoId int64) *QueryCommentListFlow {
	return &QueryCommentListFlow{userId: userId, videoId: videoId}
}

func (q *QueryCommentListFlow) Do() (*CommentsList, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.commentList, nil
}

func (q *QueryCommentListFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return fmt.Errorf("用户%d处于登出状态", q.userId)
	}
	if !models.NewVideoDAO().IsVideoExistById(q.videoId) {
		return fmt.Errorf("视频%d不存在或已经被删除", q.videoId)
	}
	return nil
}

func (q *QueryCommentListFlow) prepareData() error {
	err := models.NewCommentDAO().QueryCommentListByVideoId(q.videoId, &q.comments)
	if err != nil {
		return err
	}
	err = util.FillCommentListFields(&q.comments)
	if err != nil {
		return errors.New("暂时还没有人评论")
	}
	return nil
}

func (q *QueryCommentListFlow) packData() error {
	q.commentList = &CommentsList{Comments: q.comments}
	return nil
}
