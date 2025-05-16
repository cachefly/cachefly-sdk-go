package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
)

type Account struct {
	ID                       string   `json:"_id"`
	UpdatedAt                string   `json:"updateAt"`
	CreatedAt                string   `json:"createdAt"`
	CompanyName              string   `json:"companyName"`
	Website                  string   `json:"website"`
	Address1                 string   `json:"address1"`
	Address2                 string   `json:"address2"`
	City                     string   `json:"city"`
	Country                  string   `json:"country"`
	State                    string   `json:"state"`
	Phone                    string   `json:"phone"`
	Zip                      string   `json:"zip"`
	TwoFactorAuthEnabled     bool     `json:"twoFactorAuthEnabled"`
	TwoFactorAuthGracePeriod int      `json:"twoFactorAuthGracePeriod"`
	Users                    []string `json:"users"`
	Services                 []string `json:"services"`
	Origins                  []string `json:"origins"`
	Certificates             []string `json:"certificates"`
}

// communication with the accounts endpoint
type AccountsService struct {
	Client *httpclient.Client
}

type ListAccountsResponse struct {
	Meta     MetaInfo  `json:"meta"`
	Accounts []Account `json:"data"`
}

type ListAccountsOptions struct {
	IsChild      bool
	IsParent     bool
	Status       string
	Offset       int
	Limit        int
	ResponseType string
}

// UpdateAccountRequest is the payload for updating an account.
type UpdateAccountRequest struct {
	CompanyName              string `json:"companyName"`
	Website                  string `json:"website"`
	Address1                 string `json:"address1"`
	Address2                 string `json:"address2"`
	City                     string `json:"city"`
	Country                  string `json:"country"`
	State                    string `json:"state"`
	Phone                    string `json:"phone"`
	Email                    string `json:"email"`
	TwoFactorAuthGracePeriod int    `json:"twoFactorAuthGracePeriod"`
	SAMLRequired             bool   `json:"samlRequired"`
	DefaultDeliveryRegion    string `json:"defaultDeliveryRegion"`
}

// Get current account
func (a *AccountsService) Get(ctx context.Context, responseType string) (*Account, error) {
	endpoint := "/accounts/me"

	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var result Account
	err := a.Client.Get(ctx, fullURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// List all accounts
func (a *AccountsService) List(ctx context.Context, opts ListAccountsOptions) (*ListAccountsResponse, error) {
	endpoint := "/accounts"
	params := url.Values{}

	params.Set("isChild", strconv.FormatBool(opts.IsChild))
	params.Set("isParent", strconv.FormatBool(opts.IsParent))

	if opts.Status != "" {
		params.Set("status", opts.Status)
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

	var result ListAccountsResponse
	err := a.Client.Get(ctx, fullURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Get account by ID
func (a *AccountsService) GetByID(ctx context.Context, id string, responseType string) (*Account, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := "/accounts"

	params := url.Values{}
	params.Set("id", id)

	if responseType != "" {
		params.Set("responseType", responseType)
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var result Account
	err := a.Client.Get(ctx, fullURL, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// UpdateAccountByID updates the fields of an existing account.
// id is required.
func (a *AccountsService) UpdateAccountByID(ctx context.Context, id string, req UpdateAccountRequest) (*Account, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("/accounts/%s", id)

	var updated Account
	if err := a.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}

	return &updated, nil
}

// ActivateAccountByID sends a PUT to /accounts/{id}/activate.
// id is required. Returns the updated Account.
func (a *AccountsService) ActivateAccountByID(ctx context.Context, id string) (*Account, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/accounts/%s/activate", id)

	var updated Account
	// use an empty JSON body so that Content-Type: application/json is set
	if err := a.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeactivateAccountByID sends a PUT to /accounts/{id}/deactivate.
// id is required. Returns the updated Account.
func (a *AccountsService) DeactivateAccountByID(ctx context.Context, id string) (*Account, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/accounts/%s/deactivate", id)

	var updated Account
	if err := a.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// UpdateCurrentAccount updates the authenticated account (me) via PUT /accounts/me.
func (a *AccountsService) UpdateCurrentAccount(ctx context.Context, req UpdateAccountRequest) (*Account, error) {
	endpoint := "/accounts/me"

	var updated Account
	if err := a.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}
