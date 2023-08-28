package util

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"tiktok/conf"
	"tiktok/internal/repository/cache"
	"tiktok/internal/repository/models"
	"time"
)

func GetFileUrl(fileName string) string {
	config := conf.NewConfig()

	base := fmt.Sprintf("http://%s:%d/static/%s", config.Server.IP, config.Server.Port, fileName)
	return base
}

func NewFileName(userId int64) string {
	var count int64

	err := models.NewVideoDAO().QueryVideoCountByUserId(userId, &count)
	if err != nil {
		log.Println(err)
	}
	return fmt.Sprintf("%d-%d", userId, count)
}

func FillVideoListFields(userId int64, videos *[]*models.Video) (*time.Time, error) {
	size := len(*videos)
	if videos == nil || size == 0 {
		return nil, errors.New("util.FillVideoListFields videos为空")
	}
	dao := models.NewUserInfoDAO()
	p := cache.NewProxyIndexMap()

	latestTime := (*videos)[size-1].CreatedAt
	for i := 0; i < size; i++ {
		var userInfo models.UserInfo
		err := dao.QueryUserInfoById((*videos)[i].UserInfoId, &userInfo)
		if err != nil {
			continue
		}
		userInfo.IsFollow = p.GetUserRelation(userId, userInfo.Id)
		(*videos)[i].Author = userInfo
		if userId > 0 {
			(*videos)[i].IsFavorite = p.GetVideoFavorState(userId, (*videos)[i].Id)
		}
	}
	return &latestTime, nil
}

func SaveImageFromVideo(name string, isDebug bool) error {
	v2i := NewVideo2Image()
	if isDebug {
		v2i.Debug()
	}
	config := conf.NewConfig()
	v2i.InputPath = filepath.Join(config.Path.StaticSourcePath, name+defaultVideoSuffix)
	v2i.OutputPath = filepath.Join(config.Path.StaticSourcePath, name+defaultImageSuffix)
	v2i.FrameCount = 1
	queryString, err := v2i.GetQueryString()
	if err != nil {
		return err
	}
	return v2i.ExecCommand(queryString)
}
