package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
)

func TestTLSProfilesService_List(t *testing.T) {
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
	svc := &api.TLSProfilesService{Client: client}

	// List TLS profiles with pagination & grouping options
	resp, err := svc.List(context.Background(), api.ListTLSProfilesOptions{
		Offset: 0,
		Limit:  10,
		Group:  "",
	})
	if err != nil {
		t.Fatalf("expected no error listing TLS profiles, got %v", err)
	}

	if len(resp.Profiles) == 0 {
		t.Skip("no TLS profiles found; skipping further assertions")
	}

	// Basic assertions on the first profile
	prof := resp.Profiles[0]
	if prof.ID == "" {
		t.Errorf("expected profile ID to be set, got empty string")
	}
	if prof.Name == "" {
		t.Errorf("expected profile Name to be set, got empty string")
	}

}
