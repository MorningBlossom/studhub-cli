---------------------------------------------------------------
# StudHub CLI
StudHub CLI is a Go-based scaffolding tool for generating new microservices for the **StudHub inter-college student networking platform**.

The CLI automates the creation of a standard microservice structure, including a basic Go HTTP server, Docker configuration, Git support, Makefile, GitHub Actions CI/CD workflow, and common development files.

The goal is to provide developers with a **consistent, ready-to-use starting point** for every new StudHub microservice while reducing repetitive setup work.

---

## ✨ Features

* 🚀 Generate a new StudHub microservice with a single command
* 🐹 Automatically initialize a Go module
* 🐳 Docker and Docker Compose support
* 🛠️ Makefile with common development commands
* 🔧 Git and `.gitignore` configuration
* 🔄 GitHub Actions CI/CD workflow
* 🏷️ GitHub Labels & Service Organization
* 👥 CODEOWNERS support for controlled PR approvals
* 📄 Automatically generate project documentation
* 📦 Initialize required Go dependencies
* 🎨 Automatically format generated Go code
* 💻 Supports installation through **Homebrew** and **Scoop**

---

# 📁 Project Structure

```text
studhub-cli/
├── cmd/
│   └── studhub-generate/
│       ├── cmd/
│       │   ├── root.go
│       │   └── service.go
│       └── main.go
│
├── pkg/
│   └── studhub-generate/
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
│
├── go.mod
├── LICENSE
└── README.md
```

---

# 🏗️ Architecture

The CLI is organized into two main areas:

## `cmd/studhub-generate`

This folder contains the command-line interface implementation.

### `cmd/root.go`

`root.go` contains the basic CLI configuration.

It defines the root command and initializes the fundamental CLI properties such as:

* `Use`
* `Short`
* `Long`
* Command execution through the `Execute()` function

The root command acts as the entry point for the CLI command structure.

### `cmd/service.go`

`service.go` defines the `service` command.

It contains the logic required to accept a service name and trigger the microservice generation process.

For example:

```bash
studhub-generate service auth-service
```

The command passes the service information to the generator, which then creates the required project structure.

### `main.go`

The `main.go` file is the outer entry point of the application.

It calls the `Execute()` function defined in the CLI command package to start the application.

The overall flow is:

```text
main.go
   ↓
Execute()
   ↓
root command
   ↓
service command
   ↓
generator
   ↓
templates
   ↓
generated microservice
```

---

# 🧩 `pkg/studhub-generate`

This package contains the actual microservice generation logic.

## `generator.go`

`generator.go` is responsible for executing the generation process.

It handles the commands and operations required to create the new service, including:

* Creating the service directory
* Creating the required project structure
* Initializing the Go module
* Installing required dependencies
* Running `go mod tidy`
* Formatting the generated Go code
* Generating files from templates

The generator essentially coordinates the complete scaffolding process.

---

## `template.go`

`template.go` is responsible for creating the generated files from the predefined templates.

It loads the appropriate template files, applies the required service-specific information, and generates the final files inside the newly created microservice.

This keeps the generation logic separate from the actual template content.

---

# 📄 Templates

The `templates` directory contains the default files that every generated StudHub microservice should have.

```text
templates/
├── ci-cd.yaml.tmpl
├── docker_file.tmpl
├── docker-compose.yml.tmpl
├── dockerignore.tmpl
├── gitignore.tmpl
├── main.go.tmpl
├── makefile.tmpl
└── README.md.tmpl
```

These files act as blueprints for the generated service.

Instead of manually creating the same configuration files for every new microservice, the CLI uses these templates to generate them automatically.

---

# 🚀 Installation

## Prerequisites

Before installing StudHub CLI, make sure you have:

* Go 1.27+
* Git
* Homebrew for macOS **or**
* Scoop for Windows

---

## Clone and Build

Clone the repository:

```bash
git clone https://github.com/MorningBlossom/studhub-cli.git
```

Move into the project:

```bash
cd studhub-cli
```

Build the CLI:

```bash
go build ./cmd/studhub-generate
```

---

## Install Globally Using Go

You can also install the CLI globally using:

```bash
go install ./cmd/studhub-generate
```

After installation, the `studhub-generate` command can be used from your terminal.

---

# 🪄 Usage

To generate a new microservice:

```bash
studhub-generate service auth-service
```

This creates a new directory named `auth-service` in the current working directory.

---

# 📦 Generated Service Structure

For example:

```text
auth-service/
├── cmd/
│   └── service/
│       └── main.go
│
├── .github/
│   ├── CODEOWNERS
│   └── workflows/
│       └── ci-cd.yaml
│
├── Dockerfile
├── docker-compose.yml
├── .dockerignore
├── .gitignore
├── Makefile
├── README.md
├── go.mod
└── go.sum
```

The generated project provides a common baseline for development and deployment.

---

# 💻 Local Development

After generating a service, move into the newly created project:

```bash
cd auth-service
```

The generated Makefile provides common development commands.

For example:

```bash
make run
```

The generated project also performs the standard Go module setup and formatting operations:

```bash
go mod init github.com/MorningBlossom/<service-name>
go get github.com/jackc/pgx/v5
go mod tidy
go fmt ./...
```

---

# 🐳 Docker Support

Every generated service includes Docker-related configuration.

The generator provides:

* `Dockerfile`
* `docker-compose.yml`
* `.dockerignore`

This gives each microservice a consistent containerization setup and makes it easier to run services locally or prepare them for deployment.

---

# 🔄 CI/CD

The generated project includes a GitHub Actions workflow:

```text
.github/
└── workflows/
    └── ci-cd.yaml
```

The workflow provides a standard foundation for automating build, validation, and deployment-related processes.

---

# 🏷️ GitHub Labels & Service Organization

The StudHub project uses **GitHub Labels** to organize and categorize work based on the service or component it belongs to.

Each microservice can have its own dedicated label, making it easier to identify and filter:

* Issues
* Feature requests
* Bug reports
* Pull requests
* Development tasks

For the StudHub CLI, the label used is:

```text
studhub-generate
```

This allows all issues and development activities related to the **StudHub CLI** to be easily identified and separated from work belonging to other StudHub services.

For example:

```text
GitHub Project
│
├── studhub-generate
│   ├── CLI development
│   ├── Template changes
│   ├── Generator improvements
│   ├── Bug fixes
│   └── Documentation
│
├── <service-label>
│   ├── Service-specific issues
│   └── Service-specific development
│
└── <another-service-label>
    ├── Service-specific issues
    └── Service-specific development
```

Using service-specific labels provides a consistent way to **segregate and track work across the different StudHub microservices** without mixing issues or tasks between projects.

The `studhub-generate` label therefore acts as the identifier for all GitHub work associated with the StudHub CLI.


# 👥 CODEOWNERS & Pull Request Approval

The`.github` configuration also includes **CODEOWNERS**.

CODEOWNERS is used to define the specific developers or groups responsible for particular parts of the repository.

This helps ensure that pull requests are reviewed and approved by the appropriate owner or team.

For example, a CODEOWNERS configuration can specify a particular person or group as the owner of the repository:

```text
* @studhub-team
```

When a pull request modifies files covered by a CODEOWNERS rule, GitHub can automatically request reviews from the specified owner.

This provides better ownership and review control across StudHub microservices.

> The exact approval requirements depend on the repository's GitHub branch-protection and review settings.

---

# 📦 Package Manager Installation

To make installation easier for developers, StudHub CLI can also be distributed through package managers.

The goal is to provide a **zero-setup experience** where developers can install and update the CLI without manually downloading binaries or configuring environment variables.

## Why Homebrew and Scoop?

### 1. No manual PATH configuration

Developers don't need to manually configure system variables such as:

```text
PATH
~/.zshrc
```

The package managers handle the required setup.

### 2. Easy updates

Once a newer version of the CLI is released, developers can update it through the package manager instead of downloading a new binary manually.

For example:

```bash
brew upgrade studhub-generate
```

or:

```bash
scoop update studhub-generate
```

### 3. Available globally

After installation, the CLI can be used from any directory:

```bash
studhub-generate service auth-service
```

---

# 🪟 Installing Scoop on Windows

Scoop is currently used for the Windows installation flow.

Open **PowerShell** and run:

```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

Then install Scoop:

```powershell
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
```

---

# 🪄 Installing StudHub CLI Using Scoop

Add the StudHub Scoop bucket:

```powershell
scoop bucket add studhub https://github.com/MorningBlossom/scoop-studhub-cli
```

Install the CLI:

```powershell
scoop install studhub-generate
```

After installation:

```powershell
studhub-generate service <service-name>
```

Example:

```powershell
studhub-generate service auth-service
```

---

# 🍺 Homebrew

For macOS, the CLI can be distributed through Homebrew.

Once the StudHub Homebrew formula/tap is configured, developers can install the CLI through Homebrew and keep it updated using the standard Homebrew commands.

# 🍺 Installing StudHub CLI Using Homebrew

Add the StudHub Homebrew tap:

```bash
brew tap MorningBlossom/studhub-cli
```

Install the CLI:

```bash
brew install studhub-generate
```

After installation:

```bash
studhub-generate service <service-name>
```

Example:

```bash
studhub-generate service auth-service
```

---




# 🔁 Generation Workflow

The complete generation process can be summarized as:

```text
Developer
    │
    │ studhub-generate service auth-service
    ▼
┌──────────────────────┐
│      main.go         │
│   Execute()          │
└──────────┬───────────┘
           ▼
┌──────────────────────┐
│     root.go          │
│  CLI configuration   │
└──────────┬───────────┘
           ▼
┌──────────────────────┐
│    service.go        │
│  Service command     │
└──────────┬───────────┘
           ▼
┌──────────────────────┐
│   generator.go       │
│ Generation process   │
└──────────┬───────────┘
           ▼
┌──────────────────────┐
│    template.go       │
│ Template processing  │
└──────────┬───────────┘
           ▼
┌──────────────────────┐
│      templates/      │
│  Default file        │
│  blueprints          │
└──────────┬───────────┘
           ▼
┌──────────────────────┐
│  Generated Service   │
│                      │
│ Go + Docker + Git    │
│ Makefile + CI/CD     │
│ CODEOWNERS + README  │
└──────────────────────┘
```

---

# 🎯 Why StudHub CLI?

Creating a new microservice usually requires repeating the same setup steps:

* Creating directories
* Initializing Go modules
* Adding dependencies
* Creating a basic HTTP server
* Configuring Docker
* Creating a Makefile
* Setting up CI/CD
* Adding Git configuration
* Creating documentation
* Configuring repository ownership

StudHub CLI automates these repetitive tasks and ensures that every StudHub microservice starts with a **consistent project structure and development baseline**.

Instead of spending time setting up infrastructure, developers can immediately start working on the actual business functionality of their service.

---

# 📝 Notes

StudHub CLI is intentionally focused on **microservice bootstrapping** for the StudHub platform.

The generated project provides a common baseline for:

* Local development
* Version control
* Containerization
* CI/CD
* Code ownership
* Documentation
* Go development

The templates can be extended over time as StudHub's development standards evolve.

---------------------------------------------------------------
