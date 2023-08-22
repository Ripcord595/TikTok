package service

import (
	"TikTok/data"
	"TikTok/handler"
	"TikTok/util"
	_ "config/github.com/go-sql-driver/mysql"
	"net/http"
)

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	// 连接数据库
	db, err := data.DbConnect(writer)

	//获取请求体数据
	requestData := handler.HandleRequest(writer, request)

	// 将注册信息插入数据库表
	// 查询数据库检查用户名是否已存在
	query := `
		SELECT COUNT(*) FROM user WHERE username = ?
	`
	var count int
	err = db.QueryRow(query, requestData.Username).Scan(&count)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if count > 0 {
		http.Error(writer, "用户名已存在", http.StatusBadRequest)
		return
	}

	token := util.GenerateToken() // 生成令牌
	insertQuery := `
		INSERT INTO user (username, password, token)
		VALUES (?, ?, ?)
	`
	result, err := db.Exec(insertQuery, requestData.Username, requestData.Password, token)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// 发送响应数据
	handler.HandleResponse(userID, token, writer, "注册成功！")
}
