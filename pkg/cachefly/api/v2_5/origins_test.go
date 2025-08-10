package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// CREATE - Test Create method
func TestOriginsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/origins" {
			t.Errorf("Expected path /api/2.5/origins, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"_id":"origin-123","hostname":"example.com","type":"http"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &OriginsService{Client: client}

	hostname := "example.com"
	req := CreateOriginRequest{Type: "http", Hostname: &hostname}
	result, err := svc.Create(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "origin-123" {
		t.Errorf("Expected origin ID origin-123, got %s", result.ID)
	}
}

// READ - Test List method
func TestOriginsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/origins" {
			t.Errorf("Expected path /api/2.5/origins, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"origin-123","hostname":"example.com"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &OriginsService{Client: client}

	opts := ListOriginsOptions{Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Origins) != 1 {
		t.Errorf("Expected 1 origin, got %d", len(result.Origins))
	}
	if result.Origins[0].ID != "origin-123" {
		t.Errorf("Expected origin ID origin-123, got %s", result.Origins[0].ID)
	}
}

// READ - Test GetByID method
func TestOriginsService_GetByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/origins/origin-123" {
			t.Errorf("Expected path /api/2.5/origins/origin-123, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"origin-123","hostname":"example.com","type":"http"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &OriginsService{Client: client}

	result, err := svc.GetByID(context.Background(), "origin-123", "")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "origin-123" {
		t.Errorf("Expected origin ID origin-123, got %s", result.ID)
	}
}

// UPDATE - Test UpdateByID method
func TestOriginsService_UpdateByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/origins/origin-123" {
			t.Errorf("Expected path /api/2.5/origins/origin-123, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"origin-123","hostname":"updated.com","type":"http"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &OriginsService{Client: client}

	updatedHostname := "updated.com"
	req := UpdateOriginRequest{Hostname: &updatedHostname}
	result, err := svc.UpdateByID(context.Background(), "origin-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "origin-123" {
		t.Errorf("Expected origin ID origin-123, got %s", result.ID)
	}
}

// DELETE - Test Delete method
func TestOriginsService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/origins/origin-123" {
			t.Errorf("Expected path /api/2.5/origins/origin-123, got %s", r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &OriginsService{Client: client}

	err := svc.Delete(context.Background(), "origin-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Error handling test - missing ID
func TestOriginsService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &OriginsService{Client: client}

	_, err := svc.GetByID(context.Background(), "", "")

	if err == nil {
		t.Error("Expected error for missing ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}
}
