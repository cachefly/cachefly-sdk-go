package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

type CertificatesService struct {
	Client *httpclient.Client
}

// TLS/SSL certificate in CacheFly.
type Certificate struct {
	ID                string   `json:"_id"`
	CreatedAt         string   `json:"createdAt"`
	SubjectCommonName string   `json:"subjectCommonName"`
	SubjectNames      []string `json:"subjectNames"`
	Expired           bool     `json:"expired"`
	Expiring          bool     `json:"expiring"`
	InUse             bool     `json:"inUse"`
	Managed           bool     `json:"managed"`
	Services          []string `json:"services"`
	Domains           []string `json:"domains"`
	NotBefore         string   `json:"notBefore"`
	NotAfter          string   `json:"notAfter"`
}

type ListCertificatesResponse struct {
	Meta         MetaInfo      `json:"meta"`
	Certificates []Certificate `json:"data"`
}

type ListCertificatesOptions struct {
	ResponseType string
	Search       string
	Offset       int
	Limit        int
}

type CreateCertificateRequest struct {
	Certificate    string `json:"certificate"`        // required, PEM-encoded certificate
	CertificateKey string `json:"certificateKey"`     // required, PEM-encoded private key
	Password       string `json:"password,omitempty"` // optional password for key
}

// List fetches all certificates with optional filters.
func (s *CertificatesService) List(ctx context.Context, opts ListCertificatesOptions) (*ListCertificatesResponse, error) {
	endpoint := "/certificates"
	params := url.Values{}
	if opts.ResponseType != "" {
		params.Set("responseType", opts.ResponseType)
	}
	if opts.Search != "" {
		params.Set("search", opts.Search)
	}
	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}
	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var resp ListCertificatesResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create uploads a new certificate.
func (s *CertificatesService) Create(ctx context.Context, req CreateCertificateRequest) (*Certificate, error) {
	// validate required fields
	if req.Certificate == "" || req.CertificateKey == "" {
		return nil, fmt.Errorf("certificate and certificateKey are required")
	}
	endpoint := "/certificates"

	var created Certificate
	if err := s.Client.Post(ctx, endpoint, req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// GetByID retrieves a certificate by its ID.
func (s *CertificatesService) GetByID(ctx context.Context, id, responseType string) (*Certificate, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/certificates/%s", id)
	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var cert Certificate
	if err := s.Client.Get(ctx, fullURL, &cert); err != nil {
		return nil, err
	}
	return &cert, nil
}

// Delete removes a certificate by ID.
func (s *CertificatesService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/certificates/%s", id)
	return s.Client.Delete(ctx, endpoint, nil)
}
