package v2_6

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// CREATE - Test CreateChildAccount method
func TestAccountsService_CreateChildAccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/accounts" {
			t.Errorf("Expected path /api/2.6/accounts, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"_id":"account-123","companyName":"Test Company","email":"test@example.com"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &AccountsService{Client: client}

	req := CreateChildAccountRequest{
		CompanyName: "Test Company",
		Username:    "testuser",
		Password:    "password123",
		FullName:    "Test User",
		Email:       "test@example.com",
	}
	result, err := svc.CreateChildAccount(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "account-123" {
		t.Errorf("Expected account ID account-123, got %s", result.ID)
	}
}

// READ - Test Get method (current account)
func TestAccountsService_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/accounts/me" {
			t.Errorf("Expected path /api/2.6/accounts/me, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"current-123","companyName":"My Company","email":"me@example.com"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &AccountsService{Client: client}

	result, err := svc.Get(context.Background(), "")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "current-123" {
		t.Errorf("Expected account ID current-123, got %s", result.ID)
	}
}

// READ - Test List method
func TestAccountsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/accounts" {
			t.Errorf("Expected path /api/2.6/accounts, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"account-123","companyName":"Test Company"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &AccountsService{Client: client}

	opts := ListAccountsOptions{Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Accounts) != 1 {
		t.Errorf("Expected 1 account, got %d", len(result.Accounts))
	}
	if result.Accounts[0].ID != "account-123" {
		t.Errorf("Expected account ID account-123, got %s", result.Accounts[0].ID)
	}
}

// UPDATE - Test UpdateCurrentAccount method
func TestAccountsService_UpdateCurrentAccount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/accounts/me" {
			t.Errorf("Expected path /api/2.6/accounts/me, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"current-123","companyName":"Updated Company","email":"updated@example.com"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &AccountsService{Client: client}

	req := UpdateAccountRequest{CompanyName: "Updated Company"}
	result, err := svc.UpdateCurrentAccount(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "current-123" {
		t.Errorf("Expected account ID current-123, got %s", result.ID)
	}
}

// DELETE - Test DeactivateAccountByID method
func TestAccountsService_DeactivateAccountByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/accounts/account-123/deactivate" {
			t.Errorf("Expected path /api/2.6/accounts/account-123/deactivate, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"account-123","companyName":"Test Company","status":"inactive"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &AccountsService{Client: client}

	result, err := svc.DeactivateAccountByID(context.Background(), "account-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "account-123" {
		t.Errorf("Expected account ID account-123, got %s", result.ID)
	}
}

// Error handling test - missing required fields
func TestAccountsService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &AccountsService{Client: client}

	// Test missing required fields for CreateChildAccount
	req := CreateChildAccountRequest{CompanyName: "Test"} // Missing other required fields
	_, err := svc.CreateChildAccount(context.Background(), req)

	if err == nil {
		t.Error("Expected error for missing required fields")
	}

	// Test missing ID for GetByID
	_, err = svc.GetByID(context.Background(), "", "")
	if err == nil {
		t.Error("Expected error for missing ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}
}
