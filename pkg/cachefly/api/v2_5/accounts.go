// Package v2_5 provides types and services for CacheFly API v2.5.
package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// AccountsService handles account-related API operations.
type AccountsService struct {
	Client *httpclient.Client
}

// Account represents a CacheFly account with all configuration and metadata.
type Account struct {
	ID          string `json:"_id"`
	Uid         int    `json:"uid"`
	CompanyName string `json:"companyName"`
	Status      string `json:"status"`

	// Parent relationship
	Parent   *string `json:"parent"`
	IsParent bool    `json:"isParent"`
	IsChild  bool    `json:"isChild"`

	// Signup and user info
	SignupCountry string `json:"signupCountry"`
	SignupIp      string `json:"signupIp"`
	SignupDate    string `json:"signupDate"`
	Email         string `json:"email"`

	// Timestamps
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	// Two-factor authentication
	TwoFactorAuthEnabled     bool `json:"twoFactorAuthEnabled"`
	TwoFactorAuthGracePeriod int  `json:"twoFactorAuthGracePeriod"`

	// TLS and SSL settings
	DefaultTlsProfile           string `json:"defaultTlsProfile"`
	AutoSslEnabled              bool   `json:"autoSslEnabled"`
	DefaultAutoSsl              bool   `json:"defaultAutoSsl"`
	DefaultDomainValidationMode string `json:"defaultDomainValidationMode"`
	CertificatesEnabled         bool   `json:"certificatesEnabled"`

	// SAML settings
	SamlEnabled  bool `json:"samlEnabled"`
	SamlRequired bool `json:"samlRequired"`
	SamlActive   bool `json:"samlActive"`

	// Caching and delivery
	DefaultDeliveryRegion string `json:"defaultDeliveryRegion"`
	CacheWarmingEnabled   bool   `json:"cacheWarmingEnabled"`

	// Storage limits
	MaxNumOfActiveStorageAccounts int `json:"maxNumOfActiveStorageAccounts"`

	// Address info
	Website  string `json:"website"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	State    string `json:"state"`
	Country  string `json:"country"`
	Zip      string `json:"zip"`
	Phone    string `json:"phone"`

	// Relations and collections
	Origins                 []string `json:"origins"`
	Certificates            []string `json:"certificates"`
	ScriptConfigDefinitions []string `json:"scriptConfigDefinitions"`
	LogTargets              []string `json:"logTargets"`
	Users                   []string `json:"users"`
	Services                []string `json:"services"`
	V1Account               bool     `json:"v1Account"`
}

// CreateChildAccountRequest contains the required fields for creating a child account.
// Most companies will need only one account. Child accounts are only required
// to separate invoicing concerns. refer api doc for more detail.
type CreateChildAccountRequest struct {
	CompanyName string `json:"companyName"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	Website     string `json:"website,omitempty"`
	Address1    string `json:"address1,omitempty"`
	Address2    string `json:"address2,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	State       string `json:"state,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Zip         string `json:"zip,omitempty"`
}

// ListAccountsResponse contains paginated account results.
type ListAccountsResponse struct {
	Meta     MetaInfo  `json:"meta"`
	Accounts []Account `json:"data"`
}

// ListAccountsOptions specifies filters and pagination for listing accounts.
type ListAccountsOptions struct {
	IsChild      bool
	IsParent     bool
	Status       string
	Offset       int
	Limit        int
	ResponseType string
}

// UpdateAccountRequest contains fields for updating an existing account.
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

// ChildAccountAuthResponse contains authentication token for child account access.
type ChildAccountAuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
}

// Get retrieves the current authenticated account.
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

// List retrieves accounts with optional filtering and pagination.
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

// GetByID retrieves an account by its ID.
func (a *AccountsService) GetByID(ctx context.Context, id string, responseType string) (*Account, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("/accounts/%s", url.PathEscape(id))

	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}

	fullURL := endpoint
	if len(params) > 0 {
		fullURL = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var result Account
	if err := a.Client.Get(ctx, fullURL, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateCurrentAccount updates the authenticated account.
func (a *AccountsService) UpdateCurrentAccount(ctx context.Context, req UpdateAccountRequest) (*Account, error) {
	endpoint := "/accounts/me"

	var updated Account
	if err := a.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// UpdateAccountByID updates an existing account by ID.
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

// ActivateAccountByID activates an account.
func (a *AccountsService) ActivateAccountByID(ctx context.Context, id string) (*Account, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/accounts/%s/activate", id)

	var updated Account
	if err := a.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeactivateAccountByID deactivates an account.
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

// CreateChildAccount creates a new child account.
func (a *AccountsService) CreateChildAccount(ctx context.Context, req CreateChildAccountRequest) (*Account, error) {
	if req.CompanyName == "" || req.Username == "" || req.Password == "" ||
		req.FullName == "" || req.Email == "" {
		return nil, fmt.Errorf("companyName, username, password, fullName and email are required")
	}

	var created Account
	if err := a.Client.Post(ctx, "/accounts", req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// GetChildAccountAuthToken generates an authentication token for a child account.
// Parent accounts can use this token to manage child account services.
func (a *AccountsService) GetChildAccountAuthToken(ctx context.Context, id string) (*ChildAccountAuthResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("/accounts/%s/auth", url.PathEscape(id))

	var resp ChildAccountAuthResponse
	if err := a.Client.Post(ctx, endpoint, struct{}{}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// Enable2FAForCurrentAccount enables two-factor authentication for the current account.
func (a *AccountsService) Enable2FAForCurrentAccount(ctx context.Context) (*Account, error) {
	endpoint := "/accounts/me/enable2FA"

	var updated Account
	if err := a.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// Disable2FAForCurrentAccount disables two-factor authentication for the current account.
func (a *AccountsService) Disable2FAForCurrentAccount(ctx context.Context) (*Account, error) {
	endpoint := "/accounts/me/disable2FA"

	var updated Account
	if err := a.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}
