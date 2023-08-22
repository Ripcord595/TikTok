package data

import (
	"log"
)

func InsertUser(username, password string) error {
	db := data.GetDB() // 获取数据库连接
	_, err := db.Exec("INSERT INTO user (username, password) VALUES (?, ?)", username, password)
	if err != nil {
		log.Println("Failed to insert user:", err)
		return err
	}
	return nil
}
