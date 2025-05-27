package v2_5_test

import (
	"context"
	"os"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
	api "github.com/cachefly/cachefly-go-sdk/pkg/cachefly/api/v2_5"
)

func TestServiceDomainsService_List(t *testing.T) {
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

	// 1) First list Services to grab a valid serviceID
	svcSvc := &api.ServicesService{Client: client}
	svcList, err := svcSvc.List(context.Background(), api.ListOptions{Limit: 1})
	if err != nil {
		t.Fatalf("cannot list services: %v", err)
	}
	if len(svcList.Services) == 0 {
		t.Skip("no services available to test Service Domains")
	}
	serviceID := svcList.Services[0].ID

	// 2) Now list the domains for that service
	domSvc := &api.ServiceDomainsService{Client: client}
	domList, err := domSvc.List(context.Background(), serviceID, api.ListServiceDomainsOptions{Limit: 10})
	if err != nil {
		t.Fatalf("expected no error listing domains: %v", err)
	}
	if len(domList.Domains) == 0 {
		t.Fatalf("expected at least one domain for service %s, got none", serviceID)
	}
}
