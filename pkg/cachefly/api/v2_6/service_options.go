package v2_6

import (
	"context"
	"fmt"
	"net/url"
	"slices"
	"strconv"

	"github.com/cachefly/cachefly-sdk-go/internal/httpclient"
)

// OptionProperty represents detailed metadata about an option property
type OptionProperty struct {
	Label      string      `json:"label"`
	ID         string      `json:"_id"`
	Name       string      `json:"name"`
	Type       string      `json:"type"` // "boolean", "integer", "enum", "bitfield", "strings"
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

// UpdateOptions updates service options with strict validation and handles special cases
func (s *ServiceOptionsService) UpdateOptions(ctx context.Context, id string, options ServiceOptions) (ServiceOptions, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
	}

	var protectServeKeyEnabled *bool

	if val, exists := options["protectServeKeyEnabled"]; exists {
		if boolVal, ok := val.(bool); ok {
			protectServeKeyEnabled = &boolVal
			// Remove from options since it's handled separately
			delete(options, "protectServeKeyEnabled")
		}
	}

	var updated ServiceOptions
	if len(options) > 0 {
		metadata, err := s.GetOptionsMetadata(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to get options metadata: %w", err)
		}

		if err := s.validateOptions(options, metadata); err != nil {
			return nil, err
		}

		endpoint := fmt.Sprintf("/services/%s/options", id)
		if err := s.Client.Put(ctx, endpoint, options, &updated); err != nil {
			return nil, err
		}
	} else {
		var err error
		updated, err = s.GetOptions(ctx, id)
		if err != nil {
			return nil, err
		}
	}

	if protectServeKeyEnabled != nil {
		if *protectServeKeyEnabled {
			_, err := s.RecreateProtectServeKey(ctx, id, "REGENERATE")
			if err != nil {
				return nil, fmt.Errorf("failed to regenerate ProtectServe key: %w", err)
			}
			updated["protectServeKeyEnabled"] = true
		} else {
			// Delete ProtectServe key
			err := s.DeleteProtectServeKey(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("failed to delete ProtectServe key: %w", err)
			}
			updated["protectServeKeyEnabled"] = false
		}
	}

	return updated, nil
}

// validateOptions performs strict validation against metadata
func (s *ServiceOptionsService) validateOptions(options ServiceOptions, metadata *ServiceOptionsMetadata) error {
	var validationErrors []ValidationError

	dynamicOptions := make(map[string]OptionMetadata)
	standardOptions := make(map[string]OptionMetadata)

	for _, opt := range metadata.Data {
		if opt.Type == "dynamic" && opt.Property != nil {
			dynamicOptions[opt.Property.Name] = opt
		} else if opt.Type == "standard" {
			var optName string
			switch opt.Name {
			case "Reverse Proxy":
				optName = "reverseProxy"
			case "ProtectServe":
				optName = "protectServeKeyEnabled"
			case "CORS Override":
				optName = "cors"
			case "Expiry Overrides":
				optName = "expiryHeaders"
			case "Referrer Blocking":
				optName = "referrerBlocking"
			case "Auto HTTPS Redirect":
				optName = "autoRedirect"
			default:
				optName = opt.Name
			}
			standardOptions[optName] = opt
		}
	}

	for optionName, value := range options {
		var optMeta OptionMetadata
		var exists bool
		var isDynamic bool

		if optMeta, exists = dynamicOptions[optionName]; exists {
			isDynamic = true
		} else if optMeta, exists = standardOptions[optionName]; exists {
			isDynamic = false
		} else {
			validationErrors = append(validationErrors, ValidationError{
				Field:   optionName,
				Message: fmt.Sprintf("Option '%s' is not available for this service", optionName),
				Code:    "OPTION_NOT_AVAILABLE",
			})
			continue
		}

		if optMeta.ReadOnly {
			validationErrors = append(validationErrors, ValidationError{
				Field:   optionName,
				Message: fmt.Sprintf("Option '%s' is read-only and cannot be modified", optionName),
				Code:    "OPTION_READ_ONLY",
			})
			continue
		}

		if isDynamic {
			if err := s.validateDynamicOptionValue(optionName, optMeta, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field:   optionName,
					Message: err.Error(),
					Code:    "INVALID_VALUE",
				})
			}
		} else {
			if err := s.validateStandardOptionValue(optionName, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field:   optionName,
					Message: err.Error(),
					Code:    "INVALID_VALUE",
				})
			}
		}
	}

	if len(validationErrors) > 0 {
		return ServiceOptionsValidationError{
			Message: fmt.Sprintf("Validation failed for %d option(s)", len(validationErrors)),
			Errors:  validationErrors,
		}
	}

	return nil
}

func (s *ServiceOptionsService) validateOptions_without_metadata(options ServiceOptions) error {
	var validationErrors []ValidationError

	for optionName, value := range options {

		standardOptions := []string{"reverseProxy", "protectServeKeyEnabled", "cors", "referrerBlocking", "autoRedirect"}

		if slices.Contains(standardOptions, optionName) {
			if err := s.validateStandardOptionValue(optionName, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field:   optionName,
					Message: err.Error(),
					Code:    "INVALID_VALUE",
				})
			}
		}

	}

	if len(validationErrors) > 0 {
		return ServiceOptionsValidationError{
			Message: fmt.Sprintf("Validation failed for %d option(s)", len(validationErrors)),
			Errors:  validationErrors,
		}
	}

	return nil
}

func (s *ServiceOptionsService) validateDynamicOptionValue(optionName string, opt OptionMetadata, value interface{}) error {
	if opt.Property == nil {
		return nil // No validation possible without property metadata
	}

	prop := opt.Property

	if objVal, ok := value.(map[string]interface{}); ok {
		if _, hasEnabled := objVal["enabled"]; hasEnabled {
			return s.validateEnabledValueStructure(optionName, objVal, prop)
		}
	}

	return s.validatePropertyValue(prop, value)
}

func (s *ServiceOptionsService) validateEnabledValueStructure(optionName string, objVal map[string]interface{}, prop *OptionProperty) error {
	enabled, hasEnabled := objVal["enabled"]
	val, hasValue := objVal["value"]

	if !hasEnabled {
		return fmt.Errorf("option '%s' is missing required 'enabled' field", optionName)
	}

	enabledBool, ok := enabled.(bool)
	if !ok {
		return fmt.Errorf("'enabled' field must be a boolean for option '%s'", optionName)
	}

	if enabledBool && !hasValue {
		return fmt.Errorf("option '%s' requires 'value' field when enabled is true", optionName)
	}

	if hasValue {
		return s.validatePropertyValue(prop, val)
	}

	return nil
}

func (s *ServiceOptionsService) validateStandardOptionValue(optionName string, value interface{}) error {
	switch optionName {
	case "protectServeKeyEnabled", "cors", "referrerBlocking", "autoRedirect":
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("option '%s' expects a boolean value, got %T", optionName, value)
		}

	case "reverseProxy":
		return s.validateReverseProxyOption(value)

	default:
		return s.validateGenericStandardOption(optionName, value)
	}

	return nil
}

func (s *ServiceOptionsService) validateReverseProxyOption(value interface{}) error {
	objVal, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("reverseProxy must be an object, got %T", value)
	}

	enabled, hasEnabled := objVal["enabled"]
	if !hasEnabled {
		return fmt.Errorf("reverseProxy is missing required 'enabled' field")
	}

	enabledBool, ok := enabled.(bool)
	if !ok {
		return fmt.Errorf("reverseProxy 'enabled' field must be a boolean")
	}

	if enabledBool {
		requiredFields := []string{"hostname", "originScheme", "useRobotsTxt", "ttl", "cacheByQueryParam"}

		for _, field := range requiredFields {
			if _, exists := objVal[field]; !exists {
				return fmt.Errorf("reverseProxy configuration missing required field: %s", field)
			}
		}

		if mode, ok := objVal["mode"]; ok {
			if modeStr, ok := mode.(string); ok {
				if modeStr == "OBJECT_STORAGE" {
					objectStorageFields := []string{"accessKey", "secretKey", "region"}
					for _, field := range objectStorageFields {
						if _, exists := objVal[field]; !exists {
							return fmt.Errorf("reverseProxy configuration with mode OBJECT_STORAGE missing required field: %s", field)
						}
					}
				} else if modeStr != "WEB" {
					return fmt.Errorf("reverseProxy mode must be one of: %v, got '%s'", []string{"WEB", "OBJECT_STORAGE"}, modeStr)
				}
			} else {
				return fmt.Errorf("reverseProxy mode must be a string")
			}
		}

		if hostname, ok := objVal["hostname"]; ok {
			if _, ok := hostname.(string); !ok {
				return fmt.Errorf("reverseProxy hostname must be a string")
			}
		}

		if val, exists := objVal["cacheByQueryParam"]; exists {
			if _, ok := val.(bool); !ok {
				return fmt.Errorf("reverseProxy cacheByQueryParam must be a boolean")
			}
		}

		if ttl, exists := objVal["ttl"]; exists {
			if !s.isNumeric(ttl) {
				return fmt.Errorf("reverseProxy ttl must be a number")
			}
		}

		if originScheme, exists := objVal["originScheme"]; exists {
			if schemeStr, ok := originScheme.(string); ok {
				validSchemes := []string{"FOLLOW", "HTTP", "HTTPS"}
				if !s.isValidEnumValue(schemeStr, validSchemes) {
					return fmt.Errorf("reverseProxy originScheme must be one of: %v", validSchemes)
				}
			} else {
				return fmt.Errorf("reverseProxy originScheme must be a string")
			}
		}

		for _, field := range []string{"useRobotsTxt", "cacheByQueryParam"} {
			if val, exists := objVal[field]; exists {
				if _, ok := val.(bool); !ok {
					return fmt.Errorf("reverseProxy %s must be a boolean", field)
				}
			}
		}
	}

	return nil
}

func (s *ServiceOptionsService) validateGenericStandardOption(optionName string, value interface{}) error {
	switch v := value.(type) {
	case bool:
		return nil
	case map[string]interface{}:
		if enabled, hasEnabled := v["enabled"]; hasEnabled {
			if _, ok := enabled.(bool); !ok {
				return fmt.Errorf("option '%s' enabled field must be a boolean", optionName)
			}
		}
		return nil
	case []interface{}:
		return nil
	default:
		return nil
	}
}

func (s *ServiceOptionsService) validatePropertyValue(prop *OptionProperty, value interface{}) error {
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
	case "string":
		if _, ok := value.(string); !ok {
			return fmt.Errorf("expected string value, got %T", value)
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

	// Check both dynamic and standard options
	for _, opt := range metadata.Data {
		if opt.Type == "dynamic" && opt.Property != nil && opt.Property.Name == optionName {
			return true, &opt, nil
		} else if opt.Type == "standard" {
			var mappedName string
			switch opt.Name {
			case "Reverse Proxy":
				mappedName = "reverseProxy"
			case "ProtectServe":
				mappedName = "protectServeKeyEnabled"
			case "CORS Override":
				mappedName = "cors"
			case "Expiry Overrides":
				mappedName = "expiryHeaders"
			case "Referrer Blocking":
				mappedName = "referrerBlocking"
			case "Auto HTTPS Redirect":
				mappedName = "autoRedirect"
			default:
				mappedName = opt.Name
			}
			if mappedName == optionName {
				return true, &opt, nil
			}
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
		} else if opt.Type == "standard" {
			var mappedName string
			switch opt.Name {
			case "Reverse Proxy":
				mappedName = "reverseProxy"
			case "ProtectServe":
				mappedName = "protectServeKeyEnabled"
			case "CORS Override":
				mappedName = "cors"
			case "Expiry Overrides":
				mappedName = "expiryHeaders"
			case "Referrer Blocking":
				mappedName = "referrerBlocking"
			case "Auto HTTPS Redirect":
				mappedName = "autoRedirect"
			default:
				mappedName = opt.Name
			}
			names = append(names, mappedName)
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

// Helper functions
func (s *ServiceOptionsService) isValidEnumValue(value string, validValues []string) bool {
	for _, valid := range validValues {
		if value == valid {
			return true
		}
	}
	return false
}

func (s *ServiceOptionsService) isNumeric(value interface{}) bool {
	switch value.(type) {
	case int, int32, int64, float32, float64:
		return true
	default:
		return false
	}
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
