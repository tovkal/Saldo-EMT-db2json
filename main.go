package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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
	languageId := getLanguageId(db)

	var buffer bytes.Buffer
	time := time.Now()
	buffer.WriteString("{")
	buffer.WriteString("\"timestamp\":" + time.Format("20060102") + ",")
	buffer.WriteString(getFares(db, languageId))
	buffer.WriteString(getBusLines(db))
	buffer.WriteString("}")

	var output bytes.Buffer
	err = json.Indent(&output, buffer.Bytes(), "", "  ")
	checkError(err)

	uploadFile(output)

	log.Println("Done!")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/* Parse fares */

func getLanguageId(db *sql.DB) string {
	var languageId string

	err := db.QueryRow("SELECT id FROM Language WHERE code = ?", language_export).Scan(&languageId)
	checkError(err)

	return languageId
}

func getFares(db *sql.DB, languageId string) string {
	var buffer bytes.Buffer

	rows, err := db.Query("SELECT Fare.id, Fare.days, Fare.rides, FareNameTranslation.name FROM Fare INNER JOIN FareNameTranslation ON Fare.id = FareNameTranslation.fareId WHERE FareNameTranslation.language = ?", languageId)
	checkError(err)
	defer rows.Close()
	buffer.WriteString("\"fares\": [")

	firstFare := true
	fareId := 1
	for rows.Next() {
		if firstFare {
			firstFare = false
		} else {
			buffer.WriteString(",")
		}

		var id int
		var days sql.NullInt64
		var rides sql.NullInt64
		var name string

		err = rows.Scan(&id, &days, &rides, &name)
		checkError(err)

		busLineForPriceMap := getFarePriceAndBusLines(db, id)

		first := true
		for price, busLines := range busLineForPriceMap {
			if first {
				first = false
			} else {
				buffer.WriteString(",")
			}
			buffer.WriteString(buildFare(fareId, days, rides, name, price, busLines))
			fareId++
		}
	}

	buffer.WriteString("],")

	return buffer.String()
}

func buildFare(id int, days sql.NullInt64, rides sql.NullInt64, name string, price string, busLines []int) string {
	var buffer bytes.Buffer

	buffer.WriteString("{\"" + strconv.Itoa(id) + "\": " + "{")

	firstAttribute := true

	// days
	if days.Valid {
		buffer.WriteString("\"days\":" + strconv.Itoa(int(days.Int64)))
		firstAttribute = false
	}

	// rides
	if rides.Valid {
		if !firstAttribute {
			buffer.WriteString(",")
		} else {
			firstAttribute = false
		}

		buffer.WriteString("\"rides\":" + strconv.Itoa(int(rides.Int64)))
	}

	// name
	if !firstAttribute {
		buffer.WriteString(",")
	} else {
		firstAttribute = false
	}
	buffer.WriteString("\"name\": \"" + name + "\"")

	// price
	buffer.WriteString(",\"price\": " + price)

	// busLines
	firstBusLine := true
	buffer.WriteString(",\"lines\": [")
	for _, busLine := range busLines {

		if firstBusLine {
			firstBusLine = false
		} else {
			buffer.WriteString(",")
		}

		buffer.WriteString(strconv.Itoa(busLine))
	}
	buffer.WriteString("]")

	buffer.WriteString("}")
	buffer.WriteString("}")

	return buffer.String()
}

func getFarePriceAndBusLines(db *sql.DB, id int) map[string][]int {
	rows, err := db.Query("SELECT busLineId, price FROM BusLineFarePrice WHERE fareId = ?", id)
	checkError(err)
	defer rows.Close()

	busLineForPriceMap := make(map[string][]int)

	var busLineId int
	var price string
	for rows.Next() {
		err = rows.Scan(&busLineId, &price)
		checkError(err)

		busLineForPriceMap[price] = append(busLineForPriceMap[price], busLineId)
	}

	return busLineForPriceMap
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

/* Upload file to aws S3 */

func uploadFile(output bytes.Buffer) {
	bucket := "saldo-emt"
	key := "fares_" + language_export + ".json"

	log.Println("Uploading json to s3...")

	svc := s3.New(session.New(&aws.Config{Region: aws.String("eu-central-1")}))

	_, err := svc.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader(output.String()),
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		log.Printf("Failed to upload data to %s/%s, %s\n", bucket, key, err)
		return
	}

	log.Printf("Successfully uploaded file %s to bucket %s\n", key, bucket)
}
