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
	jwtKey := os.Getenv("JWT_KEY")

	databaseConnection := framework.GetDatabaseConnection()
	userRepository := adapter.NewUserRepository(databaseConnection)

	authMiddleware := framework.NewAuthMiddleware(usecase.NewAuthUsecase(jwtKey, userRepository))
	expensesGroup := router.Group("/expenses", authMiddleware.Authorization())

	framework.CreateUserRoutes(usecase.NewUserUseCase(userRepository, jwtKey), router)
	framework.CreateExpenseRoutes(
		usecase.NewExpenseUseCase(adapter.NewExpenseRepository(databaseConnection),
			userRepository,
			jwtKey),
		expensesGroup)
	router.Run()
}
