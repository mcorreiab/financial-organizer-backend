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
		{"TryToSignInUserDontExists", testTryToSignInUserDontExists},
		{"TryToSignInUserWrongCredentials", testTryToSignInUserWrongCredentials},
	}
	newSuite(t, tests).run()
}

func testSignInSuccess(t *testing.T) {
	usecase := createUseCase()
	userPayload := framework.UserPayload{Username: "username", Password: "password"}

	createUserOnDatabase(usecase, userPayload, t)

	executeCallToApi(t, usecase, userPayload).checkStatusCode(201)
}

func testTryToSignInUserDontExists(t *testing.T) {
	usecase := createUseCase()
	userOnDatabase := framework.UserPayload{Username: "username", Password: "password"}

	createUserOnDatabase(usecase, userOnDatabase, t)

	userOnSignin := framework.UserPayload{Username: "username2", Password: "password"}

	executeCallToApi(t, usecase, userOnSignin).checkStatusCode(403)
}

func testTryToSignInUserWrongCredentials(t *testing.T) {
	usecase := createUseCase()
	userOnDatabase := framework.UserPayload{Username: "username", Password: "password"}

	createUserOnDatabase(usecase, userOnDatabase, t)

	userOnSignin := framework.UserPayload{Username: userOnDatabase.Username, Password: "password2"}

	executeCallToApi(t, usecase, userOnSignin).checkStatusCode(403)
}

func createUseCase() usecase.UserUseCase {
	return usecase.NewUserUseCase(adapter.NewUserRepository(databaseConnection))
}

func createUserOnDatabase(uc usecase.UserUseCase, userPayload framework.UserPayload, t *testing.T) {
	_, err := uc.SaveUser(userPayload.Username, userPayload.Password)

	if err != nil {
		t.Fatal("Error creating user on database", err)
	}
}

func executeCallToApi(t *testing.T, usecase usecase.UserUseCase, payload framework.UserPayload) *apiTest {
	return newApiTest(t, usecase).
		setRequest(http.MethodPost, "/signin", payload).
		execute()
}
