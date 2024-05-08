package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jsjsljs92/transferSystem/src/components/account"
	_ "github.com/lib/pq"
)

type Server struct {
	server *gin.Engine
}

func main() {

	// setup db and tables
	db := setupDB()
	setupTables(db)

	//setup gin server and routes
	s := CreateServer()
	s.AddRoutes(db)
	s.Run()
}

func CreateServer() *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	server := &Server{engine}
	return server
}

func (s *Server) AddRoutes(db *sql.DB) {
	CreateRoutes(s.server, db)
}

func (s *Server) Run() {
	err := s.server.Run(fmt.Sprintf(`:%v`, 8080))

	if err != nil {
		log.Printf(`Failed to start server with error: %v`, err)
	}
}

func setupDB() *sql.DB {
	// Connection parameters
	connStr := "user=postgres password=password dbname=postgres sslmode=disable"

	log.Println("Test connection")

	// Connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// defer db.Close()

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}
	return db
}

func setupTables(db *sql.DB) {
	// Open database connection
	// db, err := sql.Open("sqlite3", "test.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()

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

func CreateRoutes(route *gin.Engine, db *sql.DB) {
	{
		v1 := route.Group("/v1")

		// create account
		{
			accountApi := v1.Group("/accounts")
			accountController := account.NewAccountController(db)

			accountApi.POST("", accountController.CreateAccount)
			accountApi.GET("/:id", accountController.GetAccountByID)
		}
	}
}
