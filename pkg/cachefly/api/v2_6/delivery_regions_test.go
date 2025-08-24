package v2_6

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// READ - Test List method
func TestDeliveryRegionsService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/deliveryregions" {
			t.Errorf("Expected path /api/2.6/deliveryregions, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":2},"data":[{"_id":"region-123","name":"North America","description":"North American delivery region"},{"_id":"region-456","name":"Europe","description":"European delivery region"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &DeliveryRegionsService{Client: client}

	opts := ListDeliveryRegionsOptions{Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Regions) != 2 {
		t.Errorf("Expected 2 delivery regions, got %d", len(result.Regions))
	}
	if result.Regions[0].ID != "region-123" {
		t.Errorf("Expected delivery region ID region-123, got %s", result.Regions[0].ID)
	}
	if result.Regions[0].Name != "North America" {
		t.Errorf("Expected delivery region name 'North America', got %s", result.Regions[0].Name)
	}
	if result.Regions[0].Description != "North American delivery region" {
		t.Errorf("Expected delivery region description 'North American delivery region', got %s", result.Regions[0].Description)
	}
}

// Test List method with search parameter
func TestDeliveryRegionsService_ListWithSearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/deliveryregions" {
			t.Errorf("Expected path /api/2.6/deliveryregions, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		// Check for search parameter
		searchParam := r.URL.Query().Get("search")
		if searchParam != "america" {
			t.Errorf("Expected search parameter 'america', got %s", searchParam)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"region-123","name":"North America","description":"North American delivery region"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &DeliveryRegionsService{Client: client}

	opts := ListDeliveryRegionsOptions{Search: "america", Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Regions) != 1 {
		t.Errorf("Expected 1 delivery region, got %d", len(result.Regions))
	}
	if result.Regions[0].Name != "North America" {
		t.Errorf("Expected delivery region name 'North America', got %s", result.Regions[0].Name)
	}
}

// Test List method with sort parameters
func TestDeliveryRegionsService_ListWithSort(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/deliveryregions" {
			t.Errorf("Expected path /api/2.6/deliveryregions, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		// Check for sortBy parameters
		sortParams := r.URL.Query()["sortBy"]
		if len(sortParams) != 2 {
			t.Errorf("Expected 2 sortBy parameters, got %d", len(sortParams))
		}
		if sortParams[0] != "name" || sortParams[1] != "description" {
			t.Errorf("Expected sortBy parameters ['name', 'description'], got %v", sortParams)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":5,"offset":10,"count":1},"data":[{"_id":"region-123","name":"Asia Pacific","description":"Asia Pacific delivery region"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &DeliveryRegionsService{Client: client}

	opts := ListDeliveryRegionsOptions{
		SortBy: []string{"name", "description"},
		Limit:  5,
		Offset: 10,
	}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Meta.Limit != 5 {
		t.Errorf("Expected meta limit 5, got %d", result.Meta.Limit)
	}
	if result.Meta.Offset != 10 {
		t.Errorf("Expected meta offset 10, got %d", result.Meta.Offset)
	}
}

// Error handling test - search too short
func TestDeliveryRegionsService_ErrorHandling_SearchTooShort(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &DeliveryRegionsService{Client: client}

	opts := ListDeliveryRegionsOptions{Search: "a"}
	_, err := svc.List(context.Background(), opts)

	if err == nil {
		t.Error("Expected error for short search term")
	}
	if err.Error() != "search must be at least 2 characters" {
		t.Errorf("Expected 'search must be at least 2 characters' error, got %s", err.Error())
	}
}

// Error handling test - empty search
func TestDeliveryRegionsService_ErrorHandling_EmptySearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/deliveryregions" {
			t.Errorf("Expected path /api/2.6/deliveryregions, got %s", r.URL.Path)
		}

		// Should not have search parameter when empty
		searchParam := r.URL.Query().Get("search")
		if searchParam != "" {
			t.Errorf("Expected no search parameter, got %s", searchParam)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"region-123","name":"Global","description":"Global delivery region"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &DeliveryRegionsService{Client: client}

	opts := ListDeliveryRegionsOptions{Search: "", Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error for empty search, got %v", err)
	}
	if len(result.Regions) != 1 {
		t.Errorf("Expected 1 delivery region, got %d", len(result.Regions))
	}
}

// Test List method with all parameters
func TestDeliveryRegionsService_ListWithAllParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		// Verify all parameters are present and correct
		if query.Get("search") != "europe" {
			t.Errorf("Expected search 'europe', got %s", query.Get("search"))
		}
		if query.Get("limit") != "20" {
			t.Errorf("Expected limit '20', got %s", query.Get("limit"))
		}
		if query.Get("offset") != "5" {
			t.Errorf("Expected offset '5', got %s", query.Get("offset"))
		}

		sortParams := query["sortBy"]
		if len(sortParams) != 1 || sortParams[0] != "name" {
			t.Errorf("Expected sortBy ['name'], got %v", sortParams)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":20,"offset":5,"count":1},"data":[{"_id":"region-456","name":"Europe","description":"European delivery region"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &DeliveryRegionsService{Client: client}

	opts := ListDeliveryRegionsOptions{
		Search: "europe",
		SortBy: []string{"name"},
		Offset: 5,
		Limit:  20,
	}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Meta.Limit != 20 {
		t.Errorf("Expected meta limit 20, got %d", result.Meta.Limit)
	}
	if result.Meta.Offset != 5 {
		t.Errorf("Expected meta offset 5, got %d", result.Meta.Offset)
	}
	if result.Regions[0].Name != "Europe" {
		t.Errorf("Expected region name 'Europe', got %s", result.Regions[0].Name)
	}
}
