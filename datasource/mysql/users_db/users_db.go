package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {

	log.Println("insiden init")
	var connErr error
	Client, connErr = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=default", username, password, host, schema))

	if connErr != nil {
		panic(connErr)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Datasource configured successfully")
}
