package service

import (
	"database/sql"
	"net/http"
)

func DbConnect(writer http.ResponseWriter) (db *sql.DB) {
	/*
		连接数据库只需要写：
			db := service.DbConnect(writer)
		其中writer是网络响应，在这里用来返回错误状态码
	*/
	dsn := "root:123456@tcp(localhost:3306)/tiktok" // 要改成自己的数据源
	// 连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return db
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	}(db)

	// 检查是否连接成功
	err = db.Ping()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return db
	}
	return db
}
