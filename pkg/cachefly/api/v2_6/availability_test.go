package v2_6

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

func TestAvailability_Domains(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/availability/domains" {
			t.Errorf("Expected path /api/2.6/availability/domains, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"available":true}`))
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "token"})
	svc := &AvailabilityService{Client: client}

	available, err := svc.Domains(context.Background(), CheckDomainRequest{Name: "example.com"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !available {
		t.Errorf("Expected available=true")
	}
}

func TestAvailability_Users(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/availability/users" {
			t.Errorf("Expected path /api/2.6/availability/users, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"available":false}`))
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "token"})
	svc := &AvailabilityService{Client: client}

	available, err := svc.Users(context.Background(), CheckUserRequest{Username: "test-user"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if available {
		t.Errorf("Expected available=false")
	}
}

func TestAvailability_Services(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/availability/services" {
			t.Errorf("Expected path /api/2.6/availability/services, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"available":true}`))
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "token"})
	svc := &AvailabilityService{Client: client}

	available, err := svc.Services(context.Background(), CheckServiceRequest{UniqueName: "svc-name"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !available {
		t.Errorf("Expected available=true")
	}
}

func TestAvailability_SAML(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.6/availability/saml" {
			t.Errorf("Expected path /api/2.6/availability/saml, got %s", r.URL.Path)
		}
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST method, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"available":false}`))
	}))
	defer server.Close()

	client := httpclient.New(httpclient.Config{BaseURL: server.URL + "/api/2.6", AuthToken: "token"})
	svc := &AvailabilityService{Client: client}

	available, err := svc.SAML(context.Background(), CheckSAMLRequest{Name: "sso-name"})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if available {
		t.Errorf("Expected available=false")
	}
}
