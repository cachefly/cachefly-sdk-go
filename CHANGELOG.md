# Changelog

All notable changes to this project will be documented here.

## [v1.0.5] - 2025-06-17

### Refactor
- Refactor Dynamic service option handling
- update docs

## [v1.0.4] - 2025-06-10

### Added
- Dynamic service option handling, use /options/metadata endpoint to discover available options per service
- Validate service option against metadata before sending

## [v1.0.3] - 2025-06-09

### Added
- Added unit tests for core CRUD resource endpoints
- Added code comments for pksite resource definitions

## [v1.0.2] - 2025-06-02

### Fixed
- Added `omitempty` to UpdateServiceRequest so empty values no longer appear in JSON payload, fixes terraform service apply

## [v1.0.1] - 2025-05-27

### Changed
- Removed unused directories
- Cleaned up internal project structure

## [v1.0.0] - 2025-05-27

### Added
- Initial release of SDK for CacheFly API v2.5