package main

import (
	"TikTok/controller"
	_ "config/github.com/go-sql-driver/mysql"
	"database/sql"
	"net/http"
)

func LoginHandler(writer http.ResponseWriter, request *http.Request) {

	// 连接数据库
	dsn := "root:123456@tcp(localhost:3306)/tiktok" // 要改成自己的数据源
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	//获取请求体数据
	requestData := controller.HandleRequest(writer, request)

	// 执行查询操作，验证用户名和密码
	query := `
		SELECT id, token FROM user WHERE username = ? AND password = ?
	`
	var userID int64
	var token string
	err = db.QueryRow(query, requestData.Username, requestData.Password).Scan(&userID, &token)
	if err != nil {
		http.Error(writer, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	//发送响应数据
	controller.HandleResponse(userID, token, writer, "登录成功！")
}

func main() {
	http.HandleFunc("/douyin/user/register/", RegisterHandler)
	http.HandleFunc("/douyin/user/login/", LoginHandler)
	http.ListenAndServe(":8080", nil)
}
