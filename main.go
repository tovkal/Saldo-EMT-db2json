package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"os"
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

	// Get language id for value defined in constatnt language_export
	//languageId := getLanguageId(db)

	var buffer bytes.Buffer
	buffer.WriteString("{")
	buffer.WriteString(getBusLines(db))
	buffer.WriteString("}")

	var output bytes.Buffer
	err = json.Indent(&output, buffer.Bytes(), "", "  ")
	checkError(err)

	// Create file
	f, err := os.Create("out.json")
	checkError(err)
	defer f.Close()

	f.Write(output.Bytes())
	f.Sync()

	log.Println("Done!")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLanguageId(db *sql.DB) string {
	var languageId string

	err := db.QueryRow("SELECT id FROM Language WHERE code = ?", language_export).Scan(&languageId)
	checkError(err)

	return languageId
}

func getBusLines(db *sql.DB) string {
	var buffer bytes.Buffer

	rows, err := db.Query("SELECT BusLine.id, BusLine.hexColor, BusLineNameTranslation.name FROM BusLine INNER JOIN BusLineNameTranslation ON BusLine.id = BusLineNameTranslation.busLineId ")
	checkError(err)
	buffer.WriteString("\"lines\": [")

	firstTime := true
	defer rows.Close()
	for rows.Next() {
		var id int
		var hexColor string
		var name string
		err = rows.Scan(&id, &hexColor, &name)
		checkError(err)

		if firstTime {
			firstTime = false
		} else {
			buffer.WriteString(",")
		}

		buffer.WriteString("{")
		buffer.WriteString("\"" + strconv.Itoa(id) + "\" : { \"color\" : \"" + hexColor + "\", \"name\" : \"" + name + "\"}")
		buffer.WriteString("}")
	}
	buffer.WriteString("]")

	return buffer.String()
}