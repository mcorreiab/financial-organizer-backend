package main

import (
	"mcorreiab/financial-organizer-backend/internal/framework"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	jwtKey := os.Getenv("JWT_KEY")
	databaseConnection := framework.GetDatabaseConnection()
	router := framework.CreateRoutes(jwtKey, databaseConnection)

	router.Run()
}
