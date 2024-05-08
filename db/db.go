package db

// func setupDB() *sql.DB {
// 	// Connection parameters
// 	connStr := "user=postgres password=password dbname=postgres sslmode=disable"

// 	log.Println("Test connection")

// 	// Connect to the database
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	}
// 	// defer db.Close()

// 	rows, err := db.Query("CREATE TABLE Persons (id int);")

// 	log.Println(rows)
// 	log.Println(err)

// 	// Ping the database to verify the connection
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("Failed to ping the database: %v", err)
// 	}
// 	return db
// }

// func setupTables(db *sql.DB) {
// 	// Open database connection
// 	// db, err := sql.Open("sqlite3", "test.db")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	defer db.Close()

// 	// Create account table
// 	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS account (
// 		id INTEGER PRIMARY KEY,
// 		acc_id INTEGER UNIQUE,
// 		balance FLOAT,
// 		version INTEGER,
// 		timestamp TIMESTAMP
// 	)`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create payin table
// 	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS payin (
// 		id INTEGER PRIMARY KEY,
// 		to_acc_id INTEGER,
// 		amount FLOAT,
// 		timestamp TIMESTAMP
// 	)`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Create payout table
// 	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS payout (
// 		id INTEGER PRIMARY KEY,
// 		from_acc_id INTEGER,
// 		amount FLOAT,
// 		timestamp TIMESTAMP
// 	)`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("All tables created successfully")
// }
