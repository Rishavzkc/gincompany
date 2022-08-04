package routes

import (
	"final/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	host   string
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		host:   config.GetConfig().ServerHost,
		port:   config.GetConfig().ServerPort,
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := ConfigRoutes(s.server)

	log.Printf("Server running at port: %v", s.port)
	log.Fatal(router.Run(":", s.port))

}
