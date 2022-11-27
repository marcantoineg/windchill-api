package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"gin-api-template/api"
	"gin-api-template/models"
	"gin-api-template/responses"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/marcantoineg/windchill/windchill/speed"
	temp "github.com/marcantoineg/windchill/windchill/temperature"
	"github.com/stretchr/testify/assert"
)

type testrun struct {
	given    models.WindchillRequest
	expected responses.Response
}

func TestPOSTWindchill(t *testing.T) {
	testruns := []testrun{
		{
			given: models.WindchillRequest{
				Temperature: temp.F(40.),
				WindSpeed:   speed.Mph(10.),
			},
			expected: *responses.OK(temp.F(33.613135223155105)),
		},
		{
			given: models.WindchillRequest{
				Temperature: temp.C(10.),
				WindSpeed:   speed.Kph(10.),
			},
			expected: *responses.OK(temp.C(8.63151849762641)),
		},
		{
			given: models.WindchillRequest{
				Temperature: temp.Temperature{Unit: "not a unit"},
				WindSpeed:   speed.Kph(10.),
			},
			expected: *responses.BadRequest(errors.New("could not parse Temperature Unit: 'not a unit'")),
		},
		{
			given: models.WindchillRequest{
				Temperature: temp.C(10.),
				WindSpeed:   speed.Speed{Unit: "not a unit"},
			},
			expected: *responses.BadRequest(errors.New("could not parse Speed Unit: 'not a unit'")),
		},
	}

	r := api.SetupRouter()

	for _, tr := range testruns {
		body, err := json.Marshal(tr.given)
		if err != nil {
			t.Fatal(err)
		}
		req, _ := http.NewRequest("POST", "/windchill", bytes.NewReader(body))
		recorder := httptest.NewRecorder()
		r.ServeHTTP(recorder, req)

		responseData, _ := ioutil.ReadAll(recorder.Body)
		assert.Equal(t, tr.expected.ToString(), string(responseData))
		assert.Equal(t, tr.expected.Status, recorder.Code)
	}
}
