package data

//数据库连接和配置初始化
import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitDB() {
	// 数据库连接信息
	connStr := "root:008625@tcp(localhost:3306)/TikTok"
	database, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// 验证数据库连接
	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db = database
}

func GetDB() *sql.DB {
	return db
}
