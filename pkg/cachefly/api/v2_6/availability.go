package v2_6

import (
	"context"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// AvailabilityService provides endpoints to check name availability across resources.
type AvailabilityService struct {
	Client *httpclient.Client
}

// availabilityResponse represents the common response for availability checks.
type availabilityResponse struct {
	Available bool `json:"available"`
}

// CheckDomainRequest is the payload to check domain availability.
type CheckDomainRequest struct {
	Name string `json:"name"`
}

// CheckUserRequest is the payload to check username availability.
type CheckUserRequest struct {
	Username string `json:"username"`
}

// CheckServiceRequest is the payload to check service uniqueName availability.
type CheckServiceRequest struct {
	UniqueName string `json:"uniqueName"`
}

// CheckSAMLRequest is the payload to check SAML configuration name availability.
type CheckSAMLRequest struct {
	Name string `json:"name"`
}

// Domains checks if a domain name is available.
func (s *AvailabilityService) Domains(ctx context.Context, req CheckDomainRequest) (bool, error) {
	var out availabilityResponse
	if err := s.Client.Post(ctx, "/availability/domains", req, &out); err != nil {
		return false, err
	}
	return out.Available, nil
}

// Users checks if a username is available.
func (s *AvailabilityService) Users(ctx context.Context, req CheckUserRequest) (bool, error) {
	var out availabilityResponse
	if err := s.Client.Post(ctx, "/availability/users", req, &out); err != nil {
		return false, err
	}
	return out.Available, nil
}

// Services checks if a service uniqueName is available.
func (s *AvailabilityService) Services(ctx context.Context, req CheckServiceRequest) (bool, error) {
	var out availabilityResponse
	if err := s.Client.Post(ctx, "/availability/services", req, &out); err != nil {
		return false, err
	}
	return out.Available, nil
}

// SAML checks if a SAML configuration name is available.
func (s *AvailabilityService) SAML(ctx context.Context, req CheckSAMLRequest) (bool, error) {
	var out availabilityResponse
	if err := s.Client.Post(ctx, "/availability/saml", req, &out); err != nil {
		return false, err
	}
	return out.Available, nil
}
