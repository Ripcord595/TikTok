package test

import (
	"TikTok/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	// 构造测试请求的JSON数据
	jsonData := `{"username": "test_username", "password": "test_password"}`

	// 创建请求体
	requestBody := strings.NewReader(jsonData)

	// 创建模拟请求
	request := httptest.NewRequest(http.MethodPost, "/douyin/user/login/", requestBody)

	// 创建模拟响应写入器
	responseWriter := httptest.NewRecorder()

	// 调用被测试的处理函数
	service.LoginHandler(responseWriter, request)

	// 获取响应
	response := responseWriter.Result()

	// 检查状态码
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}
}
