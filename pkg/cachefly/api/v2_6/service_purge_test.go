package v2_6

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

func TestServicesService_Purge_Paths(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/purge" {
			t.Errorf("Expected path /api/2.6/services/svc-123/purge, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		var payload map[string]interface{}
		if err := json.Unmarshal(body, &payload); err != nil {
			t.Fatalf("failed to unmarshal body: %v", err)
		}
		if _, ok := payload["paths"]; !ok {
			t.Errorf("expected 'paths' in payload, got: %s", string(body))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	req := PurgeRequest{Paths: []string{"/index.html", "/images/"}}
	if err := svc.Purge(context.Background(), "svc-123", req); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestServicesService_Purge_All(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/services/svc-123/purge" {
			t.Errorf("Expected path /api/2.6/services/svc-123/purge, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("Expected PUT method, got %s", r.Method)
		}

		body, _ := io.ReadAll(r.Body)
		defer r.Body.Close()

		if !bytes.Contains(body, []byte("\"all\":true")) {
			t.Errorf("expected 'all': true in payload, got: %s", string(body))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	req := PurgeRequest{All: true}
	if err := svc.Purge(context.Background(), "svc-123", req); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestServicesService_Purge_InvalidRequest(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://example.com/api/2.6", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &ServicesService{Client: client}

	// Missing id
	if err := svc.Purge(context.Background(), "", PurgeRequest{All: true}); err == nil {
		t.Errorf("expected error for missing id")
	}

	// Missing both all and paths
	if err := svc.Purge(context.Background(), "svc-123", PurgeRequest{}); err == nil {
		t.Errorf("expected error for missing all/paths")
	}
}
