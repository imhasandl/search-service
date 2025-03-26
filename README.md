Search Service Documentation


Overview


The Search Service is a microservice component of a larger social media application built using gRPC. It provides efficient search capabilities across various entities including users, posts, and reports. The service uses PostgreSQL for data storage and implements a clean, maintainable architecture.


Architecture Tech Stack


Language: Go
Communication Protocol: gRPC
Database: PostgreSQL
SQL Generation: sqlc
Database Migration: Goose
Containerization: Docker

Service Structure

search-service/
├── protos/              # Protocol buffer definitions
├── sql/
│   ├── schema/          # Database schema definitions
│   └── queries/         # SQL queries for sqlc
├── internal/
│   └── database/        # Generated database code
├── compose.yaml         # Docker Compose configuration
└── sqlc.yaml            # sqlc configuration


Setup Instructions
Prerequisites
Go 1.19 or later
PostgreSQL
Docker and Docker Compose (for containerized setup)
Protocol Buffer compiler (protoc)