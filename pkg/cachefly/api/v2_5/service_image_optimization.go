package v2_5

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// ServiceImageOptimizationService handles account-related API operations.
type ServiceImageOptimizationService struct {
	Client *httpclient.Client
}

// CreateImageOptimizationOptions specifies the payload for creating an image optimization configuration.
// refere for correct paylaod from updated documentation.
type CreateImageOptimizationOptions struct {
	Enabled        bool     `json:"enabled"`                  // enable/disable optimization
	Formats        []string `json:"formats,omitempty"`        // e.g. ["webp","avif"]
	DefaultQuality int      `json:"defaultQuality,omitempty"` // quality level (0–100)
}

// GetConfiguration fetches the current image optimization configuration (YAML or JSON string).
// GET /services/{id}/imageopt4
func (s *ServiceImageOptimizationService) GetConfiguration(ctx context.Context, serviceID string) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("serviceID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4", serviceID)

	var configStr string
	if err := s.Client.Get(ctx, endpoint, &configStr); err != nil {
		return "", err
	}
	return configStr, nil
}

// CreateConfiguration creates a new configuration; body is YAML or JSON string.
// POST /services/{id}/imageopt4
func (s *ServiceImageOptimizationService) CreateConfiguration(ctx context.Context, serviceID string, configStr CreateImageOptimizationOptions) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("serviceID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4", serviceID)

	var createdStr string
	if err := s.Client.Post(ctx, endpoint, configStr, &createdStr); err != nil {
		return "", err
	}
	return createdStr, nil
}

// UpdateConfiguration updates an existing configuration; body is YAML or JSON string.
// PUT /services/{id}/imageopt4
func (s *ServiceImageOptimizationService) UpdateConfiguration(ctx context.Context, serviceID string, configStr string) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("serviceID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4", serviceID)

	var updatedStr string
	if err := s.Client.Put(ctx, endpoint, configStr, &updatedStr); err != nil {
		return "", err
	}
	return updatedStr, nil
}

// DeleteConfiguration removes the existing configuration.
// DELETE /services/{id}/imageopt4
func (s *ServiceImageOptimizationService) DeleteConfiguration(ctx context.Context, serviceID string) error {
	if serviceID == "" {
		return fmt.Errorf("serviceID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4", serviceID)
	return s.Client.Delete(ctx, endpoint, nil)
}

// GetSchema fetches the validation schema for image optimization config.
// GET /services/{id}/imageopt4/schema
func (s *ServiceImageOptimizationService) GetSchema(ctx context.Context, serviceID string) (map[string]interface{}, error) {
	if serviceID == "" {
		return nil, fmt.Errorf("serviceID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4/schema", serviceID)

	var schema map[string]interface{}
	if err := s.Client.Get(ctx, endpoint, &schema); err != nil {
		return nil, err
	}
	return schema, nil
}

// GetDefaults fetches the default config for image optimization.
// GET /services/{id}/imageopt4/defaults
func (s *ServiceImageOptimizationService) GetDefaults(ctx context.Context, serviceID string) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("serviceID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4/default", serviceID)

	var defStr string
	if err := s.Client.Get(ctx, endpoint, &defStr); err != nil {
		return "", err
	}
	return defStr, nil
}

func (s *ServiceImageOptimizationService) GetDetail(ctx context.Context, serviceID string) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("serviceID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4/details", serviceID)

	var exStr string
	if err := s.Client.Get(ctx, endpoint, &exStr); err != nil {
		return "", err
	}
	return exStr, nil
}

// ValidateConfiguration validates a config string against the schema.
// POST /services/{id}/imageopt4/validate
func (s *ServiceImageOptimizationService) ValidateConfiguration(ctx context.Context, serviceID string, configStr string) (map[string]interface{}, error) {
	if serviceID == "" {
		return nil, fmt.Errorf("serviceID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4/validate", serviceID)

	var result map[string]interface{}
	if err := s.Client.Post(ctx, endpoint, configStr, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ActivateConfiguration enables the image optimization configuration for a service.
func (s *ServiceImageOptimizationService) ActivateConfiguration(ctx context.Context, serviceID string) error {
	if serviceID == "" {
		return fmt.Errorf("service ID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4/activate", url.PathEscape(serviceID))

	emptyBody := struct{}{}

	// Perform POST with no request body
	if err := s.Client.Put(ctx, endpoint, emptyBody, nil); err != nil {
		return err
	}
	return nil
}

// DeactivateConfiguration disables the image optimization configuration for a service.
func (s *ServiceImageOptimizationService) DeactivateConfiguration(ctx context.Context, serviceID string) error {
	if serviceID == "" {
		return fmt.Errorf("service ID is required")
	}
	endpoint := fmt.Sprintf("/services/%s/imageopt4/deactivate", url.PathEscape(serviceID))
	// Perform PUT with no request body
	if err := s.Client.Put(ctx, endpoint, nil, nil); err != nil {
		return err
	}
	return nil
}
