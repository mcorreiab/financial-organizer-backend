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
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *SaveUser) Save(c *gin.Context) {
	var payload UserPayload
	c.BindJSON(&payload)

	_, err := s.saveUserUsecase.SaveUser(payload.Username, payload.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, nil)
	}
}
