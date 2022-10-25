package app

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/y-sugiyama654/banking/dto"
	"github.com/y-sugiyama654/banking/errs"
	"github.com/y-sugiyama654/banking/mocks/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var ch CustomerHandlers
var mockService *service.MockCustomerService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockCustomerService(ctrl)
	ch = CustomerHandlers{mockService}
	router = mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomer)
	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	dummyCustomers := []dto.CustomerResponse{
		{"1001", "Yuta Sugiyama", "Tokyo", "1000001", "1994-05-04", "1"},
		{"1002", "Vinton Gray Cerf", "NewYork", "1000002", "1943-06-23", "1"},
		{"1003", "Robert Elliot Kahn", "Los Angeles", "1000003", "1938-12-23", "1"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Failed while testing the status code.")
	}
}

func Test_should_return_status_code_500_with_error_message(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewUnexpectedError("some database error."))
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Failed while testing the status code.")
	}
}
