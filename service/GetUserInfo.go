package service

import (
	"TikTok/data"
	"TikTok/model"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUserInfo(c *gin.Context) {
	// 获取 Query 参数
	userID := c.Query("id")
	token := c.Query("token")

	// 将 userID 转换为 int64 类型
	id, err := strconv.ParseInt(userID, 10, 64)
	// 根据用户id和token从数据库或其他存储中获取用户信息
	user, err := getUserFromDB(id, token)
	if err != nil {
		// 处理获取用户信息失败的情况，比如返回错误信息或者默认值
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取用户信息",
		})
		return
	}
	/*
		// 模拟获取到的用户信息
		user = model.User{
			Avatar:          "ss",
			BackgroundImage: "小明",
			FavoriteCount:   100,
			FollowCount:     200,
			FollowerCount:   111,
			ID:              11123334452453,
			IsFollow:        true,
			Name:            "222",
			Signature:       "简介",
			TotalFavorited:  234,
			WorkCount:       500,
		}
	*/
	// 构造响应数据
	response := gin.H{
		"status_code": 0,
		"status_msg":  "成功",
		"user":        user,
	}

	c.JSON(http.StatusOK, response)
}

func getUserFromDB(userID int64, token string) (*model.User, error) {
	// 连接数据库
	var writer http.ResponseWriter
	db, err := data.DbConnect(writer)

	// 执行查询语句
	query := fmt.Sprintf("SELECT * FROM user WHERE id = '%d' AND token = '%s'", userID, token)
	row := db.QueryRow(query)

	// 读取查询结果
	var user model.User
	err = row.Scan(&user.ID, &user.Name, &user.FollowCount, &user.FollowerCount, &user.IsFollow, &user.Avatar, &user.BackgroundImage, &user.Signature, &user.TotalFavorited, &user.WorkCount, &user.FavoriteCount)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有找到对应用户
			return nil, fmt.Errorf("用户不存在")
		}
		log.Println("读取查询结果失败:", err)
		return nil, err
	}

	return &user, nil
}
