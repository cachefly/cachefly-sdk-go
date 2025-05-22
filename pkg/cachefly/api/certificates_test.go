package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
)

func TestCertificatesService_List(t *testing.T) {
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
	svc := &api.CertificatesService{Client: client}

	// List certificates with a small page size
	resp, err := svc.List(context.Background(), api.ListCertificatesOptions{Offset: 0, Limit: 10})
	if err != nil {
		t.Fatalf("expected no error listing certificates, got %v", err)
	}

	if len(resp.Certificates) == 0 {
		t.Skip("no certificates found in account; skipping further assertions")
	}

	// Basic assertions on the first certificate
	cert := resp.Certificates[0]
	if cert.ID == "" {
		t.Errorf("expected certificate ID to be set, got empty string")
	}
	if len(cert.Domains) > 0 {
		t.Errorf("expected certificate Domain names, got empty string")
	}
}
