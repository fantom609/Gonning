package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func connectionDatabase() *sql.DB {
	db, err := sql.Open("postgres", "host=surus.db.elephantsql.com user=ogtmmrkn password=cYQHw9URHn5KIY26dXhP_CEdHWxpipuQ "+
		"dbname=ogtmmrkn sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateEvent() {
	var id int
	db := connectionDatabase()
	err := db.QueryRow("INSERT INTO Event (title,startDate,endDate,location,tag) VALUES ($1,$2,$3,$4,$5) RETURNING id", "Alice", "2023-03-11", "2023-03-17", "la", "ici").Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d", id)
}

func PatchEvent() {}

func DeleteEvent() {}

func GetEvent() {}

func GetEvents() {}
