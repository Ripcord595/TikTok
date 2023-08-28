package favorite

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/internal/repository/models"
	"tiktok/internal/services"
)

type FavorVideoListResponse struct {
	models.CommonResponse
	*services.FavorList
}

func QueryFavorVideoListHandler(c *gin.Context) {
	NewProxyFavorVideoListHandler(c).Do()
}

type ProxyFavorVideoListHandler struct {
	*gin.Context

	userId int64
}

func NewProxyFavorVideoListHandler(c *gin.Context) *ProxyFavorVideoListHandler {
	return &ProxyFavorVideoListHandler{Context: c}
}

func (p *ProxyFavorVideoListHandler) Do() {
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}

	favorVideoList, err := services.QueryFavorVideoList(p.userId)
	if err != nil {
		p.SendError(err.Error())
		return
	}

	p.SendOk(favorVideoList)
}

func (p *ProxyFavorVideoListHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyFavorVideoListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyFavorVideoListHandler) SendOk(favorList *services.FavorList) {
	p.JSON(http.StatusOK, FavorVideoListResponse{CommonResponse: models.CommonResponse{StatusCode: 0},
		FavorList: favorList,
	})
}
