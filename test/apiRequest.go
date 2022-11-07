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

func newApiRequest(t *testing.T, routerBuilder *routerBuilder) *apiRequest {
	return &apiRequest{routerBuilder.Build(), nil, t, nil}
}

func (test *apiRequest) setRequest(method string, path string, body any) *apiRequest {
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
