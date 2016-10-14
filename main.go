package main

import (
	"bytes"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

const db_host string = "localhost"
const db_user string = "root"
const db_name string = "saldo_emt"
const language_export string = "es"

func main() {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", db_user+":@"+"/"+db_name)
	checkError(err)

	defer db.Close()

	var languageId string

	err = db.QueryRow("SELECT id FROM Language WHERE code = ?", language_export).Scan(&languageId)
	checkError(err)

	var buffer bytes.Buffer

	buffer.WriteString("{")

	buffer.WriteString(getBusLines(db))

	buffer.WriteString("}")
	log.Println("Resulting JSON:")
	log.Print(buffer.String())
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getBusLines(db *sql.DB) string {
	var buffer bytes.Buffer

	rows, err := db.Query("SELECT * FROM BusLine")
	checkError(err)
	buffer.WriteString("\"lines\": [")

	firstTime := true
	defer rows.Close()
	for rows.Next() {
		var id int
		var hexColor string
		err = rows.Scan(&id, &hexColor)
		checkError(err)

		if firstTime {
			firstTime = false
		} else {
			buffer.WriteString(",")
		}

		buffer.WriteString("{")
		buffer.WriteString("\"" + strconv.Itoa(id) + "\" : { \"color\" : \"" + hexColor + "\"}")
		//buffer.WriteString("\"" + id + "\" : { \"color\" : \"" + hexColor + "\", \"name\" : " + name + "\"}")
		buffer.WriteString("}")

		//fmt.Printf("id: %v, hexColor: %v\n", id, hexColor)
	}
	buffer.WriteString("]")

	return buffer.String()
}
