package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// CREATE - Test Create method
func TestServiceOptionsRefererRulesService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options/refererrules" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options/refererrules, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"_id":"rule-123","directory":"/images","defaultAction":"allow"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsRefererRulesService{Client: client}

	req := CreateRefererRuleRequest{
		Directory:     "/images",
		DefaultAction: "allow",
		Exceptions:    []string{"example.com"},
	}
	result, err := svc.Create(context.Background(), "svc-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "rule-123" {
		t.Errorf("Expected rule ID rule-123, got %s", result.ID)
	}
}

// READ - Test List method
func TestServiceOptionsRefererRulesService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options/refererrules" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options/refererrules, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"rule-123","directory":"/images"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsRefererRulesService{Client: client}

	opts := ListRefererRulesOptions{Limit: 10}
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

// READ - Test GetByID method
func TestServiceOptionsRefererRulesService_GetByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options/refererrules/rule-123" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options/refererrules/rule-123, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"rule-123","directory":"/images","defaultAction":"allow"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsRefererRulesService{Client: client}

	result, err := svc.GetByID(context.Background(), "svc-123", "rule-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "rule-123" {
		t.Errorf("Expected rule ID rule-123, got %s", result.ID)
	}
}

// UPDATE - Test Update method
func TestServiceOptionsRefererRulesService_Update(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options/refererrules/rule-123" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options/refererrules/rule-123, got %s", r.URL.Path)
		}
		if r.Method != "PUT" {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"rule-123","directory":"/updated","defaultAction":"deny"}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsRefererRulesService{Client: client}

	req := UpdateRefererRuleRequest{
		Directory:     "/updated",
		DefaultAction: "deny",
	}
	result, err := svc.Update(context.Background(), "svc-123", "rule-123", req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "rule-123" {
		t.Errorf("Expected rule ID rule-123, got %s", result.ID)
	}
}

// DELETE - Test Delete method
func TestServiceOptionsRefererRulesService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/services/svc-123/options/refererrules/rule-123" {
			t.Errorf("Expected path /api/2.5/services/svc-123/options/refererrules/rule-123, got %s", r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsRefererRulesService{Client: client}

	err := svc.Delete(context.Background(), "svc-123", "rule-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Error handling test - missing service ID and rule ID
func TestServiceOptionsRefererRulesService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsRefererRulesService{Client: client}

	// Test missing service ID
	_, err := svc.List(context.Background(), "", ListRefererRulesOptions{})
	if err == nil {
		t.Error("Expected error for missing service ID")
	}
	if err.Error() != "service ID is required" {
		t.Errorf("Expected 'service ID is required' error, got %s", err.Error())
	}

	// Test missing both service ID and rule ID
	_, err = svc.GetByID(context.Background(), "", "")
	if err == nil {
		t.Error("Expected error for missing service ID and rule ID")
	}
	if err.Error() != "service ID and rule ID are required" {
		t.Errorf("Expected 'service ID and rule ID are required' error, got %s", err.Error())
	}
}
