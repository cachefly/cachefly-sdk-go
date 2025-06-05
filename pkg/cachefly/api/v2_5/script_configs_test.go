package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// CREATE - Test Create method
func TestScriptConfigsService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/scriptConfigs" {
			t.Errorf("Expected path /api/2.5/scriptConfigs, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"_id":"config-123","name":"Test Config","mimeType":"application/json"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptConfigsService{Client: client}

	req := CreateScriptConfigRequest{Name: "Test Config", MimeType: "application/json"}
	result, err := svc.Create(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "config-123" {
		t.Errorf("Expected config ID config-123, got %s", result.ID)
	}
}

// READ - Test List method
func TestScriptConfigsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/scriptConfigs" {
			t.Errorf("Expected path /api/2.5/scriptConfigs, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"config-123","name":"Test Config"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptConfigsService{Client: client}

	opts := ListScriptConfigsOptions{Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Configs) != 1 {
		t.Errorf("Expected 1 config, got %d", len(result.Configs))
	}
	if result.Configs[0].ID != "config-123" {
		t.Errorf("Expected config ID config-123, got %s", result.Configs[0].ID)
	}
}

// READ - Test GetByID method
func TestScriptConfigsService_GetByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/scriptConfigs/config-123" {
			t.Errorf("Expected path /api/2.5/scriptConfigs/config-123, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"config-123","name":"Test Config","mimeType":"application/json"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptConfigsService{Client: client}

	result, err := svc.GetByID(context.Background(), "config-123", "")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "config-123" {
		t.Errorf("Expected config ID config-123, got %s", result.ID)
	}
}

// UPDATE - Test UpdateByID method
func TestScriptConfigsService_UpdateByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/scriptConfigs/config-123" {
			t.Errorf("Expected path /api/2.5/scriptConfigs/config-123, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"config-123","name":"Updated Config","mimeType":"application/json"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptConfigsService{Client: client}

	req := UpdateScriptConfigRequest{Name: "Updated Config"}
	result, err := svc.UpdateByID(context.Background(), "config-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "config-123" {
		t.Errorf("Expected config ID config-123, got %s", result.ID)
	}
}

// DELETE - Test DeactivateByID method (we don't have actual delete)
func TestScriptConfigsService_DeactivateByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/scriptConfigs/config-123/deactivate" {
			t.Errorf("Expected path /api/2.5/scriptConfigs/config-123/deactivate, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"config-123","name":"Test Config","status":"inactive"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptConfigsService{Client: client}

	result, err := svc.DeactivateByID(context.Background(), "config-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "config-123" {
		t.Errorf("Expected config ID config-123, got %s", result.ID)
	}
}

// Error handling test - missing ID
func TestScriptConfigsService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptConfigsService{Client: client}

	_, err := svc.GetByID(context.Background(), "", "")

	if err == nil {
		t.Error("Expected error for missing ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}
}
