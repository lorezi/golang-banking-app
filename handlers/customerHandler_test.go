package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/lorezi/golang-bank-app/dto"
	"github.com/lorezi/golang-bank-app/errs"
	"github.com/lorezi/golang-bank-app/mocks"
)

var r *mux.Router
var ch CustomerHandler
var mck *mocks.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	mck = mocks.NewMockCustomerService(ctrl)
	ch = CustomerHandler{CustomerService: mck}
	r = mux.NewRouter()

	r.HandleFunc("/customers", ch.GetAllCustomers)

	return func() {
		r = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {

	// Arrange ---> set up mock services
	tearDown := setup(t)
	defer tearDown()

	dummyCustomers := []dto.CustomerResponse{
		{
			Name: "John Doe", City: "New York", Zipcode: "1100034", DateofBirth: "2000-01-04", Status: "true", Id: "100001",
		},
		{
			Name: "Jane Doe", City: "Silicon Valley", Zipcode: "8900001", DateofBirth: "2000-01-08", Status: "false", Id: "100034",
		},
	}
	mck.EXPECT().GetAllCustomers("").Return(dummyCustomers, nil)
	req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	// Assert
	if rec.Code != http.StatusOK {
		t.Error("Failed while testing the status code")
	}

}

func Test_should_return_customers_with_status_code_500_with_error_message(t *testing.T) {

	// Arrange ---> set up mock services
	tearDown := setup(t)
	defer tearDown()

	mck.EXPECT().GetAllCustomers("").Return(nil, errs.UnExpectedError("some database error", "fail"))

	r.HandleFunc("/customers", ch.GetAllCustomers)
	req, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	// Assert
	if rec.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code")
	}
}
