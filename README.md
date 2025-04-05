# Search Service

A microservice for handling searcjing in a social media application, built with Go and gRPC.

## Overview

The Search Service is a core microservice in our social media application architecture that provides fast and efficient search capabilities. It's built with Go and uses gRPC for communication with other services. This service enables users to search for:

- Other users by username, name, or other profile attributes
- Posts and content by keywords, hashtags, or topics
- Trending topics and content based on popularity metrics

The service implements sophisticated indexing and query optimization to ensure search results are relevant and returned quickly, even at scale. It connects to a PostgreSQL database for storing search indexes and communicates with other microservices via RabbitMQ to maintain up-to-date searchable content. The service employs a combination of full-text search capabilities and custom ranking algorithms to deliver personalized search experiences.

## Prerequisites

- Go 1.20 or later
- PostgreSQL database

## Configuration

Create a `.env` file in the root directory with the following variables:

```
PORT=":YOUR_GRPC_PORT"
DB_URL="postgres://username:password@host:port/database?sslmode=disable"
# DB_URL="postgres://username:password@db:port/database?sslmode=disable" // FOR DOCKER COMPOSE
TOKEN_SECRET="YOUR_JWT_SECRET_KEY"
```

## Database Migrations

This service uses Goose for database migrations:

```bash
# Install Goose
go install github.com/pressly/goose/v3/cmd/goose@latest

# Run migrations
goose -dir migrations postgres "YOUR_DB_CONNECTION_STRING" up
```
## gRPC Methods

The service implements the following gRPC methods:

### SearchUsers

Searches users with a specific query.

```sql
-- name: SearchUsers :many
SELECT * FROM users
WHERE username LIKE $1 || '%';
```

The query searches for users whose usernames begin with the provided search string.

### Request Format

```json
{
   "query": "Some characters to find any users with that characters"
}
```
### Response

```json
{
   "users": [
      {
         "id": "user UUID",
         "created_at": "timestamp",
         "updated_at": "timestamp",
         "email": "user email",
         "username": "username",
         "is_premium": true/false,
         "verification_code": 12345,
         "is_verified": true/false
      }
   ]
}
```

### RegisterDeviceToken

Registers a user's device for push notifications.

### Request Format

```json
{
   "user_id": "UUID of the user",
   "device_token": "Device-specific token for push notifications",
   "device_type": "Device platform (e.g., 'android', 'ios', 'web')"
}

### Response Format
```json
{
   "device_token": {
      "id": "UUID of the device token record",
      "user_id": "UUID of the user",
      "device_token": "The device token string",
      "device_type": "Device platform type",
      "created_at": "Timestamp when the record was created",
      "updated_at": "Timestamp when the record was last updated"
   }
}
```

> **Note:** This method delivers notifications via Firebase Cloud Messaging if the user has a registered device token. If Firebase isn't initialized or no device token exists, the method will log this situation but still return a successful response.


## Running the Service

```bash
go run cmd/main.go
```

## Docker Support

The service can be run as part of a Docker Compose setup along with other microservices. When using Docker, make sure to use the Docker Compose specific DB_URL configuration.