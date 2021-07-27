package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"ken.com.tw/demo/internal/router"
)


func TestUserSignin(t *testing.T) {
	r := router.Initial()

	bodyReader := strings.NewReader(`{"username": "test", "password": "abcd1234"}`)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/v1/users/signin", bodyReader)

	r.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
}