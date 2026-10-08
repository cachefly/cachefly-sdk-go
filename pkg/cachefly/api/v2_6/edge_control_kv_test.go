package v2_6

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

func TestEdgeControlKVService_GetAccount(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", "/api/2.6/edgecontrol/kv",
		`{"_id":"kv-1","type":"ACCOUNT","data":{"name":"x","ttl":300,"enabled":true,"empty":null}}`, nil)
	svc := &EdgeControlKVService{Client: client}

	kv, err := svc.GetAccount(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if kv.Type != EdgeControlKVTypeAccount {
		t.Errorf("Expected type ACCOUNT, got %s", kv.Type)
	}
	if kv.Data["name"] != "x" || kv.Data["ttl"] != float64(300) || kv.Data["enabled"] != true {
		t.Errorf("Unexpected data: %v", kv.Data)
	}
	if value, ok := kv.Data["empty"]; !ok || value != nil {
		t.Errorf("Expected null value for key empty, got %v (present: %v)", value, ok)
	}
}

func TestEdgeControlKVService_ReplaceAccount(t *testing.T) {
	client := newEdgeControlTestClient(t, "PUT", "/api/2.6/edgecontrol/kv",
		`{"_id":"kv-1","type":"ACCOUNT","data":{"ttl":300,"enabled":false}}`,
		func(r *http.Request) {
			data, ok := decodeJSONBody(t, r)["data"].(map[string]interface{})
			if !ok || data["ttl"] != float64(300) || data["enabled"] != false {
				t.Errorf("Unexpected request data: %v", data)
			}
		})
	svc := &EdgeControlKVService{Client: client}

	kv, err := svc.ReplaceAccount(context.Background(), map[string]interface{}{"ttl": 300, "enabled": false})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if kv.Data["enabled"] != false {
		t.Errorf("Unexpected data: %v", kv.Data)
	}
}

func TestEdgeControlKVService_ReplaceNilData(t *testing.T) {
	client := newEdgeControlTestClient(t, "PUT", "/api/2.6/edgecontrol/kv",
		`{"_id":"kv-1","type":"ACCOUNT","data":{}}`,
		func(r *http.Request) {
			data, ok := decodeJSONBody(t, r)["data"].(map[string]interface{})
			if !ok || len(data) != 0 {
				t.Errorf("Expected an empty data object, got %v", data)
			}
		})
	svc := &EdgeControlKVService{Client: client}

	if _, err := svc.ReplaceAccount(context.Background(), nil); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestEdgeControlKVService_ClearAccount(t *testing.T) {
	client := newEdgeControlTestClient(t, "DELETE", "/api/2.6/edgecontrol/kv",
		`{"_id":"kv-1","type":"ACCOUNT","data":{}}`, nil)
	svc := &EdgeControlKVService{Client: client}

	kv, err := svc.ClearAccount(context.Background())
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(kv.Data) != 0 {
		t.Errorf("Expected empty data, got %v", kv.Data)
	}
}

func TestEdgeControlKVService_GetService(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", "/api/2.6/edgecontrol/services/svc-123/kv",
		`{"_id":"kv-2","type":"SERVICE","data":{"region":"eu"}}`, nil)
	svc := &EdgeControlKVService{Client: client}

	kv, err := svc.GetService(context.Background(), "svc-123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if kv.Type != EdgeControlKVTypeService || kv.Data["region"] != "eu" {
		t.Errorf("Unexpected store: %+v", kv)
	}
}

func TestEdgeControlKVService_ReplaceService(t *testing.T) {
	client := newEdgeControlTestClient(t, "PUT", "/api/2.6/edgecontrol/services/svc-123/kv",
		`{"_id":"kv-2","type":"SERVICE","data":{"region":"us"}}`,
		func(r *http.Request) {
			data, ok := decodeJSONBody(t, r)["data"].(map[string]interface{})
			if !ok || data["region"] != "us" {
				t.Errorf("Unexpected request data: %v", data)
			}
		})
	svc := &EdgeControlKVService{Client: client}

	kv, err := svc.ReplaceService(context.Background(), "svc-123", map[string]interface{}{"region": "us"})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if kv.Data["region"] != "us" {
		t.Errorf("Unexpected data: %v", kv.Data)
	}
}

func TestEdgeControlKVService_ClearService(t *testing.T) {
	client := newEdgeControlTestClient(t, "DELETE", "/api/2.6/edgecontrol/services/svc-123/kv",
		`{"_id":"kv-2","type":"SERVICE","data":{}}`, nil)
	svc := &EdgeControlKVService{Client: client}

	if _, err := svc.ClearService(context.Background(), "svc-123"); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestEdgeControlKVService_GetMerged(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", "/api/2.6/edgecontrol/services/svc-123/kv/merged",
		`{"data":{"region":"us","ttl":300}}`, nil)
	svc := &EdgeControlKVService{Client: client}

	merged, err := svc.GetMerged(context.Background(), "svc-123")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if merged["region"] != "us" || merged["ttl"] != float64(300) {
		t.Errorf("Unexpected merged data: %v", merged)
	}
}

func TestEdgeControlKVService_RequiresServiceID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("Unexpected request to %s", r.URL.Path)
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"})
	svc := &EdgeControlKVService{Client: client}
	ctx := context.Background()

	if _, err := svc.GetService(ctx, ""); err == nil {
		t.Error("Expected error for empty service ID")
	}
	if _, err := svc.ReplaceService(ctx, "", map[string]interface{}{}); err == nil {
		t.Error("Expected error for empty service ID")
	}
	if _, err := svc.ClearService(ctx, ""); err == nil {
		t.Error("Expected error for empty service ID")
	}
	if _, err := svc.GetMerged(ctx, ""); err == nil {
		t.Error("Expected error for empty service ID")
	}
}
