package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/conf"
	"tiktok/internal/api/comment"
	"tiktok/internal/api/favorite"
	"tiktok/internal/api/feed"
	"tiktok/internal/api/publish"
	"tiktok/internal/api/user/user_login"
	"tiktok/internal/repository/models"
	"tiktok/pkg/middleware"
)

func Init() *gin.Engine {
	models.InitDB()
	r := gin.Default()
	config := conf.NewConfig()
	r.Static("static", config.Path.StaticSourcePath)

	baseGroup := r.Group("/douyin")
	baseGroup.POST("/user/login/", middleware.SHAMiddleWare(), user.UserLoginHandler)
	baseGroup.POST("/user/register/", middleware.SHAMiddleWare(), user.UserRegisterHandler)

	baseGroup.GET("/feed/", feed.FeedVideoListHandler)
	baseGroup.GET("/user/", middleware.JWTMiddleWare(), user.UserInfoHandler)

	baseGroup.POST("/publish/action/", middleware.JWTMiddleWare(), publish.PublishVideoHandler)
	baseGroup.GET("/publish/list/", middleware.NoAuthToGetUserId(), publish.QueryVideoListHandler)

	baseGroup.POST("/favorite/action/", middleware.JWTMiddleWare(), favorite.PostFavorHandler)
	baseGroup.GET("/favorite/list/", middleware.NoAuthToGetUserId(), favorite.QueryFavorVideoListHandler)
	baseGroup.POST("/comment/action/", middleware.JWTMiddleWare(), comment.PostCommentHandler)
	baseGroup.GET("/comment/list/", middleware.JWTMiddleWare(), comment.QueryCommentListHandler)
	return r
}
