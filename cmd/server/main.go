package main

import (
	"log"

	"go-fridge/internal/config"
	"go-fridge/internal/database"
	"go-fridge/internal/router"

	_ "go-fridge/internal/docs"
)

// @title Fridge API
// @version 1.0
// @description Simple inventory tracking REST API built with Gin.
// @host localhost:8080
// @BasePath /api
func main() {
	cfg := config.Load()

	db, err := database.Init(cfg)
	if err != nil {
		log.Fatalf("could not connect to db %v", err)
	}

	r := router.Setup(db)
	log.Println("Starting server on :", cfg.Port)
	r.Run(":" + cfg.Port)
}
