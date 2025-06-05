package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// CREATE - Test Create method
func TestServicesService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services" {
			t.Errorf("Expected path /api/2.5/services, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"_id":"new-123","name":"New Service"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	req := CreateServiceRequest{Name: "New Service", UniqueName: "new-service"}
	result, err := svc.Create(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "new-123" {
		t.Errorf("Expected service ID new-123, got %s", result.ID)
	}
}

// READ - Test Get method
func TestServicesService_Get(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/test-123" {
			t.Errorf("Expected path /api/2.5/services/test-123, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"test-123","name":"Test Service"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	result, err := svc.Get(context.Background(), "test-123", "", true)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "test-123" {
		t.Errorf("Expected service ID test-123, got %s", result.ID)
	}
}

// READ - Test List method
func TestServicesService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services" {
			t.Errorf("Expected path /api/2.5/services, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"list-123","name":"Listed Service"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	opts := ListOptions{Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Services) != 1 {
		t.Errorf("Expected 1 service, got %d", len(result.Services))
	}
	if result.Services[0].ID != "list-123" {
		t.Errorf("Expected service ID list-123, got %s", result.Services[0].ID)
	}
}

// UPDATE - Test UpdateServiceByID method
func TestServicesService_UpdateServiceByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/update-123" {
			t.Errorf("Expected path /api/2.5/services/update-123, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"update-123","name":"Updated Service"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	req := UpdateServiceRequest{Description: "Updated description"}
	result, err := svc.UpdateServiceByID(context.Background(), "update-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "update-123" {
		t.Errorf("Expected service ID update-123, got %s", result.ID)
	}
}

// DELETE - Test DeactivateServiceByID method
func TestServicesService_DeactivateServiceByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/deactivate-123/deactivate" {
			t.Errorf("Expected path /api/2.5/services/deactivate-123/deactivate, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"deactivate-123","name":"Deactivated Service","status":"inactive"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	result, err := svc.DeactivateServiceByID(context.Background(), "deactivate-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "deactivate-123" {
		t.Errorf("Expected service ID deactivate-123, got %s", result.ID)
	}
}

// Error handling test
func TestServicesService_ErrorHandling(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Invalid request"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	_, err := svc.Get(context.Background(), "error-test", "", false)

	if err == nil {
		t.Error("Expected error for 400 response")
	}
}
