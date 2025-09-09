package v2_6

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

func TestCacheWarming_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/cachewarming" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"cw1","name":"Warm Home","targets":["https://example.com/"],"regions":["global"]}]}`))
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "t"})
	svc := &CacheWarmingService{Client: client}

	resp, err := svc.List(context.Background(), ListCacheWarmingTasksOptions{Limit: 10})
	if err != nil {
		t.Fatalf("List returned error: %v", err)
	}
	if resp.Meta.Count != 1 {
		t.Fatalf("expected count 1, got %d", resp.Meta.Count)
	}
	if len(resp.Data) != 1 || resp.Data[0].ID != "cw1" {
		t.Fatalf("unexpected data: %+v", resp.Data)
	}
}

func TestCacheWarming_Create_Get_Delete(t *testing.T) {
	// Track created id
	createdID := "cw123"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/api/2.6/cachewarming":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"_id":"` + createdID + `","name":"Warm","targets":["https://example.com/a"],"regions":["global"]}`))
		case r.Method == http.MethodGet && r.URL.Path == "/api/2.6/cachewarming/"+createdID:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"_id":"` + createdID + `","name":"Warm","targets":["https://example.com/a"],"regions":["global"],"status":"QUEUED"}`))
		case r.Method == http.MethodDelete && r.URL.Path == "/api/2.6/cachewarming/"+createdID:
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "t"})
	svc := &CacheWarmingService{Client: client}

	created, err := svc.Create(context.Background(), CreateCacheWarmingTaskRequest{
		Name:    "Warm",
		Targets: []string{"https://example.com/a"},
		Regions: []string{"global"},
	})
	if err != nil {
		t.Fatalf("Create returned error: %v", err)
	}
	if created.ID != createdID {
		t.Fatalf("unexpected id: %s", created.ID)
	}

	got, err := svc.GetByID(context.Background(), createdID)
	if err != nil {
		t.Fatalf("GetByID returned error: %v", err)
	}
	if got.Status != "QUEUED" {
		t.Fatalf("expected status QUEUED, got %s", got.Status)
	}

	if err := svc.DeleteByID(context.Background(), createdID); err != nil {
		t.Fatalf("DeleteByID returned error: %v", err)
	}
}

func TestCacheWarming_Validation(t *testing.T) {
	client := httpclient.New(httpclient.Config{BaseURL: "http://x", AuthToken: "t"})
	svc := &CacheWarmingService{Client: client}

	if _, err := svc.Create(context.Background(), CreateCacheWarmingTaskRequest{}); err == nil {
		t.Fatalf("expected error for missing name and targets")
	}
	if _, err := svc.GetByID(context.Background(), ""); err == nil {
		t.Fatalf("expected error for missing id")
	}
	if err := svc.DeleteByID(context.Background(), ""); err == nil {
		t.Fatalf("expected error for missing id on delete")
	}
}
