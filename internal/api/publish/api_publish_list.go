package publish

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/internal/repository/models"
	"tiktok/internal/services"
)

type ListResponse struct {
	models.CommonResponse
	*services.VideoList
}

func QueryVideoListHandler(c *gin.Context) {
	p := NewProxyQueryVideoList(c)
	rawId, _ := c.Get("user_id")
	err := p.DoQueryVideoListByUserId(rawId)
	if err != nil {
		p.QueryVideoListError(err.Error())
	}
}

type ProxyQueryVideoList struct {
	c *gin.Context
}

func NewProxyQueryVideoList(c *gin.Context) *ProxyQueryVideoList {
	return &ProxyQueryVideoList{c: c}
}

func (p *ProxyQueryVideoList) DoQueryVideoListByUserId(rawId interface{}) error {
	userId, ok := rawId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}

	videoList, err := services.QueryVideoListByUserId(userId)
	if err != nil {
		return err
	}

	p.QueryVideoListOk(videoList)
	return nil
}

func (p *ProxyQueryVideoList) QueryVideoListError(msg string) {
	p.c.JSON(http.StatusOK, ListResponse{CommonResponse: models.CommonResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	}})
}

func (p *ProxyQueryVideoList) QueryVideoListOk(videoList *services.VideoList) {
	p.c.JSON(http.StatusOK, ListResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 0,
		},
		VideoList: videoList,
	})
}
