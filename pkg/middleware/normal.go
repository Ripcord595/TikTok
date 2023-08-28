package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/internal/repository/models"
)

func NoAuthToGetUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawId := c.Query("user_id")
		if rawId == "" {
			rawId = c.PostForm("user_id")
		}
		if rawId == "" {
			c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 401, StatusMsg: "用户不存在"})
			c.Abort()
			return
		}
		userId, err := strconv.ParseInt(rawId, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, models.CommonResponse{StatusCode: 401, StatusMsg: "用户不存在"})
			c.Abort()
		}
		c.Set("user_id", userId)
		c.Next()
	}
}
