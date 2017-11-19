package main

import (
	"bytes"
	"database/sql"
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
const db_name string = "SaldoEMT"
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
	buffer.WriteString("}")

	uploadFile(buffer)

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

	rows, err := db.Query("SELECT t.name 'fareName', bt.name 'busLineTypeName', bltf.cost, f.days, f.rides, b.imageUrl, b.displayBusLineTypeName FROM Fare f INNER JOIN FareNameTranslation t ON f.id = t.fareId INNER JOIN BusLineTypeFare bltf ON f.id = bltf.fareId INNER JOIN BusLineTypeTranslation bt ON bltf.busLineTypeId = bt.busLineTypeId AND bt.language = ? INNER JOIN BusLineType b ON bt.busLineTypeId = b.id WHERE t.language = ? ORDER BY bt.id, f.id;", languageId, languageId)
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

		var name string
		var busLineType string
		var cost float64
		var days sql.NullInt64
		var rides sql.NullInt64
		var imageUrl string
		var displayBusLineTypeName string

		err = rows.Scan(&name, &busLineType, &cost, &days, &rides, &imageUrl, &displayBusLineTypeName)
		checkError(err)

		buffer.WriteString(buildFare(fareId, name, busLineType, cost, days, rides, imageUrl, displayBusLineTypeName))
		fareId++
	}

	buffer.WriteString("]")

	return buffer.String()
}

func buildFare(id int, name string, busLineType string, cost float64, days sql.NullInt64, rides sql.NullInt64, imageUrl string, displayBusLineTypeName string) string {
	var buffer bytes.Buffer

	buffer.WriteString("{\"" + strconv.Itoa(id) + "\": " + "{")

	buffer.WriteString("\"name\": \"" + name + "\"")
	buffer.WriteString(",\"busLineType\": \"" + busLineType + "\"")
	buffer.WriteString(",\"cost\": \"" + strconv.FormatFloat(cost, 'f', 3, 64) + "\"")
	buffer.WriteString(",\"imageUrl\": \"" + imageUrl + "\"")
	buffer.WriteString(",\"displayBusLineTypeName\": " + strconv.FormatBool(displayBusLineTypeName == "\x01"))

	// days
	if days.Valid {
		buffer.WriteString(",\"days\":" + strconv.Itoa(int(days.Int64)))
	}

	// rides
	if rides.Valid {
		buffer.WriteString(",\"rides\":" + strconv.Itoa(int(rides.Int64)))
	}

	buffer.WriteString("}")
	buffer.WriteString("}")

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
