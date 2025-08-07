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

### Building and Running a Single Service

You can build and run any individual service directly using its Dockerfile. For example, to run the `menu-service`:

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