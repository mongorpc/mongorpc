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
- **Service Layer**: Introduced robust service layer with resilience patterns.
  - **Resilience**: Added Circuit Breaker, Caching, Retry mechanisms, and Pagination for reliable operations.
  - **Query Builder**: Added support for complex queries including `limit`, `sort`, `filters`, and `text search`.
- **CRUD Operations**: Expanded CRUD capabilities.
  - Added `InsertMany` and `BulkInsert` for high-throughput data ingestion.
  - Added specific methods for `ReplaceDocument`, `UpdateDocument`, and `DeleteDocument`.
- **Authentication**: Added API Key authentication for client connections.
- **Observability**: Added Prometheus metrics server (default port `:9099`).
- **CI/CD**: Added Nightly GitHub Workflow and automated Docker Hub release process.

### Changed
- **Real-time Sync**: Enhanced Change Streams architecture for reliability.
  - Added support for `updateLookup` to ensure full document availability in snapshots.
  - Server-side rule evaluation for secure change streams.
- **Dependencies**: Updated MongoDB Driver and gRPC dependencies.
- **Refactoring**: Significant refactoring of Server and Client for better modularity (v2 proto, middleware integration).

### Fixed
- **Port Conflicts**: Resolved issues with dynamic port allocation and binding.
- **Client Stability**: Fixed nil pointer dereference issues in client.
- **Bugs**: Fixed issue `RVV-B0012`.
