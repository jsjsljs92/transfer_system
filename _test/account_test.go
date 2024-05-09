package transfer_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	account "github.com/jsjsljs92/transferSystem/src/components/account"
	"github.com/jsjsljs92/transferSystem/src/models"
)

type MockAccountService struct {
	ValidateCreateAccountReqFn func(req account.CreateAccountRequest) error
	CreateAccountFn            func(req account.CreateAccountRequest) error
	GetAccountByIDFn           func(id int) (*models.Account, error)
}

func (m *MockAccountService) ValidateCreateAccountReq(req account.CreateAccountRequest) error {
	if m.ValidateCreateAccountReqFn != nil {
		return m.ValidateCreateAccountReqFn(req)
	}
	return nil
}

func (m *MockAccountService) CreateAccount(req account.CreateAccountRequest) error {
	if m.CreateAccountFn != nil {
		return m.CreateAccountFn(req)
	}
	return nil
}

func (m *MockAccountService) GetAccountByID(id int) (*models.Account, error) {
	if m.GetAccountByIDFn != nil {
		return m.GetAccountByIDFn(id)
	}
	return nil, nil
}

// Scenario 1: Successful Account Creation
func TestCreateAccountHandler(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.Default()

	// Mock the dependencies for the controller
	accountService := &MockAccountService{} // Implement a mock for your AccountService

	// Create an instance of your controller with the mocked dependencies
	controller := &account.AccountController{AccountService: accountService}

	// Set up the route to handle POST requests to /accounts
	router.POST("/accounts", controller.CreateAccount)

	// Create a JSON request body
	reqBody := account.CreateAccountRequest{
		AccountId: 123,
		Balance:   "100.00",
	}
	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Error marshaling JSON request body: %v", err)
	}

	// Create a new HTTP POST request with the JSON request body
	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatalf("Error creating HTTP request: %v", err)
	}

	// Create a response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Serve the HTTP request and record the response
	router.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusCreated)
	}
}

// Scenario 2: Missing Parameter
func TestCreateAccountMissingParamHandler(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create a new Gin router
	router := gin.Default()

	// Mock the dependencies for the controller
	accountService := &MockAccountService{} // Implement a mock for your AccountService

	// Create an instance of your controller with the mocked dependencies
	controller := &account.AccountController{AccountService: accountService}

	// Set up the route to handle POST requests to /accounts
	router.POST("/accounts", controller.CreateAccount)

	// Create a JSON request body
	reqBody := account.CreateAccountRequest{
		AccountId: 123,
	}
	jsonBytes, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatalf("Error marshaling JSON request body: %v", err)
	}

	// Create a new HTTP POST request with the JSON request body
	req, err := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Fatalf("Error creating HTTP request: %v", err)
	}

	// Create a response recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Serve the HTTP request and record the response
	router.ServeHTTP(rr, req)

	// Check the status code of the response
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusBadRequest)
	}
}
