package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
)

func TestServiceOptionsService_GetBasicOptions(t *testing.T) {
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
	svcSvc := &api.ServicesService{Client: client}
	svcList, err := svcSvc.List(context.Background(), api.ListOptions{Limit: 1})
	if err != nil {
		t.Fatalf("cannot list services: %v", err)
	}
	if len(svcList.Services) == 0 {
		t.Skip("no services available to test GetBasicOptions")
	}
	serviceID := svcList.Services[0].ID

	//Call GetBasicOptions for that service
	optSvc := &api.ServiceOptionsService{Client: client}
	opts, err := optSvc.GetBasicOptions(context.Background(), serviceID)
	if err != nil {
		t.Fatalf("expected no error getting basic options for service %s, got %v", serviceID, err)
	}

	// Basic assertions
	if opts == nil {
		t.Fatal("expected non-nil ServiceOptions, got nil")
	}

	//(Optional) for inspection
	//_ = json.NewEncoder(os.Stdout).Encode(opts)
}
