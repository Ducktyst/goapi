package server

import (
	"fmt"
	"github.com/ducktyst/goapi/internal/config"
	"github.com/ducktyst/goapi/internal/database"
	"github.com/gin-gonic/gin"
)

// Server struct
type Server struct {
	router *gin.Engine
	port   int
	db     database.Database
}

// New create new Server instance
func New(cfg config.ServerConfig, db database.Database) *Server {
	return &Server{
		router: gin.Default(),
		port:   cfg.Port,
		db:     db,
	}
}

// Run server
func (s *Server) Run() {
	s.initHandlers()
	s.router.Run(fmt.Sprintf(":%d", s.port))
}

func (s *Server) initHandlers() {

	group := s.router.Group("/api")

	group.POST("/create", s.CreateExchangeRate)
	group.POST("/convert", s.Convert)
}
