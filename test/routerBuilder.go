package integration

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"mcorreiab/financial-organizer-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type routerBuilder struct {
	router             *gin.Engine
	databaseConnection *sql.DB
	jwtKey             string
}

func newRouterBuilder(databaseConnection *sql.DB, jwtKey string) *routerBuilder {
	return &routerBuilder{gin.Default(), databaseConnection, jwtKey}
}

func (builder *routerBuilder) BuildExpensesRoutes() *routerBuilder {
	userRepository := adapter.NewUserRepository(builder.databaseConnection)
	authMiddleware := framework.NewAuthMiddleware(usecase.NewAuthUsecase(builder.jwtKey, userRepository))
	framework.CreateExpenseRoutes(
		usecase.NewExpenseUseCase(adapter.NewExpenseRepository(builder.databaseConnection),
			userRepository,
			builder.jwtKey),
		builder.router,
		authMiddleware)

	return builder
}

func (builder *routerBuilder) BuildUserRoutes() *routerBuilder {
	userRepository := adapter.NewUserRepository(builder.databaseConnection)
	framework.CreateUserRoutes(usecase.NewUserUseCase(userRepository, builder.jwtKey), builder.router)

	return builder
}

func (builder *routerBuilder) Build() *gin.Engine {
	return builder.router
}
