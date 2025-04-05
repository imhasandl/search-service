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

> **Note:** Make sure that you use same token secret in every services

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

#### Request Format

```json
{
   "query": "Some characters to find any users with that characters"
}
```
#### Response

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

### SearchUsersByDate

Searches for users who registered within a specific date range.

```sql
-- name: SearchUsersByDate :many
SELECT * FROM users
WHERE username LIKE $1 || '%'
ORDER BY created_at;
```

The query returns users whose accounts were created between the specified start and end dates.

#### Request Format

```json
{
   "query": "Some characters to find any users with that characters depending on create_at field"
}
```
#### Response

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

### SearchPosts

Searches for posts containing specific keywords or phrases.

```sql
-- name: SearchPosts :many
SELECT * FROM posts
WHERE body LIKE '%' || $1 || '%';
```

The query searches for posts whose content contains the provided search string.

#### Request Format

```json
{
   "query": "Search keyword or phrase"
}
```

#### Response

```json
{
   "post": [
      {
         "id": "post UUID",
         "created_at": "timestamp",
         "updated_at": "timestamp",
         "posted_by": "user UUID",
         "body": "post content",
         "views": 42,
         "likes": 10,
         "liked_by": ["user1 UUID", "user2 UUID"]
      }
   ]
}
```

### SearchPostsByDate

Searches for posts containing specific keywords or phrases, ordered by creation date.

```sql
-- name: SearchPostsByDate :many
SELECT * FROM posts
WHERE body LIKE '%' || $1 || '%'
ORDER BY created_at;
```

The query returns posts containing the search string, sorted by creation date.

#### Request Format

```json
{
   "query": "Search keyword or phrase"
}
```

#### Response

```json
{
   "post": [
      {
         "id": "post UUID",
         "created_at": "timestamp",
         "updated_at": "timestamp",
         "posted_by": "user UUID",
         "body": "post content",
         "views": 42,
         "likes": 10,
         "liked_by": ["user1 UUID", "user2 UUID"]
      }
   ]
}
```

### SearchReports

Searches for reports based on specified criteria.

```sql
-- name: SearchReports :many
SELECT * FROM reports
WHERE reason LIKE '%' || $1 || '%';
```

The query searches for reports whose reason contains the provided search string.

#### Request Format

```json
{
   "query": "Report reason keyword"
}
```

#### Response

```json
{
   "report": [
      {
         "id": "report UUID",
         "reported_at": "timestamp",
         "reported_by": "user UUID",
         "reason": "reason for report"
      }
   ]
}
```

### SearchReportsByDate

Searches for reports based on specified criteria, ordered by report date.

```sql
-- name: SearchReportsByDate :many
SELECT * FROM reports
WHERE reason LIKE '%' || $1 || '%'
ORDER BY reported_at;
```

The query searches for reports containing the search string in their reason, sorted by when they were reported.

#### Request Format

```json
{
   "query": "Report reason keyword"
}
```

#### Response

```json
{
   "report": [
      {
         "id": "report UUID",
         "reported_at": "timestamp",
         "reported_by": "user UUID",
         "reason": "reason for report"
      }
   ]
}
```

## Running the Service

```bash
go run cmd/main.go
```

## Docker Support

The service can be run as part of a Docker Compose setup along with other microservices. When using Docker, make sure to use the Docker Compose specific DB_URL configuration.
