package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateToken() string {
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
