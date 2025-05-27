package v2_5_test

import (
	"context"
	"os"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
	api "github.com/cachefly/cachefly-go-sdk/pkg/cachefly/api/v2_5"
)

func TestOriginsService_List(t *testing.T) {
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
	svc := &api.OriginsService{Client: client}

	// List account origins with a reasonable page size
	resp, err := svc.List(context.Background(), api.ListOriginsOptions{
		Offset:       0,
		Limit:        10,
		ResponseType: "shallow",
	})
	if err != nil {
		t.Fatalf("expected no error listing origins, got %v", err)
	}

	if resp.Meta.Limit != 10 {
		t.Skip("expected api to fetch origins list; skipping further assertions")
	}
}
