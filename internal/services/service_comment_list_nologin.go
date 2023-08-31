package services

import (
	"errors"
	"fmt"
	"tiktok/internal/repository/models"
	"tiktok/pkg/util"
)

type QueryCommentListFlowNologin struct {
	videoId int64

	comments []*models.Comment

	commentList *CommentsList
}

func QueryCommentListNologin(videoId int64) (*CommentsList, error) {
	return NewQueryCommentListFlowNologin(videoId).DoNologin()
}

func NewQueryCommentListFlowNologin(videoId int64) *QueryCommentListFlowNologin {
	return &QueryCommentListFlowNologin{videoId: videoId}
}

func (q *QueryCommentListFlowNologin) DoNologin() (*CommentsList, error) {
	if err := q.checkNumNologin(); err != nil {
		return nil, err
	}
	if err := q.prepareDataNologin(); err != nil {
		return nil, err
	}
	if err := q.packDataNologin(); err != nil {
		return nil, err
	}
	return q.commentList, nil
}

func (q *QueryCommentListFlowNologin) checkNumNologin() error {
	if !models.NewVideoDAO().IsVideoExistById(q.videoId) {
		return fmt.Errorf("视频%d不存在或已经被删除", q.videoId)
	}
	return nil
}

func (q *QueryCommentListFlowNologin) prepareDataNologin() error {
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

func (q *QueryCommentListFlowNologin) packDataNologin() error {
	q.commentList = &CommentsList{Comments: q.comments}
	return nil
}
