# StudHub CLI

StudHub CLI is a Go-based scaffolding tool for generating new microservices for the StudHub inter-college student networking platform. It creates a starter project with a basic HTTP service, Docker support, CI/CD workflow, and common local-development files.

## Features

- Generates a new service directory with a Go module
- Creates a minimal HTTP server with a health endpoint
- Adds Docker, Git, and Makefile support
- Includes a GitHub Actions CI/CD workflow template
- Initializes dependencies and runs formatting automatically

## Project Structure

```text
studhub-cli/
├── cmd/
│   └── studhub-generate/
│       ├── cmd/
│       │   ├── root.go
│       │   └── service.go
│       └── main.go
├── pkg/
│   └── studhub-generator/
│       ├── generator.go
│       ├── template.go
│       └── templates/
│           ├── ci-cd.yaml.tmpl
│           ├── docker_file.tmpl
│           ├── docker-compose.yml.tmpl
│           ├── dockerignore.tmpl
│           ├── gitignore.tmpl
│           ├── main.go.tmpl
│           ├── makefile.tmpl
│           └── README.md.tmpl
├── go.mod
├── LICENSE
└── README.md
```

## Prerequisites

- Go 1.27+
- Git
- Docker (optional, for local container-based workflows)

## Installation

Clone the repository and build the CLI:

```bash
git clone https://github.com/MorningBlossom/studhub-cli.git
cd studhub-cli
go build ./cmd/studhub-generate
```

You can also install it globally:

```bash
go install ./cmd/studhub-generate
```

## Usage

Generate a new StudHub microservice:

```bash
studhub-generate service auth-service
```

This creates a folder named `auth-service` in the current directory with the generated scaffold.

## Generated Service Layout

A generated service includes:

```text
my-service/
├── cmd/
│   └── service/
│       └── main.go
├── .github/
│   └── workflows/
│       └── ci-cd.yaml
├── Dockerfile
├── .dockerignore
├── .gitignore
├── Makefile
├── README.md
├── go.mod
└── go.sum
```

## Generated Application Behavior

The generated service exposes a simple HTTP health endpoint:

- `GET /health` returns a 200 OK status
- The app listens on port `:8080`
- It is intended as a starting point for a StudHub microservice behind the main API gateway

## Local Development

After generating a service, move into the new project and use the generated Makefile:

```bash
cd my-service
make run
```

The generated project also includes the standard Go module setup and formatting commands used during scaffolding:

```bash
go mod init
go get github.com/jackc/pgx/v5
go mod tidy
go fmt ./...
```

## CI/CD and Containerization

The generator includes templates for:

- GitHub Actions workflow automation
- Docker image build support
- Container ignore rules
- Initial repository documentation and project health checks

## Notes

This project is intentionally focused on bootstrapping microservices for the StudHub platform and providing a common baseline for local development and deployment.
