package main

import (
	"ECCO2K/internal/database"
	"ECCO2K/internal/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	r := gin.Default()

	router.RouterSetup(r, database.DB)
	
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("ошибка при запуске сервера")
	}
}
