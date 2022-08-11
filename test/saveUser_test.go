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

func TestSaveUser(t *testing.T) {
	db := getDatabaseConnection(t)

	requestToSaveUser(t, db)

	checkIfUserGotInserted(t, db)
}

func getDatabaseConnection(t *testing.T) *sql.DB {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}

	return framework.GetDatabaseConnection()
}

func requestToSaveUser(t *testing.T, databasConnection *sql.DB) {
	uc := usecase.NewSaveUserUseCase(adapter.NewUserRepository(databasConnection))
	saveUserRoute := framework.NewSaveUserController(uc)

	router := gin.Default()
	router.POST("/users", saveUserRoute.Save)

	req := createRequest(t)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func createRequest(t *testing.T) *http.Request {
	p, err := json.Marshal(framework.UserPayload{Username: username, Password: "password"})
	if err != nil {
		t.Error("Error creating the payload", err)
	}

	req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(p))
	if err != nil {
		t.Fatal("Failed to create request: ", err)
	}

	return req
}

func checkIfUserGotInserted(t *testing.T, databaseConnection *sql.DB) {
	var u string
	err := databaseConnection.
		QueryRow("SELECT username from users where username = $1", username).
		Scan(&u)

	if err != nil {
		t.Fatal("Failed to query database: ", err)
	}

	if u != username {
		t.Errorf("Expected username %s, got %s", username, u)
	}
}
