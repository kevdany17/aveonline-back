package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Server struct{
	Gin *gin.Engine
}

func(server *Server) InitServer(){	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	port := os.Getenv("SERVER_PORT")
	server.Gin.Run(":"+port)
}

func NewServer() *Server{
	server := Server{}
	server.Gin = gin.New()
	return &server
}