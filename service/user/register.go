package user

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//_ "github.com/go-sql-driver/mysql"
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

func main() {
	// 连接数据库
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	url := "douyin/user/register/?username=your_username&password=your_password"
	method := "POST"

	// 构建请求的数据
	requestData := RegisterRequest{
		Username: "your_username",
		Password: "your_password",
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Set("Content-Type", "application/json")

	// 将注册信息插入数据库表
	insertQuery := `
		INSERT INTO user (username, password)
		VALUES (?, ?)
	`

	_, err = db.Exec(insertQuery, requestData.Username, requestData.Password)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 解析响应数据
	responseData := RegisterResponse{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 处理响应数据
	fmt.Printf("Status Code: %d\n", responseData.StatusCode)
	fmt.Printf("Status Message: %s\n", responseData.StatusMsg)
	fmt.Printf("User ID: %d\n", responseData.UserID)
	fmt.Printf("Token: %s\n", responseData.Token)

	fmt.Println("注册成功！")
}
