package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root", "", "127.0.0.1:3306", "users_db")
	var err error
	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database configured successfully")
}
