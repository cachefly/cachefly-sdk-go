/*
Core SDK responsibilities tested:

constructs correct API endpoints
Uses proper HTTP methods for all CRUD operations
Validates required service ID parameter
Handles string responses (configurations) and JSON responses (schemas)
Parses JSON schema data as map[string]interface{}
Handles error scenarios appropriately
*/

package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// CREATE - Test CreateConfiguration method
func TestServiceImageOptimizationService_CreateConfiguration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/imageopt4" {
			t.Errorf("Expected path /api/2.5/services/svc-123/imageopt4, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`"enabled: true\nformats: [webp, avif]\ndefaultQuality: 85"`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceImageOptimizationService{Client: client}

	req := CreateImageOptimizationOptions{
		Enabled:        true,
		Formats:        []string{"webp", "avif"},
		DefaultQuality: 85,
	}
	result, err := svc.CreateConfiguration(context.Background(), "svc-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == "" {
		t.Error("Expected configuration string, got empty")
	}
}

// READ - Test GetConfiguration method
func TestServiceImageOptimizationService_GetConfiguration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/imageopt4" {
			t.Errorf("Expected path /api/2.5/services/svc-123/imageopt4, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"enabled: true\nformats: [webp]\ndefaultQuality: 80"`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceImageOptimizationService{Client: client}

	result, err := svc.GetConfiguration(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == "" {
		t.Error("Expected configuration string, got empty")
	}
}

// UPDATE - Test UpdateConfiguration method
func TestServiceImageOptimizationService_UpdateConfiguration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/imageopt4" {
			t.Errorf("Expected path /api/2.5/services/svc-123/imageopt4, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`"enabled: true\nformats: [webp, avif]\ndefaultQuality: 90"`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceImageOptimizationService{Client: client}

	configStr := "enabled: true\nformats: [webp, avif]\ndefaultQuality: 90"
	result, err := svc.UpdateConfiguration(context.Background(), "svc-123", configStr)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == "" {
		t.Error("Expected updated configuration string, got empty")
	}
}

// DELETE - Test DeleteConfiguration method
func TestServiceImageOptimizationService_DeleteConfiguration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/imageopt4" {
			t.Errorf("Expected path /api/2.5/services/svc-123/imageopt4, got %s", r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceImageOptimizationService{Client: client}

	err := svc.DeleteConfiguration(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// READ - Test GetSchema method
func TestServiceImageOptimizationService_GetSchema(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/imageopt4/schema" {
			t.Errorf("Expected path /api/2.5/services/svc-123/imageopt4/schema, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"type":"object","properties":{"enabled":{"type":"boolean"}}}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceImageOptimizationService{Client: client}

	result, err := svc.GetSchema(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result["type"] != "object" {
		t.Errorf("Expected schema type 'object', got %v", result["type"])
	}
}

// Error handling test - missing service ID
func TestServiceImageOptimizationService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceImageOptimizationService{Client: client}

	_, err := svc.GetConfiguration(context.Background(), "")

	if err == nil {
		t.Error("Expected error for missing service ID")
	}
	if err.Error() != "serviceID is required" {
		t.Errorf("Expected 'serviceID is required' error, got %s", err.Error())
	}
}
