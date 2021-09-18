package unitTest

import (
	controller2 "W/controller"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	//"os"
	"testing"
)

var controller = controller2.Controller{}

// routing testing and view response
func TestRoutesAndResponse(t *testing.T) {
	t.Run("search route", func(t *testing.T) {
		srv := httptest.NewServer(controller.Server.Router)
		defer srv.Close()
		res, err := http.Get(fmt.Sprintf("%s/api/search", srv.URL))
		if err != nil {
			t.Fatal(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		bodyString := string(body)
		println(bodyString)
		AssertEquals(http.StatusOK, res.StatusCode, t)
	})
	t.Run("show customer route", func(t *testing.T) {
		customer, _ := InitRandomCustomer()
		insertedCustomer, _ := repo.InsertCustomer(&customer)

		req, _ := http.NewRequest("GET", fmt.Sprintf("/api/id=%v", insertedCustomer.ID), nil)
		responseRecorder := httptest.NewRecorder()
		controller.Server.Router.ServeHTTP(responseRecorder, req)

		s := responseRecorder.Body.String()

		AssertContains(insertedCustomer.FirstName, s, t)
		AssertEquals(http.StatusOK, responseRecorder.Code, t)
	})

}

//func TestRoutingShowCustomer(t *testing.T) {
//	customer, _ := InitRandomCustomer()
//	insertedCustomer, _ := repo.InsertCustomer(&customer)
//
//	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/id=%v", insertedCustomer.ID), nil)
//	responseRecorder := httptest.NewRecorder()
//	controller.Server.Router.ServeHTTP(responseRecorder, req)
//
//	s := responseRecorder.Body.String()
//
//	println(s)
//	AssertContains(insertedCustomer.FirstName, s, t)
//	AssertEquals(http.StatusOK, responseRecorder.Code, t)
//
//}
