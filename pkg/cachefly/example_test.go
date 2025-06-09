package cachefly_test

import (
	"context"
	"fmt"
	"log"

	"github.com/cachefly/cachefly-go-sdk/pkg/cachefly"
	api "github.com/cachefly/cachefly-go-sdk/pkg/cachefly/api/v2_5"
)

// ============================================================================
// Account Examples
// ============================================================================

// ExampleAccountsService_CreateChildAccount demonstrates creating a child account.
func ExampleAccountsService_CreateChildAccount() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	createReq := api.CreateChildAccountRequest{
		CompanyName: "Subsidiary Corp",
		Username:    "john.doe",
		Password:    "SecurePass123!",
		FullName:    "John Doe",
		Email:       "john.doe@subsidiary.com",
		Website:     "https://subsidiary.com",
		Address1:    "123 Main Street",
		City:        "New York",
		State:       "NY",
		Country:     "US",
		Zip:         "10001",
		Phone:       "+1-212-555-0100",
	}

	account, err := client.Accounts.CreateChildAccount(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create child account: %v", err)
	}

	fmt.Printf("Created account ID: %s\n", account.ID)
	fmt.Printf("Company: %s\n", account.CompanyName)
	fmt.Printf("Status: %s\n", account.Status)
}

// ExampleAccountsService_GetByID demonstrates retrieving an account by ID.
func ExampleAccountsService_GetByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Retrieve account with full details
	account, err := client.Accounts.GetByID(context.Background(), "acc_123456789", "full")
	if err != nil {
		log.Fatalf("Failed to get account: %v", err)
	}

	fmt.Printf("Account ID: %s\n", account.ID)
	fmt.Printf("Company: %s\n", account.CompanyName)
	fmt.Printf("Status: %s\n", account.Status)
	fmt.Printf("Is Child: %v\n", account.IsChild)
	fmt.Printf("Parent: %v\n", account.Parent)
}

// ExampleAccountsService_UpdateAccountByID demonstrates updating an account.
func ExampleAccountsService_UpdateAccountByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	updateReq := api.UpdateAccountRequest{
		CompanyName:           "Updated Subsidiary Corp",
		Website:               "https://updated-subsidiary.com",
		Address1:              "456 New Street",
		City:                  "San Francisco",
		State:                 "CA",
		Country:               "US",
		Phone:                 "+1-415-555-0200",
		Email:                 "admin@updated-subsidiary.com",
		DefaultDeliveryRegion: "us-west",
	}

	account, err := client.Accounts.UpdateAccountByID(context.Background(), "acc_123456789", updateReq)
	if err != nil {
		log.Fatalf("Failed to update account: %v", err)
	}

	fmt.Printf("Updated account: %s\n", account.CompanyName)
	fmt.Printf("New website: %s\n", account.Website)
	fmt.Printf("Delivery region: %s\n", account.DefaultDeliveryRegion)
}

// ExampleAccountsService_DeactivateAccountByID demonstrates deactivating an account.
func ExampleAccountsService_DeactivateAccountByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Deactivate account (soft delete)
	account, err := client.Accounts.DeactivateAccountByID(context.Background(), "acc_123456789")
	if err != nil {
		log.Fatalf("Failed to deactivate account: %v", err)
	}

	fmt.Printf("Account deactivated: %s\n", account.ID)
	fmt.Printf("Status: %s\n", account.Status)
}

// ExampleAccountsService_ActivateAccountByID demonstrates reactivating an account.
func ExampleAccountsService_ActivateAccountByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Activate a previously deactivated account
	account, err := client.Accounts.ActivateAccountByID(context.Background(), "acc_123456789")
	if err != nil {
		log.Fatalf("Failed to activate account: %v", err)
	}

	fmt.Printf("Account activated: %s\n", account.ID)
	fmt.Printf("Status: %s\n", account.Status)
}

// ExampleAccountsService_List demonstrates listing accounts with filters.
func ExampleAccountsService_List() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// List child accounts with pagination
	opts := api.ListAccountsOptions{
		IsChild: true,
		Status:  "active",
		Limit:   10,
		Offset:  0,
	}

	resp, err := client.Accounts.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("Failed to list accounts: %v", err)
	}

	fmt.Printf("Total accounts: %d\n", resp.Meta.Count)
	for _, account := range resp.Accounts {
		fmt.Printf("- %s (ID: %s, Status: %s)\n",
			account.CompanyName,
			account.ID,
			account.Status,
		)
	}
}

// ExampleAccountsService_GetChildAccountAuthToken demonstrates getting auth token for child account management.
func ExampleAccountsService_GetChildAccountAuthToken() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Get authentication token to manage child account
	authResp, err := client.Accounts.GetChildAccountAuthToken(context.Background(), "acc_123456789")
	if err != nil {
		log.Fatalf("Failed to get child account auth token: %v", err)
	}

	fmt.Printf("Auth token: %s\n", authResp.Token)
	fmt.Printf("Expires at: %s\n", authResp.ExpiresAt)

	// Use the token to create a new client for child account management
	childClient := cachefly.NewClient(
		cachefly.WithToken(authResp.Token),
	)
	_ = childClient
}

// ============================================================================
// Service Examples
// ============================================================================

// ExampleServicesService_Create demonstrates creating a new service.
func ExampleServicesService_Create() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	createReq := api.CreateServiceRequest{
		Name:        "Production CDN",
		UniqueName:  "prod-cdn-example",
		Description: "Main production CDN service",
	}

	service, err := client.Services.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	fmt.Printf("Created service ID: %s\n", service.ID)
	fmt.Printf("Service name: %s\n", service.Name)
	fmt.Printf("Unique name: %s\n", service.UniqueName)
	fmt.Printf("Status: %s\n", service.Status)
}

// ExampleServicesService_GetByID demonstrates retrieving a service by ID.
func ExampleServicesService_GetByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	service, err := client.Services.GetByID(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to get service: %v", err)
	}

	fmt.Printf("Service ID: %s\n", service.ID)
	fmt.Printf("Name: %s\n", service.Name)
	fmt.Printf("Unique name: %s\n", service.UniqueName)
	fmt.Printf("Status: %s\n", service.Status)
	fmt.Printf("Auto SSL: %v\n", service.AutoSSL)
	fmt.Printf("Configuration mode: %s\n", service.ConfigurationMode)
}

// ExampleServicesService_UpdateServiceByID demonstrates updating a service.
func ExampleServicesService_UpdateServiceByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	updateReq := api.UpdateServiceRequest{
		Description:       "Updated production CDN service",
		AutoSSL:           true,
		DeliveryRegion:    "global",
		ConfigurationMode: "advanced",
	}

	service, err := client.Services.UpdateServiceByID(context.Background(), "srv_123456789", updateReq)
	if err != nil {
		log.Fatalf("Failed to update service: %v", err)
	}

	fmt.Printf("Updated service: %s\n", service.Name)
	fmt.Printf("Auto SSL enabled: %v\n", service.AutoSSL)
	fmt.Printf("Configuration mode: %s\n", service.ConfigurationMode)
}

// ExampleServicesService_List demonstrates listing services with filters.
func ExampleServicesService_List() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	opts := api.ListOptions{
		Status:          "active",
		IncludeFeatures: true,
		Limit:           20,
		Offset:          0,
	}

	resp, err := client.Services.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("Failed to list services: %v", err)
	}

	fmt.Printf("Total services: %d\n", resp.Meta.Count)
	for _, service := range resp.Services {
		fmt.Printf("- %s (ID: %s, Status: %s)\n",
			service.Name,
			service.ID,
			service.Status,
		)
	}
}

// ExampleServicesService_ActivateServiceByID demonstrates activating a service.
func ExampleServicesService_ActivateServiceByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	service, err := client.Services.ActivateServiceByID(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to activate service: %v", err)
	}

	fmt.Printf("Service activated: %s\n", service.ID)
	fmt.Printf("Status: %s\n", service.Status)
}

// ExampleServicesService_DeactivateServiceByID demonstrates deactivating a service.
func ExampleServicesService_DeactivateServiceByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	service, err := client.Services.DeactivateServiceByID(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to deactivate service: %v", err)
	}

	fmt.Printf("Service deactivated: %s\n", service.ID)
	fmt.Printf("Status: %s\n", service.Status)
}

// ExampleServicesService_EnableAccessLogging demonstrates enabling access logs for a service.
func ExampleServicesService_EnableAccessLogging() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	logReq := api.EnableAccessLogsRequest{
		LogTarget: "log_target_123",
	}

	service, err := client.Services.EnableAccessLogging(context.Background(), "srv_123456789", logReq)
	if err != nil {
		log.Fatalf("Failed to enable access logging: %v", err)
	}

	fmt.Printf("Access logging enabled for service: %s\n", service.ID)
}

// ExampleServicesService_DeleteAccessLoggingByID demonstrates disabling access logs for a service.
func ExampleServicesService_DeleteAccessLoggingByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	service, err := client.Services.DeleteAccessLoggingByID(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to disable access logging: %v", err)
	}

	fmt.Printf("Access logging disabled for service: %s\n", service.ID)
}

// ExampleServicesService_EnableOriginLogging demonstrates enabling origin logs for a service.
func ExampleServicesService_EnableOriginLogging() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	logReq := api.EnableOriginLogsRequest{
		LogTarget: "log_target_456",
	}

	service, err := client.Services.EnableOriginLogging(context.Background(), "srv_123456789", logReq)
	if err != nil {
		log.Fatalf("Failed to enable origin logging: %v", err)
	}

	fmt.Printf("Origin logging enabled for service: %s\n", service.ID)
}

// ExampleServicesService_DeleteOriginLoggingByID demonstrates disabling origin logs for a service.
func ExampleServicesService_DeleteOriginLoggingByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	service, err := client.Services.DeleteOriginLoggingByID(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to disable origin logging: %v", err)
	}

	fmt.Printf("Origin logging disabled for service: %s\n", service.ID)
}

// ============================================================================
// Service Domain Examples
// ============================================================================

// ExampleServiceDomainsService_Create demonstrates adding a domain to a service.
func ExampleServiceDomainsService_Create() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	createReq := api.CreateServiceDomainRequest{
		Name:           "www.example.com",
		Description:    "Primary domain for production service",
		ValidationMode: "dns",
	}

	domain, err := client.ServiceDomains.Create(context.Background(), "srv_123456789", createReq)
	if err != nil {
		log.Fatalf("Failed to create service domain: %v", err)
	}

	fmt.Printf("Created domain ID: %s\n", domain.ID)
	fmt.Printf("Domain name: %s\n", domain.Name)
	fmt.Printf("Validation mode: %s\n", domain.ValidationMode)
	fmt.Printf("Validation status: %s\n", domain.ValidationStatus)
}

// ExampleServiceDomainsService_GetByID demonstrates retrieving a specific domain.
func ExampleServiceDomainsService_GetByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	domain, err := client.ServiceDomains.GetByID(context.Background(), "srv_123456789", "dom_987654321", "full")
	if err != nil {
		log.Fatalf("Failed to get service domain: %v", err)
	}

	fmt.Printf("Domain ID: %s\n", domain.ID)
	fmt.Printf("Name: %s\n", domain.Name)
	fmt.Printf("Description: %s\n", domain.Description)
	fmt.Printf("Service: %s\n", domain.Service)
	fmt.Printf("Validation status: %s\n", domain.ValidationStatus)
	fmt.Printf("Validation target: %s\n", domain.ValidationTarget)
}

// ExampleServiceDomainsService_UpdateByID demonstrates updating a domain configuration.
func ExampleServiceDomainsService_UpdateByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	updateReq := api.UpdateServiceDomainRequest{
		Description:    "Updated production domain",
		ValidationMode: "http",
	}

	domain, err := client.ServiceDomains.UpdateByID(context.Background(), "srv_123456789", "dom_987654321", updateReq)
	if err != nil {
		log.Fatalf("Failed to update service domain: %v", err)
	}

	fmt.Printf("Updated domain: %s\n", domain.Name)
	fmt.Printf("New description: %s\n", domain.Description)
	fmt.Printf("Validation mode: %s\n", domain.ValidationMode)
}

// ExampleServiceDomainsService_List demonstrates listing all domains for a service.
func ExampleServiceDomainsService_List() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	opts := api.ListServiceDomainsOptions{
		Search: "example.com",
		Limit:  20,
		Offset: 0,
	}

	resp, err := client.ServiceDomains.List(context.Background(), "srv_123456789", opts)
	if err != nil {
		log.Fatalf("Failed to list service domains: %v", err)
	}

	fmt.Printf("Total domains: %d\n", resp.Meta.Count)
	for _, domain := range resp.Domains {
		fmt.Printf("- %s (ID: %s, Status: %s)\n",
			domain.Name,
			domain.ID,
			domain.ValidationStatus,
		)
	}
}

// ExampleServiceDomainsService_ValidationReady demonstrates marking a domain ready for validation.
func ExampleServiceDomainsService_ValidationReady() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	domain, err := client.ServiceDomains.ValidationReady(context.Background(), "srv_123456789", "dom_987654321")
	if err != nil {
		log.Fatalf("Failed to mark domain validation ready: %v", err)
	}

	fmt.Printf("Domain validation initiated: %s\n", domain.Name)
	fmt.Printf("Validation status: %s\n", domain.ValidationStatus)
	fmt.Printf("Validation target: %s\n", domain.ValidationTarget)
}

// ExampleServiceDomainsService_DeleteByID demonstrates removing a domain from a service.
func ExampleServiceDomainsService_DeleteByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	err := client.ServiceDomains.DeleteByID(context.Background(), "srv_123456789", "dom_987654321")
	if err != nil {
		log.Fatalf("Failed to delete service domain: %v", err)
	}

	fmt.Println("Domain removed from service successfully")
}

// ============================================================================
// Origin Examples
// ============================================================================

// ExampleOriginsService_Create demonstrates creating a new origin.
func ExampleOriginsService_Create() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	createReq := api.CreateOriginRequest{
		Type:                   "http",
		Name:                   "Primary Origin",
		Hostname:               "origin.example.com",
		Scheme:                 "https",
		Gzip:                   true,
		CacheByQueryParam:      true,
		TTL:                    3600,
		MissedTTL:              60,
		ConnectionTimeout:      30,
		TimeToFirstByteTimeout: 10,
	}

	origin, err := client.Origins.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create origin: %v", err)
	}

	fmt.Printf("Created origin ID: %s\n", origin.ID)
	fmt.Printf("Origin name: %s\n", origin.Name)
	fmt.Printf("Hostname: %s\n", origin.Hostname)
	fmt.Printf("Scheme: %s\n", origin.Scheme)
}

// ExampleOriginsService_GetByID demonstrates retrieving an origin by ID.
func ExampleOriginsService_GetByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	origin, err := client.Origins.GetByID(context.Background(), "org_123456789", "full")
	if err != nil {
		log.Fatalf("Failed to get origin: %v", err)
	}

	fmt.Printf("Origin ID: %s\n", origin.ID)
	fmt.Printf("Name: %s\n", origin.Name)
	fmt.Printf("Hostname: %s\n", origin.Hostname)
	fmt.Printf("Type: %s\n", origin.Type)
	fmt.Printf("TTL: %d seconds\n", origin.TTL)
	fmt.Printf("Gzip enabled: %v\n", origin.Gzip)
}

// ExampleOriginsService_UpdateByID demonstrates updating an origin configuration.
func ExampleOriginsService_UpdateByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	updateReq := api.UpdateOriginRequest{
		Name:                   "Updated Primary Origin",
		TTL:                    7200,
		MissedTTL:              120,
		Gzip:                   false,
		TimeToFirstByteTimeout: 15,
	}

	origin, err := client.Origins.UpdateByID(context.Background(), "org_123456789", updateReq)
	if err != nil {
		log.Fatalf("Failed to update origin: %v", err)
	}

	fmt.Printf("Updated origin: %s\n", origin.Name)
	fmt.Printf("New TTL: %d seconds\n", origin.TTL)
	fmt.Printf("Gzip enabled: %v\n", origin.Gzip)
}

// ExampleOriginsService_List demonstrates listing origins with filters.
func ExampleOriginsService_List() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	opts := api.ListOriginsOptions{
		Type:   "http",
		Limit:  20,
		Offset: 0,
	}

	resp, err := client.Origins.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("Failed to list origins: %v", err)
	}

	fmt.Printf("Total origins: %d\n", resp.Meta.Count)
	for _, origin := range resp.Origins {
		fmt.Printf("- %s (%s) - %s\n",
			origin.Name,
			origin.ID,
			origin.Hostname,
		)
	}
}

// ExampleOriginsService_Create_s3 demonstrates creating an S3 origin.
func ExampleOriginsService_Create_s3() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	createReq := api.CreateOriginRequest{
		Type:             "s3",
		Name:             "S3 Bucket Origin",
		Hostname:         "my-bucket.s3.amazonaws.com",
		Scheme:           "https",
		AccessKey:        "AKIAIOSFODNN7EXAMPLE",
		SecretKey:        "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		Region:           "us-east-1",
		SignatureVersion: "v4",
		TTL:              3600,
		MissedTTL:        60,
	}

	origin, err := client.Origins.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create S3 origin: %v", err)
	}

	fmt.Printf("Created S3 origin ID: %s\n", origin.ID)
	fmt.Printf("Bucket hostname: %s\n", origin.Hostname)
	fmt.Printf("Region: %s\n", origin.Region)
}

// ExampleOriginsService_Delete demonstrates removing an origin.
func ExampleOriginsService_Delete() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	err := client.Origins.Delete(context.Background(), "org_123456789")
	if err != nil {
		log.Fatalf("Failed to delete origin: %v", err)
	}

	fmt.Println("Origin deleted successfully")
}

// ============================================================================
// Script Config Examples
// ============================================================================

// ExampleScriptConfigsService_Create demonstrates creating a new script configuration.
func ExampleScriptConfigsService_Create() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	createReq := api.CreateScriptConfigRequest{
		Name:                   "Custom Redirect Rules",
		Services:               []string{"srv_123456789", "srv_987654321"},
		ScriptConfigDefinition: "def_redirect_rules",
		MimeType:               "application/json",
		Value: map[string]interface{}{
			"rules": []map[string]string{
				{
					"from": "/old-path",
					"to":   "/new-path",
					"code": "301",
				},
			},
		},
	}

	config, err := client.ScriptConfigs.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create script config: %v", err)
	}

	fmt.Printf("Created config ID: %s\n", config.ID)
	fmt.Printf("Config name: %s\n", config.Name)
	fmt.Printf("Purpose: %s\n", config.Purpose)
	fmt.Printf("Data mode: %s\n", config.DataMode)
}

// ExampleScriptConfigsService_GetByID demonstrates retrieving a script configuration by ID.
func ExampleScriptConfigsService_GetByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	config, err := client.ScriptConfigs.GetByID(context.Background(), "cfg_123456789", "full")
	if err != nil {
		log.Fatalf("Failed to get script config: %v", err)
	}

	fmt.Printf("Config ID: %s\n", config.ID)
	fmt.Printf("Name: %s\n", config.Name)
	fmt.Printf("Definition: %s\n", config.ScriptConfigDefinition)
	fmt.Printf("MIME type: %s\n", config.MimeType)
	fmt.Printf("Services: %v\n", config.Services)
	fmt.Printf("Use schema: %v\n", config.UseSchema)
}

// ExampleScriptConfigsService_UpdateByID demonstrates updating a script configuration.
func ExampleScriptConfigsService_UpdateByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	updateReq := api.UpdateScriptConfigRequest{
		Name:                   "Updated Redirect Rules",
		ScriptConfigDefinition: "def_redirect_rules",
		Services:               []string{"srv_123456789", "srv_555555555"},
		Value: map[string]interface{}{
			"rules": []map[string]string{
				{
					"from": "/old-path",
					"to":   "/new-path",
					"code": "301",
				},
				{
					"from": "/legacy",
					"to":   "/modern",
					"code": "302",
				},
			},
		},
	}

	config, err := client.ScriptConfigs.UpdateByID(context.Background(), "cfg_123456789", updateReq)
	if err != nil {
		log.Fatalf("Failed to update script config: %v", err)
	}

	fmt.Printf("Updated config: %s\n", config.Name)
	fmt.Printf("Services: %v\n", config.Services)
}

// ExampleScriptConfigsService_List demonstrates listing script configurations with filters.
func ExampleScriptConfigsService_List() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	opts := api.ListScriptConfigsOptions{
		Search:          "redirect",
		Status:          "active",
		IncludeFeatures: true,
		Limit:           20,
		Offset:          0,
	}

	resp, err := client.ScriptConfigs.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("Failed to list script configs: %v", err)
	}

	fmt.Printf("Total configs: %d\n", resp.Meta.Count)
	for _, config := range resp.Configs {
		fmt.Printf("- %s (ID: %s, Purpose: %s)\n",
			config.Name,
			config.ID,
			config.Purpose,
		)
	}
}

// ExampleScriptConfigsService_ActivateByID demonstrates activating a script configuration.
func ExampleScriptConfigsService_ActivateByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	config, err := client.ScriptConfigs.ActivateByID(context.Background(), "cfg_123456789")
	if err != nil {
		log.Fatalf("Failed to activate script config: %v", err)
	}

	fmt.Printf("Config activated: %s\n", config.ID)
	fmt.Printf("Name: %s\n", config.Name)
}

// ExampleScriptConfigsService_DeactivateByID demonstrates deactivating a script configuration.
func ExampleScriptConfigsService_DeactivateByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	config, err := client.ScriptConfigs.DeactivateByID(context.Background(), "cfg_123456789")
	if err != nil {
		log.Fatalf("Failed to deactivate script config: %v", err)
	}

	fmt.Printf("Config deactivated: %s\n", config.ID)
}

// ExampleScriptConfigsService_GetSchemaByID demonstrates retrieving a configuration schema.
func ExampleScriptConfigsService_GetSchemaByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	schema, err := client.ScriptConfigs.GetSchemaByID(context.Background(), "cfg_123456789")
	if err != nil {
		log.Fatalf("Failed to get schema: %v", err)
	}

	fmt.Printf("Schema type: %v\n", schema["type"])
	fmt.Printf("Schema properties: %v\n", schema["properties"])
}

// ExampleScriptConfigsService_UpdateValueAsFile demonstrates updating config value using raw file content.
func ExampleScriptConfigsService_UpdateValueAsFile() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Example: VCL configuration file content
	vclContent := []byte(`
		sub vcl_recv {
			if (req.url ~ "^/api/") {
				set req.backend_hint = api_backend;
			}
		}
	`)

	config, err := client.ScriptConfigs.UpdateValueAsFile(context.Background(), "cfg_123456789", vclContent)
	if err != nil {
		log.Fatalf("Failed to update config value: %v", err)
	}

	fmt.Printf("Updated config: %s\n", config.Name)
	fmt.Printf("MIME type: %s\n", config.MimeType)
}

// ExampleScriptConfigsService_GetValueAsFile demonstrates retrieving raw config file content.
func ExampleScriptConfigsService_GetValueAsFile() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	content, err := client.ScriptConfigs.GetValueAsFile(context.Background(), "cfg_123456789")
	if err != nil {
		log.Fatalf("Failed to get config file: %v", err)
	}

	fmt.Printf("Retrieved config content: %v\n", content)
}

// ExampleScriptConfigsService_GetDefinitionByID demonstrates retrieving a script config definition.
func ExampleScriptConfigsService_GetDefinitionByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	definition, err := client.ScriptConfigs.GetDefinitionByID(context.Background(), "def_redirect_rules")
	if err != nil {
		log.Fatalf("Failed to get definition: %v", err)
	}

	fmt.Printf("Definition ID: %s\n", definition.ID)
	fmt.Printf("Definition name: %s\n", definition.Name)
	fmt.Printf("Purpose: %s\n", definition.Purpose)
}

// ============================================================================
// Service Options Examples
// ============================================================================

// ExampleServiceOptionsService_GetOptions demonstrates retrieving service options.
func ExampleServiceOptionsService_GetOptions() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	options, err := client.ServiceOptions.GetOptions(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to get service options: %v", err)
	}

	// Access options using type assertions since ServiceOptions is map[string]interface{}
	if ftp, ok := options["ftp"].(bool); ok {
		fmt.Printf("FTP enabled: %v\n", ftp)
	}

	if cors, ok := options["cors"].(bool); ok {
		fmt.Printf("CORS enabled: %v\n", cors)
	}

	if brotli, ok := options["brotli_compression"].(bool); ok {
		fmt.Printf("Brotli compression: %v\n", brotli)
	}

	if geocountry, ok := options["cachebygeocountry"].(bool); ok {
		fmt.Printf("Cache by geo country: %v\n", geocountry)
	}

	if nocache, ok := options["nocache"].(bool); ok {
		fmt.Printf("No cache: %v\n", nocache)
	}

	if apikey, ok := options["apiKeyEnabled"].(bool); ok {
		fmt.Printf("API key enabled: %v\n", apikey)
	}

	// Access complex nested options
	if reverseProxy, ok := options["reverseProxy"].(map[string]interface{}); ok {
		if enabled, ok := reverseProxy["enabled"].(bool); ok {
			fmt.Printf("Reverse proxy enabled: %v\n", enabled)
		}
	}
}

// ExampleServiceOptionsService_GetOptionsMetadata demonstrates getting available options.
func ExampleServiceOptionsService_GetOptionsMetadata() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	metadata, err := client.ServiceOptions.GetOptionsMetadata(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to get options metadata: %v", err)
	}

	fmt.Printf("Available options count: %d\n", metadata.Meta.Count)

	// Show available options grouped by category
	for _, option := range metadata.Data {
		if option.Type == "dynamic" && option.Property != nil && !option.ReadOnly {
			fmt.Printf("- %s (%s): %s\n", option.Property.Name, option.Property.Type, option.Description)
		}
	}
}

// ExampleServiceOptionsService_UpdateOptions demonstrates updating service options with validation.
func ExampleServiceOptionsService_UpdateOptions() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Create options update request
	updateReq := api.ServiceOptions{
		"ftp":                true,
		"cors":               true,
		"autoRedirect":       false,
		"brotli_compression": true,
		"brotli_support":     true,
		"nocache":            false,
		"cachebygeocountry":  true,
		"followredirect":     true,
		"send-xff":           true,
		"apiKeyEnabled":      true,
		"reverseProxy": map[string]interface{}{
			"enabled":           true,
			"hostname":          "origin.example.com",
			"ttl":               3600,
			"cacheByQueryParam": true,
			"originScheme":      "https",
		},
		"mimeTypesOverrides": []map[string]interface{}{
			{
				"extension": "woff2",
				"mimeType":  "font/woff2",
			},
		},
		"expiryHeaders": []map[string]interface{}{
			{
				"path":       "/static/",
				"expiryTime": 31536000, // 1 year
			},
		},
		"error_ttl": map[string]interface{}{
			"enabled": true,
			"value":   60,
		},
		"ttfb_timeout": map[string]interface{}{
			"enabled": true,
			"value":   30,
		},
	}

	// Update options (automatically validates against metadata)
	options, err := client.ServiceOptions.UpdateOptions(context.Background(), "srv_123456789", updateReq)
	if err != nil {
		// Handle validation errors gracefully
		if validationErr, ok := err.(api.ServiceOptionsValidationError); ok {
			fmt.Printf("Validation failed: %s\n", validationErr.Message)
			for _, fieldErr := range validationErr.Errors {
				fmt.Printf("- %s: %s\n", fieldErr.Field, fieldErr.Message)
			}
			return
		}
		log.Fatalf("Failed to update service options: %v", err)
	}

	// Access updated values
	if cors, ok := options["cors"].(bool); ok {
		fmt.Printf("Updated CORS: %v\n", cors)
	}

	if reverseProxy, ok := options["reverseProxy"].(map[string]interface{}); ok {
		if enabled, ok := reverseProxy["enabled"].(bool); ok {
			fmt.Printf("Reverse proxy enabled: %v\n", enabled)
		}
	}
}

// ExampleServiceOptionsService_UpdateSpecificOption demonstrates updating a single option.
func ExampleServiceOptionsService_UpdateSpecificOption() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Update just the CORS setting
	options, err := client.ServiceOptions.UpdateSpecificOption(
		context.Background(),
		"srv_123456789",
		"cors",
		true,
	)
	if err != nil {
		if validationErr, ok := err.(api.ServiceOptionsValidationError); ok {
			fmt.Printf("Validation failed: %s\n", validationErr.Message)
			return
		}
		log.Fatalf("Failed to update CORS option: %v", err)
	}

	if cors, ok := options["cors"].(bool); ok {
		fmt.Printf("CORS updated to: %v\n", cors)
	}
}

// ExampleServiceOptionsService_IsOptionAvailable demonstrates checking option availability.
func ExampleServiceOptionsService_IsOptionAvailable() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Check if livestreaming option is available for this service
	available, metadata, err := client.ServiceOptions.IsOptionAvailable(
		context.Background(),
		"srv_123456789",
		"livestreaming",
	)
	if err != nil {
		log.Fatalf("Failed to check option availability: %v", err)
	}

	if available {
		fmt.Printf("Livestreaming option is available\n")
		fmt.Printf("Description: %s\n", metadata.Description)
		if metadata.Property != nil {
			fmt.Printf("Type: %s\n", metadata.Property.Type)
		}
	} else {
		fmt.Printf("Livestreaming option is not available for this service\n")
	}
}

// ExampleServiceOptionsService_ValidationErrorHandling demonstrates handling validation errors.
func ExampleServiceOptionsService_ValidationErrorHandling() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Try to update with invalid options
	invalidOptions := api.ServiceOptions{
		"nonexistent_option": true,      // This option doesn't exist
		"cors":               "invalid", // Wrong type (should be boolean)
	}

	_, err := client.ServiceOptions.UpdateOptions(context.Background(), "srv_123456789", invalidOptions)
	if err != nil {
		if validationErr, ok := err.(api.ServiceOptionsValidationError); ok {
			fmt.Printf("Validation failed with %d errors:\n", len(validationErr.Errors))

			for _, fieldErr := range validationErr.Errors {
				fmt.Printf("Field: %s\n", fieldErr.Field)
				fmt.Printf("Error: %s\n", fieldErr.Message)
				fmt.Printf("Code: %s\n", fieldErr.Code)
				fmt.Println("---")
			}

			// Get available options to help user
			availableOptions, _ := client.ServiceOptions.GetAvailableOptionNames(
				context.Background(),
				"srv_123456789",
			)
			fmt.Printf("Available options: %v\n", availableOptions)
		} else {
			log.Fatalf("Unexpected error: %v", err)
		}
	}
}

// ExampleServiceOptionsService_GetLegacyAPIKey demonstrates retrieving the legacy API key.
func ExampleServiceOptionsService_GetLegacyAPIKey() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	apiKeyResp, err := client.ServiceOptions.GetLegacyAPIKey(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to get legacy API key: %v", err)
	}

	fmt.Printf("Legacy API key: %s\n", apiKeyResp.APIKey)
}

// ExampleServiceOptionsService_RegenerateLegacyAPIKey demonstrates regenerating the API key.
func ExampleServiceOptionsService_RegenerateLegacyAPIKey() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	apiKeyResp, err := client.ServiceOptions.RegenerateLegacyAPIKey(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to regenerate legacy API key: %v", err)
	}

	fmt.Printf("New legacy API key: %s\n", apiKeyResp.APIKey)
}

// ExampleServiceOptionsService_DeleteLegacyAPIKey demonstrates deleting the API key.
func ExampleServiceOptionsService_DeleteLegacyAPIKey() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	err := client.ServiceOptions.DeleteLegacyAPIKey(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to delete legacy API key: %v", err)
	}

	fmt.Println("Legacy API key deleted successfully")
}

// ExampleServiceOptionsService_GetProtectServeKey demonstrates retrieving ProtectServe key.
func ExampleServiceOptionsService_GetProtectServeKey() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Get key without hiding secrets
	protectServeResp, err := client.ServiceOptions.GetProtectServeKey(context.Background(), "srv_123456789", false)
	if err != nil {
		log.Fatalf("Failed to get ProtectServe key: %v", err)
	}

	fmt.Printf("ProtectServe key: %s\n", protectServeResp.ProtectServeKey)
	fmt.Printf("Force ProtectServe: %s\n", protectServeResp.ForceProtectServe)
}

// ExampleServiceOptionsService_RecreateProtectServeKey demonstrates regenerating ProtectServe key.
func ExampleServiceOptionsService_RecreateProtectServeKey() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Regenerate the key
	protectServeResp, err := client.ServiceOptions.RecreateProtectServeKey(context.Background(), "srv_123456789", "regenerate")
	if err != nil {
		log.Fatalf("Failed to regenerate ProtectServe key: %v", err)
	}

	fmt.Printf("New ProtectServe key: %s\n", protectServeResp.ProtectServeKey)
}

// ExampleServiceOptionsService_UpdateProtectServeOptions demonstrates updating ProtectServe settings.
func ExampleServiceOptionsService_UpdateProtectServeOptions() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	updateReq := api.UpdateProtectServeRequest{
		ForceProtectServe: "enabled",
		ProtectServeKey:   "new-secret-key-123",
	}

	protectServeResp, err := client.ServiceOptions.UpdateProtectServeOptions(context.Background(), "srv_123456789", updateReq)
	if err != nil {
		log.Fatalf("Failed to update ProtectServe options: %v", err)
	}

	fmt.Printf("Updated ProtectServe key: %s\n", protectServeResp.ProtectServeKey)
	fmt.Printf("Force ProtectServe: %s\n", protectServeResp.ForceProtectServe)
}

// ExampleServiceOptionsService_DeleteProtectServeKey demonstrates deleting ProtectServe key.
func ExampleServiceOptionsService_DeleteProtectServeKey() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	err := client.ServiceOptions.DeleteProtectServeKey(context.Background(), "srv_123456789")
	if err != nil {
		log.Fatalf("Failed to delete ProtectServe key: %v", err)
	}

	fmt.Println("ProtectServe key deleted successfully")
}

// ExampleServiceOptionsService_GetFTPSettings demonstrates retrieving FTP settings.
func ExampleServiceOptionsService_GetFTPSettings() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Get FTP settings without hiding secrets
	ftpResp, err := client.ServiceOptions.GetFTPSettings(context.Background(), "srv_123456789", false)
	if err != nil {
		log.Fatalf("Failed to get FTP settings: %v", err)
	}

	fmt.Printf("FTP password: %s\n", ftpResp.FTPPassword)
}

// ExampleServiceOptionsService_RegenerateFTPPassword demonstrates regenerating FTP password.
func ExampleServiceOptionsService_RegenerateFTPPassword() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	ftpResp, err := client.ServiceOptions.RegenerateFTPPassword(context.Background(), "srv_123456789", false)
	if err != nil {
		log.Fatalf("Failed to regenerate FTP password: %v", err)
	}

	fmt.Printf("New FTP password: %s\n", ftpResp.FTPPassword)
}

// ============================================================================
// Users Examples
// ============================================================================

// ExampleUsersService_Create demonstrates creating a new user.
func ExampleUsersService_Create() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	createReq := api.CreateUserRequest{
		Username:               "john.smith",
		Password:               "SecurePass123!",
		Email:                  "john.smith@example.com",
		FullName:               "John Smith",
		Phone:                  "+1-555-0123",
		PasswordChangeRequired: true,
		Services:               []string{"srv_123456789", "srv_987654321"},
		Permissions: []string{
			"services.read",
			"services.write",
			"users.read",
		},
	}

	user, err := client.Users.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}

	fmt.Printf("Created user ID: %s\n", user.ID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Status: %s\n", user.Status)
}

// ExampleUsersService_GetByID demonstrates retrieving a user by ID.
func ExampleUsersService_GetByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	user, err := client.Users.GetByID(context.Background(), "usr_123456789", "full")
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}

	fmt.Printf("User ID: %s\n", user.ID)
	fmt.Printf("Username: %s\n", user.Username)
	fmt.Printf("Full name: %s\n", user.FullName)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Status: %s\n", user.Status)
	fmt.Printf("Password change required: %v\n", user.PasswordChangeRequired)
	fmt.Printf("Permissions: %v\n", user.Permissions)
}

// ExampleUsersService_UpdateByID demonstrates updating a user.
func ExampleUsersService_UpdateByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	passwordChangeRequired := false
	showDeactivatedServices := true

	updateReq := api.UpdateUserRequest{
		Email:                   "john.smith.updated@example.com",
		FullName:                "John M. Smith",
		Phone:                   "+1-555-9999",
		PasswordChangeRequired:  &passwordChangeRequired,
		ShowDeactivatedServices: &showDeactivatedServices,
		Services:                []string{"srv_123456789", "srv_555555555"},
		Permissions: []string{
			"services.read",
			"services.write",
			"users.read",
			"users.write",
		},
	}

	user, err := client.Users.UpdateByID(context.Background(), "usr_123456789", updateReq)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	fmt.Printf("Updated user: %s\n", user.Username)
	fmt.Printf("New email: %s\n", user.Email)
	fmt.Printf("New full name: %s\n", user.FullName)
}

// ExampleUsersService_List demonstrates listing users with search.
func ExampleUsersService_List() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	opts := api.ListUsersOptions{
		Search: "john",
		Limit:  20,
		Offset: 0,
	}

	resp, err := client.Users.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}

	fmt.Printf("Total users: %d\n", resp.Meta.Count)
	for _, user := range resp.Users {
		fmt.Printf("- %s (%s) - %s\n",
			user.Username,
			user.FullName,
			user.Status,
		)
	}
}

// ExampleUsersService_GetCurrentUser demonstrates retrieving the authenticated user.
func ExampleUsersService_GetCurrentUser() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	user, err := client.Users.GetCurrentUser(context.Background())
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	fmt.Printf("Current user: %s\n", user.Username)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Permissions: %v\n", user.Permissions)
}

// ExampleUsersService_UpdateCurrentUser demonstrates updating the authenticated user.
func ExampleUsersService_UpdateCurrentUser() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	walkthroughVisible := false

	updateReq := api.UpdateUserRequest{
		Phone:              "+1-555-7777",
		WalkthroughVisible: &walkthroughVisible,
	}

	user, err := client.Users.UpdateCurrentUser(context.Background(), updateReq)
	if err != nil {
		log.Fatalf("Failed to update current user: %v", err)
	}

	fmt.Printf("Updated phone: %s\n", user.Phone)
}

// ExampleUsersService_ActivateByID demonstrates activating a user account.
func ExampleUsersService_ActivateByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	user, err := client.Users.ActivateByID(context.Background(), "usr_123456789")
	if err != nil {
		log.Fatalf("Failed to activate user: %v", err)
	}

	fmt.Printf("User activated: %s\n", user.Username)
	fmt.Printf("Status: %s\n", user.Status)
}

// ExampleUsersService_DeactivateByID demonstrates deactivating a user account.
func ExampleUsersService_DeactivateByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	user, err := client.Users.DeactivateByID(context.Background(), "usr_123456789")
	if err != nil {
		log.Fatalf("Failed to deactivate user: %v", err)
	}

	fmt.Printf("User deactivated: %s\n", user.Username)
	fmt.Printf("Status: %s\n", user.Status)
}

// ExampleUsersService_DeleteByID demonstrates removing a user.
func ExampleUsersService_DeleteByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	err := client.Users.DeleteByID(context.Background(), "usr_123456789")
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}

	fmt.Println("User deleted successfully")
}

// ExampleUsersService_GetAllowedPermissions demonstrates retrieving allowed permissions.
func ExampleUsersService_GetAllowedPermissions() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	permissions, err := client.Users.GetAllowedPermissions(context.Background(), "usr_123456789")
	if err != nil {
		log.Fatalf("Failed to get allowed permissions: %v", err)
	}

	fmt.Println("Allowed permissions:")
	for _, perm := range permissions {
		fmt.Printf("- %s\n", perm)
	}
}

// ExampleUsersService_EnableTwoFactorAuth demonstrates enabling 2FA for current user.
func ExampleUsersService_EnableTwoFactorAuth() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	user, err := client.Users.EnableTwoFactorAuth(context.Background())
	if err != nil {
		log.Fatalf("Failed to enable 2FA: %v", err)
	}

	fmt.Printf("2FA enabled for user: %s\n", user.Username)
}

// ExampleUsersService_DisableTwoFactorAuth demonstrates disabling 2FA for current user.
func ExampleUsersService_DisableTwoFactorAuth() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	user, err := client.Users.DisableTwoFactorAuth(context.Background())
	if err != nil {
		log.Fatalf("Failed to disable 2FA: %v", err)
	}

	fmt.Printf("2FA disabled for user: %s\n", user.Username)
}

// ============================================================================
// Certificate Examples
// ============================================================================

// ExampleCertificatesService_Create demonstrates uploading a new TLS/SSL certificate.
func ExampleCertificatesService_Create() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Example PEM-encoded certificate and key (these would be real in production)
	certPEM := `-----BEGIN CERTIFICATE-----
MIIDXTCCAkWgAwIBAgIJAKl4mM7N3BrOMA0GCSqGSIb3DQEBCwUAMEUxCzAJBgNV
... (certificate content) ...
-----END CERTIFICATE-----`

	keyPEM := `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC7VJTUt9Us8cKj
... (private key content) ...
-----END PRIVATE KEY-----`

	createReq := api.CreateCertificateRequest{
		Certificate:    certPEM,
		CertificateKey: keyPEM,
	}

	cert, err := client.Certificates.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create certificate: %v", err)
	}

	fmt.Printf("Created certificate ID: %s\n", cert.ID)
	fmt.Printf("Common name: %s\n", cert.SubjectCommonName)
	fmt.Printf("Subject names: %v\n", cert.SubjectNames)
	fmt.Printf("Expires: %s\n", cert.NotAfter)
}

// ExampleCertificatesService_GetByID demonstrates retrieving a certificate by ID.
func ExampleCertificatesService_GetByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	cert, err := client.Certificates.GetByID(context.Background(), "cert_123456789", "full")
	if err != nil {
		log.Fatalf("Failed to get certificate: %v", err)
	}

	fmt.Printf("Certificate ID: %s\n", cert.ID)
	fmt.Printf("Common name: %s\n", cert.SubjectCommonName)
	fmt.Printf("Subject names: %v\n", cert.SubjectNames)
	fmt.Printf("Valid from: %s\n", cert.NotBefore)
	fmt.Printf("Valid until: %s\n", cert.NotAfter)
	fmt.Printf("Expired: %v\n", cert.Expired)
	fmt.Printf("Expiring soon: %v\n", cert.Expiring)
	fmt.Printf("In use: %v\n", cert.InUse)
	fmt.Printf("Managed: %v\n", cert.Managed)
	fmt.Printf("Services: %v\n", cert.Services)
	fmt.Printf("Domains: %v\n", cert.Domains)
}

// ExampleCertificatesService_List demonstrates listing certificates with search.
func ExampleCertificatesService_List() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	opts := api.ListCertificatesOptions{
		Search: "example.com",
		Limit:  20,
		Offset: 0,
	}

	resp, err := client.Certificates.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("Failed to list certificates: %v", err)
	}

	fmt.Printf("Total certificates: %d\n", resp.Meta.Count)
	for _, cert := range resp.Certificates {
		fmt.Printf("- %s (CN: %s, Expires: %s)\n",
			cert.ID,
			cert.SubjectCommonName,
			cert.NotAfter,
		)
		if cert.Expired {
			fmt.Println("  ⚠️  EXPIRED")
		} else if cert.Expiring {
			fmt.Println("  ⚠️  EXPIRING SOON")
		}
	}
}

// ExampleCertificatesService_Create_withPassword demonstrates uploading a password-protected certificate.
func ExampleCertificatesService_Create_withPassword() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	certPEM := `-----BEGIN CERTIFICATE-----
... (certificate content) ...
-----END CERTIFICATE-----`

	encryptedKeyPEM := `-----BEGIN ENCRYPTED PRIVATE KEY-----
... (encrypted private key content) ...
-----END ENCRYPTED PRIVATE KEY-----`

	createReq := api.CreateCertificateRequest{
		Certificate:    certPEM,
		CertificateKey: encryptedKeyPEM,
		Password:       "keyPassword123",
	}

	cert, err := client.Certificates.Create(context.Background(), createReq)
	if err != nil {
		log.Fatalf("Failed to create certificate with password: %v", err)
	}

	fmt.Printf("Created certificate: %s\n", cert.SubjectCommonName)
}

// ExampleCertificatesService_List_expiring demonstrates finding expiring certificates.
func ExampleCertificatesService_List_expiring() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	opts := api.ListCertificatesOptions{
		Limit: 100,
	}

	resp, err := client.Certificates.List(context.Background(), opts)
	if err != nil {
		log.Fatalf("Failed to list certificates: %v", err)
	}

	fmt.Println("Certificates requiring attention:")
	for _, cert := range resp.Certificates {
		if cert.Expired || cert.Expiring {
			status := "EXPIRING"
			if cert.Expired {
				status = "EXPIRED"
			}
			fmt.Printf("- %s: %s (expires %s)\n",
				status,
				cert.SubjectCommonName,
				cert.NotAfter,
			)
		}
	}
}

// ExampleCertificatesService_Delete demonstrates removing a certificate.
func ExampleCertificatesService_Delete() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	err := client.Certificates.Delete(context.Background(), "cert_123456789")
	if err != nil {
		log.Fatalf("Failed to delete certificate: %v", err)
	}

	fmt.Println("Certificate deleted successfully")
}

// ============================================================================
// Service Referer Rules Examples
// ============================================================================

// ExampleServiceOptionsRefererRulesService_Create demonstrates creating a referer rule.
func ExampleServiceOptionsRefererRulesService_Create() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	createReq := api.CreateRefererRuleRequest{
		Directory:     "/images/",
		Extension:     "jpg,png,gif",
		DefaultAction: "deny",
		Exceptions: []string{
			"*.mycompany.com",
			"*.trustedpartner.com",
			"myapp.example.org",
		},
	}

	rule, err := client.ServiceOptionsRefererRules.Create(context.Background(), "srv_123456789", createReq)
	if err != nil {
		log.Fatalf("Failed to create referer rule: %v", err)
	}

	fmt.Printf("Created rule ID: %s\n", rule.ID)
	fmt.Printf("Directory: %s\n", rule.Directory)
	fmt.Printf("Extensions: %s\n", rule.Extension)
	fmt.Printf("Default action: %s\n", rule.DefaultAction)
	fmt.Printf("Allowed domains: %v\n", rule.Exceptions)
}

// ExampleServiceOptionsRefererRulesService_GetByID demonstrates retrieving a specific referer rule.
func ExampleServiceOptionsRefererRulesService_GetByID() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	rule, err := client.ServiceOptionsRefererRules.GetByID(context.Background(), "srv_123456789", "rule_987654321")
	if err != nil {
		log.Fatalf("Failed to get referer rule: %v", err)
	}

	fmt.Printf("Rule ID: %s\n", rule.ID)
	fmt.Printf("Directory: %s\n", rule.Directory)
	fmt.Printf("Extension: %s\n", rule.Extension)
	fmt.Printf("Default action: %s\n", rule.DefaultAction)
	fmt.Printf("Exceptions: %v\n", rule.Exceptions)
	fmt.Printf("Order: %d\n", rule.Order)
}

// ExampleServiceOptionsRefererRulesService_Update demonstrates updating a referer rule.
func ExampleServiceOptionsRefererRulesService_Update() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	updateReq := api.UpdateRefererRuleRequest{
		DefaultAction: "allow",
		Exceptions: []string{
			"*.malicious-site.com",
			"*.content-scraper.net",
		},
		Order: 10,
	}

	rule, err := client.ServiceOptionsRefererRules.Update(context.Background(), "srv_123456789", "rule_987654321", updateReq)
	if err != nil {
		log.Fatalf("Failed to update referer rule: %v", err)
	}

	fmt.Printf("Updated rule: %s\n", rule.ID)
	fmt.Printf("New default action: %s\n", rule.DefaultAction)
	fmt.Printf("New exceptions: %v\n", rule.Exceptions)
	fmt.Printf("New order: %d\n", rule.Order)
}

// ExampleServiceOptionsRefererRulesService_List demonstrates listing all referer rules for a service.
func ExampleServiceOptionsRefererRulesService_List() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	opts := api.ListRefererRulesOptions{
		Limit:  20,
		Offset: 0,
	}

	resp, err := client.ServiceOptionsRefererRules.List(context.Background(), "srv_123456789", opts)
	if err != nil {
		log.Fatalf("Failed to list referer rules: %v", err)
	}

	fmt.Printf("Total rules: %d\n", resp.Meta.Count)
	for _, rule := range resp.Rules {
		fmt.Printf("- Rule %s: %s (action: %s)\n",
			rule.ID,
			rule.Directory,
			rule.DefaultAction,
		)
	}
}

// ExampleServiceOptionsRefererRulesService_Create_blockAll demonstrates blocking all referers except specific domains.
func ExampleServiceOptionsRefererRulesService_Create_blockAll() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Block all referers by default, only allow specific domains
	createReq := api.CreateRefererRuleRequest{
		Directory:     "/",
		DefaultAction: "deny",
		Exceptions: []string{
			"*.mycompany.com",
			"www.mycompany.net",
			"app.mycompany.io",
		},
	}

	rule, err := client.ServiceOptionsRefererRules.Create(context.Background(), "srv_123456789", createReq)
	if err != nil {
		log.Fatalf("Failed to create blocking rule: %v", err)
	}

	fmt.Printf("Created hotlink protection rule: %s\n", rule.ID)
	fmt.Printf("Protected directory: %s\n", rule.Directory)
}

// ExampleServiceOptionsRefererRulesService_Create_allowWithBlacklist demonstrates allowing all except specific domains.
func ExampleServiceOptionsRefererRulesService_Create_allowWithBlacklist() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Allow all referers by default, block specific domains
	createReq := api.CreateRefererRuleRequest{
		Directory:     "/api/",
		DefaultAction: "allow",
		Exceptions: []string{
			"*.competitor.com",
			"*.scraper-bot.net",
			"badactor.example.org",
		},
	}

	rule, err := client.ServiceOptionsRefererRules.Create(context.Background(), "srv_123456789", createReq)
	if err != nil {
		log.Fatalf("Failed to create blacklist rule: %v", err)
	}

	fmt.Printf("Created blacklist rule: %s\n", rule.ID)
	fmt.Printf("Blocked domains: %v\n", rule.Exceptions)
}

// ExampleServiceOptionsRefererRulesService_Create_mediaProtection demonstrates protecting media files.
func ExampleServiceOptionsRefererRulesService_Create_mediaProtection() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	// Protect video and audio files from hotlinking
	createReq := api.CreateRefererRuleRequest{
		Directory:     "/media/",
		Extension:     "mp4,webm,mp3,m4a",
		DefaultAction: "deny",
		Exceptions: []string{
			"*.mycompany.com",
			"*.mobile-app.mycompany.com",
			"localhost:*", // For development
		},
	}

	rule, err := client.ServiceOptionsRefererRules.Create(context.Background(), "srv_123456789", createReq)
	if err != nil {
		log.Fatalf("Failed to create media protection rule: %v", err)
	}

	fmt.Printf("Media files protected in: %s\n", rule.Directory)
	fmt.Printf("Protected extensions: %s\n", rule.Extension)
}

// ExampleServiceOptionsRefererRulesService_Delete demonstrates removing a referer rule.
func ExampleServiceOptionsRefererRulesService_Delete() {
	client := cachefly.NewClient(
		cachefly.WithToken("your-api-token"),
	)

	err := client.ServiceOptionsRefererRules.Delete(context.Background(), "srv_123456789", "rule_987654321")
	if err != nil {
		log.Fatalf("Failed to delete referer rule: %v", err)
	}

	fmt.Println("Referer rule deleted successfully")
}
