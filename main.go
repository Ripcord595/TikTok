package main

import (
	"fmt"
	"tiktok/conf"
	"tiktok/internal/repository/models"
	"tiktok/internal/router"
)

func main() {
	//migration初始化
	models.InitDB()

	//初始化路由引擎
	config := conf.NewConfig()

	err := router.Init().Run(fmt.Sprintf(":%d", config.Server.Port))
	if err != nil {
		fmt.Println("Faild to run Engine!", err)
	}

}
