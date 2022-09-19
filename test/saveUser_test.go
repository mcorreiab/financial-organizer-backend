//go:build integration

package integration

import (
	"database/sql"
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"
	"testing"

	_ "github.com/lib/pq"
)

const username = "username"
const password = "password"

func TestSaveNewUser(t *testing.T) {
	tests := []testRunner{
		{"SaveUser", testSaveUser},
		{"UsernameMissing", testReturnBadRequestWhenUsernameIsMissing},
		{"PasswordMissing", testReturnBadRequestWhenPasswordIsMissing},
		{"InsertExistentUser", testTryToInsertUserThatAlreadyExists},
	}
	newSuite(t, tests).run()
}

func testSaveUser(t *testing.T) {
	executeSaveUserIntegrationTest(t, username, password).checkStatusCode(http.StatusCreated)
	checkIfUserGotInserted(t, databaseConnection)
}

func checkIfUserGotInserted(t *testing.T, databaseConnection *sql.DB) {
	repository := adapter.NewUserRepository(databaseConnection)
	user, err := repository.FindUserByUsername(username)

	if err != nil {
		t.Fatal("Failed to query database: ", err)
	}

	if user == nil {
		t.Errorf("Could not find the user on database")
	}
}

func testReturnBadRequestWhenUsernameIsMissing(t *testing.T) {
	username := ""
	executeSaveUserIntegrationTest(t, username, password).checkStatusCode(http.StatusBadRequest)
	checkThatUserIsNotOnDatabase(t, username)
}

func testReturnBadRequestWhenPasswordIsMissing(t *testing.T) {
	executeSaveUserIntegrationTest(t, username, "").checkStatusCode(http.StatusBadRequest)
	checkThatUserIsNotOnDatabase(t, username)
}

func checkThatUserIsNotOnDatabase(t *testing.T, username string) {
	repository := adapter.NewUserRepository(databaseConnection)
	user, err := repository.FindUserByUsername(username)

	if err != nil {
		t.Fatal("Failed to query database: ", err)
	}

	if user != nil {
		t.Errorf("User shouldn't be on database")
	}
}

func testTryToInsertUserThatAlreadyExists(t *testing.T) {
	executeSaveUserIntegrationTest(t, username, password).checkStatusCode(http.StatusCreated)
	executeSaveUserIntegrationTest(t, username, password).checkStatusCode(http.StatusConflict)
}

func executeSaveUserIntegrationTest(t *testing.T, username string, password string) *apiTest {
	return newApiTest(t, usecase.NewUserUseCase(adapter.NewUserRepository(databaseConnection))).
		setRequest(http.MethodPost, "/users", createBody(username, password)).
		execute()
}

func createBody(username string, password string) framework.UserPayload {
	return framework.UserPayload{Username: username, Password: password}
}
