package main

import (
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	router := gin.Default()
	framework.CreateUserRoutes(
		usecase.NewUserUseCase(
			adapter.NewUserRepository(framework.GetDatabaseConnection()),
			os.Getenv("JWT_KEY"),
		),
		router,
	)

	router.Run()
}
