package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// READ - Test List method
func TestServiceRulesService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/rules" {
			t.Errorf("Expected path /api/2.5/services/svc-123/rules, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"rule-123","createdAt":"2023-01-01T00:00:00Z"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceRulesService{Client: client}

	opts := ListServiceRulesOptions{Limit: 10}
	result, err := svc.List(context.Background(), "svc-123", opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Rules) != 1 {
		t.Errorf("Expected 1 rule, got %d", len(result.Rules))
	}
	if result.Rules[0].ID != "rule-123" {
		t.Errorf("Expected rule ID rule-123, got %s", result.Rules[0].ID)
	}
}

// UPDATE - Test Update method (bulk update)
func TestServiceRulesService_Update(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/rules" {
			t.Errorf("Expected path /api/2.5/services/svc-123/rules, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":2},"data":[{"_id":"rule-123"},{"_id":"rule-456"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceRulesService{Client: client}

	req := UpdateServiceRulesRequest{
		Rules: []ServiceRule{
			{ID: "rule-123"},
			{ID: "rule-456"},
		},
	}
	result, err := svc.Update(context.Background(), "svc-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Rules) != 2 {
		t.Errorf("Expected 2 rules, got %d", len(result.Rules))
	}
	if result.Rules[0].ID != "rule-123" {
		t.Errorf("Expected rule ID rule-123, got %s", result.Rules[0].ID)
	}
}

// READ - Test GetSchema method
func TestServiceRulesService_GetSchema(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/rules/schema" {
			t.Errorf("Expected path /api/2.5/services/svc-123/rules/schema, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"type":"object","properties":{"rules":{"type":"array"}}}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceRulesService{Client: client}

	result, err := svc.GetSchema(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result["type"] != "object" {
		t.Errorf("Expected schema type 'object', got %v", result["type"])
	}
}

// Error handling test - missing service ID
func TestServiceRulesService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceRulesService{Client: client}

	_, err := svc.List(context.Background(), "", ListServiceRulesOptions{})

	if err == nil {
		t.Error("Expected error for missing service ID")
	}
	if err.Error() != "serviceID is required" {
		t.Errorf("Expected 'serviceID is required' error, got %s", err.Error())
	}
}
