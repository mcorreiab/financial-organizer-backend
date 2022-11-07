//go:build integration

package integration

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

const username = "username"
const password = "password"

type SaveUserSuite struct {
	suite.Suite
	routerBuilder      *routerBuilder
	databaseConnection *sql.DB
}

func TestSaveUser(t *testing.T) {
	suite.Run(t, new(SaveUserSuite))
}

func (suite *SaveUserSuite) SetupSuite() {
	suite.databaseConnection = initLocalDatabase(suite.T())

	suite.routerBuilder = newRouterBuilder(suite.databaseConnection, "mockKey").BuildUserRoutes().BuildExpensesRoutes()
}

func (suite *SaveUserSuite) TearDownTest() {
	_, err := suite.databaseConnection.Exec("DELETE from users")
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *SaveUserSuite) TestSaveUser() {
	executeSaveUserIntegrationTest(suite, username, password).checkStatusCode(http.StatusCreated)
	checkIfUserGotInserted(suite.T(), suite.databaseConnection)
}

func checkIfUserGotInserted(t *testing.T, databaseConnection *sql.DB) {
	repository := adapter.NewUserRepository(databaseConnection)
	user, err := repository.FindUserByUsername(username)

	if err != nil {
		t.Fatal("Failed to query database: ", err)
	}

	require.NotNil(t, user)
}

func (suite *SaveUserSuite) TestReturnBadRequestWhenUsernameIsMissing() {
	username := ""
	executeSaveUserIntegrationTest(suite, username, password).checkStatusCode(http.StatusBadRequest)
	suite.checkThatUserIsNotOnDatabase(username)
}

func (suite *SaveUserSuite) TestReturnBadRequestWhenPasswordIsMissing() {
	executeSaveUserIntegrationTest(suite, username, "").checkStatusCode(http.StatusBadRequest)
	suite.checkThatUserIsNotOnDatabase(username)
}

func (suite *SaveUserSuite) checkThatUserIsNotOnDatabase(username string) {
	repository := adapter.NewUserRepository(suite.databaseConnection)
	user, err := repository.FindUserByUsername(username)

	if err != nil {
		suite.T().Fatal("Failed to query database: ", err)
	}

	require.Nil(suite.T(), user)
}

func (suite *SaveUserSuite) TestTryToInsertUserThatAlreadyExists() {
	executeSaveUserIntegrationTest(suite, username, password).checkStatusCode(http.StatusCreated)
	executeSaveUserIntegrationTest(suite, username, password).checkStatusCode(http.StatusConflict)
}

func executeSaveUserIntegrationTest(suite *SaveUserSuite, username string, password string) *apiRequest {
	userPayload := createBody(username, password)
	return newApiRequest(suite.T(), suite.routerBuilder).setRequest(http.MethodPost, "/users", userPayload).execute()
}

func createBody(username string, password string) framework.UserPayload {
	return framework.UserPayload{Username: username, Password: password}
}
