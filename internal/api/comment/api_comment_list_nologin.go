package comment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/internal/api/favorite"
	"tiktok/internal/repository/models"
	"tiktok/internal/services"
)

func QueryCommentListHandlerNologin(c *gin.Context) {
	NewProxyCommentListHandlerNologin(c).DoNologin()
}

type ProxyCommentListHandlerNologin struct {
	*gin.Context

	videoId int64
}

func NewProxyCommentListHandlerNologin(context *gin.Context) *ProxyCommentListHandlerNologin {
	return &ProxyCommentListHandlerNologin{Context: context}
}

func (p *ProxyCommentListHandlerNologin) DoNologin() {
	if err := p.parseNumNologin(); err != nil {
		p.SendErrorNologin(err.Error())
		return
	}

	commentList, err := services.QueryCommentListNologin(p.videoId)
	if err != nil {
		p.SendErrorNologin(err.Error())
		return
	}

	p.SendOkNologin(commentList)
}

func (p *ProxyCommentListHandlerNologin) parseNumNologin() error {

	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	return nil
}
func (p *ProxyCommentListHandlerNologin) SendErrorNologin(msg string) {
	p.JSON(http.StatusOK, favorite.FavorVideoListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyCommentListHandlerNologin) SendOkNologin(commentList *services.CommentsList) {
	p.JSON(http.StatusOK, ListResponse{CommonResponse: models.CommonResponse{StatusCode: 0},
		CommentsList: commentList,
	})
}
