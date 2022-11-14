package framework

import (
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthMiddleware(authUsecase *usecase.AuthUsecase) AuthMiddleware {
	return AuthMiddleware{authUsecase}
}

func (m AuthMiddleware) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewError(InvalidPayload, "Authorization header is missing"))
			return
		}

		if !strings.HasPrefix(auth, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewError(InvalidPayload, "Authorization header is malformed"))
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(auth, "Bearer "))

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewError(InvalidPayload, "Authorization header is malformed"))
			return
		}

		userId, err := m.authUsecase.ValidateToken(token)
		if err != nil {
			if _, ok := err.(usecase.InvalidToken); ok {
				c.AbortWithStatusJSON(http.StatusNotFound, nil)
				return
			}

			c.AbortWithStatusJSON(http.StatusInternalServerError, NewError(UnexpectedError, err.Error()))
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
