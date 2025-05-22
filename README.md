<p align="center">
  <img src="https://www.cachefly.com/wp-content/uploads/2023/10/Thumbnail-About-Us-Video.png" alt="CacheFly Logo" width="200"/>
</p>

<h4 align="center">Go implementation of CacheFly API (2.5.0)</h4>

<hr style="width: 50%; border: 1px solid #000; margin: 20px auto;">

## CacheFly SDK for Go

A Golang SDK for interacting with the [CacheFly CDN API v2.5](https://portal.cachefly.com/api/2.5/docs/).

This SDK is designed to abstract the HTTP API layer and simplify working with CacheFly resources 
and can be used independently as golang package in your project or as the backend foundation for managing CacheFly resources. 

> ⚠️ This SDK is in active development.

## CacheFly

CacheFly CDN is the only CDN built for throughput, delivering rich-media content up to 158% more than other major CDNs.

## ✨ Features

- **Accounts**  
  Manage your CacheFly account and child accounts, including 2FA settings.

- **Services**  
  Create, list, update, activate/deactivate services and control access/origin logging.

- **Service Domains**  
  Add, list, update, delete domains and signal readiness for validation.

- **Service Rules**  
  List and update routing/cache rules and fetch the JSON schema for custom rules.

- **Service Options**  
  View and save basic settings, manage legacy and ProtectServe API keys, and configure FTP.

- **Image Optimization**  
  Fetch, create, update and toggle image‐optimization configurations.

- **Certificates**  
  List, create, retrieve, and delete TLS certificates.

- **Origins**  
  Add, list, update, and remove origin servers.

- **Users**  
  Create, list, update, deactivate users and manage permissions + 2FA.

- **Script Configs**  
  Manage custom scripts: list, create, update, delete.

- **TLS Profiles**  
  Create and manage TLS profiles for your services.


## Installation

```bash
go get github.com/avvvet/cachefly-sdk-go

```

## Quick Start

Copy the snippet below into your project to get started:

```go
client := cachefly.NewClient(cachefly.WithToken("YOUR_API_TOKEN"))
resp, _ := client.Accounts.List(ctx, api.ListAccountsOptions{Offset: 0, Limit: 5})
for _, a := range resp.Accounts {
    fmt.Println(a.ID, a.CompanyName)
}
```

## Example Usage

Below is an example of how to use the CacheFly SDK in your Go project:

1. Create a `.env` file in your project root containing:

   ```dotenv
   CACHEFLY_API_TOKEN=your_real_api_token_here
   ```
2. Run with:

   ```bash
   go run examples/<resource>/<example>.go
   ```

### Accounts

* [List Accounts](examples/accounts/list/main.go)
* [Get Current Account](examples/accounts/get/main.go)
* [Get Account By ID](examples/accounts/getbyid/main.go)
* [Update Current Account](examples/accounts/update/main.go)
* [Update Account By ID](examples/accounts/updatebyid/main.go)
* [Create Child Account](examples/accounts/create/main.go)

### Services

* [List Services](examples/services/list/main.go)
* [Get Service By ID](examples/services/getbyid/main.go)
* [Create Service](examples/services/create/main.go)
* [Update Service By ID](examples/services/updatebyid/main.go)
* [Activate Service](examples/services/activate/main.go)
* [Deactivate Service](examples/services/deactivate/main.go)

### Service Domains

* [List Service Domains](examples/service_domains/list/main.go)  
* [Get Service Domain By ID](examples/service_domains/getbyid/main.go)  
* [Create Service Domain](examples/service_domains/create/main.go)  
* [Update Service Domain](examples/service_domains/update/main.go)  
* [Delete Service Domain](examples/service_domains/delete/main.go)  
* [Signal Domain Validation Ready](examples/service_domains/validationready/main.go)  


### Service Rules

* [List Service Rules](examples/service_rules/list/main.go)  
* [Update Service Rules](examples/service_rules/update/main.go)  
* [Fetch Service Rules JSON Schema](examples/service_rules/schema/main.go)  


### Service Options

* [Get Basic Service Options](examples/service_options/get_basic/main.go)  
* [Save Basic Service Options](examples/service_options/save/main.go)  
* [Get Legacy API Key](examples/service_options/get_legacy_apikey/main.go)  
* [Regenerate Legacy API Key](examples/service_options/regenerate_legacy_apikey/main.go)  
* [Delete Legacy API Key](examples/service_options/delete_legacy_apikey/main.go)  
* [Get ProtectServe Key](examples/service_options/get_protectserve_key/main.go)  
* [Regenerate ProtectServe Key](examples/service_options/recreate_protectserve_key/main.go)  
* [Update ProtectServe Key Options](examples/service_options/update_protectserve_key_options/main.go)  
* [Delete ProtectServe Key](examples/service_options/delete_protectserve_key/main.go)  
* [Get FTP Settings](examples/service_options/get_ftp_settings/main.go)  
* [Regenerate FTP Password](examples/service_options/regenerate_ftp_password/main.go)  

### Service Options – Referer Rules

* [List Service Referer Rules](examples/referer_rules/list/main.go)  
* [Get Service Referer Rule By ID](examples/referer_rules/getbyid/main.go)  
* [Create Service Referer Rule](examples/referer_rules/create/main.go)  
* [Update Service Referer Rule](examples/referer_rules/update/main.go)  
* [Delete Service Referer Rule](examples/referer_rules/delete/main.go)  

### Image Optimization

* [Fetch Configuration](examples/image_optimization/fetch_configuration/main.go)  
* [Create Configuration](examples/image_optimization/create/main.go)  
* [Activate Configuration](examples/image_optimization/activate/main.go)  
* [Deactivate Configuration](examples/image_optimization/deactivate/main.go)  
* [Fetch Default Configuration](examples/image_optimization/fetch_default/main.go)  
* [Fetch Validation Schema](examples/image_optimization/fetch_schema/main.go)  


### Certificates

* [List Certificates](examples/certificates/list/main.go)  
* [Create Certificate](examples/certificates/create/main.go)  
* [Get Certificate By ID](examples/certificates/getbyid/main.go)  
* [Delete Certificate By ID](examples/certificates/delete/main.go)  

### Origins

* [List All Origins](examples/origins/list/main.go)  
* [Get Origin By ID](examples/origins/getbyid/main.go)  
* [Create Origin](examples/origins/create/main.go)  
* [Update Origin By ID](examples/origins/update/main.go)  
* [Delete Origin By ID](examples/origins/delete/main.go)  

### Users

* [List Users](examples/users/list/main.go)  
* [Get User By ID](examples/users/getbyid/main.go)  
* [Create User](examples/users/create/main.go)  
* [Update User By ID](examples/users/update/main.go)  
* [Delete User By ID](examples/users/delete/main.go)  
* [Get Current User](examples/users/me/main.go)  
* [Update Current User](examples/users/update_me/main.go)  
* [Activate User](examples/users/activate/main.go)  
* [Deactivate User](examples/users/deactivate/main.go)  
* [List Allowed Permissions](examples/users/allowed_permissions/main.go)  

### User Security

* [Enable Two-Factor Authentication](examples/users/enable2fa/main.go)  
* [Disable Two-Factor Authentication](examples/users/disable2fa/main.go) 

### Script Configs
* [List Script Configs](examples/script_configs/list/main.go)  
* [Get Script Config By ID](examples/script_configs/getbyid/main.go)  
* [Create Script Config](examples/script_configs/create/main.go)  
* [Update Script Config By ID](examples/script_configs/update/main.go)  
* [Update Script Config Value As File](examples/script_configs/update_file/main.go)  
* [Get Script Config Value As File](examples/script_configs/get_file/main.go)  
* [Fetch Script Config JSON Schema](examples/script_configs/schema/main.go)  
* [Activate Script Config](examples/script_configs/activate/main.go)
* [Deactivate Script Config](examples/script_configs/deactivate/main.go)

### Script Config Definitions
* [List Account Script Config Definitions](examples/script_configs/list_account_definitions/main.go)  
* [Get Script Config Definition By ID](examples/script_configs/definitions_getbyid/main.go)  
* [List Promo Script Config Definitions](examples/script_configs/list_promo/main.go)  

### TLS Profiles

* [List TLS Profiles](examples/tls_profiles/list/main.go)  
* [Get TLS Profile By ID](examples/tls_profiles/getbyid/main.go)  


## Running the Examples with Make

simplify running example scripts.

1. **Ensure** you have a `.env` in your project root with your API token:
   ```dotenv
   CACHEFLY_API_TOKEN=your_real_api_token_here
   ```

  Run example, to list services:

  ```
    make service-list
  ```

  Run example, to list service by ID, You’ll be asked to enter the service ID interactively.

  ```
    make service-getbyid
  ```

## Acceptance Tests

The SDK includes one or two end-to-end (acceptance) tests per resources that hit real CacheFly API endpoints. 

- Require a live API token in a `.env` file at the project root  
- Automatically skip if `.env` or `CACHEFLY_API_TOKEN` is not found  
- Discover valid resource IDs at runtime (no other setup needed)  
- Run against your account/staging environment to verify full CRUD behavior  

Create a `.env` in your repo root with:

```dotenv
CACHEFLY_API_TOKEN=your_real_api_token_here
```


```run acceptance tests 
  go test ./pkg/cachefly/api -timeout 30s
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.