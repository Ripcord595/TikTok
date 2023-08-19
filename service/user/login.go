package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "config/github.com/go-sql-driver/mysql"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserID     int    `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}

func LoginHandler(writer http.ResponseWriter, request *http.Request) {
	// 检查请求方法是否为POST
	if request.Method != http.MethodPost {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// 定义登录接口的URL
	//url := "/douyin/user/login/"

	// 读取请求体数据
	requestBody, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// 解析JSON数据
	var requestData LoginRequest
	err = json.Unmarshal(requestBody, &requestData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/tiktok")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// 执行查询操作，验证用户名和密码
	query := `
		SELECT id, token FROM user WHERE username = ? AND password = ?
	`
	var userID int
	var token string
	err = db.QueryRow(query, requestData.Username, requestData.Password).Scan(&userID, &token)
	if err != nil {
		http.Error(writer, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	// 构造登录响应
	response := LoginResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "登录成功",
		UserID:     userID,
		Token:      token,
	}

	// 将响应转换为JSON并写入响应体中
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

//func main() {
//	http.HandleFunc("/douyin/user/login/", LoginHandler)
//	http.ListenAndServe(":8080", nil)
//}
