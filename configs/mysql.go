package configs

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var MYSQL_DB *sql.DB

func InitDB() {
	var err error
	MYSQL_DB, err = sql.Open("mysql", "root:123456QWE@tcp(192.168.242.246:3306)/test")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 测试数据库连接
	err = MYSQL_DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connected successfully")
}
