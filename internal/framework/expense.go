package framework

import (
	"math/big"
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	expenseUc usecase.ExpenseUseCase
}

func NewExpenseRouter(usecase usecase.ExpenseUseCase, router *gin.Engine) {
	c := ExpenseController{usecase}
	router.POST("/expenses", c.Save)
}

type ExpensePayload struct {
	Name  string    `json:"name"`
	Value big.Float `json:"value"`
}

func (ec ExpenseController) Save(c *gin.Context) {
	var payload ExpensePayload
	err := c.ShouldBindJSON(payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(InvalidPayload, err.Error()))
		return
	}

	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		c.JSON(http.StatusBadRequest, NewError(InvalidPayload, "Authorization header is missing"))
		return
	}

	token := strings.TrimPrefix(auth, "Bearer ")

	if token == "" {
		c.JSON(http.StatusBadRequest, NewError(InvalidPayload, "Authorization header is malformed"))
	}

	_, err = ec.expenseUc.SaveExpense(payload.Name, payload.Value, token)

	if err != nil {
		if _, ok := err.(usecase.InvalidToken); ok {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusInternalServerError, NewError(UnexpectedError, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, nil)
}
