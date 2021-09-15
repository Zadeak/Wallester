package unitTest

import (
	controller "W/controller"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var con = controller.Controller{}

func TestTest(t *testing.T) {
	//a.Initialize()
	con.InitController()
	con.InitiateRoutes()

	req, _ := http.NewRequest("GET", "/api/search", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, req)
	assert.Equal(t, 200, response.Code, "OK response is expected")

}

//
//func executeRequest(req *http.Request) *httptest.ResponseRecorder {
//	rr := httptest.NewRecorder()
//	a.Controller.Server.Router.ServeHTTP(rr, req)
//	return rr
//}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/create", CreateEndpoint).Methods("GET")
	return router
}
func CreateEndpoint(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Item Created"))
}
