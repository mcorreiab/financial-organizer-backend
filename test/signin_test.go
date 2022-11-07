//go:build integration

package integration

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SignInSuite struct {
	suite.Suite
	routerBuilder      *routerBuilder
	databaseConnection *sql.DB
}

var userOnDB = framework.UserPayload{Username: "username", Password: "password"}

func TestSignin(t *testing.T) {
	suite.Run(t, new(SignInSuite))
}

func (suite *SignInSuite) SetupSuite() {
	suite.databaseConnection = initLocalDatabase(suite.T())
	suite.cleanUpDatabase()

	suite.routerBuilder = newRouterBuilder(suite.databaseConnection, "mockKey").BuildUserRoutes().BuildExpensesRoutes()
	suite.createUser(userOnDB)
}

func (suite *SignInSuite) cleanUpDatabase() {
	_, err := suite.databaseConnection.Exec("DELETE from users")
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *SignInSuite) createUser(userPayload framework.UserPayload) {
	newApiRequest(suite.T(), suite.routerBuilder).
		setRequest(http.MethodPost, "/users", userPayload).
		execute().checkStatusCode(201)
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

	suite.executeCallToApi(userOnSignin).checkStatusCode(403)
}

func (suite *SignInSuite) testTryToSignInUserWrongCredentials() {
	userOnSignin := framework.UserPayload{Username: userOnDB.Username, Password: "password2"}

	suite.executeCallToApi(userOnSignin).checkStatusCode(403)
}

func (suite *SignInSuite) executeCallToApi(payload framework.UserPayload) *apiRequest {
	return newApiRequest(suite.T(), suite.routerBuilder).
		setRequest(http.MethodPost, "/signin", payload).
		execute()
}
