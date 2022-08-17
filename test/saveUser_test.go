//go:build integration

package integration

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const username = "username"
const password = "password"

var databaseConnection *sql.DB

func TestSaveNewUser(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatal(err)
	}

	databaseConnection = framework.GetDatabaseConnection()

	cleanDatabase(t)
	t.Run("SaveUser", testSaveUser)
	cleanDatabase(t)
	t.Run("UsernameMissing", testReturnBadRequestWhenUsernameIsMissing)
	cleanDatabase(t)
	t.Run("PasswordMissing", testReturnBadRequestWhenPasswordIsMissing)
	cleanDatabase(t)
	t.Run("InsertExistentUser", testTryToInsertUserThatAlreadyExists)
}

func cleanDatabase(t *testing.T) {
	_, err := databaseConnection.Exec("DELETE from users")
	if err != nil {
		t.Fatal(err)
	}
}

func testSaveUser(t *testing.T) {
	createSaveUserIntegrationTest(t, username, password).Execute().CheckStatusCode(http.StatusCreated)
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
	createSaveUserIntegrationTest(t, username, password).Execute().CheckStatusCode(http.StatusBadRequest)
	checkThatUserIsNotOnDatabase(t, username)
}

func testReturnBadRequestWhenPasswordIsMissing(t *testing.T) {
	createSaveUserIntegrationTest(t, username, "").Execute().CheckStatusCode(http.StatusBadRequest)
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
	createSaveUserIntegrationTest(t, username, password).Execute().CheckStatusCode(http.StatusCreated)
	createSaveUserIntegrationTest(t, username, password).Execute().CheckStatusCode(http.StatusConflict)
}

func createSaveUserIntegrationTest(t *testing.T, username string, password string) *saveUserIntegrationTest {
	return &saveUserIntegrationTest{
		createSaveUserRoute(databaseConnection),
		createRequest(t, username, password),
		t,
		nil,
	}
}

func createSaveUserRoute(databaseConnection *sql.DB) *gin.Engine {
	uc := usecase.NewSaveUserUseCase(adapter.NewUserRepository(databaseConnection))
	saveUserRoute := framework.NewSaveUserController(uc)
	router := gin.Default()
	router.POST("/users", saveUserRoute.Save)

	return router
}

func createRequest(t *testing.T, username string, password string) *http.Request {
	p, err := json.Marshal(framework.UserPayload{Username: username, Password: password})
	if err != nil {
		t.Error("Error creating the payload", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(p))
	if err != nil {
		t.Fatal("Failed to create request: ", err)
	}

	return req
}

type saveUserIntegrationTest struct {
	router   *gin.Engine
	request  *http.Request
	t        *testing.T
	response *httptest.ResponseRecorder
}

func (s *saveUserIntegrationTest) Execute() *saveUserIntegrationTest {
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, s.request)

	s.response = w
	return s
}

func (s *saveUserIntegrationTest) CheckStatusCode(expected int) *saveUserIntegrationTest {
	if s.response.Code != expected {
		s.t.Errorf("Expected status code %d, got %d", expected, s.response.Code)
	}
	return s
}
