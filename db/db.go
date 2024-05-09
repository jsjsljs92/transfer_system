package db

import (
	"database/sql"
	"fmt"
	"log"
)

func SetupDB() *sql.DB {
	// Connection parameters
	connStr := "user=postgres password=password dbname=postgres sslmode=disable"

	log.Println("Test connection")

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	return db
}

func SetupTables(db *sql.DB) {

	// Create account table
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		acc_id INT NOT NULL UNIQUE,
		balance FLOAT NOT NULL,
		version INT NOT NULL,
		timestamp TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Create payin table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS payin (
		id SERIAL PRIMARY KEY,
		to_acc_id INT NOT NULL,
		amount FLOAT NOT NULL,
		timestamp TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Create payout table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS payout (
		id SERIAL PRIMARY KEY NOT NULL,
		from_acc_id INT NOT NULL,
		amount FLOAT NOT NULL,
		timestamp TIMESTAMP NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All tables created successfully")
}
