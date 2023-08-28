package util

import (
	"errors"
	"tiktok/internal/repository/models"
)

func FillCommentListFields(comments *[]*models.Comment) error {
	size := len(*comments)
	if comments == nil || size == 0 {
		return errors.New("util.FillCommentListFields comments为空")
	}
	dao := models.NewUserInfoDAO()
	for _, v := range *comments {
		_ = dao.QueryUserInfoById(v.UserInfoId, &v.User)
		v.CreateDate = v.CreatedAt.Format("1-2")
	}
	return nil
}

func FillCommentFields(comment *models.Comment) error {
	if comment == nil {
		return errors.New("FillCommentFields comments为空")
	}
	comment.CreateDate = comment.CreatedAt.Format("1-2")
	return nil
}
