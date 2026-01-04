# MongoRPC

**MongoDB + gRPC = MongoRPC**

A high-performance gRPC proxy for MongoDB, enabling seamless database access from any platform.

[![Go Version](https://img.shields.io/badge/go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/license-Apache%202.0-blue.svg)](LICENSE)

## Features

- ✅ **Full MongoDB Type Support** - ObjectId, Decimal128, GeoPoint, UUID, Regex, and more
- ✅ **gRPC API** - High-performance, strongly-typed API with streaming support
- ✅ **Aggregation Pipelines** - Full MongoDB aggregation framework support
- ✅ **Transactions** - ACID transaction support
- ✅ **Change Streams** - Real-time data change notifications
- ✅ **Modern Go** - Built with Go 1.22+, mongo-driver v2, and structured logging

## Quick Start

### Prerequisites

- Go 1.22+
- MongoDB 6.0+
- [buf](https://buf.build) (for proto generation)

### Installation

```bash
# Clone the repository
git clone https://github.com/mongorpc/mongorpc.git
cd mongorpc

# Install dependencies
go mod download

# Build the server
go build -o bin/mongorpc ./cmd/server
```

### Running the Server

```bash
# Start MongoDB (if not running)
docker run -d -p 27017:27017 --name mongodb mongo:7.0

# Run MongoRPC server
./bin/mongorpc --address localhost:50051 --mongodb-uri mongodb://localhost:27017
```

### Configuration

| Flag | Environment | Default | Description |
|------|-------------|---------|-------------|
| `--address` | `SERVER_ADDRESS` | `localhost:50051` | gRPC server address |
| `--mongodb-uri` | `MONGODB_URI` | `mongodb://localhost:27017` | MongoDB connection URI |
| `--database` | `MONGODB_DATABASE` | `mongorpc` | Default database |
| `--debug` | `DEBUG` | `false` | Enable debug logging |

## Client Libraries

- **Web/JavaScript**: [mongorpc-js](https://github.com/mongorpc/mongorpc-js)
- **Dart/Flutter**: [mongorpc-dart](https://github.com/mongorpc/mongorpc-dart)

### JavaScript Example

```typescript
const client = new MongoRPC("http://localhost:50051");

const document = await client
  .database("sample_mflix")
  .collection("movies")
  .document("573a13b0f29313caabd35231")
  .get();
```

### Dart Example

```dart
var client = await mongorpc("localhost", port: 50051);
var collection = client.database("sample_mflix").collection("movies");

var documents = await collection
    .documents()
    .limit(10)
    .sort(by: "title")
    .get();
```

## Project Structure

```
mongorpc/
├── cmd/server/           # Server entry point
├── gen/                  # Generated protobuf/gRPC code
├── internal/
│   ├── api/mongorpc/    # gRPC service implementations
│   ├── config/          # Configuration management
│   └── repository/      # Data access layer
│       └── mongodb/     # MongoDB client
├── proto/               # Protocol buffer definitions
│   └── mongorpc/v1/     # v1 API protos
├── buf.yaml             # Buf configuration
└── buf.gen.yaml         # Code generation config
```

## Development

### Generate Proto Code

```bash
buf generate proto
```

### Build

```bash
go build -o bin/mongorpc ./cmd/server
```

### Test with grpcurl

```bash
# List services
grpcurl -plaintext localhost:50051 list

# Create a document
grpcurl -plaintext -d '{
  "database": "test",
  "collection": "users",
  "document": {
    "fields": {
      "name": {"string_value": "Alice"},
      "age": {"int32_value": 30}
    }
  }
}' localhost:50051 mongorpc.v1.MongoRPC/CreateDocument
```

## API Overview

### Document Operations
- `GetDocument` - Get a single document by ID
- `ListDocuments` - List documents with filtering and pagination
- `CreateDocument` - Create a new document
- `UpdateDocument` - Update an existing document
- `DeleteDocument` - Delete a document
- `BatchGetDocuments` - Get multiple documents (streaming)

### Query Operations
- `RunQuery` - Execute queries with streaming results
- `Aggregate` - Run aggregation pipelines
- `CountDocuments` - Count matching documents
- `Distinct` - Get distinct field values

### Transaction Operations
- `BeginTransaction` - Start a new transaction
- `CommitTransaction` - Commit a transaction
- `AbortTransaction` - Abort a transaction

### Change Streams
- `Watch` - Watch collection for changes
- `WatchDatabase` - Watch database for changes

## License

Copyright 2021-2025 MongoRPC

Licensed under the Apache License, Version 2.0
