package unitTest

import (
	s "W/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server = s.Server{}

func TestEmptyTable(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/customer", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}

}
func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	//server.InitializeRouter()
	//server.StartServer()
	server.Router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
