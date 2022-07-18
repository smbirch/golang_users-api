package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB
	//these may be used later for cloud hosting if I do not use godotenv.
	//leaving here to remind myself what I was doing

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {
	// this is for godotenv.
	// godoterr := godotenv.Load()
	// if godoterr != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// username := os.Getenv("mysql_users_username")
	// password := os.Getenv("mysql_users_password")
	// host := os.Getenv("mysql_users_host")
	// schema := os.Getenv("mysql_users_schema")

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema)

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
