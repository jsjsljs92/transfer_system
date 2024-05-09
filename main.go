package main

import (
	_ "github.com/lib/pq"

	"github.com/jsjsljs92/transferSystem/db"
	"github.com/jsjsljs92/transferSystem/server"
)

func main() {

	// setup db and tables
	DB := db.SetupDB()
	db.SetupTables(DB)

	//setup gin server and routes
	s := server.CreateServer()
	s.AddRoutes(DB)
	s.Run()
}
