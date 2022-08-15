package framework

import (
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SaveUser struct {
	saveUserUsecase usecase.SaveUserUseCase
}

func NewSaveUserController(saveUserUsecase usecase.SaveUserUseCase) SaveUser {
	return SaveUser{saveUserUsecase}
}

type UserPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *SaveUser) Save(c *gin.Context) {
	var payload UserPayload
	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err = s.saveUserUsecase.SaveUser(payload.Username, payload.Password)

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
