package v2_5

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cachefly/cachefly-go-sdk/internal/httpclient"
)

// OptionProperty represents detailed metadata about an option property
type OptionProperty struct {
	Label      string      `json:"label"`
	ID         string      `json:"_id"`
	Name       string      `json:"name"`
	Type       string      `json:"type"` // "boolean", "integer", "enum", "bitfield", "strings" this is as the per the api doc
	MaxValue   *int        `json:"maxValue,omitempty"`
	MinValue   *int        `json:"minValue,omitempty"`
	Default    interface{} `json:"default,omitempty"`
	EnumValues []EnumValue `json:"enumValues,omitempty"`
	BitFields  []BitField  `json:"bitFields,omitempty"`
	UpdatedAt  string      `json:"updatedAt"`
	CreatedAt  string      `json:"createdAt"`
}

// EnumValue represents possible values for enum type options
type EnumValue struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

// BitField represents bitfield options (like HTTP methods)
type BitField struct {
	BitPosition int    `json:"bitPosition"`
	Key         string `json:"key"`
	Label       string `json:"label"`
}

// PromoInfo contains promotional/UI information
type PromoInfo struct {
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}

// OptionMetadata describes a complete service option with all its metadata
type OptionMetadata struct {
	ID          string          `json:"_id"`
	Name        string          `json:"name"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Template    string          `json:"template"`
	Group       string          `json:"group"`
	Scope       string          `json:"scope"`
	ReadOnly    bool            `json:"readOnly"`
	Type        string          `json:"type"`               // "standard", "dynamic"
	Property    *OptionProperty `json:"property,omitempty"` // Only present for dynamic types
	Promo       PromoInfo       `json:"promo"`
	UpdatedAt   string          `json:"updatedAt"`
	CreatedAt   string          `json:"createdAt"`
}

// ServiceOptionsMetadata contains the complete metadata response
type ServiceOptionsMetadata struct {
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
	Data []OptionMetadata `json:"data"`
}

// ValidationError represents a validation error for a specific field
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
	Code    string `json:"code"`
}

// ServiceOptionsValidationError represents multiple validation errors
type ServiceOptionsValidationError struct {
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors"`
}

func (e ServiceOptionsValidationError) Error() string {
	return e.Message
}

// ServiceOptions represents service options as a flexible map
type ServiceOptions map[string]interface{}

// ServiceOptionsService handles service options endpoints.
type ServiceOptionsService struct {
	Client *httpclient.Client
}

// GetOptionsMetadata retrieves metadata about available options for a service
func (s *ServiceOptionsService) GetOptionsMetadata(ctx context.Context, id string) (*ServiceOptionsMetadata, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/metadata", id)

	var metadata ServiceOptionsMetadata
	if err := s.Client.Get(ctx, endpoint, &metadata); err != nil {
		return nil, err
	}
	return &metadata, nil
}

// GetOptions retrieves current options for a service
func (s *ServiceOptionsService) GetOptions(ctx context.Context, id string) (ServiceOptions, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options", id)

	var opts ServiceOptions
	if err := s.Client.Get(ctx, endpoint, &opts); err != nil {
		return nil, err
	}
	return opts, nil
}

// UpdateOptions updates service options with strict validation
func (s *ServiceOptionsService) UpdateOptions(ctx context.Context, id string, options ServiceOptions) (ServiceOptions, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	// we are gettimg metadata for validation
	metadata, err := s.GetOptionsMetadata(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get options metadata: %w", err)
	}

	// validate options against metadata
	if err := s.validateOptions(options, metadata); err != nil {
		return nil, err
	}

	// lets transform options to match api expectations
	transformedOptions := s.transformOptionsForAPI(options)

	// finally we Update options
	endpoint := fmt.Sprintf("/services/%s/options", id)
	var updated ServiceOptions
	if err := s.Client.Put(ctx, endpoint, transformedOptions, &updated); err != nil {
		return nil, err
	}
	return updated, nil
}

// validateOptions performs strict validation against metadata
func (s *ServiceOptionsService) validateOptions(options ServiceOptions, metadata *ServiceOptionsMetadata) error {
	var validationErrors []ValidationError

	// Create a map of available options for quick lookup
	availableOptions := make(map[string]OptionMetadata)
	for _, opt := range metadata.Data {
		// For dynamic options, use the property name
		if opt.Type == "dynamic" && opt.Property != nil {
			availableOptions[opt.Property.Name] = opt
		}
	}

	// Validate each option in the request
	for optionName, value := range options {
		if s.isFeatureEnablementOption(optionName) {

			if err := s.validateEnabledValueStructure(optionName, value, availableOptions); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field:   optionName,
					Message: err.Error(),
					Code:    "INVALID_STRUCTURE",
				})
			}
			continue
		}

		optMeta, exists := availableOptions[optionName]
		if !exists {
			validationErrors = append(validationErrors, ValidationError{
				Field:   optionName,
				Message: fmt.Sprintf("Option '%s' is not available for this service", optionName),
				Code:    "OPTION_NOT_AVAILABLE",
			})
			continue
		}

		// Check if option is read-only
		if optMeta.ReadOnly {
			validationErrors = append(validationErrors, ValidationError{
				Field:   optionName,
				Message: fmt.Sprintf("Option '%s' is read-only and cannot be modified", optionName),
				Code:    "OPTION_READ_ONLY",
			})
			continue
		}

		// Validate value based on metadata
		if err := s.validateOptionValue(optMeta, value); err != nil {
			validationErrors = append(validationErrors, ValidationError{
				Field:   optionName,
				Message: err.Error(),
				Code:    "INVALID_VALUE",
			})
		}
	}

	// Return validation errors if any
	if len(validationErrors) > 0 {
		return ServiceOptionsValidationError{
			Message: fmt.Sprintf("Validation failed for %d option(s)", len(validationErrors)),
			Errors:  validationErrors,
		}
	}

	return nil
}

// validateEnabledValueStructure validates the enabled/value structure for complex options
func (s *ServiceOptionsService) validateEnabledValueStructure(optionName string, value interface{}, availableOptions map[string]OptionMetadata) error {
	// Must be an object
	objVal, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("option '%s' must be an object", optionName)
	}

	// Special handling for complex objects that don't use enabled/value structure
	if optionName == "reverseProxy" || optionName == "rawLogs" {
		enabled, hasEnabled := objVal["enabled"]
		if !hasEnabled {
			return fmt.Errorf("option '%s' is missing required 'enabled' field", optionName)
		}

		// Validate enabled field is boolean
		if _, ok := enabled.(bool); !ok {
			return fmt.Errorf("'enabled' field must be a boolean for option '%s'", optionName)
		}

		// Special validation for reverseProxy
		if optionName == "reverseProxy" {
			if enabledBool, ok := enabled.(bool); ok && enabledBool {
				return s.validateReverseProxyStructure(objVal)
			}
		}
		return nil
	}

	// Standard enabled/value structure for other options
	enabled, hasEnabled := objVal["enabled"]
	val, hasValue := objVal["value"]

	// Must have enabled field
	if !hasEnabled {
		return fmt.Errorf("option '%s' is missing required 'enabled' field", optionName)
	}

	// Validate enabled field is boolean
	enabledBool, ok := enabled.(bool)
	if !ok {
		return fmt.Errorf("'enabled' field must be a boolean for option '%s'", optionName)
	}

	// Value field is required only when enabled is true
	if enabledBool && !hasValue {
		return fmt.Errorf("option '%s' requires 'value' field when enabled is true", optionName)
	}

	// Validate the value field against metadata if available and present
	if hasValue {
		if optMeta, exists := availableOptions[optionName]; exists {
			if err := s.validateOptionValue(optMeta, val); err != nil {
				return fmt.Errorf("invalid value in enabled/value structure for '%s': %w", optionName, err)
			}
		}
	}

	return nil
}

// validateReverseProxyStructure validates reverse proxy configuration
func (s *ServiceOptionsService) validateReverseProxyStructure(config map[string]interface{}) error {
	requiredFields := []string{"mode", "hostname"}

	for _, field := range requiredFields {
		if _, exists := config[field]; !exists {
			return fmt.Errorf("reverse proxy configuration missing required field: %s", field)
		}
	}
	return nil
}

// isFeatureEnablementOption checks if an option is a feature enablement option
func (s *ServiceOptionsService) isFeatureEnablementOption(optionName string) bool {
	// Define options that require enabled/value structure
	featureOptions := []string{
		"reverseProxy",
		"rawLogs",
		"error_ttl",
		"ttfb_timeout",
		"contimeout",
		"maxcons",
		"bwthrottle",
		"sharedshield",
		"originhostheader",
		"purgemode",
		"dirpurgeskip",
		"httpmethods",
		"skip_pserve_ext",
		"skip_encoding_ext",
		"redirect",
		"slice",
		"bwthrottlequery",
	}

	// Check if this is a feature option
	for _, feature := range featureOptions {
		if optionName == feature {
			return true
		}
	}

	return false
}

func (s *ServiceOptionsService) transformOptionsForAPI(options ServiceOptions) ServiceOptions {
	transformed := make(ServiceOptions)

	// Since we only accept enabled/value format, pass through as-is
	for optionName, value := range options {
		transformed[optionName] = value
	}

	return transformed
}

// validateOptionValue validates a value against option metadata
func (s *ServiceOptionsService) validateOptionValue(opt OptionMetadata, value interface{}) error {
	if opt.Property == nil {
		return nil // No validation for standard options
	}

	prop := opt.Property
	switch prop.Type {
	case "boolean":
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("expected boolean value, got %T", value)
		}
	case "integer":
		var intVal int
		switch v := value.(type) {
		case int:
			intVal = v
		case float64:
			intVal = int(v)
		default:
			return fmt.Errorf("expected integer value, got %T", value)
		}

		if prop.MinValue != nil && intVal < *prop.MinValue {
			return fmt.Errorf("value %d is below minimum %d", intVal, *prop.MinValue)
		}
		if prop.MaxValue != nil && intVal > *prop.MaxValue {
			return fmt.Errorf("value %d is above maximum %d", intVal, *prop.MaxValue)
		}
	case "enum":
		strVal, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string value for enum, got %T", value)
		}

		validValues := make([]string, len(prop.EnumValues))
		for i, enumVal := range prop.EnumValues {
			validValues[i] = enumVal.Value
			if enumVal.Value == strVal {
				return nil // Valid enum value
			}
		}
		return fmt.Errorf("value '%s' is not valid, must be one of: %v", strVal, validValues)
	case "bitfield":
		// Bitfield should be a map[string]bool
		bitMap, ok := value.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected object for bitfield, got %T", value)
		}

		validKeys := make(map[string]bool)
		for _, field := range prop.BitFields {
			validKeys[field.Key] = true
		}

		for key := range bitMap {
			if !validKeys[key] {
				return fmt.Errorf("invalid bitfield key '%s'", key)
			}
		}
	case "strings":
		// Should be an array of strings
		switch v := value.(type) {
		case []interface{}:
			for _, item := range v {
				if _, ok := item.(string); !ok {
					return fmt.Errorf("all items in strings array must be strings")
				}
			}
		case []string:
			// Already valid
		default:
			return fmt.Errorf("expected array of strings, got %T", value)
		}
	}

	return nil
}

// UpdateSpecificOption updates a single option by name with validation
func (s *ServiceOptionsService) UpdateSpecificOption(ctx context.Context, id string, optionName string, value interface{}) (ServiceOptions, error) {
	options := ServiceOptions{
		optionName: value,
	}
	return s.UpdateOptions(ctx, id, options)
}

// IsOptionAvailable checks if a specific option is available for the service
func (s *ServiceOptionsService) IsOptionAvailable(ctx context.Context, id string, optionName string) (bool, *OptionMetadata, error) {
	metadata, err := s.GetOptionsMetadata(ctx, id)
	if err != nil {
		return false, nil, err
	}

	for _, opt := range metadata.Data {
		if opt.Type == "dynamic" && opt.Property != nil && opt.Property.Name == optionName {
			return true, &opt, nil
		}
	}
	return false, nil, nil
}

// GetAvailableOptionNames returns a list of all available option names
func (s *ServiceOptionsService) GetAvailableOptionNames(ctx context.Context, id string) ([]string, error) {
	metadata, err := s.GetOptionsMetadata(ctx, id)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, opt := range metadata.Data {
		if opt.Type == "dynamic" && opt.Property != nil {
			names = append(names, opt.Property.Name)
		}
	}
	return names, nil
}

// GetOptionsByGroup returns options grouped by their group field
func (s *ServiceOptionsService) GetOptionsByGroup(ctx context.Context, id string) (map[string][]OptionMetadata, error) {
	metadata, err := s.GetOptionsMetadata(ctx, id)
	if err != nil {
		return nil, err
	}

	groups := make(map[string][]OptionMetadata)
	for _, opt := range metadata.Data {
		groups[opt.Group] = append(groups[opt.Group], opt)
	}
	return groups, nil
}

// LegacyAPIKeyResponse represents API key payload.
type LegacyAPIKeyResponse struct {
	APIKey string `json:"apiKey"`
}

// ProtectServeKeyResponse for protectserve.
type ProtectServeKeyResponse struct {
	ProtectServeKey   string `json:"protectServeKey"`
	ForceProtectServe string `json:"forceProtectserve"`
}

// UpdateProtectServeRequest updates protectserve options.
type UpdateProtectServeRequest struct {
	ForceProtectServe string `json:"forceProtectServe"`
	ProtectServeKey   string `json:"protectServeKey"`
}

// FTPSettingsResponse represents FTP settings.
type FTPSettingsResponse struct {
	FTPPassword string `json:"ftpPassword"`
}

// GetLegacyAPIKey returns the legacy API key for a service.
func (s *ServiceOptionsService) GetLegacyAPIKey(ctx context.Context, id string) (*LegacyAPIKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/apikey", id)

	var res LegacyAPIKeyResponse
	if err := s.Client.Get(ctx, endpoint, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RegenerateLegacyAPIKey creates a new legacy API key.
func (s *ServiceOptionsService) RegenerateLegacyAPIKey(ctx context.Context, id string) (*LegacyAPIKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/apikey", id)

	//it needs empty body
	emptyBody := struct{}{}

	var res LegacyAPIKeyResponse
	if err := s.Client.Post(ctx, endpoint, emptyBody, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteLegacyAPIKey deletes the legacy API key for a service.
func (s *ServiceOptionsService) DeleteLegacyAPIKey(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/apikey", id)
	return s.Client.Delete(ctx, endpoint, nil)
}

// GetProtectServeKey retrieves the protectserve key (optional hideSecrets).
func (s *ServiceOptionsService) GetProtectServeKey(ctx context.Context, id string, hideSecrets bool) (*ProtectServeKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/protectserve", id)

	params := url.Values{}
	params.Set("hideSecrets", "false")

	if strconv.FormatBool(hideSecrets) != "" {
		params.Set("hideSecrets", strconv.FormatBool(hideSecrets))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var res ProtectServeKeyResponse
	if err := s.Client.Get(ctx, fullURL, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RecreateProtectServeKey regenerates or reverts the protectserve key.
func (s *ServiceOptionsService) RecreateProtectServeKey(ctx context.Context, id, action string) (*ProtectServeKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/protectserve", id)
	params := url.Values{}
	if action != "" {
		params.Set("action", action)
	}
	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	//it needs empty body
	emptyBody := struct{}{}

	var res ProtectServeKeyResponse
	if err := s.Client.Post(ctx, fullURL, emptyBody, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// UpdateProtectServeOptions updates protectserve key and options.
func (s *ServiceOptionsService) UpdateProtectServeOptions(ctx context.Context, id string, req UpdateProtectServeRequest) (*ProtectServeKeyResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/protectserve", id)

	var res ProtectServeKeyResponse
	if err := s.Client.Put(ctx, endpoint, req, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// DeleteProtectServeKey deletes the ProtectServe key for the specified service.
func (s *ServiceOptionsService) DeleteProtectServeKey(ctx context.Context, serviceID string) error {
	if serviceID == "" {
		return fmt.Errorf("service ID is required")
	}

	// Build endpoint path: DELETE /services/{id}/options/protectserve
	endpoint := fmt.Sprintf("/services/%s/options/protectserve", url.PathEscape(serviceID))

	// Perform DELETE request. No request body expected.
	if err := s.Client.Delete(ctx, endpoint, nil); err != nil {
		return err
	}
	return nil
}

// GetFTPSettings retrieves FTP settings for a service (optional hideSecrets).
func (s *ServiceOptionsService) GetFTPSettings(ctx context.Context, id string, hideSecrets bool) (*FTPSettingsResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/ftp", id)
	params := url.Values{}

	params.Set("hideSecrets", "false")
	if strconv.FormatBool(hideSecrets) != "" {
		params.Set("hideSecrets", strconv.FormatBool(hideSecrets))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var res FTPSettingsResponse
	if err := s.Client.Get(ctx, fullURL, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// RegenerateFTPPassword regenerates the FTP password for a service.
func (s *ServiceOptionsService) RegenerateFTPPassword(ctx context.Context, id string, hideSecrets bool) (*FTPSettingsResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}
	endpoint := fmt.Sprintf("/services/%s/options/ftp", id)
	params := url.Values{}

	params.Set("hideSecrets", "false")
	if strconv.FormatBool(hideSecrets) != "" {
		params.Set("hideSecrets", strconv.FormatBool(hideSecrets))
	}

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	emptyBody := struct{}{}

	var res FTPSettingsResponse
	if err := s.Client.Post(ctx, fullURL, emptyBody, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
