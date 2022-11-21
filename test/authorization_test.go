package integration

import (
	"fmt"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AuthorizationTest struct {
	ConcreteTestSuite
}

func TestAuthorization(t *testing.T) {
	testSuite := AuthorizationTest{}
	testSuite.Init()
	suite.Run(t, &testSuite)
}

func (suite *AuthorizationTest) TestRequestWithoutAuthHeader() {
	suite.createRequestToSaveExpense().execute().checkStatusCode(401)
}

func (suite *AuthorizationTest) TestRequestWithWrongAuthHeader() {
	apiRequest := suite.createRequestToSaveExpense()
	apiRequest.request.Header.Set("Authorization", "Bears Beet Battlestar Galactica")
	apiRequest.execute().checkStatusCode(401)
}

func (suite *AuthorizationTest) TestRequestWithIncompleteAuthHeader() {
	apiRequest := suite.createRequestToSaveExpense()
	apiRequest.request.Header.Set("Authorization", "Bearer      ")
	apiRequest.execute().checkStatusCode(401)
}

func (suite *AuthorizationTest) TestRequestWithInvalidToken() {
	apiRequest := suite.createRequestToSaveExpense()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	apiRequest.request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	apiRequest.execute().checkStatusCode(401)
}

func (suite *AuthorizationTest) createRequestToSaveExpense() *apiRequest {
	ctx := apiRequestContext{suite: suite, path: "/expenses", body: framework.ExpensePayload{Name: "name", Value: 5}}
	return createApiRequest(ctx)
}
