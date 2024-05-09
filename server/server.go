package server

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *gin.Engine
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
