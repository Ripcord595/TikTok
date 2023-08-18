package main

import (
	_ "config/github.com/go-sql-driver/mysql"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int    `json:"user_id"`
	Token      string `json:"token"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/tiktok")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	//检查是否连接成功
	err = db.Ping()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 读取请求体数据
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 解析请求体数据
	var requestData RegisterRequest
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 将注册信息插入数据库表
	insertQuery := `
		INSERT INTO user (username, password)
		VALUES (?, ?)
	`

	result, err := db.Exec(insertQuery, requestData.Username, requestData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 构建响应数据
	responseData := RegisterResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "注册成功！",
		UserID:     int(userID),
		Token:      generateToken(),
	}

	// 将响应数据转换为 JSON 格式
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头部
	w.Header().Set("Content-Type", "application/json")

	// 发送响应数据
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
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

func main() {
	http.HandleFunc("/douyin/user/register/", RegisterHandler)
	http.ListenAndServe(":8080", nil)
}
