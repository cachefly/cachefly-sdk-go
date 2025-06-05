package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// READ - Test List method
func TestTLSProfilesService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/tlsprofiles" {
			t.Errorf("Expected path /api/2.5/tlsprofiles, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"tls-123","name":"Modern TLS","createdAt":"2023-01-01T00:00:00Z"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &TLSProfilesService{Client: client}

	opts := ListTLSProfilesOptions{Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Profiles) != 1 {
		t.Errorf("Expected 1 TLS profile, got %d", len(result.Profiles))
	}
	if result.Profiles[0].ID != "tls-123" {
		t.Errorf("Expected TLS profile ID tls-123, got %s", result.Profiles[0].ID)
	}
}

// READ - Test GetByID method
func TestTLSProfilesService_GetByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/tlsprofiles/tls-123" {
			t.Errorf("Expected path /api/2.5/tlsprofiles/tls-123, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"tls-123","name":"Modern TLS","createdAt":"2023-01-01T00:00:00Z"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &TLSProfilesService{Client: client}

	result, err := svc.GetByID(context.Background(), "tls-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "tls-123" {
		t.Errorf("Expected TLS profile ID tls-123, got %s", result.ID)
	}
	if result.Name != "Modern TLS" {
		t.Errorf("Expected TLS profile name 'Modern TLS', got %s", result.Name)
	}
}

// Error handling test - missing ID
func TestTLSProfilesService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &TLSProfilesService{Client: client}

	_, err := svc.GetByID(context.Background(), "")

	if err == nil {
		t.Error("Expected error for missing ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}
}
