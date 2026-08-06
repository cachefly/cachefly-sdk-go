package v2_6

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// LogTarget represents a CacheFly log target configuration.
//
// The fields are a union across all log target types (S3_BUCKET,
// GOOGLE_BUCKET, AZURE_BLOB, HTTP); only the fields relevant to Type are
// populated by the API.
type LogTarget struct {
	ID        string  `json:"_id"`
	UpdatedAt string  `json:"updatedAt"`
	CreatedAt string  `json:"createdAt"`
	Type      string  `json:"type"`
	Name      *string `json:"name,omitempty"`

	// Common log delivery options.
	Format      *string `json:"format,omitempty"`      // "JSON" | "NDJSON"
	Compression *string `json:"compression,omitempty"` // "NONE" | "GZIP" | "ZSTD"
	Sampling    *int    `json:"sampling,omitempty"`    // 0..100

	// S3_BUCKET fields.
	Endpoint         *string `json:"endpoint,omitempty"`
	Region           *string `json:"region,omitempty"`
	Bucket           *string `json:"bucket,omitempty"` // also used by GOOGLE_BUCKET
	AccessKey        *string `json:"accessKey,omitempty"`
	SecretKey        *string `json:"secretKey,omitempty"`
	SignatureVersion *string `json:"signatureVersion,omitempty"` // always "v4"

	// GOOGLE_BUCKET fields.
	JsonKey *string `json:"jsonKey,omitempty"`

	// AZURE_BLOB fields.
	EndpointProtocol *string `json:"endpointProtocol,omitempty"` // "HTTP" | "HTTPS"
	EndpointSuffix   *string `json:"endpointSuffix,omitempty"`
	AccountName      *string `json:"accountName,omitempty"`
	AccountKey       *string `json:"accountKey,omitempty"`
	ContainerName    *string `json:"containerName,omitempty"`
	Prefix           *string `json:"prefix,omitempty"`

	// HTTP fields.
	Uri      *string `json:"uri,omitempty"`
	Method   *string `json:"method,omitempty"` // "POST" | "PUT"
	Auth     *string `json:"auth,omitempty"`   // "NONE" | "BASIC" | "BEARER"
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Token    *string `json:"token,omitempty"`

	// Services with logging enabled. These are not documented in the current
	// API responses; they are populated only when the API returns them.
	AccessLogsServices *[]string `json:"accessLogsServices,omitempty"`
	OriginLogsServices *[]string `json:"originLogsServices,omitempty"`
}

// ListLogTargetsResponse contains paginated log target results.
type ListLogTargetsResponse struct {
	Meta       MetaInfo    `json:"meta"`
	LogTargets []LogTarget `json:"data"`
}

// ListLogTargetsOptions allows filtering & pagination for log targets.
type ListLogTargetsOptions struct {
	Type         string
	Offset       int
	Limit        int
	ResponseType string
}

// CreateLogTargetRequest contains the fields for creating a new log target.
//
// Type is required. The remaining fields depend on the chosen type:
//   - S3_BUCKET: region, bucket, accessKey, secretKey required; endpoint,
//     signatureVersion optional.
//   - GOOGLE_BUCKET: bucket, jsonKey required.
//   - AZURE_BLOB: accountName, accountKey, containerName required;
//     endpointProtocol, endpointSuffix, prefix optional.
//   - HTTP: uri required; method, auth, username, password, token optional.
//
// name, format, compression and sampling are optional for every type. The
// API rejects fields that do not belong to the chosen type.
type CreateLogTargetRequest struct {
	Type string  `json:"type"`
	Name *string `json:"name,omitempty"`

	// Common log delivery options.
	Format      *string `json:"format,omitempty"`
	Compression *string `json:"compression,omitempty"`
	Sampling    *int    `json:"sampling,omitempty"`

	// S3_BUCKET fields.
	Endpoint         *string `json:"endpoint,omitempty"`
	Region           *string `json:"region,omitempty"`
	Bucket           *string `json:"bucket,omitempty"` // also used by GOOGLE_BUCKET
	AccessKey        *string `json:"accessKey,omitempty"`
	SecretKey        *string `json:"secretKey,omitempty"`
	SignatureVersion *string `json:"signatureVersion,omitempty"`

	// GOOGLE_BUCKET fields.
	JsonKey *string `json:"jsonKey,omitempty"`

	// AZURE_BLOB fields.
	EndpointProtocol *string `json:"endpointProtocol,omitempty"`
	EndpointSuffix   *string `json:"endpointSuffix,omitempty"`
	AccountName      *string `json:"accountName,omitempty"`
	AccountKey       *string `json:"accountKey,omitempty"`
	ContainerName    *string `json:"containerName,omitempty"`
	Prefix           *string `json:"prefix,omitempty"`

	// HTTP fields.
	Uri      *string `json:"uri,omitempty"`
	Method   *string `json:"method,omitempty"`
	Auth     *string `json:"auth,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Token    *string `json:"token,omitempty"`
}

// UpdateLogTargetRequest contains the fields for updating an existing log
// target. All fields are optional; only the provided fields are changed.
//
// Services logging is managed separately via SetLogging: the update endpoint
// rejects accessLogsServices/originLogsServices.
type UpdateLogTargetRequest struct {
	Type *string `json:"type,omitempty"`
	Name *string `json:"name,omitempty"`

	// Common log delivery options.
	Format      *string `json:"format,omitempty"`
	Compression *string `json:"compression,omitempty"`
	Sampling    *int    `json:"sampling,omitempty"`

	// S3_BUCKET fields.
	Endpoint         *string `json:"endpoint,omitempty"`
	Region           *string `json:"region,omitempty"`
	Bucket           *string `json:"bucket,omitempty"` // also used by GOOGLE_BUCKET
	AccessKey        *string `json:"accessKey,omitempty"`
	SecretKey        *string `json:"secretKey,omitempty"`
	SignatureVersion *string `json:"signatureVersion,omitempty"`

	// GOOGLE_BUCKET fields.
	JsonKey *string `json:"jsonKey,omitempty"`

	// AZURE_BLOB fields.
	EndpointProtocol *string `json:"endpointProtocol,omitempty"`
	EndpointSuffix   *string `json:"endpointSuffix,omitempty"`
	AccountName      *string `json:"accountName,omitempty"`
	AccountKey       *string `json:"accountKey,omitempty"`
	ContainerName    *string `json:"containerName,omitempty"`
	Prefix           *string `json:"prefix,omitempty"`

	// HTTP fields.
	Uri      *string `json:"uri,omitempty"`
	Method   *string `json:"method,omitempty"`
	Auth     *string `json:"auth,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Token    *string `json:"token,omitempty"`
}

// SetLoggingRequest contains the services to set logging for.
type SetLoggingRequest struct {
	AccessLogsServices []string `json:"accessLogsServices"`
	OriginLogsServices []string `json:"originLogsServices"`
}

// LogTargetsService handles log target-related API operations.
type LogTargetsService struct {
	Client *httpclient.Client
}

// List returns all log targets for the current account.
func (s *LogTargetsService) List(ctx context.Context, opts ListLogTargetsOptions) (*ListLogTargetsResponse, error) {
	endpoint := "/logtargets"

	params := url.Values{}

	if opts.Type != "" {
		params.Set("type", opts.Type)
	}

	if opts.Offset >= 0 {
		params.Set("offset", strconv.Itoa(opts.Offset))
	}

	if opts.Limit > 0 {
		params.Set("limit", strconv.Itoa(opts.Limit))
	}

	if opts.ResponseType != "" {
		params.Set("responseType", opts.ResponseType)
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	var resp ListLogTargetsResponse
	if err := s.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create creates a new log target.
func (s *LogTargetsService) Create(ctx context.Context, req CreateLogTargetRequest) (*LogTarget, error) {
	endpoint := "/logtargets"

	var created LogTarget
	if err := s.Client.Post(ctx, endpoint, req, &created); err != nil {
		return nil, err
	}

	return &created, nil
}

// UpdateByID updates an existing log target by its ID.
func (s *LogTargetsService) UpdateByID(ctx context.Context, id string, req UpdateLogTargetRequest) (*LogTarget, error) {
	if id == "" {
		return nil, fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s", id)

	var updated LogTarget
	if err := s.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// GetByID retrieves a log target by its ID.
func (s *LogTargetsService) GetByID(ctx context.Context, id string) (*LogTarget, error) {
	if id == "" {
		return nil, fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s", id)

	var logTarget LogTarget
	if err := s.Client.Get(ctx, endpoint, &logTarget); err != nil {
		return nil, err
	}
	return &logTarget, nil
}

// DeleteByID deletes a log target by its ID.
func (s *LogTargetsService) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s", id)
	return s.Client.Delete(ctx, endpoint, nil)
}

// SetLogging sets services logging for a log target.
func (s *LogTargetsService) SetLogging(ctx context.Context, id string, req SetLoggingRequest) (*LogTarget, error) {
	if id == "" {
		return nil, fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s/logging", id)

	var result LogTarget
	if err := s.Client.Put(ctx, endpoint, req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// TestConnection tests the connection to the log target with the given id
// (POST /logtargets/{id}/test). A nil error means the connection test
// succeeded.
func (s *LogTargetsService) TestConnection(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("log target ID is required")
	}
	endpoint := fmt.Sprintf("/logtargets/%s/test", id)

	// The success response body is unspecified and may be empty; tolerate an
	// empty body but surface API errors.
	var out map[string]interface{}
	if err := s.Client.Post(ctx, endpoint, struct{}{}, &out); err != nil && err != io.EOF {
		return err
	}
	return nil
}
