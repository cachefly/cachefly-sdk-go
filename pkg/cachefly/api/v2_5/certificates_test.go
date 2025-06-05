package v2_5

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// CREATE - Test Create method
func TestCertificatesService_Create(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/certificates" {
			t.Errorf("Expected path /api/2.5/certificates, got %s", r.URL.Path)
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"_id":"cert-123","subjectCommonName":"example.com","managed":false}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &CertificatesService{Client: client}

	req := CreateCertificateRequest{
		Certificate:    "-----BEGIN CERTIFICATE-----\nMIIExample...\n-----END CERTIFICATE-----",
		CertificateKey: "-----BEGIN PRIVATE KEY-----\nMIIExample...\n-----END PRIVATE KEY-----",
	}
	result, err := svc.Create(context.Background(), req)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "cert-123" {
		t.Errorf("Expected certificate ID cert-123, got %s", result.ID)
	}
}

// READ - Test List method
func TestCertificatesService_List(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/certificates" {
			t.Errorf("Expected path /api/2.5/certificates, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"meta":{"limit":10,"offset":0,"count":1},"data":[{"_id":"cert-123","subjectCommonName":"example.com"}]}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &CertificatesService{Client: client}

	opts := ListCertificatesOptions{Limit: 10}
	result, err := svc.List(context.Background(), opts)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(result.Certificates) != 1 {
		t.Errorf("Expected 1 certificate, got %d", len(result.Certificates))
	}
	if result.Certificates[0].ID != "cert-123" {
		t.Errorf("Expected certificate ID cert-123, got %s", result.Certificates[0].ID)
	}
}

// READ - Test GetByID method
func TestCertificatesService_GetByID(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/certificates/cert-123" {
			t.Errorf("Expected path /api/2.5/certificates/cert-123, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected GET method, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"_id":"cert-123","subjectCommonName":"example.com","managed":false}`))
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &CertificatesService{Client: client}

	result, err := svc.GetByID(context.Background(), "cert-123", "")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.ID != "cert-123" {
		t.Errorf("Expected certificate ID cert-123, got %s", result.ID)
	}
}

// DELETE - Test Delete method
func TestCertificatesService_Delete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/2.5/certificates/cert-123" {
			t.Errorf("Expected path /api/2.5/certificates/cert-123, got %s", r.URL.Path)
		}
		if r.Method != "DELETE" {
			t.Errorf("Expected DELETE method, got %s", r.Method)
		}

		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()

	cfg := httpclient.Config{BaseURL: server.URL + "/api/2.5", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &CertificatesService{Client: client}

	err := svc.Delete(context.Background(), "cert-123")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

// Error handling test - missing required fields and ID
func TestCertificatesService_ErrorHandling(t *testing.T) {
	cfg := httpclient.Config{BaseURL: "http://test.com", AuthToken: "test-token"}
	client := httpclient.New(cfg)
	svc := &CertificatesService{Client: client}

	// Test missing required fields for Create
	req := CreateCertificateRequest{Certificate: "cert-data"} // Missing CertificateKey
	_, err := svc.Create(context.Background(), req)

	if err == nil {
		t.Error("Expected error for missing required fields")
	}
	if err.Error() != "certificate and certificateKey are required" {
		t.Errorf("Expected 'certificate and certificateKey are required' error, got %s", err.Error())
	}

	// Test missing ID for GetByID
	_, err = svc.GetByID(context.Background(), "", "")
	if err == nil {
		t.Error("Expected error for missing ID")
	}
	if err.Error() != "id is required" {
		t.Errorf("Expected 'id is required' error, got %s", err.Error())
	}
}
