package feed

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/internal/repository/models"
	"tiktok/internal/services"
	"tiktok/pkg/middleware"
	"time"
)

type FeedResponse struct {
	models.CommonResponse
	*services.FeedVideoList
}

func FeedVideoListHandler(c *gin.Context) {

	p := NewProxyFeedVideoList(c)

	token, ok := c.GetQuery("token")
	fmt.Printf("1======")
	if !ok {
		err := p.DoNoToken()
		fmt.Printf("2======")
		if err != nil {
			fmt.Printf("3======")
			p.FeedVideoListError(err.Error())
		}
		fmt.Printf("4======")
		return
	}
	fmt.Printf("5======")
	err := p.DoHasToken(token)
	if err != nil {
		p.FeedVideoListError(err.Error())
	}
}

type ProxyFeedVideoList struct {
	*gin.Context
}

func NewProxyFeedVideoList(c *gin.Context) *ProxyFeedVideoList {
	return &ProxyFeedVideoList{Context: c}
}

func (p *ProxyFeedVideoList) DoNoToken() error {
	rawTimestamp := p.Query("latest_time")
	var latestTime time.Time
	intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
	if err == nil {
		latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
	}
	fmt.Printf("6======")
	videoList, err := services.QueryFeedVideoList(0, latestTime)
	if err != nil {
		fmt.Printf("7======")
		return err
	}
	fmt.Printf("8======")
	p.FeedVideoListOk(videoList)
	return nil
}

func (p *ProxyFeedVideoList) DoHasToken(token string) error {

	if claim, ok := middleware.ParseToken(token); ok {
		if time.Now().Unix() > claim.ExpiresAt {
			return errors.New("token超时")
		}
		rawTimestamp := p.Query("latest_time")
		var latestTime time.Time
		intTime, err := strconv.ParseInt(rawTimestamp, 10, 64)
		if err != nil {
			latestTime = time.Unix(0, intTime*1e6) //注意：前端传来的时间戳是以ms为单位的
		}

		videoList, err := services.QueryFeedVideoList(claim.UserId, latestTime)
		if err != nil {
			return err
		}
		p.FeedVideoListOk(videoList)
		return nil
	}
	return errors.New("token不正确")
}

func (p *ProxyFeedVideoList) FeedVideoListError(msg string) {
	p.JSON(http.StatusOK, FeedResponse{CommonResponse: models.CommonResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	}})
}

func (p *ProxyFeedVideoList) FeedVideoListOk(videoList *services.FeedVideoList) {
	fmt.Printf("9======")
	p.JSON(http.StatusOK, FeedResponse{

		CommonResponse: models.CommonResponse{
			StatusCode: 0,
		},
		FeedVideoList: videoList,
	},
	)
}
