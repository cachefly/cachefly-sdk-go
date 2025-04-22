<p align="center">
  <img src="https://www.cachefly.com/wp-content/uploads/2023/10/Thumbnail-About-Us-Video.png" alt="CacheFly Logo" width="200"/>
</p>

<h4 align="center">Go implementation of CacheFly API (2.5.0)</h4>

<hr style="width: 50%; border: 1px solid #000; margin: 20px auto;">

## CacheFly SDK for Go

A Golang SDK for interacting with the [CacheFly CDN API v2.5](https://portal.cachefly.com/api/2.5/docs/).

This SDK is designed to abstract the HTTP API layer and simplify working with CacheFly resources 
and can be used independently as golang package in your project or as the backend foundation for managing CacheFly resources. 

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

✅ LS Profiles
  - Manage TLS profiles

## Installation

```bash
go get github.com/avvvet/cachefly-sdk-go

```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.