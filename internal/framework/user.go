package framework

import (
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	userUsecase usecase.UserUseCase
}

func CreateUserRoutes(userUsecase usecase.UserUseCase, router *gin.Engine) {
	userController := User{userUsecase}

	router.POST("/users", userController.Save)
	router.POST("/signin", userController.GenerateToken)
}

type UserPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *User) Save(c *gin.Context) {
	var payload UserPayload
	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(InvalidPayload, err.Error()))
		return
	}

	_, err = s.userUsecase.SaveUser(payload.Username, payload.Password)

	if _, ok := err.(usecase.UserExistsError); ok {
		c.JSON(http.StatusConflict, NewError(UserExistsError, err.Error()))
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, NewError(UnexpectedError, err.Error()))
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
			c.JSON(http.StatusForbidden, NewError(AuthenticationError, err.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, NewError(UnexpectedError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"access_token": token.AccessToken, "expires_in": token.ExpiresIn})
}
