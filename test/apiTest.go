package integration

import (
	"encoding/json"
	"mcorreiab/financial-organizer-backend/internal/framework"
	"mcorreiab/financial-organizer-backend/internal/usecase"
	"testing"

	"net/http"
	"net/http/httptest"

	"bytes"

	"github.com/gin-gonic/gin"
)

type apiTest struct {
	router   *gin.Engine
	request  *http.Request
	t        *testing.T
	response *httptest.ResponseRecorder
}

func newApiTest(t *testing.T, usecase usecase.UserUseCase) *apiTest {
	return &apiTest{framework.NewUserRouter(usecase), nil, t, nil}
}

func (test *apiTest) setRequest(method string, path string, body any) *apiTest {
	p, err := json.Marshal(body)

	if err != nil {
		test.t.Fatal("Failed to marshal payload: ", err)
	}

	req, err := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(p))
	if err != nil {
		test.t.Fatal("Failed to create request: ", err)
	}

	test.request = req
	return test
}

func (test *apiTest) execute() *apiTest {
	if test.request == nil {
		panic("No request was defined")
	}

	w := httptest.NewRecorder()
	test.router.ServeHTTP(w, test.request)

	test.response = w
	return test
}

func (test *apiTest) checkStatusCode(expected int) *apiTest {
	if test.response.Code != expected {
		test.t.Errorf("Expected status code %d, got %d", expected, test.response.Code)
	}
	return test
}
