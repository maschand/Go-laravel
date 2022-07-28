# Fiber backend template for [Create Go App CLI](https://github.com/chand19-af/Go-laravel)

<img src="https://img.shields.io/badge/Go-1.17+-00ADD8?style=for-the-badge&logo=go" alt="go version" />&nbsp;<a href="https://goreportcard.com/report/gitlab.com/d6825/golang_template" target="_blank"><img src="https://img.shields.io/badge/Go_report-A+-success?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;<img src="https://img.shields.io/badge/license-Apache_2.0-red?style=for-the-badge&logo=none" alt="license" />

[Go-laravel](https://github.com/chand19-af/Go-laravel) is an Laravel Framework inspired backend template build on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for **fast** development with **zero memory allocation** and **performance** in mind.

## ⚡️ Quick start

1. Create a new project with Fiber:

```bash
git clone git@gitlab.com:d6825/golang_template.git
```

2. Rename `.env.example` to `.env` and fill it with your environment values.
3. Install [Docker](https://www.docker.com/get-started) and the following useful Go tools to your system:

   - [golang-migrate/migrate](https://github.com/golang-migrate/migrate#cli-usage) for apply migrations
   - [github.com/swaggo/swag](https://github.com/swaggo/swag) for auto-generating Swagger API docs
   - [github.com/securego/gosec](https://github.com/securego/gosec) for checking Go security issues
   - [github.com/go-critic/go-critic](https://github.com/go-critic/go-critic) for checking Go the best practice issues
   - [github.com/golangci/golangci-lint](https://github.com/golangci/golangci-lint) for checking Go linter issues
   - [gorm.io/gorm](gorm.io/gorm) for checking GORM issues
   - [github.com/robfig/cron/v3](github.com/robfig/cron/v3) for checking Go cron issues

## 📦 Used packages

| Name                                                                  | Version    | Type       |
| --------------------------------------------------------------------- | ---------- | ---------- |
| [gofiber/fiber](https://github.com/gofiber/fiber)                     | `v2.34.0`  | core       |
| [gofiber/jwt](https://github.com/gofiber/jwt)                         | `v2.2.7`   | middleware |
| [arsmn/fiber-swagger](https://github.com/arsmn/fiber-swagger)         | `v2.31.1`  | middleware |
| [stretchr/testify](https://github.com/stretchr/testify)               | `v1.7.1`   | tests      |
| [golang-jwt/jwt](https://github.com/golang-jwt/jwt)                   | `v4.4.1`   | auth       |
| [joho/godotenv](https://github.com/joho/godotenv)                     | `v1.4.0`   | config     |
| [jmoiron/sqlx](https://github.com/jmoiron/sqlx)                       | `v1.3.5`   | database   |
| [jackc/pgx](https://github.com/jackc/pgx)                             | `v4.16.1`  | database   |
| [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)         | `v1.6.0`   | database   |
| [go-redis/redis](https://github.com/go-redis/redis)                   | `v8.11.5`  | cache      |
| [swaggo/swag](https://github.com/swaggo/swag)                         | `v1.8.2`   | utils      |
| [google/uuid](https://github.com/google/uuid)                         | `v1.3.0`   | utils      |
| [go-playground/validator](https://github.com/go-playground/validator) | `v10.10.0` | utils      |

## 🗄 Template structure

### ./app

**Folder with business logic only**. This directory doesn't care about _what database driver you're using_ or _which caching solution your choose_ or any third-party things.

- `./app/controllers` folder for functional controllers (used in routes)
- `./app/models` folder for describe business models and methods of your project
- `./app/queries` folder for describe queries for models of your project
- `./app/helpers` folder for describe helpers of your project
- `./app/interfaces` folder for describe interfaces for repository, service, and controller of your project
- `./app/providers` folder for describe providers and binding interface of your project
- `./app/services` folder for describe services for controller of your project
- `./app/repositories` folder for describe repository for service of your project
- `./app/tests` folder for describe unit testing for your project

### ./docs

**Folder with API Documentation**. This directory contains config files for auto-generated API Docs by Swagger.

### ./pkg

**Folder with project-specific functionality**. This directory contains all the project-specific code tailored only for your business use case, like _configs_, _middleware_, _routes_ or _utils_.

- `./pkg/configs` folder for configuration functions
- `./pkg/middleware` folder for add middleware (Fiber built-in and yours)
- `./pkg/const` folder for describe `const` of your project
- `./pkg/routes` folder for describe routes of your project
- `./pkg/utils` folder with utility functions (server starter, error checker, etc)

### ./platform

**Folder with platform-level logic**. This directory contains all the platform-level logic that will build up the actual project, like _setting up the database_ or _cache server instance_ and _storing migrations_.

- `./platform/cache` folder with in-memory cache setup functions (by default, Redis)
- `./platform/database` folder with database setup functions (by default, PostgreSQL)
- `./platform/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)

## ⚙️ Configuration

```ini
# .env

# Stage status to start server:
#   - "dev", for start server without graceful shutdown
#   - "prod", for start server with graceful shutdown
STAGE_STATUS="dev"

# Server settings:
SERVER_HOST="127.0.0.1"
SERVER_PORT=3000
SERVER_READ_TIMEOUT=60

# JWT settings:
JWT_SECRET_KEY="\secret"
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=15
JWT_REFRESH_KEY="refresh"
JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT=720

# Database settings:
DB_TYPE="mysql"   # pgx or mysql
DB_HOST="127.0.0.0"
DB_PORT=3306
DB_USER="root"
DB_PASSWORD=""
DB_NAME="digitels_api"
DB_SSL_MODE="disable"
DB_MAX_CONNECTIONS=100
DB_MAX_IDLE_CONNECTIONS=10
DB_MAX_LIFETIME_CONNECTIONS=2

# Gorm settings:
GORM_MYSQL="root:@tcp(127.0.0.1:3306)/midtrans?charset=utf8mb4&parseTime=True&loc=Local"

# Redis settings:
REDIS_HOST="cgapp-redis"
REDIS_PORT=6379
REDIS_PASSWORD=""
REDIS_DB_NUMBER=0

# Goose settings:
GOOSE_DRIVER="mysql"

## ⚠️ License

```
