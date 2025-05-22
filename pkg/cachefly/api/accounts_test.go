package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
	"github.com/joho/godotenv"
)

func loadEnv(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Log(".env not found, relying on existing environment variables")
	}
}

func TestAccountsService_Get(t *testing.T) {
	loadEnv(t)

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		t.Skip("CACHEFLY_API_TOKEN not set")
	}

	config := httpclient.Config{
		BaseURL:   "https://api.cachefly.com/api/2.5",
		AuthToken: token,
	}
	client := httpclient.New(config)
	service := &api.AccountsService{Client: client}

	account, err := service.Get(context.Background(), "shallow")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if account.ID == "" {
		t.Errorf("expected account ID, got empty")
	}
}

func TestAccountsService_List(t *testing.T) {
	loadEnv(t)

	token := os.Getenv("CACHEFLY_API_TOKEN")
	if token == "" {
		t.Skip("CACHEFLY_API_TOKEN not set")
	}

	config := httpclient.Config{
		BaseURL:   "https://api.cachefly.com/api/2.5",
		AuthToken: token,
	}
	client := httpclient.New(config)
	service := &api.AccountsService{Client: client}

	opts := api.ListAccountsOptions{
		IsChild:      false,
		IsParent:     false,
		Status:       "ACTIVE",
		Offset:       0,
		Limit:        10,
		ResponseType: "shallow",
	}

	resp, err := service.List(context.Background(), opts)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(resp.Accounts) == 0 {
		t.Log("no accounts found — verify with live data or staging API")
	}

	if resp.Meta.Count < 0 {
		t.Errorf("expected non-negative count, got %d", resp.Meta.Count)
	}
}
