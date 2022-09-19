package framework

import (
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	userUsecase usecase.UserUseCase
}

func NewUserRouter(userUsecase usecase.UserUseCase) *gin.Engine {
	router := gin.Default()

	userController := User{userUsecase}

	router.POST("/users", userController.Save)
	router.POST("/signin", userController.GenerateToken)
	return router
}

type UserPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *User) Save(c *gin.Context) {
	var payload UserPayload
	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err = s.userUsecase.SaveUser(payload.Username, payload.Password)

	if _, ok := err.(usecase.UserExistsError); ok {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (s *User) GenerateToken(c *gin.Context) {
	var payload UserPayload
	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := s.userUsecase.GenerateLoginToken(payload.Username, payload.Password)

	if err != nil {
		if _, ok := err.(usecase.InvalidCredentialsError); ok {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}
