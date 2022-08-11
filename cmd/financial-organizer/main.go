package main

import (
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"mcorreiab/financial-organizer-backend/internal/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	saveUserController := framework.NewSaveUserController(
		usecase.NewSaveUserUseCase(adapter.NewUserRepository(framework.GetDatabaseConnection())),
	)

	router := gin.Default()
	router.POST("/users", saveUserController.Save)

	router.Run()
}
