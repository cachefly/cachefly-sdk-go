// Package v2_5 implements CacheFly API v2.5 service interfaces.
//
// This package contains all the service implementations for interacting
// with CacheFly's API version 2.5. Each service group handles specific
// functionality within the CacheFly platform.
//
// Service Groups:
//
// - ServicesService: Manages CDN services and configurations
// - AccountsService: Handles account-level operations
// - ServiceDomainsService: Manages domain configurations
// - ServiceRulesService: Controls caching and delivery rules
// - ServiceOptionsService: Manages service-specific options and settings
// - ScriptConfigsService: Handles script configurations and definitions
// - CertificatesService: Manages SSL/TLS certificates
// - OriginsService: Configures origin server settings
// - UsersService: Handles user management and permissions
// - TLSProfilesService: Manages TLS profile configurations
//
// This package is typically not imported directly. Instead, use the
// main cachefly package which provides a unified client interface.
//
// Example usage through main client:
//
//	import "github.com/cachefly/cachefly-go-sdk"
//
//	client := cachefly.NewClient(cachefly.WithToken("your-token"))
//	services, err := client.Services.List()
package v2_5
