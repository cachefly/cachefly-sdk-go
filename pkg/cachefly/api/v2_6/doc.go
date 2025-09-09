// package v2_6 implements CacheFly API v2.6 service interfaces.
//
// This package contains all the service implementations for interacting
// with CacheFly's API version 2.6. Each service group handles specific
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
// - CacheWarmingService: Manages cache warming tasks (list/create/get/delete)
// - AccountStatsService: Provides account-level statistics endpoints
// - ServiceStatsService: Provides service-level statistics endpoints
// - AvailabilityService: Checks availability of domains, usernames, services, SAML
// - SAMLService: Manages SAML configuration operations
//
// This package is typically not imported directly. Instead, use the
// main cachefly package which provides a unified client interface.
//
// Example usage through main client:
//
//	import "github.com/cachefly/cachefly-sdk-go"
//
//	client := cachefly.NewClient(cachefly.WithToken("your-token"))
//	services, err := client.Services.List()
package v2_6
