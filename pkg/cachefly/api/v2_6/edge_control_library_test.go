package v2_6

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

func TestEdgeControlLibraryService_List(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", "/api/2.6/edgecontrol/library",
		`{"meta":{"limit":50,"offset":10,"count":11},"data":[{"_id":"lib-1","type":"SYSTEM","kind":"REQUEST","name":"Geo block","code":"function handler(e) { return e; }"}]}`,
		func(r *http.Request) {
			query := r.URL.Query()
			if query.Get("type") != "SYSTEM" || query.Get("kind") != "REQUEST" || query.Get("search") != "geo" {
				t.Errorf("Unexpected filters: %s", r.URL.RawQuery)
			}
			if sortBy := query["sortBy"]; len(sortBy) != 2 || sortBy[0] != "name" || sortBy[1] != "-createdAt" {
				t.Errorf("Unexpected sortBy: %v", sortBy)
			}
			if query.Get("offset") != "10" || query.Get("limit") != "50" {
				t.Errorf("Unexpected pagination: %s", r.URL.RawQuery)
			}
		})
	svc := &EdgeControlLibraryService{Client: client}

	resp, err := svc.List(context.Background(), ListEdgeControlLibraryOptions{
		Type:   EdgeControlLibraryTypeSystem,
		Kind:   EdgeControlScriptKindRequest,
		Search: "geo",
		SortBy: []string{"name", "-createdAt"},
		Offset: 10,
		Limit:  50,
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Meta.Count != 11 || len(resp.Scripts) != 1 || resp.Scripts[0].Name != "Geo block" {
		t.Errorf("Unexpected response: %+v", resp)
	}
}

func TestEdgeControlLibraryService_GetByID(t *testing.T) {
	client := newEdgeControlTestClient(t, "GET", "/api/2.6/edgecontrol/library/lib-1",
		`{"_id":"lib-1","type":"USER","kind":"RESPONSE","name":"Headers","code":"function handler(e) { return e; }"}`, nil)
	svc := &EdgeControlLibraryService{Client: client}

	script, err := svc.GetByID(context.Background(), "lib-1")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if script.Type != EdgeControlLibraryTypeUser || script.Kind != "RESPONSE" || script.Name != "Headers" {
		t.Errorf("Unexpected script: %+v", script)
	}
}

func TestEdgeControlLibraryService_Create(t *testing.T) {
	client := newEdgeControlTestClient(t, "POST", "/api/2.6/edgecontrol/library",
		`{"_id":"lib-2","type":"USER","kind":"REQUEST","name":"Redirects","code":"function handler(e) { return e; }"}`,
		func(r *http.Request) {
			body := decodeJSONBody(t, r)
			if body["name"] != "Redirects" || body["kind"] != "REQUEST" || body["code"] != "function handler(e) { return e; }" {
				t.Errorf("Unexpected request body: %v", body)
			}
		})
	svc := &EdgeControlLibraryService{Client: client}

	created, err := svc.Create(context.Background(), EdgeControlLibraryScriptRequest{
		Name: "Redirects",
		Kind: EdgeControlScriptKindRequest,
		Code: "function handler(e) { return e; }",
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if created.ID != "lib-2" || created.Type != EdgeControlLibraryTypeUser {
		t.Errorf("Unexpected script: %+v", created)
	}
}

func TestEdgeControlLibraryService_Update(t *testing.T) {
	client := newEdgeControlTestClient(t, "PUT", "/api/2.6/edgecontrol/library/lib-2",
		`{"_id":"lib-2","type":"USER","kind":"RESPONSE","name":"Renamed","code":"function handler(e) { return e; }"}`,
		func(r *http.Request) {
			body := decodeJSONBody(t, r)
			if body["name"] != "Renamed" || body["kind"] != "RESPONSE" {
				t.Errorf("Unexpected request body: %v", body)
			}
		})
	svc := &EdgeControlLibraryService{Client: client}

	updated, err := svc.Update(context.Background(), "lib-2", EdgeControlLibraryScriptRequest{
		Name: "Renamed",
		Kind: EdgeControlScriptKindResponse,
		Code: "function handler(e) { return e; }",
	})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if updated.Name != "Renamed" {
		t.Errorf("Unexpected script: %+v", updated)
	}
}

func TestEdgeControlLibraryService_Delete(t *testing.T) {
	client := newEdgeControlTestClient(t, "DELETE", "/api/2.6/edgecontrol/library/lib-2",
		`{"_id":"lib-2","type":"USER","kind":"RESPONSE","name":"Renamed"}`, nil)
	svc := &EdgeControlLibraryService{Client: client}

	deleted, err := svc.Delete(context.Background(), "lib-2")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if deleted.ID != "lib-2" {
		t.Errorf("Unexpected script: %+v", deleted)
	}
}

func TestEdgeControlLibraryService_RequiresID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("Unexpected request to %s", r.URL.Path)
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "test-token"})
	svc := &EdgeControlLibraryService{Client: client}
	ctx := context.Background()

	if _, err := svc.GetByID(ctx, ""); err == nil {
		t.Error("Expected error for empty ID")
	}
	if _, err := svc.Update(ctx, "", EdgeControlLibraryScriptRequest{}); err == nil {
		t.Error("Expected error for empty ID")
	}
	if _, err := svc.Delete(ctx, ""); err == nil {
		t.Error("Expected error for empty ID")
	}
}
