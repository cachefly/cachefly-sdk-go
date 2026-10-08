package v2_6

import (
	"context"
	"fmt"
	"net/url"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// EdgeControlKVService handles the edge control key/value stores read by edge
// scripts (/edgecontrol/kv and /edgecontrol/services/{sid}/kv).
//
// There is one store per account and one per service. At the edge the two are
// merged, with service keys overriding account keys. Each store is replaced as
// a whole on write.
type EdgeControlKVService struct {
	Client *httpclient.Client
}

// Edge control KV store scopes.
const (
	EdgeControlKVTypeAccount = "ACCOUNT"
	EdgeControlKVTypeService = "SERVICE"
)

// EdgeControlKV is an edge control key/value store.
type EdgeControlKV struct {
	ID   string `json:"_id"`
	Type string `json:"type"`
	// Data is a flat map; values are strings, numbers (float64 when decoded),
	// booleans or nil.
	Data      map[string]interface{} `json:"data"`
	CreatedAt string                 `json:"createdAt"`
	UpdatedAt string                 `json:"updatedAt"`
}

type edgeControlKVBody struct {
	Data map[string]interface{} `json:"data"`
}

// GetAccount returns the account-level store. The API creates an empty store if
// none exists yet.
func (s *EdgeControlKVService) GetAccount(ctx context.Context) (*EdgeControlKV, error) {
	return s.get(ctx, "/edgecontrol/kv")
}

// ReplaceAccount replaces the account-level store with data. Values must be
// strings, numbers, booleans or nil.
func (s *EdgeControlKVService) ReplaceAccount(ctx context.Context, data map[string]interface{}) (*EdgeControlKV, error) {
	return s.replace(ctx, "/edgecontrol/kv", data)
}

// ClearAccount removes all keys from the account-level store.
func (s *EdgeControlKVService) ClearAccount(ctx context.Context) (*EdgeControlKV, error) {
	return s.clear(ctx, "/edgecontrol/kv")
}

// GetService returns the service-level store. The API creates an empty store if
// none exists yet.
func (s *EdgeControlKVService) GetService(ctx context.Context, serviceID string) (*EdgeControlKV, error) {
	endpoint, err := edgeControlServiceKVEndpoint(serviceID, "")
	if err != nil {
		return nil, err
	}
	return s.get(ctx, endpoint)
}

// ReplaceService replaces the service-level store with data. Values must be
// strings, numbers, booleans or nil.
func (s *EdgeControlKVService) ReplaceService(ctx context.Context, serviceID string, data map[string]interface{}) (*EdgeControlKV, error) {
	endpoint, err := edgeControlServiceKVEndpoint(serviceID, "")
	if err != nil {
		return nil, err
	}
	return s.replace(ctx, endpoint, data)
}

// ClearService removes all keys from the service-level store.
func (s *EdgeControlKVService) ClearService(ctx context.Context, serviceID string) (*EdgeControlKV, error) {
	endpoint, err := edgeControlServiceKVEndpoint(serviceID, "")
	if err != nil {
		return nil, err
	}
	return s.clear(ctx, endpoint)
}

// GetMerged returns the account store merged with the service store, as seen by
// the service's edge scripts.
func (s *EdgeControlKVService) GetMerged(ctx context.Context, serviceID string) (map[string]interface{}, error) {
	endpoint, err := edgeControlServiceKVEndpoint(serviceID, "/merged")
	if err != nil {
		return nil, err
	}

	var merged edgeControlKVBody
	if err := s.Client.Get(ctx, endpoint, &merged); err != nil {
		return nil, err
	}
	return merged.Data, nil
}

func (s *EdgeControlKVService) get(ctx context.Context, endpoint string) (*EdgeControlKV, error) {
	var kv EdgeControlKV
	if err := s.Client.Get(ctx, endpoint, &kv); err != nil {
		return nil, err
	}
	return &kv, nil
}

func (s *EdgeControlKVService) replace(ctx context.Context, endpoint string, data map[string]interface{}) (*EdgeControlKV, error) {
	// The API rejects a null data object.
	if data == nil {
		data = map[string]interface{}{}
	}

	var kv EdgeControlKV
	if err := s.Client.Put(ctx, endpoint, edgeControlKVBody{Data: data}, &kv); err != nil {
		return nil, err
	}
	return &kv, nil
}

func (s *EdgeControlKVService) clear(ctx context.Context, endpoint string) (*EdgeControlKV, error) {
	var kv EdgeControlKV
	if err := s.Client.Delete(ctx, endpoint, &kv); err != nil {
		return nil, err
	}
	return &kv, nil
}

func edgeControlServiceKVEndpoint(serviceID, suffix string) (string, error) {
	if serviceID == "" {
		return "", fmt.Errorf("service ID is required")
	}
	return fmt.Sprintf("/edgecontrol/services/%s/kv%s", url.PathEscape(serviceID), suffix), nil
}
