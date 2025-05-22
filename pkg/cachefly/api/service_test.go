package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
)

func TestServicesService_List(t *testing.T) {
	loadEnv(t)

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		t.Skip("CACHEFLY_API_TOKEN not set")
	}

	cfg := httpclient.Config{
		BaseURL:   "https://api.cachefly.com/api/2.5",
		AuthToken: token,
	}
	client := httpclient.New(cfg)
	svc := &api.ServicesService{Client: client}

	resp, err := svc.List(context.Background(), api.ListOptions{Limit: 1})
	if err != nil {
		t.Fatalf("expected no error listing services, got %v", err)
	}
	if len(resp.Services) == 0 {
		t.Fatalf("expected at least one service in list, got 0")
	}
}

func TestServicesService_GetByID(t *testing.T) {
	loadEnv(t)

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		t.Skip("CACHEFLY_API_TOKEN not set")
	}

	cfg := httpclient.Config{
		BaseURL:   "https://api.cachefly.com/api/2.5",
		AuthToken: token,
	}
	client := httpclient.New(cfg)
	svc := &api.ServicesService{Client: client}

	// First, list to get a valid ID
	listResp, err := svc.List(context.Background(), api.ListOptions{Limit: 1})
	if err != nil {
		t.Fatalf("cannot list services: %v", err)
	}
	if len(listResp.Services) == 0 {
		t.Skip("no services available to test GetByID")
	}
	id := listResp.Services[0].ID

	// fetch by that ID
	single, err := svc.GetByID(context.Background(), id)
	if err != nil {
		t.Fatalf("expected no error fetching service %s, got %v", id, err)
	}
	if single.ID != id {
		t.Errorf("expected returned ID %q, got %q", id, single.ID)
	}
}
