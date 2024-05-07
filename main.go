package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connection parameters
	connStr := "user=postgres password=password dbname=postgres sslmode=disable"

	log.Println("Test connection")

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// defer db.Close()

	rows, err := db.Query("CREATE TABLE Persons (id int);")

	log.Println(rows)
	log.Println(err)

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	s := CreateServer()
	s.AddRoutes()
	s.Run()

	log.Println("Test complete")
}
