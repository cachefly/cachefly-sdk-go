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

The SDK provides support for the following CacheFly API functionalities:

✅ Accounts
  - Retrieve and update account information
  - Manage child accounts
  - Enable/disable two-factor authentication (2FA)

✅ Services
  - List, create, update, activate/deactivate services
  - Enable/disable access and origin logging

✅ Service Domains
  - Manage service domains
  - Signal domain readiness for validation

✅ Service Rules
  - List and update service rules
  - Fetch service rules JSON schema

✅ Service Options
  - Retrieve and update basic service options
  - Manage legacy API keys and ProtectServe keys
  - Handle FTP settings and child accounts

✅ Service Options - Referer Rules
  - List, create, update, and delete referer rules

✅ Service Image Optimization
  - Fetch, create, update, and activate/deactivate image optimization configurations
  - Fetch default configurations and validation schemas

✅ Certificates
  - List, create, retrieve, and delete certificates

✅ Origins
  - List, create, update, and delete origins

✅ Users
  - Retrieve and update user information
  - Manage users and their permissions
  - Activate/deactivate users

✅ User Security
  - Enable/disable two-factor authentication (2FA)

✅ Script Configs
  - List, create, retrieve, update, and delete script configurations

✅ TLS Profiles
  - Manage TLS profiles

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
 
## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.