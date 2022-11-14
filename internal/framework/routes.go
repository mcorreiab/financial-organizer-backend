package framework

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func CreateRoutes(jwtKey string, databaseConnection *sql.DB) *gin.Engine {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notEmpty", notEmpty)
	}

	userRepository := adapter.NewUserRepository(databaseConnection)

	authMiddleware := NewAuthMiddleware(usecase.NewAuthUsecase(jwtKey, userRepository))

	CreateUserRoutes(usecase.NewUserUseCase(userRepository, jwtKey), router)
	CreateExpenseRoutes(
		usecase.NewExpenseUseCase(adapter.NewExpenseRepository(databaseConnection),
			userRepository,
			jwtKey),
		router,
		authMiddleware)

	return router
}

var notEmpty validator.Func = func(fl validator.FieldLevel) bool {
	return strings.TrimSpace(fl.Field().String()) != ""
}
