package handler

import (
	"TikTok/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func HandleRequest(writer http.ResponseWriter, request *http.Request) model.UserRequest {
	var requestData model.UserRequest
	// 检查请求方法是否为POST
	if request.Method != http.MethodPost {
		http.Error(writer, "Method Not Allowed", http.StatusMethodNotAllowed)
		return requestData
	}

	// 读取请求体数据
	requestBody, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return requestData
	}

	// 解析JSON数据
	err = json.Unmarshal(requestBody, &requestData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return requestData
	}

	return requestData
}
