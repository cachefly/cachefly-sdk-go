package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// READ - Test GetBasicOptions method
func TestServiceOptionsService_GetBasicOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ftp":true,"cors":false,"autoRedirect":true}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	result, err := svc.GetBasicOptions(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !result.FTP {
		t.Error("Expected FTP to be true")
	}
}

// UPDATE - Test SaveBasicOptions method
func TestServiceOptionsService_SaveBasicOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ftp":false,"cors":true,"autoRedirect":false}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	req := ServiceOptions{FTP: false, CORS: true}
	result, err := svc.SaveBasicOptions(context.Background(), "svc-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !result.CORS {
		t.Error("Expected CORS to be true")
	}
}

// READ - Test GetLegacyAPIKey method
func TestServiceOptionsService_GetLegacyAPIKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options/apikey" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options/apikey, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"apiKey":"test-api-key-123"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	result, err := svc.GetLegacyAPIKey(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.APIKey != "test-api-key-123" {
		t.Errorf("Expected API key test-api-key-123, got %s", result.APIKey)
	}
}

// CREATE - Test RegenerateLegacyAPIKey method
func TestServiceOptionsService_RegenerateLegacyAPIKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options/apikey" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options/apikey, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"apiKey":"new-api-key-456"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	result, err := svc.RegenerateLegacyAPIKey(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.APIKey != "new-api-key-456" {
		t.Errorf("Expected API key new-api-key-456, got %s", result.APIKey)
	}
}

// DELETE - Test DeleteLegacyAPIKey method
func TestServiceOptionsService_DeleteLegacyAPIKey(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options/apikey" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options/apikey, got %s", r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	err := svc.DeleteLegacyAPIKey(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Error handling test - missing service ID
func TestServiceOptionsService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	_, err := svc.GetBasicOptions(context.Background(), "")

	if err == nil {
		t.Error("Expected error for missing service ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}
}
