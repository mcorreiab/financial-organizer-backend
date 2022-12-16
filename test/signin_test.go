//go:build integration

package integration

import (
	"mcorreiab/financial-organizer-backend/internal/framework"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SignInSuite struct {
	ConcreteTestSuite
}

var userOnDB = framework.UserPayload{Username: "username", Password: "password"}

func TestSignin(t *testing.T) {
	testSuite := &SignInSuite{}
	testSuite.Init()
	suite.Run(t, testSuite)
}

func (suite *SignInSuite) SetupSuite() {
	suite.createUser(userOnDB)
}

func (suite *SignInSuite) createUser(userPayload framework.UserPayload) {
	ctx := apiRequestContext{suite: suite, path: "/users", body: userPayload}
	createApiRequest(ctx).execute().checkStatusCode(201)
}

func (suite *SignInSuite) TearDownTest() {
	_, err := suite.databaseConnection.Exec("DELETE from users where username != $1", userOnDB.Username)
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *SignInSuite) TestSignInSuccess() {
	suite.executeCallToApi(userOnDB).checkStatusCode(201)
}

func (suite *SignInSuite) TestSignInWithMissingCredentials() {
	suite.executeCallToApi(framework.UserPayload{Username: "", Password: ""}).
		checkStatusCode(400)
}

func (suite *SignInSuite) TestTryToSignInUserDontExists() {
	userOnSignin := framework.UserPayload{Username: "username2", Password: "password"}

	suite.executeCallToApi(userOnSignin).checkStatusCode(401)
}

func (suite *SignInSuite) testTryToSignInUserWrongCredentials() {
	userOnSignin := framework.UserPayload{Username: userOnDB.Username, Password: "password2"}

	suite.executeCallToApi(userOnSignin).checkStatusCode(401)
}

func (suite *SignInSuite) executeCallToApi(payload framework.UserPayload) *apiRequest {
	ctx := apiRequestContext{suite: suite, path: "/signin", body: payload}
	return createApiRequest(ctx).execute()
}
