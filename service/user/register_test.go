package main

import (
	"bytes"
	_ "config/github.com/go-sql-driver/mysql"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type RegisterRequest struct {
	Username string
	Password string
}

type RegisterResponse struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
	UserID     int64  `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}

func TestRegisterHandler(t *testing.T) {
	// 创建一个用于测试的临时数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/tiktok")
	if err != nil {
		t.Errorf("Failed to connect to the database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			t.Errorf("Failed to close to the database: %v", err)
		}
	}(db)

	// 检查数据库连接状态
	if err := db.Ping(); err != nil {
		t.Errorf("Failed to ping the database: %v", err)
	}

	// 创建一个模拟的 HTTP 请求和响应
	reqBody := RegisterRequest{
		Username: "test_username",
		Password: "test_password",
	}
	jsonData, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/douyin/user/register/?username=your_username&password=your_password", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	// 调用被测试的处理函数
	RegisterHandler(recorder, req)

	// 检查响应状态码
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}

	// 从响应中读取内容
	body, _ := io.ReadAll(recorder.Body)

	// 解析响应数据
	var resp RegisterResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Errorf("Failed to parse response body: %v", err)
	}

	// 检查响应内容
	if resp.UserID == 0 {
		t.Error("User ID is not set in the response")
	}

	if resp.Token == "" {
		t.Error("Token is not set in the response")
	}
}
