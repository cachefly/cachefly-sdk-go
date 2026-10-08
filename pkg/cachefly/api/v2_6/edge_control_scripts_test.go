package v2_6

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// newEdgeControlTestClient starts a server that checks the request method and
// path, runs check (if set) and responds with the given JSON.
func newEdgeControlTestClient(t *testing.T, method, path, response string, check func(r *http.Request)) *httpclient.Client {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			t.Errorf("Expected %s method, got %s", method, r.Method)
		}
		if r.URL.Path != path {
			t.Errorf("Expected path %s, got %s", path, r.URL.Path)
		}
		if check != nil {
			check(r)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	t.Cleanup(server.Close)

	return httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"})
}

// decodeJSONBody decodes the request body into a map.
func decodeJSONBody(t *testing.T, r *http.Request) map[string]interface{} {
	t.Helper()
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		t.Errorf("Failed to decode request body: %v", err)
	}
	return body
}

const edgeControlScriptsTestPath = "/api/2.6/edgecontrol/services/svc-123"

func TestEdgeControlScriptsService_GetDraft(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", edgeControlScriptsTestPath+"/request",
		`{"_id":"draft-1","kind":"REQUEST","version":null,"script":"function handler(event) { return event; }","scriptSize":41,"status":"DRAFT","filename":null,"dirty":false}`, nil)
	svc := &EdgeControlScriptsService{Client: client}

	draft, err := svc.GetDraft(context.Background(), "svc-123", EdgeControlScriptKindRequest)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if draft.Status != EdgeControlScriptStatusDraft || draft.Version != 0 || draft.Filename != "" {
		t.Errorf("Unexpected draft: %+v", draft)
	}
}

func TestEdgeControlScriptsService_SaveDraft(t *testing.T) {
	client := newEdgeControlTestClient(t, "PUT", edgeControlScriptsTestPath+"/request",
		`{"_id":"draft-1","kind":"REQUEST","script":"function handler(e) { return e; }","status":"DRAFT","dirty":true}`,
		func(r *http.Request) {
			if body := decodeJSONBody(t, r); body["script"] != "function handler(e) { return e; }" {
				t.Errorf("Unexpected request body: %v", body)
			}
		})
	svc := &EdgeControlScriptsService{Client: client}

	draft, err := svc.SaveDraft(context.Background(), "svc-123", EdgeControlScriptKindRequest, "function handler(e) { return e; }")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !draft.Dirty {
		t.Errorf("Expected dirty draft, got %+v", draft)
	}
}

func TestEdgeControlScriptsService_ListVersions(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", edgeControlScriptsTestPath+"/request/versions",
		`[{"_id":"v-2","kind":"REQUEST","version":2,"status":"ACTIVE"},{"_id":"v-1","kind":"REQUEST","version":1,"status":"DEACTIVATED"}]`, nil)
	svc := &EdgeControlScriptsService{Client: client}

	versions, err := svc.ListVersions(context.Background(), "svc-123", EdgeControlScriptKindRequest)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(versions) != 2 || versions[0].Version != 2 || versions[1].Status != EdgeControlScriptStatusDeactivated {
		t.Errorf("Unexpected versions: %+v", versions)
	}
}

func TestEdgeControlScriptsService_GetActiveVersion(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", edgeControlScriptsTestPath+"/response/versions/active",
		`{"_id":"v-4","kind":"RESPONSE","version":4,"status":"ACTIVE","lastActivatedAt":"2026-10-08T12:00:00.000Z"}`, nil)
	svc := &EdgeControlScriptsService{Client: client}

	active, err := svc.GetActiveVersion(context.Background(), "svc-123", EdgeControlScriptKindResponse)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if active.Version != 4 || active.Kind != "RESPONSE" || active.LastActivatedAt == "" {
		t.Errorf("Unexpected active version: %+v", active)
	}
}

func TestEdgeControlScriptsService_GetActiveVersionNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message":"No active edge control script for this service."}`))
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"})
	svc := &EdgeControlScriptsService{Client: client}

	_, err := svc.GetActiveVersion(context.Background(), "svc-123", EdgeControlScriptKindRequest)
	if err == nil || !strings.Contains(err.Error(), "API error 404") {
		t.Errorf("Expected a 404 API error, got %v", err)
	}
}

func TestEdgeControlScriptsService_GetVersion(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", edgeControlScriptsTestPath+"/request/versions/3",
		`{"_id":"v-3","kind":"REQUEST","version":3,"script":"function handler(e) { return e; }","scriptSize":34,"status":"DEACTIVATED","filename":"abc.js"}`, nil)
	svc := &EdgeControlScriptsService{Client: client}

	version, err := svc.GetVersion(context.Background(), "svc-123", EdgeControlScriptKindRequest, 3)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if version.Version != 3 || version.ScriptSize != 34 || version.Filename != "abc.js" {
		t.Errorf("Unexpected version: %+v", version)
	}
}

func TestEdgeControlScriptsService_CreateVersion(t *testing.T) {
	client := newEdgeControlTestClient(t, "POST", edgeControlScriptsTestPath+"/request/versions",
		`{"_id":"v-5","kind":"REQUEST","version":5,"status":"DEACTIVATED"}`,
		func(r *http.Request) {
			if r.URL.RawQuery != "" {
				t.Errorf("Expected no query parameters, got %s", r.URL.RawQuery)
			}
			if body := decodeJSONBody(t, r); body["script"] != "function handler(e) { return e; }" {
				t.Errorf("Unexpected request body: %v", body)
			}
		})
	svc := &EdgeControlScriptsService{Client: client}

	created, err := svc.CreateVersion(context.Background(), "svc-123", EdgeControlScriptKindRequest, "function handler(e) { return e; }")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if created.Version != 5 || created.Status != EdgeControlScriptStatusDeactivated {
		t.Errorf("Unexpected version: %+v", created)
	}
}

func TestEdgeControlScriptsService_CreateVersionFromDraft(t *testing.T) {
	tests := []struct {
		name       string
		resetDraft bool
		wantQuery  string
	}{
		{name: "keep draft", resetDraft: false, wantQuery: "fromDraft=true"},
		{name: "reset draft", resetDraft: true, wantQuery: "fromDraft=true&resetDraft=true"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := newEdgeControlTestClient(t, "POST", edgeControlScriptsTestPath+"/request/versions",
				`{"_id":"v-6","kind":"REQUEST","version":6,"status":"DEACTIVATED"}`,
				func(r *http.Request) {
					if r.URL.RawQuery != tt.wantQuery {
						t.Errorf("Expected query %s, got %s", tt.wantQuery, r.URL.RawQuery)
					}
					if body := decodeJSONBody(t, r); body == nil || len(body) != 0 {
						t.Errorf("Expected an empty JSON object body, got %v", body)
					}
				})
			svc := &EdgeControlScriptsService{Client: client}

			created, err := svc.CreateVersionFromDraft(context.Background(), "svc-123", EdgeControlScriptKindRequest, tt.resetDraft)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if created.Version != 6 {
				t.Errorf("Unexpected version: %+v", created)
			}
		})
	}
}

func TestEdgeControlScriptsService_ActivateVersion(t *testing.T) {
	client := newEdgeControlTestClient(t, "PUT", edgeControlScriptsTestPath+"/request/versions/2/activate",
		`{"_id":"v-2","kind":"REQUEST","version":2,"status":"ACTIVE"}`, nil)
	svc := &EdgeControlScriptsService{Client: client}

	activated, err := svc.ActivateVersion(context.Background(), "svc-123", EdgeControlScriptKindRequest, 2)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if activated.Status != EdgeControlScriptStatusActive {
		t.Errorf("Unexpected version: %+v", activated)
	}
}

func TestEdgeControlScriptsService_ActivateLastVersion(t *testing.T) {
	client := newEdgeControlTestClient(t, "PUT", edgeControlScriptsTestPath+"/request/activate",
		`{"_id":"v-1","kind":"REQUEST","version":1,"status":"ACTIVE"}`, nil)
	svc := &EdgeControlScriptsService{Client: client}

	activated, err := svc.ActivateLastVersion(context.Background(), "svc-123", EdgeControlScriptKindRequest)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if activated.Version != 1 || activated.Status != EdgeControlScriptStatusActive {
		t.Errorf("Unexpected version: %+v", activated)
	}
}

func TestEdgeControlScriptsService_Deactivate(t *testing.T) {
	client := newEdgeControlTestClient(t, "PUT", edgeControlScriptsTestPath+"/response/deactivate",
		`{"message":"Edge control script deactivated."}`, nil)
	svc := &EdgeControlScriptsService{Client: client}

	if err := svc.Deactivate(context.Background(), "svc-123", EdgeControlScriptKindResponse); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestEdgeControlScriptsService_LowercaseKind(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", edgeControlScriptsTestPath+"/response",
		`{"_id":"draft-2","kind":"RESPONSE","status":"DRAFT"}`, nil)
	svc := &EdgeControlScriptsService{Client: client}

	if _, err := svc.GetDraft(context.Background(), "svc-123", "response"); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestEdgeControlScriptsService_InvalidArguments(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("Unexpected request to %s", r.URL.Path)
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"})
	svc := &EdgeControlScriptsService{Client: client}
	ctx := context.Background()

	if _, err := svc.GetDraft(ctx, "", EdgeControlScriptKindRequest); err == nil {
		t.Error("Expected error for empty service ID")
	}
	if _, err := svc.GetDraft(ctx, "svc-123", "BOTH"); err == nil {
		t.Error("Expected error for invalid kind")
	}
	if _, err := svc.GetVersion(ctx, "svc-123", EdgeControlScriptKindRequest, 0); err == nil {
		t.Error("Expected error for version 0")
	}
	if _, err := svc.ActivateVersion(ctx, "svc-123", EdgeControlScriptKindRequest, 0); err == nil {
		t.Error("Expected error for version 0")
	}
	if _, err := svc.CreateVersion(ctx, "svc-123", EdgeControlScriptKindRequest, ""); err == nil {
		t.Error("Expected error for empty script")
	}
	if _, err := svc.SaveDraft(ctx, "svc-123", EdgeControlScriptKindRequest, ""); err == nil {
		t.Error("Expected error for empty script")
	}
}
