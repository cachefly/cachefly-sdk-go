package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// User is the representation of a CacheFly user.
type User struct {
	ID                     string   `json:"_id"`
	UpdatedAt              string   `json:"updateAt"`
	CreatedAt              string   `json:"createdAt"`
	Username               string   `json:"username"`
	PasswordChangeRequired bool     `json:"passwordChangeRequired"`
	Email                  string   `json:"email"`
	FullName               string   `json:"fullName"`
	Phone                  string   `json:"phone"`
	Permissions            []string `json:"permissions"`
	Services               []string `json:"services"`
	Status                 string   `json:"status"`
}

// ListUsersOptions allows filtering & pagination when listing users.
type ListUsersOptions struct {
	Search       string
	Offset       int
	Limit        int
	ResponseType string
}

// ListUsersResponse is the paged response for listing users.
type ListUsersResponse struct {
	Meta  MetaInfo `json:"meta"`
	Users []User   `json:"data"`
}

// CreateUserRequest is the payload to create a new user.
type CreateUserRequest struct {
	Username               string   `json:"username"`
	Password               string   `json:"password"`
	Services               []string `json:"services"`
	PasswordChangeRequired bool     `json:"passwordChangeRequired"`
	Email                  string   `json:"email"`
	FullName               string   `json:"fullName"`
	Phone                  string   `json:"phone"`
	Permissions            []string `json:"permissions"`
}

// UpdateUserRequest is the payload to update an existing user.
type UpdateUserRequest struct {
	Password                string   `json:"password,omitempty"`
	Services                []string `json:"services,omitempty"`
	PasswordChangeRequired  *bool    `json:"passwordChangeRequired,omitempty"`
	Email                   string   `json:"email,omitempty"`
	FullName                string   `json:"fullName,omitempty"`
	Phone                   string   `json:"phone,omitempty"`
	WalkthroughVisible      *bool    `json:"walkthroughVisible,omitempty"`
	ShowDeactivatedServices *bool    `json:"showDeactivatedServices,omitempty"`
	ShowDeactivatedScripts  *bool    `json:"showDeactivatedScripts,omitempty"`
	Permissions             []string `json:"permissions,omitempty"`
}

// UsersService handles communication with the /users endpoints.
type UsersService struct {
	Client *httpclient.Client
}

// GetCurrentUser fetches the currently authenticated user.
func (u *UsersService) GetCurrentUser(ctx context.Context) (*User, error) {
	var usr User
	if err := u.Client.Get(ctx, "/users/me", &usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

// UpdateCurrentUser updates the currently authenticated user.
func (u *UsersService) UpdateCurrentUser(ctx context.Context, req UpdateUserRequest) (*User, error) {
	var updated User
	if err := u.Client.Put(ctx, "/users/me", req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// List retrieves a paginated list of users.
func (u *UsersService) List(ctx context.Context, opts ListUsersOptions) (*ListUsersResponse, error) {
	endpoint := "/users"
	params := url.Values{}
	if opts.Search != "" {
		params.Set("search", opts.Search)
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
	var resp ListUsersResponse
	if err := u.Client.Get(ctx, fullURL, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Create adds a new user.
func (u *UsersService) Create(ctx context.Context, req CreateUserRequest) (*User, error) {
	var created User
	if err := u.Client.Post(ctx, "/users", req, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

// GetByID fetches a user by their ID.
func (u *UsersService) GetByID(ctx context.Context, id, responseType string) (*User, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/users/%s", id)
	params := url.Values{}
	if responseType != "" {
		params.Set("responseType", responseType)
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	var usr User
	if err := u.Client.Get(ctx, fullURL, &usr); err != nil {
		return nil, err
	}
	return &usr, nil
}

// UpdateByID modifies an existing user by ID.
func (u *UsersService) UpdateByID(ctx context.Context, id string, req UpdateUserRequest) (*User, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/users/%s", id)

	var updated User
	if err := u.Client.Put(ctx, endpoint, req, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeleteByID removes a user by ID.
func (u *UsersService) DeleteByID(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/users/%s", id)
	return u.Client.Delete(ctx, endpoint, nil)
}

// GetAllowedPermissions returns permissions the current token can grant to a user.
func (u *UsersService) GetAllowedPermissions(ctx context.Context, id string) ([]string, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/users/%s/allowedPermissions", id)

	var out struct {
		Permissions []string `json:"permissions"`
	}
	if err := u.Client.Get(ctx, endpoint, &out); err != nil {
		return nil, err
	}
	return out.Permissions, nil
}

// ActivateByID activates a user by ID.
func (u *UsersService) ActivateByID(ctx context.Context, id string) (*User, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/users/%s/activate", id)

	var updated User
	if err := u.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DeactivateByID deactivates a user by ID.
func (u *UsersService) DeactivateByID(ctx context.Context, id string) (*User, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/users/%s/deactivate", id)

	var updated User
	if err := u.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// EnableTwoFactorAuth enables 2FA for the current user.
// PUT /users/me/enable2FA
func (u *UsersService) EnableTwoFactorAuth(ctx context.Context) (*User, error) {
	const endpoint = "/users/me/enable2FA"

	var updated User
	// Send an empty JSON body so Content-Type: application/json is set
	if err := u.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

// DisableTwoFactorAuth disables 2FA for the current user.
// PUT /users/me/disable2FA
func (u *UsersService) DisableTwoFactorAuth(ctx context.Context) (*User, error) {
	const endpoint = "/users/me/disable2FA"

	var updated User
	if err := u.Client.Put(ctx, endpoint, struct{}{}, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}
