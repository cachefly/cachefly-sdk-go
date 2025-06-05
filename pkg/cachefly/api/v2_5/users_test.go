package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// CREATE - Test Create method
func TestUsersService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/users" {
			t.Errorf("Expected path /api/2.5/users, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"_id":"user-123","username":"testuser","email":"test@example.com"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &UsersService{Client: client}

	req := CreateUserRequest{
		Username: "testuser",
		Password: "password123",
		Email:    "test@example.com",
		FullName: "Test User",
	}
	result, err := svc.Create(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "user-123" {
		t.Errorf("Expected user ID user-123, got %s", result.ID)
	}
}

// READ - Test GetCurrentUser method
func TestUsersService_GetCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/users/me" {
			t.Errorf("Expected path /api/2.5/users/me, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"current-123","username":"currentuser","email":"current@example.com"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &UsersService{Client: client}

	result, err := svc.GetCurrentUser(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "current-123" {
		t.Errorf("Expected user ID current-123, got %s", result.ID)
	}
}

// READ - Test List method
func TestUsersService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/users" {
			t.Errorf("Expected path /api/2.5/users, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"user-123","username":"testuser"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &UsersService{Client: client}

	opts := ListUsersOptions{Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(result.Users))
	}
	if result.Users[0].ID != "user-123" {
		t.Errorf("Expected user ID user-123, got %s", result.Users[0].ID)
	}
}

// UPDATE - Test UpdateCurrentUser method
func TestUsersService_UpdateCurrentUser(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/users/me" {
			t.Errorf("Expected path /api/2.5/users/me, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"current-123","username":"currentuser","email":"updated@example.com"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &UsersService{Client: client}

	req := UpdateUserRequest{Email: "updated@example.com"}
	result, err := svc.UpdateCurrentUser(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "current-123" {
		t.Errorf("Expected user ID current-123, got %s", result.ID)
	}
}

// DELETE - Test DeleteByID method
func TestUsersService_DeleteByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/users/user-123" {
			t.Errorf("Expected path /api/2.5/users/user-123, got %s", r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &UsersService{Client: client}

	err := svc.DeleteByID(context.Background(), "user-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Error handling test - missing ID
func TestUsersService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &UsersService{Client: client}

	_, err := svc.GetByID(context.Background(), "", "")

	if err == nil {
		t.Error("Expected error for missing ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}
}
