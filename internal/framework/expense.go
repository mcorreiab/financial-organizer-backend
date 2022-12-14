package framework

import (
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	expenseUc *usecase.ExpenseUseCase
}

func CreateExpenseRoutes(usecase *usecase.ExpenseUseCase, router *gin.Engine, authMiddleware AuthMiddleware) {
	expensesGroup := router.Group("/expenses", authMiddleware.Authorization())
	c := ExpenseController{usecase}
	expensesGroup.POST("", c.Save)
}

type ExpensePayload struct {
	Name  string  `json:"name" binding:"notEmpty"`
	Value float64 `json:"value" binding:"gt=0"`
}

func (ec ExpenseController) Save(c *gin.Context) {
	var payload ExpensePayload
	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(InvalidPayload, err.Error()))
		return
	}

	if userId, ok := c.MustGet("userId").(string); ok {
		_, err = ec.expenseUc.SaveExpense(payload.Name, payload.Value, userId)

		if err != nil {
			c.JSON(http.StatusInternalServerError, NewError(UnexpectedError, err.Error()))
			return
		}

		c.JSON(http.StatusCreated, nil)
		return
	}

	panic("userId should be a string")
}
