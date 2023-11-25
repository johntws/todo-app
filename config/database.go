package route

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"todo-app/repository"
	"todo-app/service"
)

var DB *sql.DB

func InitializeMysql() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "todos",
		AllowNativePasswords: true,
	}

	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	repository.DB = DB
	service.DB = DB
}
