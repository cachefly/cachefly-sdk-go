package v2_6

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// CREATE - Test Create method
func TestServiceDomainsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/domains" {
			t.Errorf("Expected path /api/2.6/services/svc-123/domains, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"_id":"dom-123","name":"example.com","service":"svc-123"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceDomainsService{Client: client}

	req := CreateServiceDomainRequest{Name: "example.com"}
	result, err := svc.Create(context.Background(), "svc-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "dom-123" {
		t.Errorf("Expected domain ID dom-123, got %s", result.ID)
	}
}

// READ - Test List method
func TestServiceDomainsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/domains" {
			t.Errorf("Expected path /api/2.6/services/svc-123/domains, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"dom-123","name":"example.com"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceDomainsService{Client: client}

	opts := ListServiceDomainsOptions{Limit: 10}
	result, err := svc.List(context.Background(), "svc-123", opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Domains) != 1 {
		t.Errorf("Expected 1 domain, got %d", len(result.Domains))
	}
	if result.Domains[0].ID != "dom-123" {
		t.Errorf("Expected domain ID dom-123, got %s", result.Domains[0].ID)
	}
}

// READ - Test GetByID method
func TestServiceDomainsService_GetByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/domains/dom-123" {
			t.Errorf("Expected path /api/2.6/services/svc-123/domains/dom-123, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"dom-123","name":"example.com","service":"svc-123"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceDomainsService{Client: client}

	result, err := svc.GetByID(context.Background(), "svc-123", "dom-123", "")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "dom-123" {
		t.Errorf("Expected domain ID dom-123, got %s", result.ID)
	}
}

// UPDATE - Test UpdateByID method
func TestServiceDomainsService_UpdateByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/domains/dom-123" {
			t.Errorf("Expected path /api/2.6/services/svc-123/domains/dom-123, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"dom-123","name":"updated.com","service":"svc-123"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceDomainsService{Client: client}

	req := UpdateServiceDomainRequest{Name: "updated.com"}
	result, err := svc.UpdateByID(context.Background(), "svc-123", "dom-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "dom-123" {
		t.Errorf("Expected domain ID dom-123, got %s", result.ID)
	}
}

// DELETE - Test DeleteByID method
func TestServiceDomainsService_DeleteByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/domains/dom-123" {
			t.Errorf("Expected path /api/2.6/services/svc-123/domains/dom-123, got %s", r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceDomainsService{Client: client}

	err := svc.DeleteByID(context.Background(), "svc-123", "dom-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Error handling test - missing service ID
func TestServiceDomainsService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceDomainsService{Client: client}

	_, err := svc.List(context.Background(), "", ListServiceDomainsOptions{})

	if err == nil {
		t.Error("Expected error for missing service ID")
	}
	if err.Error() != "service ID is required" {
		t.Errorf("Expected 'service ID is required' error, got %s", err.Error())
	}
}
