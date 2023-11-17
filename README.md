# gRPC Service Demo

[![Go Language](https://img.shields.io/badge/Language-Go-blue?logo=go)](https://go.dev/dl/)
[![Docker Container](https://img.shields.io/badge/Container-Docker-blue?logo=docker)](https://docs.docker.com/get-docker/)
[![MIT License](https://img.shields.io/github/license/VILJkid/go-simple-grpc-demo)](/LICENSE)

This is a simple gRPC service demo written in Go. It includes user-related functionality with mock database entries.

**Disclaimer:**

This project is part of an assignment for [Totality Corp](https://www.totalitycorp.com/) and is intended for educational purposes only. It is not intended for use in any production, commercial, or real-world application. The code may not adhere to best practices, security standards, or other considerations required for production-ready software.

Use this project at your own discretion, and be aware that it may contain intentional simplifications or educational constructs that may not be suitable for real-world scenarios.

## Prerequisites

- [Go](https://go.dev/dl/) (at least Go 1.21.4)
- [Docker](https://docs.docker.com/get-docker/) (if you want to run the application in a Docker container)

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/VILJkid/go-simple-grpc-demo.git
   cd go-simple-grpc-demo
   ```

2. Build and run the application locally:

   ```bash
   go mod tidy
   go build -o app
   ./app
   ```

   The gRPC service will be accessible on port `8081`.

3. Alternatively, build and run the application using Docker:

   ```bash
   docker build -t go-simple-grpc-demo .
   docker run -p 8081:8081 go-simple-grpc-demo
   ```

   The gRPC service in the Docker container will also be accessible on port `8081`.

## Running Tests

```bash
cd service
go test -v
```

## License

This project is licensed under the [MIT License](/LICENSE).
