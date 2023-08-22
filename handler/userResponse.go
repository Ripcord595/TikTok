package handler

import (
	"TikTok/model"
	"encoding/json"
	"net/http"
)

func HandleResponse(userID int64, token string, writer http.ResponseWriter, StatusMsg string) {
	// 构造响应数据
	responseData := model.UserResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  StatusMsg,
		UserID:     userID,
		Token:      token,
	}

	// 将响应转换为JSON并写入响应体中
	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头部
	writer.Header().Set("Content-Type", "application/json")

	// 发送响应数据
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
}
