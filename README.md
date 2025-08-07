# Order Management System (OMS)

![Go Version](https://img.shields.io/badge/go-1.24-blue.svg)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)

A comprehensive, microservices-based Order Management System designed for restaurant environments like fast-food chains. This project serves as a practical, professional example of building a scalable, resilient, and maintainable backend system using modern Go and cloud-native technologies.

## Architecture Overview

This system is built on a microservices architecture, where each service is responsible for a distinct business capability. Services communicate with each other asynchronously via a message broker to ensure loose coupling and high availability. Synchronous communication for queries is handled via gRPC.

*(A proper architecture diagram will be added here once the core services are in place.)*

## Tech Stack

| Component         | Technology                               | Rationale                                                                |
| ----------------- | ---------------------------------------- | ------------------------------------------------------------------------ |
| **Language**      | Golang (1.24)                           | High performance, strong concurrency primitives, and excellent for cloud-native apps. |
| **API Layer**     | gRPC & REST (via Gorilla Mux)            | Efficient inter-service communication and standard client-facing REST APIs. |
| **Database**      | PostgreSQL                               | Robust, reliable, and excellent for relational data integrity.           |
| **Messaging**     | RabbitMQ                                 | A mature and reliable message broker for asynchronous event-driven flows. |
| **Containerization**| Docker                                   | For consistent development, testing, and production environments.        |
| **Container Build**| Multi-stage & Distroless Images          | Creates minimal, secure, and production-ready container images.         |
| **CI/CD**         | GitHub Actions                           | To automate testing, building, and deployment pipelines.                 |
| **Deployment**    | Kubernetes & Helm                        | For orchestrating and managing containerized applications at scale.      |
| **Authentication**| Role-Based Access Control (RBAC) via JWT | Secure and stateless authentication for users and services.              |
| **Local Dev**     | Docker Compose                           | To easily manage and run the entire stack on a local machine.            |


## Project Structure

The project uses a monorepo structure to simplify dependency management and cross-service changes, following standard Go project layout conventions.

```text
/
├── .github/              # CI/CD workflows (GitHub Actions)
├── cmd/                  # Main applications for each service (entry points)
├── internal/             # Private application and library code
│   ├── server/           # Shared server bootstrapping logic
│   ├── service/          # Service-specific business logic (handler, service, store)
│   └── models/           # Go structs for the business domain
├── pkg/                  # Public shared libraries (if any)
├── scripts/              # Helper scripts (e.g., database migrations)
├── docker-compose.yml    # Local development environment setup
└── README.md             # This file
```


## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.24 or higher)
- [Docker](https://docs.docker.com/get-docker/)

### Running the Full System with Docker Compose (Recommended)
Docker Compose is the primary tool for running the entire application stack—including all microservices, databases, and message brokers—on your local machine.

#### 1. Start All Services
This command builds the images for your services (if they've changed) and starts all containers defined in docker-compose.yml. You will see the logs from all services streamed to your terminal.

```sh
docker-compose up --build
```

#### 2. Start All Services in the Background
To run the entire stack in detached mode (-d), which runs the containers in the background and leaves your terminal free, use:

```sh
docker-compose up -d
```

#### 3. View Logs
If you are running in detached mode, you can view logs at any time.

To follow the logs of all running services:
```sh
docker-compose logs -f
```
To follow the logs of a specific service (e.g., menu-service):

```sh
docker-compose logs -f menu-service
```

#### 4. Check Container Status
To see a list of all containers managed by Compose and check their status (e.g., Up, Exited):
```sh
docker-compose ps
```

#### 5. Stop All Services
This command stops and removes all containers, networks, and the default volumes created by docker-compose up.
```sh
docker-compose down
```


### Building and Running a Single Service (for Isolated Testing)
While Docker Compose is recommended, you can still build and run any individual service for focused debugging. For example, to run the menu-service:
`menu-service`:

1.  **Build the Docker image:**
    From the project root directory, run:
    ```sh
    docker build -f cmd/menu-service/Dockerfile -t menu-service .
    ```

2.  **Run the container:**
    ```sh
    docker run --rm -p 8080:8080 menu-service
    ```
    The service will be available at `http://localhost:8080`.

*(Instructions for running the full system with Docker Compose will be added once it is configured.)*

## Development Workflow

This project follows the **Feature Branch Workflow** with **Conventional Commits**.

1.  All work must be done on a feature branch (e.g., `feat/my-new-feature`, `fix/a-bug`).
2.  Direct pushes to the `main` branch are disabled.
3.  All changes must be submitted via a Pull Request (PR).
4.  PRs must pass all automated CI checks (testing, linting) before being merged.
5.  PRs are merged using the **"Squash and Merge"** strategy to maintain a clean and linear history on the `main` branch.