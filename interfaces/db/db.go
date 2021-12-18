package sqldb

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*Create mysql connection*/
func CreateCon() *sql.DB {
	db, err := sql.Open("mysql", "todoDB/todo")
	//db, err := sql.Open("mysql", "devstructsdit:7pi!!arsofbit$athy@tcp(bitsathyrewardsphase1.cpmfvw3nvpsc.ap-south-1.rds.amazonaws.com:3306)/bitphaseone")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("MySQL db is not connected")
		fmt.Println(err.Error())
	}
	return db
}
