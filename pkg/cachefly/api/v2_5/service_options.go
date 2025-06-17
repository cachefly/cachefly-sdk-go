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

	// Handle special key management fields
	var apiKeyEnabled *bool
	var protectServeKeyEnabled *bool

	// Extract apiKeyEnabled
	if val, exists := options["apiKeyEnabled"]; exists {
		if boolVal, ok := val.(bool); ok {
			apiKeyEnabled = &boolVal
			// Remove from options since it's handled separately
			delete(options, "apiKeyEnabled")
		}
	}

	// Extract protectServeKeyEnabled
	if val, exists := options["protectServeKeyEnabled"]; exists {
		if boolVal, ok := val.(bool); ok {
			protectServeKeyEnabled = &boolVal
			// Remove from options since it's handled separately
			delete(options, "protectServeKeyEnabled")
		}
	}

	// Get metadata for validation (only if there are options to validate)
	var updated ServiceOptions
	if len(options) > 0 {
		metadata, err := s.GetOptionsMetadata(ctx, id)
		if err != nil {
			return nil, fmt.Errorf("failed to get options metadata: %w", err)
		}

		// Validate options against metadata
		if err := s.validateOptions(options, metadata); err != nil {
			return nil, err
		}

		// Transform options to match API expectations
		transformedOptions := s.transformOptionsForAPI(options)

		// Update options
		endpoint := fmt.Sprintf("/services/%s/options", id)
		if err := s.Client.Put(ctx, endpoint, transformedOptions, &updated); err != nil {
			return nil, err
		}
	} else {
		// If no options to update, get current options for return value
		var err error
		updated, err = s.GetOptions(ctx, id)
		if err != nil {
			return nil, err
		}
	}

	// Handle apiKeyEnabled after options update
	if apiKeyEnabled != nil {
		if *apiKeyEnabled {
			// Generate new legacy API key
			_, err := s.RegenerateLegacyAPIKey(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("failed to regenerate legacy API key: %w", err)
			}
			updated["apiKeyEnabled"] = true
		} else {
			// Delete legacy API key
			err := s.DeleteLegacyAPIKey(ctx, id)
			if err != nil {
				return nil, fmt.Errorf("failed to delete legacy API key: %w", err)
			}
			updated["apiKeyEnabled"] = false
		}
	}

	// Handle protectServeKeyEnabled after options update
	if protectServeKeyEnabled != nil {
		if *protectServeKeyEnabled {
			// Generate new ProtectServe key
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

	// Create maps for both dynamic and standard options
	dynamicOptions := make(map[string]OptionMetadata)
	standardOptions := make(map[string]OptionMetadata)

	for _, opt := range metadata.Data {
		if opt.Type == "dynamic" && opt.Property != nil {
			dynamicOptions[opt.Property.Name] = opt
		} else if opt.Type == "standard" {
			// Map standard option names to their expected field names
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

	// Validate each option in the request
	for optionName, value := range options {
		var optMeta OptionMetadata
		var exists bool
		var isDynamic bool

		// Check dynamic options first
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

		// Check if option is read-only
		if optMeta.ReadOnly {
			validationErrors = append(validationErrors, ValidationError{
				Field:   optionName,
				Message: fmt.Sprintf("Option '%s' is read-only and cannot be modified", optionName),
				Code:    "OPTION_READ_ONLY",
			})
			continue
		}

		// Validate value based on option type
		if isDynamic {
			// Dynamic options can be direct values or enabled/value structure
			if err := s.validateDynamicOptionValue(optionName, optMeta, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field:   optionName,
					Message: err.Error(),
					Code:    "INVALID_VALUE",
				})
			}
		} else {
			// Standard options have various structures
			if err := s.validateStandardOptionValue(optionName, value); err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field:   optionName,
					Message: err.Error(),
					Code:    "INVALID_VALUE",
				})
			}
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

// validateDynamicOptionValue validates dynamic option values (can be direct or enabled/value structure)
func (s *ServiceOptionsService) validateDynamicOptionValue(optionName string, opt OptionMetadata, value interface{}) error {
	if opt.Property == nil {
		return nil // No validation possible without property metadata
	}

	prop := opt.Property

	// Check if this is an enabled/value structure
	if objVal, ok := value.(map[string]interface{}); ok {
		if _, hasEnabled := objVal["enabled"]; hasEnabled {
			// This is an enabled/value structure
			return s.validateEnabledValueStructure(optionName, objVal, prop)
		}
	}

	// This is a direct value - validate it directly
	return s.validatePropertyValue(prop, value)
}

// validateEnabledValueStructure validates the enabled/value structure for dynamic options
func (s *ServiceOptionsService) validateEnabledValueStructure(optionName string, objVal map[string]interface{}, prop *OptionProperty) error {
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

	// Validate the value field if present
	if hasValue {
		return s.validatePropertyValue(prop, val)
	}

	return nil
}

// validateStandardOptionValue validates standard option values with their various structures
func (s *ServiceOptionsService) validateStandardOptionValue(optionName string, value interface{}) error {
	switch optionName {
	case "protectServeKeyEnabled", "cors", "referrerBlocking", "autoRedirect":
		// Simple boolean options
		if _, ok := value.(bool); !ok {
			return fmt.Errorf("option '%s' expects a boolean value, got %T", optionName, value)
		}

	case "reverseProxy":
		// Complex object with enabled flag and configuration
		return s.validateReverseProxyOption(value)

	case "rawLogs":
		// Complex object with enabled flag and configuration
		return s.validateRawLogsOption(value)

	case "expiryHeaders":
		// Array of expiry header configurations or enabled/value structure
		return s.validateExpiryHeadersOption(value)

	default:
		// For unknown standard options, try to infer validation
		return s.validateGenericStandardOption(optionName, value)
	}

	return nil
}

// validateReverseProxyOption validates reverse proxy configuration
func (s *ServiceOptionsService) validateReverseProxyOption(value interface{}) error {
	// Must be an object
	objVal, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("reverseProxy must be an object, got %T", value)
	}

	// Check for enabled field
	enabled, hasEnabled := objVal["enabled"]
	if !hasEnabled {
		return fmt.Errorf("reverseProxy is missing required 'enabled' field")
	}

	// Validate enabled field is boolean
	enabledBool, ok := enabled.(bool)
	if !ok {
		return fmt.Errorf("reverseProxy 'enabled' field must be a boolean")
	}

	// If enabled, validate required configuration fields
	if enabledBool {
		requiredFields := []string{"mode", "hostname"}
		optionalFields := []string{"cacheByQueryParam", "originScheme", "ttl", "useRobotsTxt"}

		// Check required fields
		for _, field := range requiredFields {
			if _, exists := objVal[field]; !exists {
				return fmt.Errorf("reverseProxy configuration missing required field: %s", field)
			}
		}

		// Validate field types
		if mode, ok := objVal["mode"]; ok {
			if modeStr, ok := mode.(string); ok {
				validModes := []string{"WEB", "API", "STORAGE"}
				if !s.isValidEnumValue(modeStr, validModes) {
					return fmt.Errorf("reverseProxy mode must be one of: %v, got '%s'", validModes, modeStr)
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

		// Validate optional fields if present
		for _, field := range optionalFields {
			if val, exists := objVal[field]; exists {
				switch field {
				case "cacheByQueryParam", "useRobotsTxt":
					if _, ok := val.(bool); !ok {
						return fmt.Errorf("reverseProxy %s must be a boolean", field)
					}
				case "ttl":
					if !s.isNumeric(val) {
						return fmt.Errorf("reverseProxy ttl must be a number")
					}
				case "originScheme":
					if schemeStr, ok := val.(string); ok {
						validSchemes := []string{"FOLLOW", "HTTP", "HTTPS"}
						if !s.isValidEnumValue(schemeStr, validSchemes) {
							return fmt.Errorf("reverseProxy originScheme must be one of: %v", validSchemes)
						}
					} else {
						return fmt.Errorf("reverseProxy originScheme must be a string")
					}
				}
			}
		}
	}

	return nil
}

// validateRawLogsOption validates raw logs configuration
func (s *ServiceOptionsService) validateRawLogsOption(value interface{}) error {
	// Must be an object
	objVal, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("rawLogs must be an object, got %T", value)
	}

	// Check for enabled field
	enabled, hasEnabled := objVal["enabled"]
	if !hasEnabled {
		return fmt.Errorf("rawLogs is missing required 'enabled' field")
	}

	// Validate enabled field is boolean
	enabledBool, ok := enabled.(bool)
	if !ok {
		return fmt.Errorf("rawLogs 'enabled' field must be a boolean")
	}

	// If enabled, validate configuration fields
	if enabledBool {
		// Validate logFormat if present
		if logFormat, exists := objVal["logFormat"]; exists {
			if logFormatStr, ok := logFormat.(string); ok {
				validFormats := []string{"combined", "common", "custom"}
				if !s.isValidEnumValue(logFormatStr, validFormats) {
					return fmt.Errorf("rawLogs logFormat must be one of: %v", validFormats)
				}
			} else {
				return fmt.Errorf("rawLogs logFormat must be a string")
			}
		}

		// Validate compression if present
		if compression, exists := objVal["compression"]; exists {
			if compressionStr, ok := compression.(string); ok {
				validCompressions := []string{"gzip", "bzip2", "none"}
				if !s.isValidEnumValue(compressionStr, validCompressions) {
					return fmt.Errorf("rawLogs compression must be one of: %v", validCompressions)
				}
			} else {
				return fmt.Errorf("rawLogs compression must be a string")
			}
		}
	}

	return nil
}

// validateExpiryHeadersOption validates expiry headers array configuration
func (s *ServiceOptionsService) validateExpiryHeadersOption(value interface{}) error {
	// Could be an array directly or an enabled/value structure
	if objVal, ok := value.(map[string]interface{}); ok {
		if enabled, hasEnabled := objVal["enabled"]; hasEnabled {
			// This is an enabled/value structure
			enabledBool, ok := enabled.(bool)
			if !ok {
				return fmt.Errorf("expiryHeaders 'enabled' field must be a boolean")
			}

			if enabledBool {
				if val, hasValue := objVal["value"]; hasValue {
					return s.validateExpiryHeadersArray(val)
				} else {
					return fmt.Errorf("expiryHeaders requires 'value' field when enabled is true")
				}
			}
			return nil
		}
	}

	// Direct array
	return s.validateExpiryHeadersArray(value)
}

// validateExpiryHeadersArray validates the actual expiry headers array
func (s *ServiceOptionsService) validateExpiryHeadersArray(value interface{}) error {
	var arrayVal []interface{}

	// Handle both []interface{} and []map[string]interface{} types
	switch v := value.(type) {
	case []interface{}:
		arrayVal = v
	case []map[string]interface{}:
		// Convert []map[string]interface{} to []interface{}
		arrayVal = make([]interface{}, len(v))
		for i, item := range v {
			arrayVal[i] = item
		}
	default:
		return fmt.Errorf("expiryHeaders must be an array, got %T", value)
	}

	// Validate each expiry header entry
	for i, entry := range arrayVal {
		entryObj, ok := entry.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expiryHeaders[%d] must be an object", i)
		}

		// Check for required fields (at least one of path or extension)
		path, hasPath := entryObj["path"]
		extension, hasExtension := entryObj["extension"]
		expiryTime, hasExpiryTime := entryObj["expiryTime"]

		if !hasPath && !hasExtension {
			return fmt.Errorf("expiryHeaders[%d] must have either 'path' or 'extension' field", i)
		}

		if !hasExpiryTime {
			return fmt.Errorf("expiryHeaders[%d] is missing required 'expiryTime' field", i)
		}

		// Validate field types
		if hasPath {
			if _, ok := path.(string); !ok {
				return fmt.Errorf("expiryHeaders[%d] path must be a string", i)
			}
		}

		if hasExtension {
			if _, ok := extension.(string); !ok {
				return fmt.Errorf("expiryHeaders[%d] extension must be a string", i)
			}
		}

		if !s.isNumeric(expiryTime) {
			return fmt.Errorf("expiryHeaders[%d] expiry_time must be a number", i)
		}
	}

	return nil
}

// validateGenericStandardOption provides generic validation for unknown standard options
func (s *ServiceOptionsService) validateGenericStandardOption(optionName string, value interface{}) error {
	// Try to determine the expected type based on the value structure
	switch v := value.(type) {
	case bool:
		// Simple boolean toggle - no additional validation needed
		return nil
	case map[string]interface{}:
		// Complex object - check for common patterns
		if enabled, hasEnabled := v["enabled"]; hasEnabled {
			// Has enabled/value or enabled/config pattern
			if _, ok := enabled.(bool); !ok {
				return fmt.Errorf("option '%s' enabled field must be a boolean", optionName)
			}
		}
		return nil
	case []interface{}:
		// Array configuration - basic validation
		return nil
	default:
		// Allow other types but no specific validation
		return nil
	}
}

// validatePropertyValue validates a simple value against property constraints
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

// transformOptionsForAPI transforms validated options to match API expectations
func (s *ServiceOptionsService) transformOptionsForAPI(options ServiceOptions) ServiceOptions {
	transformed := make(ServiceOptions)

	// Simply pass through all options as-is since the input is already in the correct format
	for optionName, value := range options {
		transformed[optionName] = value
	}

	return transformed
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
