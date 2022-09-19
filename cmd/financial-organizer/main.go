package main

import (
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"mcorreiab/financial-organizer-backend/internal/usecase"

	_ "github.com/lib/pq"
)

func main() {
	router := framework.NewUserRouter(
		usecase.NewUserUseCase(adapter.NewUserRepository(framework.GetDatabaseConnection())),
	)

	router.Run()
}
