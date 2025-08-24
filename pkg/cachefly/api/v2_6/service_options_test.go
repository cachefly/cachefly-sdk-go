package v2_6

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// READ - Test GetOptions method
func TestServiceOptionsService_GetOptions(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/options" {
			t.Errorf("Expected path /api/2.6/services/svc-123/options, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"ftp":true,"cors":false,"autoRedirect":true}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	result, err := svc.GetOptions(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if ftpVal, ok := result["ftp"].(bool); !ok || !ftpVal {
		t.Error("Expected FTP to be true")
	}
}

// READ - Test GetOptionsMetadata method
func TestServiceOptionsService_GetOptionsMetadata(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/options/metadata" {
			t.Errorf("Expected path /api/2.6/services/svc-123/options/metadata, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"meta": {"count": 2},
			"data": [
				{
					"_id": "opt1",
					"name": "cors",
					"title": "CORS Settings",
					"type": "dynamic",
					"readOnly": false,
					"property": {
						"name": "cors",
						"type": "boolean"
					}
				},
				{
					"_id": "opt2",
					"name": "ftp",
					"title": "FTP Access",
					"type": "dynamic", 
					"readOnly": false,
					"property": {
						"name": "ftp",
						"type": "boolean"
					}
				}
			]
		}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	result, err := svc.GetOptionsMetadata(context.Background(), "svc-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Meta.Count != 2 {
		t.Errorf("Expected 2 options, got %d", result.Meta.Count)
	}
	if len(result.Data) != 2 {
		t.Errorf("Expected 2 data items, got %d", len(result.Data))
	}
}

// UPDATE - Test UpdateOptions method with mocked metadata validation
func TestServiceOptionsService_UpdateOptions(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")

		// First request: GET metadata
		if requestCount == 1 {
			if r.URL.Path != "/api/2.6/services/svc-123/options/metadata" {
				t.Errorf("Expected metadata path /api/2.6/services/svc-123/options/metadata, got %s", r.URL.Path)
			}
			if r.Method != "GET" {
				t.Errorf("Expected GET method for metadata, got %s", r.Method)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"meta": {"count": 3},
				"data": [
					{
						"_id": "opt1",
						"name": "cors",
						"type": "dynamic",
						"readOnly": false,
						"property": {
							"name": "cors",
							"type": "boolean"
						}
					},
					{
						"_id": "opt2", 
						"name": "ftp",
						"type": "dynamic",
						"readOnly": false,
						"property": {
							"name": "ftp",
							"type": "boolean"
						}
					},
					{
						"_id": "opt3",
						"name": "autoRedirect", 
						"type": "dynamic",
						"readOnly": false,
						"property": {
							"name": "autoRedirect",
							"type": "boolean"
						}
					}
				]
			}`))
			return
		}

		// Second request: PUT options
		if requestCount == 2 {
			if r.URL.Path != "/api/2.6/services/svc-123/options" {
				t.Errorf("Expected options path /api/2.6/services/svc-123/options, got %s", r.URL.Path)
			}
			if r.Method != "PUT" {
				t.Errorf("Expected PUT method for update, got %s", r.Method)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"ftp":false,"cors":true,"autoRedirect":false}`))
			return
		}

		t.Errorf("Unexpected request count: %d", requestCount)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	// Test with valid options that match metadata
	req := ServiceOptions{
		"ftp":          false,
		"cors":         true,
		"autoRedirect": false,
	}

	result, err := svc.UpdateOptions(context.Background(), "svc-123", req)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the response
	if corsVal, ok := result["cors"].(bool); !ok || !corsVal {
		t.Error("Expected CORS to be true")
	}

	if ftpVal, ok := result["ftp"].(bool); !ok || ftpVal {
		t.Error("Expected FTP to be false")
	}

	// Verify both metadata and update requests were made
	if requestCount != 2 {
		t.Errorf("Expected 2 requests (metadata + update), got %d", requestCount)
	}
}

// Test validation error handling
func TestServiceOptionsService_UpdateOptions_ValidationError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/2.6/services/svc-123/options/metadata" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
				"meta": {"count": 1},
				"data": [
					{
						"_id": "opt1",
						"name": "cors",
						"type": "dynamic", 
						"readOnly": false,
						"property": {
							"name": "cors",
							"type": "boolean"
						}
					}
				]
			}`))
			return
		}
		t.Errorf("Unexpected request to %s", r.URL.Path)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	// Test with invalid option that doesn't exist in metadata
	req := ServiceOptions{
		"cors":           true,
		"invalid_option": false, // This should cause validation error
	}

	_, err := svc.UpdateOptions(context.Background(), "svc-123", req)
	if err == nil {
		t.Fatal("Expected validation error for invalid option")
	}

	// Check if it's a validation error
	validationErr, ok := err.(ServiceOptionsValidationError)
	if !ok {
		t.Fatalf("Expected ServiceOptionsValidationError, got %T", err)
	}

	// Check error details
	if len(validationErr.Errors) != 1 {
		t.Errorf("Expected 1 validation error, got %d", len(validationErr.Errors))
	}

	if validationErr.Errors[0].Field != "invalid_option" {
		t.Errorf("Expected error for 'invalid_option', got '%s'", validationErr.Errors[0].Field)
	}

	if validationErr.Errors[0].Code != "OPTION_NOT_AVAILABLE" {
		t.Errorf("Expected error code 'OPTION_NOT_AVAILABLE', got '%s'", validationErr.Errors[0].Code)
	}
}

// Error handling test - missing service ID
func TestServiceOptionsService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServiceOptionsService{Client: client}

	// Test GetOptions with missing service ID
	_, err := svc.GetOptions(context.Background(), "")
	if err == nil {
		t.Error("Expected error for missing service ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}

	// Test GetOptionsMetadata with missing service ID
	_, err = svc.GetOptionsMetadata(context.Background(), "")
	if err == nil {
		t.Error("Expected error for missing service ID in GetOptionsMetadata")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error for GetOptionsMetadata, got %s", err.Error())
	}

	// Test UpdateOptions with missing service ID
	opts := ServiceOptions{"cors": true}
	_, err = svc.UpdateOptions(context.Background(), "", opts)
	if err == nil {
		t.Error("Expected error for missing service ID in UpdateOptions")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error for UpdateOptions, got %s", err.Error())
	}
}
