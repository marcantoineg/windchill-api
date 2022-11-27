package tests

import (
	"gin-api-template/api"
	"gin-api-template/responses"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingPong(t *testing.T) {
	expectedResponse := responses.OK("pong")

	r := api.SetupRouter()

	req, _ := http.NewRequest("GET", "/ping", nil)
	recorder := httptest.NewRecorder()
	r.ServeHTTP(recorder, req)

	responseData, _ := ioutil.ReadAll(recorder.Body)
	assert.Equal(t, expectedResponse.ToString(), string(responseData))
	assert.Equal(t, expectedResponse.Status, recorder.Code)
}
