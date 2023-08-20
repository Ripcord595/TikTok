package main

import (
	"TikTok/controller"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"net/http"

	_ "config/github.com/go-sql-driver/mysql"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

func RegisterHandler(writer http.ResponseWriter, request *http.Request) {
	// 连接数据库
	dsn := "root:123456@tcp(localhost:3306)/tiktok" // 要改成自己的数据源
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// 检查是否连接成功
	err = db.Ping()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	//获取请求体数据
	requestData := controller.HandleRequest(writer, request)

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

	insertQuery := `
		INSERT INTO user (username, password, token)
		VALUES (?, ?, ?)
	`
	token := generateToken() // 生成令牌
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
	controller.HandleResponse(userID, token, writer, "注册成功！")
}

func generateToken() string {
	// 生成一个具有足够安全性的随机字节数组
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		// 生成失败时，可以返回错误或使用默认令牌
		return "default_token"
	}

	// 对随机字节数组进行Base64编码，生成字符串形式的令牌
	token := base64.URLEncoding.EncodeToString(randomBytes)

	return token
}
