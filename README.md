# go-college

RESTful API for managing colleges, courses, and enrollments, built in Go with `net/http`,
PostgreSQL (pgx), structured logging (zerolog), Prometheus metrics, and OpenTelemetry tracing.

## Requirements

| Tool         | Version                | Notes                                            |
| ------------ | ---------------------- | ------------------------------------------------ |
| Go           | 1.26.4+                | see `go.mod`                                      |
| PostgreSQL   | 18.x (any recent 14+)  | database `go_college` is expected                 |
| `swag` CLI   | v1.16.6                | only needed to regenerate Swagger docs           |
| `golangci-lint` | v2.x                | only needed to run `make lint`                   |
| `make`       | any                    | optional; wraps the build/run commands           |

## 1. Clone & install dependencies

```sh
git clone <repo-url> go-college
cd go-college
go mod download
```

> `go.sum` is gitignored, so `go mod download` (or the first `go build`) will resolve
> and pin dependencies locally.

Install the `swag` CLI (required for `make swagger` / `make build`):

```sh
go install github.com/swaggo/swag/cmd/swag@v1.16.6
```

Make sure the Go bin directory is on your `PATH` so `swag` is found:

```sh
export PATH="$PATH:$(go env GOPATH)/bin"
```

Add that line to your shell profile (`~/.zshrc` / `~/.bashrc`) to make it permanent.

## 2. Environment variables

Configuration lives in `configs/config.yaml`. A few values are injected from the
environment at startup (via `os.Expand`), so these **must be exported** before running:

| Variable                  | Used for                          | Example                  |
| ------------------------- | --------------------------------- | ------------------------ |
| `POSTGRES_HOST`           | PostgreSQL host                   | `localhost`              |
| `POSTGRES_DOCKER_PASSWORD`| Password for the `postgres` user  | `postgres`               |
| `TRACER_HOST`             | OTLP/gRPC tracing collector host  | `localhost`              |

Export them, for example:

```sh
export POSTGRES_HOST=localhost
export POSTGRES_DOCKER_PASSWORD=postgres
export TRACER_HOST=localhost
```

Other settings (ports, DB name, pool sizes, timeouts, log level, rate limiter, etc.)
are edited directly in `configs/config.yaml`. Defaults of note:

- DB user: `postgres`, DB name: `go_college`, port: `5432`
- API server port: `8181`
- Metrics server port: `9191`
- Tracing endpoint: `${TRACER_HOST}:4317`

> If tracing is not available in your environment, set `tracer.enabled: false` in
> `configs/config.yaml` to avoid connection attempts to the collector.

## 3. Set up the database

Create the database and apply the schema migrations:

```sh
createdb go_college    # or: psql -U postgres -c 'CREATE DATABASE go_college;'

# apply migrations
for f in db/migrations/*.sql; do
  psql -U postgres -d go_college -f "$f"
done
```

## 4. Build & run

Using the Makefile (also regenerates Swagger docs, then builds & runs):

```sh
make run
```

Available targets:

```sh
make help          # list targets
make install-tools # install swag, golangci-lint, staticcheck
make swagger       # regenerate Swagger docs into ./api/openapi
make build         # build binary into ./bin/api
make run           # build and run
make lint          # run golangci-lint
make tidy          # go mod tidy
make clean         # remove build artifacts
```

### Linting

The project is checked with [golangci-lint](https://golangci-lint.run) v2 (config in
`.golangci.yml`). Install the tooling once:

```sh
make install-tools
export PATH="$PATH:$(go env GOPATH)/bin"   # if not already on PATH
```

Then run:

```sh
make lint
```

Or run directly with the Go toolchain:

```sh
go run ./cmd/api
```

> Note: run from the project root — the app reads `./configs/config.yaml`,
> `./configs/queries/`, and serves Swagger files from `./api/openapi` using
> relative paths.

## 5. Verify it's running

- API base URL: `http://localhost:8181`
- Swagger UI: `http://localhost:8181/swagger/`
- Metrics (Prometheus): `http://localhost:9191/metrics`

### Main endpoints

| Method | Path                          | Description                     |
| ------ | ----------------------------- | ------------------------------- |
| POST   | `/college/create`             | Create college                  |
| GET    | `/college/all`                | List colleges                   |
| GET    | `/college/{nim}`              | Get / update / delete by NIM    |
| GET    | `/college/name/{name}`        | Find by name                    |
| GET    | `/college/semester/{semester}`| Find by semester                |
| POST   | `/course/create`              | Create course                   |
| GET    | `/course/all`                 | List courses                    |
| GET    | `/course/{code}`              | Get / update / delete by code   |
| POST   | `/enrollment/create`          | Create enrollment               |
| GET    | `/enrollment/nim/{nim}`       | Get enrollments by NIM          |
| PUT    | `/enrollment/{nim}/{course}`  | Update enrollment               |
| DELETE | `/enrollment/{id}`            | Delete enrollment               |

See the Swagger UI for the full, up-to-date list and request/response schemas.

## Project layout

```
cmd/api/            application entrypoint
configs/            config.yaml and SQL query files
db/migrations/      PostgreSQL schema migrations
api/openapi/        generated Swagger/OpenAPI docs
internal/
  app/              wiring / bootstrap
  config/           config loading + env expansion
  handler/          HTTP handlers & router
  service/          business logic
  repository/       data access
  infra/            database, http, logger, metrics, tracer, query loader
  middleware/       CORS, auth-skip, rate limiting, logging
  model/            dto, entity, errors
```
