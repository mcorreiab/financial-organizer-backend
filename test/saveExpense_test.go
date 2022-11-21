//go:build integration

package integration

import (
	"encoding/json"
	"fmt"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"testing"

	"github.com/stretchr/testify/suite"
)

type authResponse struct {
	AccessToken string `json:"access_token"`
}

type SaveExpenseTest struct {
	ConcreteTestSuite
}

var expensePayload framework.ExpensePayload = framework.ExpensePayload{Name: "expenseName", Value: 13.50}

func TestSaveExpense(t *testing.T) {
	testSuite := &SaveExpenseTest{}
	testSuite.Init()
	suite.Run(t, testSuite)
}

func (suite *SaveExpenseTest) SetupSuite() {
	userPayload := framework.UserPayload{Username: "username", Password: "pass"}
	ctx := apiRequestContext{suite: suite, path: "/users", body: userPayload}
	createApiRequest(ctx).execute().checkStatusCode(201)

	ctx.path = "/signin"

	signInRequest := createApiRequest(ctx).execute().checkStatusCode(201)

	var authResponse authResponse
	json.Unmarshal(signInRequest.response.Body.Bytes(), &authResponse)

	suite.jwtToken = authResponse.AccessToken
}

func (suite *SaveExpenseTest) TearDownTest() {
	_, err := suite.databaseConnection.Exec("DELETE from expenses")
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *SaveExpenseTest) TestSaveExpenseWithSuccess() {
	apiRequest := suite.creteRequestSaveExpense(expensePayload)

	apiRequest.execute().checkStatusCode(201)
}

func (suite *SaveExpenseTest) TestTryToSaveExpenseEmptyName() {
	payload := framework.ExpensePayload{Name: "", Value: 5.50}
	apiRequest := suite.creteRequestSaveExpense(payload)

	apiRequest.execute().checkStatusCode(400)
}

func (suite *SaveExpenseTest) TestTryToSaveExpenseInvalidValue() {
	payload := framework.ExpensePayload{Name: "expense", Value: 0}
	apiRequest := suite.creteRequestSaveExpense(payload)

	apiRequest.execute().checkStatusCode(400)
}

func (suite *SaveExpenseTest) creteRequestSaveExpense(payload framework.ExpensePayload) *apiRequest {
	ctx := apiRequestContext{suite: suite, path: "/expenses", body: payload}
	apiRequest := createApiRequest(ctx)

	apiRequest.request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", suite.jwtToken))
	return apiRequest
}
