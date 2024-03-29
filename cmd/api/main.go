package main

import (
	"log"
	"os"
	"re-go-challenge/pkg/packets"
	"re-go-challenge/pkg/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	// Get port from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Service instance
	s, err := server.New(packets.New())
	if err != nil {
		log.Fatal(err)
	}

	// Echo instance
	e := echo.New()

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	// Routes
	e.POST("/pack", s.GetPacks)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
