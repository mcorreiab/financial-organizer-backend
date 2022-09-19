//go:build integration

package integration

import (
	"mcorreiab/financial-organizer-backend/internal/adapter"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"net/http"
	"testing"
)

func TestSignIn(t *testing.T) {
	tests := []testRunner{
		{"SignInSuccess", testSignInSuccess},
	}
	newSuite(t, tests).run()
}

func testSignInSuccess(t *testing.T) {
	usecase := usecase.NewUserUseCase(adapter.NewUserRepository(databaseConnection))
	userPayload := framework.UserPayload{Username: "username", Password: "password"}

	createUserOnDatabase(usecase, userPayload, t)

	newApiTest(t, usecase).
		setRequest(http.MethodPost, "/signin", userPayload).
		execute().
		checkStatusCode(201)
}

func createUserOnDatabase(uc usecase.UserUseCase, userPayload framework.UserPayload, t *testing.T) {
	_, err := uc.SaveUser(userPayload.Username, userPayload.Password)

	if err != nil {
		t.Fatal("Error creating user on database", err)
	}
}
