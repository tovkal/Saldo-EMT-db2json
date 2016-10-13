package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const db_host string = "localhost"
const db_user string = "root"
const db_name string = "saldo_emt"
const language_export string = "es"

func main() {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", db_user+":@"+"/"+db_name)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM BusLine")

	defer rows.Close()
	for rows.Next() {
		var id int
		var hexColor string
		err = rows.Scan(&id, &hexColor)
		fmt.Printf("id: %v, hexColor: %v\n", id, hexColor)
	}
}
