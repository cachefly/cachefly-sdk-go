package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// READ - Test List method
func TestScriptDefinitionsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/scriptConfigDefinitions" {
			t.Errorf("Expected path /api/2.5/scriptConfigDefinitions, got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		// Return a single definition
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
            "meta": {"limit": 10, "offset": 0, "count": 1},
            "data": [
                {
                    "_id":"def-123",
                    "name":"Edge Function A",
                    "helpText":"",
                    "docsLink":"",
                    "scriptConfigs":["cfg-1"],
                    "available":true,
                    "linkWithService":true,
                    "requiresOptions":false,
                    "requiresRules":false,
                    "requiresPlugin":false,
                    "canCreate":true,
                    "purpose":"example",
                    "meta":{},
                    "allowedMimeTypes":["application/json"],
                    "defaultMimeType":"application/json",
                    "valueSchema":{},
                    "dataMode":"JSON",
                    "defaultValue":{}
                }
            ]
        }`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptDefinitionsService{Client: client}

	opts := ListScriptDefinitionsOptions{IncludeFeatures: true, Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Definitions) != 1 {
		t.Errorf("Expected 1 definition, got %d", len(result.Definitions))
	}
	if result.Definitions[0].ID != "def-123" {
		t.Errorf("Expected definition ID def-123, got %s", result.Definitions[0].ID)
	}
}

// READ - Test GetByID method
func TestScriptDefinitionsService_GetByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/scriptConfigDefinitions/def-123" {
			t.Errorf("Expected path /api/2.5/scriptConfigDefinitions/def-123, got %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
            "_id":"def-123",
            "name":"Edge Function A",
            "helpText":"",
            "docsLink":"",
            "scriptConfigs":["cfg-1"],
            "available":true,
            "linkWithService":true,
            "requiresOptions":false,
            "requiresRules":false,
            "requiresPlugin":false,
            "canCreate":true,
            "purpose":"example",
            "meta":{},
            "allowedMimeTypes":["application/json"],
            "defaultMimeType":"application/json",
            "valueSchema":{},
            "dataMode":"JSON",
            "defaultValue":{}
        }`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptDefinitionsService{Client: client}

	result, err := svc.GetByID(context.Background(), "def-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "def-123" {
		t.Errorf("Expected definition ID def-123, got %s", result.ID)
	}
	if result.Name != "Edge Function A" {
		t.Errorf("Expected definition name 'Edge Function A', got %s", result.Name)
	}
}

// Error handling test - missing ID
func TestScriptDefinitionsService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ScriptDefinitionsService{Client: client}

	_, err := svc.GetByID(context.Background(), "")

	if err == nil {
		t.Error("Expected error for missing ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}
}
