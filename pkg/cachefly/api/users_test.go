package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
	"github.com/avvvet/cachefly-sdk-go/pkg/cachefly/api"
)

func TestUsersService_List(t *testing.T) {
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
	svc := &api.UsersService{Client: client}

	// List users with pagination options
	resp, err := svc.List(context.Background(), api.ListUsersOptions{
		Offset:       0,
		Limit:        10,
		ResponseType: "shallow",
	})
	if err != nil {
		t.Fatalf("expected no error listing users, got %v", err)
	}

	if len(resp.Users) == 0 {
		t.Skip("no users found in account; skipping further assertions")
	}

	// Basic assertions on the first user
	user := resp.Users[0]
	if user.ID == "" {
		t.Errorf("expected user ID to be set, got empty string")
	}
	if user.Email == "" {
		t.Errorf("expected user Email to be set, got empty string")
	}

}
