package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/avvvet/cachefly-sdk-go/internal/httpclient"
)

type AccountsService struct {
	Client *httpclient.Client
}

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

/*
It is expected that most companies will need only one account.
The concept of sub accounts found in the older version 1 API has been directly replaced by the concept of services.
Opening multiple accounts is only required to separate invoicing concerns.
[as per the API doc]
*/
type CreateChildAccountRequest struct {
	CompanyName string `json:"companyName"` // required
	Username    string `json:"username"`    // required, 8–32 chars, lowercase letters/numbers, dashes
	Password    string `json:"password"`    // required, ≥8 chars
	FullName    string `json:"fullName"`    // required, ≥2 chars
	Email       string `json:"email"`       // required, must be a valid email
	Website     string `json:"website,omitempty"`
	Address1    string `json:"address1,omitempty"`
	Address2    string `json:"address2,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	State       string `json:"state,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Zip         string `json:"zip,omitempty"`
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

// GetByID retrieves an account by its ID, optionally specifying a responseType.
func (a *AccountsService) GetByID(ctx context.Context, id string, responseType string) (*Account, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	// 1) Base path with path‐param
	endpoint := fmt.Sprintf("/accounts/%s", url.PathEscape(id))

	// 2) Build up any query params
	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}

	// 3) Append the ?query=string only if we have params
	fullURL := endpoint
	if len(params) > 0 {
		fullURL = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	// 4) Call into the HTTP client and unmarshal directly into Account
	var result Account
	if err := a.Client.Get(ctx, fullURL, &result); err != nil {
		return nil, err
	}
	return &result, nil
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

// CreateChildAccount creates a new child account (POST /accounts)
func (a *AccountsService) CreateChildAccount(ctx context.Context, req CreateChildAccountRequest) (*Account, error) {
	// Validate required fields
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

// ChildAccountAuthResponse is the JSON returned by POST /accounts/{id}/auth
type ChildAccountAuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
}

// parent account can request GetChildAccountAuthToken requests of a child-account.
// then using this token, all other service mangement can be requested for the child account
// POST /accounts/{id}/auth
func (a *AccountsService) GetChildAccountAuthToken(ctx context.Context, id string) (*ChildAccountAuthResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	// 1) Build the path
	endpoint := fmt.Sprintf("/accounts/%s/auth", url.PathEscape(id))

	// 2) Call POST with an empty JSON body
	var resp ChildAccountAuthResponse
	if err := a.Client.Post(ctx, endpoint, struct{}{}, &resp); err != nil {
		return nil, err
	}

	// 3) Return token + expiry
	return &resp, nil
}
