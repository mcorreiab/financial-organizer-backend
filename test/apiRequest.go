package integration

import (
	"encoding/json"
	"testing"

	"net/http"
	"net/http/httptest"

	"bytes"

	"github.com/gin-gonic/gin"
)

type apiRequest struct {
	router   *gin.Engine
	request  *http.Request
	t        *testing.T
	response *httptest.ResponseRecorder
}

type apiRequestContext struct {
	suite TestSuite
	path  string
	body  any
}

func createApiRequest(context apiRequestContext) *apiRequest {
	p, err := json.Marshal(context.body)

	if err != nil {
		context.suite.T().Fatal("Failed to marshal payload: ", err)
	}

	req, err := http.NewRequest(http.MethodPost, context.path, bytes.NewBuffer(p))
	if err != nil {
		context.suite.T().Fatal("Failed to create request: ", err)
	}

	return &apiRequest{context.suite.GetRouter(), req, context.suite.T(), nil}
}

func (test *apiRequest) execute() *apiRequest {
	if test.request == nil {
		panic("No request was defined")
	}

	w := httptest.NewRecorder()
	test.router.ServeHTTP(w, test.request)

	test.response = w
	return test
}

func (test *apiRequest) checkStatusCode(expected int) *apiRequest {
	if test.response.Code != expected {
		test.t.Errorf("Expected status code %d, got %d", expected, test.response.Code)
	}
	return test
}
