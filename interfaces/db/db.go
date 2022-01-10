package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB //Exported Database connection
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
	user := "root"
	pass := "Sound@123"
	host := "localhost"
	credentials := fmt.Sprintf("%s:%s@(%s:3306)/gotodo?charset=utf8&parseTime=True", user, pass, host)
	var err error
	db, err = sql.Open("mysql", credentials)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connected successfully!!")
	}
	return db
}
