package router

import (
	"github.com/gin-gonic/gin"
	"tiktok/conf"
	"tiktok/internal/api/comment"
	"tiktok/internal/api/favorite"
	"tiktok/internal/api/feed"
	"tiktok/internal/api/publish"
	"tiktok/internal/api/user/user_info"
	"tiktok/internal/api/user/user_login"
	"tiktok/internal/api/user/user_register"
	"tiktok/internal/repository/models"
	"tiktok/pkg/middleware"
)

func Init() *gin.Engine {
	models.InitDB()
	r := gin.Default()
	config := conf.NewConfig()
	r.Static("static", config.Path.StaticSourcePath)

	baseGroup := r.Group("/douyin")
	baseGroup.POST("/user/login/", middleware.SHAMiddleWare(), user_login.UserLoginHandler)
	baseGroup.POST("/user/register/", middleware.SHAMiddleWare(), user_register.UserRegisterHandler)

	baseGroup.GET("/feed/", feed.FeedVideoListHandler)
	baseGroup.GET("/user/", middleware.JWTMiddleWare(), user_info.UserInfoHandler)

	baseGroup.POST("/publish/action/", middleware.JWTMiddleWare(), publish.PublishVideoHandler)
	baseGroup.GET("/publish/list/", middleware.NoAuthToGetUserId(), publish.QueryVideoListHandler)

	baseGroup.POST("/favorite/action/", middleware.JWTMiddleWare(), favorite.PostFavorHandler)
	baseGroup.GET("/favorite/list/", middleware.NoAuthToGetUserId(), favorite.QueryFavorVideoListHandler)
	baseGroup.POST("/comment/action/", middleware.JWTMiddleWare(), comment.PostCommentHandler)
	baseGroup.GET("/comment/list/", middleware.JWTMiddleWare(), comment.QueryCommentListHandler)

	baseGroup.POST("/relation/action/", middleware.JWTMiddleWare(), user_info.PostFollowActionHandler)
	baseGroup.GET("/relation/follow/list/", middleware.NoAuthToGetUserId(), user_info.QueryFollowListHandler)
	baseGroup.GET("/relation/follower/list/", middleware.NoAuthToGetUserId(), user_info.QueryFollowerHandler)
	return r
}
