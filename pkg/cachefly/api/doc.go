// Package api provides the internal API client implementations for the CacheFly SDK.
//
// This package contains the HTTP client and version-specific implementations
// for interacting with the CacheFly API. Most users should not import this
// package directly, but instead use the main cachefly package which provides
// a higher-level, more convenient interface.
//
// # Available API Versions
//
// Currently supported API versions:
//
//   - v2_6: The latest stable API version with full feature support
//
// # Internal Architecture
//
// This package provides:
//
//   - HTTP client with authentication and retry logic
//   - Request/response handling and error management
//   - Version-specific service implementations
//
// # Usage
//
//		import (
//		    	"github.com/cachefly/cachefly-sdk-go/pkg/cachefly"
//	 	api "github.com/cachefly/cachefly-sdk-go/pkg/cachefly/api/v2_6"
//		)
//
//
//		client := cachefly.NewClient(
//			cachefly.WithToken(token),
//		)
//
//
//	 	newService := api.CreateServiceRequest{
//			Name:        name,
//			UniqueName:  name,
//			Description: "This is a test service created from SDK",
//		}
//
//		ctx := context.Background()
//		resp, err := client.Services.Create(ctx, newService)
//
// For standard usage, see the main cachefly package documentation.
package api
