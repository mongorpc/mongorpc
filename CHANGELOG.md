# Changelog

All notable changes to this project will be documented in this file.

## [1.0.0] - 2026-01-05

### Added
- **Rules Engine**: Added comprehensive CEL-based rules engine for granular access control.
  - `--rules` flag to load rules from YAML file.
  - Support for `CreateDocument` and `UpdateDocument` (with operator inspection) rules.
  - `admin` role bypass for privileged access.
- **Admin API**: Implemented full Admin SDK support.
  - API Key authentication (`X-Admin-Key`, `X-Admin-Secret`).
  - Schema management (Create/Drop Collection, Create Index).
- **Authentication**: Added API Key authentication for client connections.
- **Observability**: Added Prometheus metrics server (default port `:9099`).

### Changed
- **Real-time Sync**: Enhanced Change Streams architecture for reliability.
  - Added support for `updateLookup` to ensure full document availability in snapshots.
- **Dependencies**: Updated MongoDB Driver and gRPC dependencies.

### Fixed
- **Port Conflicts**: Resolved issues with dynamic port allocation and binding.
