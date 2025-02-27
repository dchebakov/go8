# Introduction
            .,*/(#####(/*,.                               .,*((###(/*.
        .*(%%%%%%%%%%%%%%#/.                           .*#%%%%####%%%%#/.
      ./#%%%%#(/,,...,,***.           .......          *#%%%#*.   ,(%%%#/.
     .(#%%%#/.                    .*(#%%%%%%%##/,.     ,(%%%#*    ,(%%%#*.
    .*#%%%#/.    ..........     .*#%%%%#(/((#%%%%(,     ,/#%%%#(/#%%%#(,
    ./#%%%(*    ,#%%%%%%%%(*   .*#%%%#*     .*#%%%#,      *(%%%%%%%#(,.
    ./#%%%#*    ,(((##%%%%(*   ,/%%%%/.      .(%%%#/   .*#%%%#(*/(#%%%#/,
     ,#%%%#(.        ,#%%%(*   ,/%%%%/.      .(%%%#/  ,/%%%#/.    .*#%%%(,
      *#%%%%(*.      ,#%%%(*   .*#%%%#*     ./#%%%#,  ,(%%%#*      .(%%%#*
       ,(#%%%%%##(((##%%%%(*    .*#%%%%#(((##%%%%(,   .*#%%%##(///(#%%%#/.
         .*/###%%%%%%%###(/,      .,/##%%%%%##(/,.      .*(##%%%%%%##(*,
              .........                ......                .......
A starter kit for Go API development. Inspired by [How I write HTTP services after eight years](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html).

However, I wanted to use [chi router](https://github.com/go-chi/chi) which is more common in the community, [sqlx](https://github.com/jmoiron/sqlx) for database operations and design towards layered architecture (handler -> business logic -> repository).

It is still in early stages, and I do not consider it is completed until all integration tests are completed.

In short, this kit is a Go + Postgres + Chi Router + sqlx + ent + unit testing starter kit for API development.

# Motivation

On the topic of API development, there are two opposing camps between using framework (like [echo](https://github.com/labstack/echo), [gin](https://github.com/gin-gonic/gin), [buffalo](http://gobuffalo.io/)) and starting small and only adding features you need through various libraries. 

However, the second option isn't that straightforward. you will want to structure your project in such a way that there are clear separation of functionalities for your controller, business logic and database operations. Dependencies need to be injected from outside to inside. Being modular, swapping a library like a router or database library to a different one becomes much easier.

# Features

This kit is composed of standard Go library together with some well-known libraries to manage things like router, database query and migration support.

  - [x] Framework-less and net/http compatible handler
  - [x] Router/Mux with [Chi Router](https://github.com/go-chi/chi)
  - [x] Database Operations with [sqlx](https://github.com/jmoiron/sqlx)
  - [x] Database Operations with [ent](https://entgo.io/docs/getting-started)
  - [x] Database migration with [golang-migrate](https://github.com/golang-migrate/migrate/)
  - [x] Input [validation](https://github.com/go-playground/validator) that returns multiple error strings
  - [x] Read all configurations using a single `.env` file or environment variable
  - [x] Clear directory structure, so you know where to find the middleware, domain, server struct, handle, business logic, store, configuration files, migrations etc. 
  - [x] (optional) Request log that logs each user uniquely based on host address
  - [x] CORS
  - [x] Scans and auto-generate [Swagger](https://github.com/swaggo/swag) docs using a declarative comments format 
  - [x] Custom model JSON output
  - [x] Filters (input port), Resource (output port) for pagination and custom response respectively.
  - [x] Cache layer
  - [x] Uses [Task](https://taskfile.dev) to simplify various tasks like mocking, linting, test coverage, hot reload etc
  - [x] Unit testing of repository, use case, and handler
  - [ ] End-to-end test using ephemeral docker containers

# Quick Start

It is advisable to use the latest [Go version installation](#appendix) (>= v1.17). Optionally `docker` and `docker-compose` for easier start up.

Get it

    git clone https://github.com/gmhafiz/go8
    cd go8

Set database credentials by either

1. Filling in your database credentials in `.env` by making a copy of `env.example` first.
```shell
 cp env.example .env
```

2. Or by exporting into environment variable

```shellexport DB_DRIVER=postgres
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=user
export DB_PASS=password
export DB_NAME=go8_db
```

Have a database ready either by installing them yourself or the following command. The `docker-compose.yml` will use database credentials set in `.env` file which is initialized by the previous step if you chose that route. Optionally, you may want redis as well.

    docker-compose up -d postgres

Once the database is up you may run the migration with,

    go run cmd/extmigrate/main.go

Run the API with the following command. For the first time run, dependencies will be downloaded first.

    go run cmd/go8/main.go


You will see the address the API is running at as well as all registered routes.

    2021/10/31 10:49:11 Starting API version: v0.12.0
    2021/10/31 10:49:11 Connecting to database...
    2021/10/31 10:49:11 Database connected
            .,*/(#####(/*,.                               .,*((###(/*.
        .*(%%%%%%%%%%%%%%#/.                           .*#%%%%####%%%%#/.
      ./#%%%%#(/,,...,,***.           .......          *#%%%#*.   ,(%%%#/.
     .(#%%%#/.                    .*(#%%%%%%%##/,.     ,(%%%#*    ,(%%%#*.
    .*#%%%#/.    ..........     .*#%%%%#(/((#%%%%(,     ,/#%%%#(/#%%%#(,
    ./#%%%(*    ,#%%%%%%%%(*   .*#%%%#*     .*#%%%#,      *(%%%%%%%#(,.
    ./#%%%#*    ,(((##%%%%(*   ,/%%%%/.      .(%%%#/   .*#%%%#(*/(#%%%#/,
     ,#%%%#(.        ,#%%%(*   ,/%%%%/.      .(%%%#/  ,/%%%#/.    .*#%%%(,
      *#%%%%(*.      ,#%%%(*   .*#%%%#*     ./#%%%#,  ,(%%%#*      .(%%%#*
       ,(#%%%%%##(((##%%%%(*    .*#%%%%#(((##%%%%(,   .*#%%%##(///(#%%%#/.
         .*/###%%%%%%%###(/,      .,/##%%%%%##(/,.      .*(##%%%%%%##(*,
              .........                ......                .......
    2021/10/31 10:49:11 Serving at 0.0.0.0:3080



To use, follow examples in the `examples/` folder

    curl -v --location --request POST 'http://localhost:3080/api/v1/book' --header 'Content-Type: application/json' --data-raw '{"title": "Test title","image_url": "https://example.com","published_date": "2020-07-31T15:04:05.123499999Z","description": "test description"}' | jq

    curl --location --request GET 'http://localhost:3080/api/v1/book' | jq

To see all available routes, run

```shell
go run cmd/route/main.go
```

![go run cmd/routes/main.go](assets/routes.png)


# Table of Contents

- [Introduction](#introduction)
- [Motivation](#motivation)
- [Features](#features)
- [Quick Start](#quick-start)
- [Tooling](#tooling)
   * [Tools](#tools)
      + [Install](#install)
   * [Tasks](#tasks)
      + [List Routes](#list-routes)
      + [Format Code](#format-code)
      + [Sync Dependencies](#sync-dependencies)
      + [Compile Check](#compile-check)
      + [Unit tests](#unit-tests)
      + [golangci Linter](#golangci-linter)
      + [Security Checks](#security-checks)
      + [Check](#check)
      + [Hot reload](#hot-reload)
      + [Generate Swagger Documentation](#generate-swagger-documentation)
      + [Go generate](#go-generate)
      + [Test Coverage](#test-coverage)
      + [Build](#build)
      + [Clean](#clean)
- [Migration](#migration)
   * [Using Task](#using-task)
      + [Create Migration](#create-migration)
      + [Migrate up](#migrate-up)
      + [Rollback](#rollback)
   * [Without Task](#without-task)
      + [Create Migration](#create-migration-1)
      + [Migrate Up](#migrate-up)
      + [Rollback](#rollback-1)
- [Run](#run)
   * [Local](#local)
   * [Docker](#docker)
      + [docker-compose](#docker-compose)
- [Build](#build-1)
   * [With Task](#with-task)
   * [Without Task](#without-task-1)
- [Swagger docs](#swagger-docs)
- [Structure](#structure)
   * [Starting Point](#starting-point)
   * [Configurations](#configurations)
      - [.env files](#env-files)
   * [Database](#database)
   * [Router](#router)
   * [Domain](#domain)
      + [Repository](#repository)
      + [Use Case](#use-case)
      + [Handler](#handler)
      + [Initialize Domain](#initialize-domain)
   * [Middleware](#middleware)
   * [Dependency Injection](#dependency-injection)
   * [Libraries](#libraries)
- [Cache](#cache)
   * [LRU](#lru)
   * [Redis](#redis)
- [Utility](#utility)
- [Testing](#testing)
   * [Unit Testing](#unit-testing)
      - [Handler](#handler-1)
      - [Use Case](#use-case-1)
      - [Repository](#repository-1)
   * [End to End Test](#end-to-end-test)
- [TODO](#todo)
- [Acknowledgements](#acknowledgements)
- [Appendix](#appendix)
   * [Dev Environment Installation](#dev-environment-installation)


# Tooling

The above quick start is sufficient to start the API. However, we can take advantage of a tool to make task management easier. While you may run migration with `go run cmd/extmigrate/main.go`,  it is a lot easier to remember to type `task migrate` instead. Think of it as a simplified `Makefile`.

You may also choose to run sql scripts directly from `database/migrations` folder instead.

This project uses [Task](https://github.com/go-task/task) to handle various tasks such as migration, generation of swagger docs, build and run the app. It is essentially a [sh interpreter](https://github.com/mvdan/sh).

Install task runner binary bash script:

    sudo ./scripts/install-task.sh

This installs `task` to `/usr/local/bin/task` so `sudo` is needed.

`Task` tasks are defined inside `Taskfile.yml` file. A list of tasks available can be viewed with:

    task -l   # or
    task list

## Tools

Various tooling can be installed automatically by running which includes

 * [golang-ci](https://golangci-lint.run)
    * An opinionated code linter from https://golangci-lint.run/
 * [swag](https://github.com/swaggo/swag)
    * Generates swagger documentation 
 * [testify](https://github.com/stretchr/testify)
    * A testing framework
 * [gomock](https://github.com/golang/mock/mockgen)
    * Mock dependencies inside unit test
 * [golang-migrate](https://github.com/golang-migrate/migrate)
    * Migration tool
 * [ent](https://entgo.io/docs/getting-started)
    * Database ORM tool
 * [gosec](https://github.com/securego/gosec)
    * Security Checker
 * [air](https://github.com/cosmtrek/air)
    * Hot reload app 

### Install

Install the tools above with:

    task install:tools


## Tasks

Various tooling are included within the `Task` runner. Configurations are done inside `Taskfile.yml` file.

### List Routes

List all registered routes, typically done by `register.go` files by

    go run cmd/route/route.go

or

    task routes

### Format Code

    task fmt

Runs `go fmt ./...` to lint Go code

`go fmt` is part of official Go toolchain that formats your code into an opinionated format.

### Sync Dependencies

    task tidy

Runs `go mod tidy` to sync dependencies.


### Compile Check

    task vet

Quickly catches compile error.


### Unit tests

    task test

Runs unit tests.


### golangci Linter

    task golint

Runs [https://golangci-lint.run](https://golangci-lint.run/) linter.

### Security Checks

    task security

Runs opinionated security checks from [https://github.com/securego/gosec](https://github.com/securego/gosec).

### Check

    task check

Runs all the above tasks (Format Code until Security Checks)

### Hot reload

    task dev

Runs `air` which watches for file changes and rebuilds binary. Configure in `.air.toml` file.

### Generate Swagger Documentation
    
    task swagger

Reads annotations from controller and model file to create a swagger documentation file. Can be accessed from [http://localhost:3080/swagger/](http://localhost:3080/swagger/)


### Go generate

    task generate

Runs `go generate ./...`. It looks for `//go:generate` tags found in .go files. Useful for recreating mock file for unit tests.


### Test Coverage

    task coverage

Runs unit test coverage with `go test -cover ./...`

### Build

    task build

Create a statically linked executable for linux.

### Clean

    task clean

Clears all files inside `bin` directory.

# Migration

Migration is a good step towards having a versioned database and makes publishing to a production server a safe process.

All migration files are stored in `database/migrations` folder.

## Using Task

### Create Migration

Using `Task`, creating a migration file is done by the following command. Name the file after `NAME=`.

    task migrate:create NAME=create_a_tablename

Write your schema in pure sql in the 'up' version and any reversal in the 'down' version of the files.
 
### Migrate up

After you are satisfied with your `.sql` files, run the following command to migrate your database.

    task migrate

To migrate one step

    task migrate:step n=1
      
### Rollback
    
To roll back migration

    task migrate:rollback n=1

Further `golang-migrate` commands are available in its [documentation (postgres)](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)


## Without Task

### Create Migration

Once `golang-migrate` tool is [installed](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate), create a migration with

    migrate create -ext sql -dir database/migrations -format unix "{{.NAME}}"

### Migrate Up

You will need to create a data source name string beforehand. e.g.:

    postgres://postgres_user:$password@$localhost:5432/db?sslmode=false

Note: You can save the above string into an environment variable for reuse e.g.

    export DSN=postgres://postgres_user:$password@$localhost:5432/db?sslmode=false

Then migrate with the following command, specifying the path to migration files, data source name and action.

    migrate -path database/migrations -database $DSN up

To migrate 2 steps,

    migrate -path database/migrations -database $DSN up 2

### Rollback

Rollback migration by using `down` action and the number of steps

    migrate -path database/migrations -database $DSN down 1

# Run

## Local

Conventionally, all apps are placed inside the `cmd` folder.

If you have `Task` installed, the server can be run with:

    task run

or without `Task`, just like in quick start section:

    go run cmd/go8/main.go

## Docker

You can build a docker image with the app with its config files. Docker needs to be installed beforehand.

     task docker:build

This task also makes a copy of `.env`. Since Docker doesn't copy hidden file, we make a copy of it on our `src` stage before transferring it to our final `scratch` stage. It also inserts formats git tag and git hash as the API version which runs at compile time. [upx](https://upx.github.io/) is used to make the resulting binary smaller.

Note that this is a multistage Dockerfile. Since we statically compile this API, we can use `scratch` image (it is empty! - no file/folder exists).

Run the following command to build a container from this image. `--net=host` tells the container to use local's network so that it can access host database.

    docker-compose up -d postgres # If you haven't run this from quick start 
    task docker:run

### docker-compose

If you prefer to use docker-compose instead, both server and the database can be run with:

    task docker-compose:start

# Build

## With Task

If you have task installed, simply run

    task build

It does task check prior to build and puts both the binary and `.env` files into `./bin` folder

## Without Task

    go mod download
    CGO_ENABLED=0 GOOS=linux
    go build -ldflags="-X main.Version=$(git describe --abbrev=0 --tags)-$(git rev-list -1 HEAD) -w -s" -o ./server ./cmd/go8/main.go;


# Swagger docs

Swagger UI allows you to play with the API from a browser

![swagger UI](assets/swagger.png)
     
Edit `cmd/go8/go8.go` `main()` function host and BasePath  

    // @host localhost:3080
    // @BasePath /api/v1

   
Generate with

    task swagger # runs: swag init 
    
Access at

    http://localhost:3080

The command `swag init` scans the whole directory and looks for [swagger's declarative comments](https://github.com/swaggo/swag#declarative-comments-format) format.

Custom theme is obtained from [https://github.com/ostranme/swagger-ui-themes](https://github.com/ostranme/swagger-ui-themes)

# Structure

This project follows a layered architecture mainly consists of three layers:

 1. Handler
 2. Business Logic
 3. Repository

The handler is responsible to receiving requests, validating them hand over to business logic, then format the response to client.

Business logic is the meat of operations, and it calls a repository if necessary.

Database calls lives in this repository layer where data is retrieved from a store.

All of these layers are encapsulated in a domain.

## Starting Point

Starting point of project is at `cmd/go8/main.go`

![main](assets/main.png)


The `Server` struct in `internal/server/server.go` is where all important dependencies are 
registered and to give a quick glance on what your server needs.

![server](assets/server.png)

`s.Init()` in `internal/server/server.go` simply initializes server configuration, database, input validator, router, global middleware, domains, and swagger. Any new dependency added to the `Server` struct can be initialized here too.

![init](assets/init.png)


## Configurations
![configs](assets/configs.png)

All environment variables are read into specific `Configs` struct initialized in `configs/configs.go`. Each of the embedded struct are defined in its own file of the same package where its fields are read from either environment variable or `.env` file.

This approach allows code completion when accessing your configurations.

![config code completion](assets/config-code-completion.png)


#### .env files

The `.env` file defines settings for various parts of the API including the database credentials. If you choose to export the variables into environment variables for example:

    export DB_DRIVER=postgres
    export DB_HOST=localhost
    export DB_PORT=5432
    etc


To add a new type of configuration, for example for Elasticsearch
 
1. Create a new go file in `./configs`

```shell
touch configs/elasticsearch.go
```
    
2. Create a new struct for your type

```go
type Elasticsearch struct {
  Address  string
  User     string
  Password string
}
```
    
3. Add a constructor for it

```go
func ElasticSearch() Elasticsearch {
   var elasticsearch Elasticsearch
   envconfig.MustProcess("ELASTICSEARCH", &elasticsearch)

   return elasticsearch
}
``` 

A namespace is defined 

4. Add to `.env` of the new environment variables

```shell
ELASTICSEARCH_ADDRESS=http://localhost:9200
ELASTICSEARCH_USER=user
ELASTICSEARCH_PASS=password
```

Limiting the number of connection pool avoids ['time-slicing' of the CPU](https://github.com/brettwooldridge/HikariCP/wiki/About-Pool-Sizing). Use the following formula to determine a suitable number
 
    number of connections = ((core_count * 2) + effective_spindle_count)    

## Database

Migrations files are stored in `database/migrations` folder. [golang-migrate](https://github.com/golang-migrate/migrate) library is used to perform migration using `task` commands.

## Router

Router multiplexer or mux is created for use by `Domain`. While [chi](https://github.com/go-chi/chi) library is being used here, you can swap out the router tto an alternative one when assigning `s.router` field. However, you will need to adjust how you register your handlers in each domain.

## Domain

Let us look at how this project attempts at layered architecture. A domain consists of: 

  1. Handler (Controllers)
  2. Use case (Use Cases)
  3. Repository (Entities)

Let us start by looking at how `repository` is implemented.

### Repository

Starting wit, `Entities`. This is where all database operations are handled. Inside the `internal/domain/health` folder:

![book-domain](assets/domain-health.png)

Interfaces for both use case and repository are on its own file under the `health` package while its implementation are in `usecase` and `repository` package respectively.

The `health` repository has a single method

`internal/domain/health/repository.go`

```go
 type Repository interface {
     Readiness() error
 }
````    

And it is implemented in a package called `postgres` in `internal/domain/health/repository/postgres/postgres.go`

```go
func (r *repository) Readiness() error {
  return r.db.Ping()
}
```

### Use Case

This is where all business logic lives. By having repository layer underneath in a separate layer, those functions are reusable in other use case layers.

### Handler

This layer is responsible in handling request from outside world and into the `use case` layer. It does the following:

 1. Parse request into private 'request' struct
 2. Sanitize and validates said struct
 3. Pass into `use case` layer
 4. Process results from coming from `use case` layer and decide how the payload is going to be formatted to the outside world.
  
Route API are defined in `RegisterHTTPEndPoints` in their respective `register.go` file. 


### Initialize Domain

Finally, a domain is initialized by wiring up all dependencies in server/initDomains.go. Here, any dependencies can be injected such as a custom logger.

```go
func (s *Server) initBook() {
   newBookRepo := bookRepo.New(s.GetDB())
   newBookUseCase := bookUseCase.New(newBookRepo)
   bookHandler.RegisterHTTPEndPoints(s.router, newBookUseCase)
}
```

## Middleware

A middleware is just a handler that returns a handler as can bee seen in the `internal/middleware/cors.go`

```go
func Cors(next http.Handler) http.Handler {
   return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
           
           // do something before going into Handler
   
           next.ServerHTTP(w, r)
           
           // do something after handler has been served
       }
   }
```

Then you may choose to have this middleware to affect all routes by registering it in`initGlobalMiddleware()` or only a specific domain at `RegisterHTTPEndPoints()` function in its `register.go` file. 


### Middleware External Dependency

Sometimes you need to add an external dependency to the middleware which is often the case for 
authorization be that a config or a database. That middleware can be wrapped around by that 
dependency by first aliasing `http.Handler` with:

```go
type Adapter func(http.Handler) http.Handler
```
Then:

```go
func Auth(cfg configs.Configs) Adapter {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, err := getClaims(r, cfg.Jwt.SecretKey)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
```

## Dependency Injection

How does dependency injection happens? It starts with `InitDomains()` method. 

```go
healthHandler.RegisterHTTPEndPoints(s.router, usecase.NewHealthUseCase(postgres.NewHealthRepository(s.db)))
```

The repository gets access to a pointer to `sql.DB` to perform database operations. This layer also knows nothing of layers above it. `NewBookUseCase` depends on that repository and finally the handler depends on the use case.

## Libraries

Initialization of external libraries are located in `third_party/`

Since `sqlx` is a third party library, it is initialized in `/third_party/database/sqlx.go`

# Cache

The three most significant bottlenecks are 

  1. Input output (I/O) like disk access including database.
  2. Network calls - like calling another API.
  3. Serialization - like serializing or deserializing JSON

We demonstrate how caching results can speed up API response: 

## LRU

To make this work, we introduce another layer that sits between use case and database layer.

`internal/author/repository/cache/lru.go` shows an example of using an LRU cache to tackle the biggest bottleneck. Once we get a result for the first time, we store it by using the requesting URL as its key. Subsequent requests of the same URL will return the result from the cache instead of from the database. 

To make this work, we store the requesting URL in the handler layer.
```go
ctx := context.WithValue(r.Context(), author.CacheURL, r.URL.String())
```

Then in the cache layer, we retrieve it
```go
url := ctx.Value(author.CacheURL).(string)
```

We try and retrieve the key,
```go
val, ok := c.lru.Get(url)
```

If it doesn't exist, we can simply add it to our cache. 
```go
c.lru.Add(url, res)
```

Avoiding I/O bottleneck results in an amazing speed, **11x** more requests/second (328 bytes response size) compared to an already blazing fast endpoint as shown by `wrk` benchmark:

CPU: AMD 3600 3.6Ghz
Storage: SSD

```shell
wrk -t2 -d60 -c200  'http://localhost:3080/api/v1/author?page=1&size=3'
Running 1m test @ http://localhost:3080/api/v1/author?page=1&size=3
  2 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.23ms    5.07ms  71.75ms   83.36%
    Req/Sec    40.64k     3.55k   52.91k    68.45%
  4847965 requests in 1.00m, 1.48GB read
Requests/sec:  80775.66
Transfer/sec:     25.27MB
```

Compared to calling database layer:
```shell
wrk -t2 -d60 -c200  'http://localhost:3080/api/v1/author?page=1&size=3'
Running 1m test @ http://localhost:3080/api/v1/author?page=1&size=3
  2 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    70.66ms  116.57ms   1.24s    88.09%
    Req/Sec     3.66k   276.15     4.53k    70.50%
  437285 requests in 1.00m, 136.79MB read
Requests/sec:   7280.82
Transfer/sec:      2.28MB
```

Since a cache stays in the store if it is frequently accessed, invalidating the cache must be done if there are any changes to the stored value in the event of update and deletion. Thus, we need to delete the cache that starts with the base URL of this domain endpoint. 

For example:
```go
func (c *AuthorLRU) Update(ctx context.Context, toAuthor *models.Author) (*models.Author, error) {
	c.invalidate(ctx)

	return c.service.Update(ctx, toAuthor)
}

func (c *AuthorLRU) invalidate(ctx context.Context) {
	url := ctx.Value(author.CacheURL)
	split := strings.Split(url.(string), "/")
	baseURL := strings.Join(split[:4], "/")

	keys := c.lru.Keys()
	for _, key := range keys {
		if strings.HasPrefix(key.(string), baseURL) {
			c.lru.Remove(key)
		}
	}
}
```
## Redis

By using Redis as a cache, you can potentially take advantage of a cluster architecture for more RAM instead of relying on the RAM on current server your API is hosted. Also, the cache won't be cleared like in-memory `LRU` when a new API is deployed.

Similar to LRU implementation above, this Redis layer sits in between use case and database layer.

This Redis library requires payload in a binary format. You may choose the builtin `encoding/json` package or `msgpack` for smaller payload and **7x** higher speed than without a cache. Using `msgpack` over `json` tackles serialization bottleneck.

```go
// marshal 
cacheEntry, err := msgpack.Marshal(res)
// unmarshal
err = msgpack.Unmarshal([]byte(val), &res)
```

```shell
wrk -t2 -d60 -c200  'http://localhost:3080/api/v1/author?page=1&size=3'
Running 1m test @ http://localhost:3080/api/v1/author?page=1&size=3
  2 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.05ms    2.56ms  37.48ms   73.63%
    Req/Sec    25.48k     1.45k   30.73k    71.29%
  3039522 requests in 1.00m, 0.93GB read
Requests/sec:  50638.73
Transfer/sec:     15.84MB
````

# Utility

Common tasks like retrieving query parameters or `filters` are done inside `utility` folder. It serves as one place abstract functionalities used across packages.

# Testing

## Unit Testing

Unit testing can be run with

    task test
    
Which runs `go test -v ./...`

In Go, unit test file is handled by appending _test to a file's name. For example, to test `/internal/domain.book/handler/http/handler.go`, we add unit test file by creating `/internal/domain.book/handler/http/handler_test.go`


To perform a unit test we take advantage of go's interface. Our interfaces are defined in:

      internal/domain/book/handler.go
      internal/domain/book/usecase.go
      internal/domain/book/repository.go

The implementation if these interfaces are in separate files. For example our concrete 
implementation for use case of `Create` is in `internal/book/usecase/http/usecase.go`.

### Repository

In this database unit test, we only concern with the behaviour of our database operations, not the actual interaction with a real database. Mocking a database call allows us to simulate that interaction. Here, we use a library from Data Dog called [go-sqlmock](github.com/DATA-DOG/go-sqlmock). It, can be installed with:


      task install:sqlmock


Then in our postgres_test.go file, we create a mock with

```go
db, mock, err := sqlmock.New()
```

And then we get a `DB` struct simply by

```go
sqlxDB := sqlx.NewDb(db, "sqlmock")
```

In each of our unit test, we get a mock repository with

```go
db, mock := NewMock()
repo := New(db)
```

The basic idea is the same as use case unit tests. We
1. create what we expect - in this case the SQL query
2. create what the response should be

In  `Create()` function in `postgres.go` file, we perform a few things
1. Prepares an sql statement
2. Perform sql insert
3. Pass in sql arguments
4. Returns a book id

To mock these,
1. Call `ExpectPrepare()` with the expected SQL query. `sqlmock` allows usage of regex

    ExpectPrepare("^INSERT INTO books")

2. Call `ExpectQuery()` because we are doing an insertion

3. Tell the library what values are inserted

    WithArgs(bookTest.Title, bookTest.PublishedDate, bookTest.ImageURL, bookTest.Description)

4. Returns a book ID

    WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

Next, call the `repo.Create()` function and perform assertions.

### Use Case

Our use case unit tests may need to retrieve or store data from and into a database with the help of repository (or repositories). Thus, there isn't any direct usage of database object. So instead, we need to mock our repository with the help of [gomock](https://github.com/golang/mock/). It can be installed with

      task install:gomock
or

    go install github.com/golang/mock/mockgen@v1.6.0            # for go >= v1.6
    GO111MODULE=on go get github.com/golang/mock/mockgen@v1.6.0 # for go < v1.6


We tell what response to be expected (using `EXPECT()` function) and then call our use case function. If the use case function returns a response of what we have expected as defined in `EXPECT()` function, then our unit test passes.

Both use case and repository are created with

```go
uc, repo := newUseCase(t)
```

The `newUseCase()` function returns mock controller and repository generated by `sqlmock` and `gomock` library respectively. Generate these mocks by calling either `go generate ./...` or `task generate` (if you use `Task`). `go generate ./...` command looks at `//go:generate` tags in all `.go` files - notice the lack of a space between `//` and `go:generate` which is required for the tag to work.

```go
//go:generate mockgen -package mock -source ../usecase.go -destination=../mock/mock_usecase.go
```

When we look at `usecase.go` file, two repository calls were made, `Create()` and `Read()`.

We expect `Create()` to return an ID and an error tuple while `Read()` function return a book model and error tuple.

Knowing this, we tell the library what to expect in each of these calls.

1. In `Create()`:
* Receives a book model (without ID)
* Returns one book ID

Thus,

```go
repo.EXPECT().Create(ctx, gomock.Eq(expected)).Return(bookID, err).Times(1)
```

2. In `Read()`:
* Receives a book ID
* Return a book model with id == 1

Thus,

```go
repo.EXPECT().Read(ctx, gomock.Eq(bookID)).Return(expectedCreated, err).Times(1)
```

Then we initiate the use case function with `bookGot, err := uc.Create(ctx, request)`
and we can to several assertions to confirm what we've got is the one we expected.

```go
assert.NotEqual(t, bookGot.BookID, 0)
assert.Equal(t, bookGot.Title, request.Title)
assert.Equal(t, bookGot.PublishedDate.String(), request.PublishedDate.String())
assert.Equal(t, bookGot.Description, request.Description)
assert.Equal(t, bookGot.ImageURL.String, request.ImageURL.String)
```


### Handler

Handler unit testing is done by mocking any use cases done in it. The interface we are testing against is defined in `internal/domain/book.handler.go` file.

```go
type HTTP interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
```

In general, these are the steps to make a handler unit test:

  1. Generate stubs using `mockgen`.
  2. Create a `gomock` controller.
  3. Create a mock use case from the generated stub.
  4. Create test input and expected output. Then, from these input and output, tell `gomock` what to take, and what to expect.
  5. Create our test handler.
  6. Call our method to be tested.
  7. Assert results.

Whew, this is a lot of steps! Let's dive in. 

#### Generate stubs using `mockgen`

Mocks are generated using `mockgen` and is generated with either running `go generate ./...` or `task generate`.
The `//go:generate mockgen <options>...` tag tells `go generate` command to run this specific command in all `.go` files.

The `mockgen` can take relative path like so:

```go
//go:generate mockgen -package mock -source ../../handler.go -destination=../../mock/mock_handler.go
```

#### `gomock` controller

For our usecase, we need to create a mock controller from the library. `defer ctrl.Finish()` can be omitted if you are using Go >= 1.14 and `mockgen` >= 1.5.0.

```go
ctrl := gomock.NewController(t)
defer ctrl.Finish()
```

#### Mock use case

Since our handler calls one use case, specifically `h.useCase.Create()`, we need to be able to mock the input and output. We don't want to call the actual implementation because this unit test should only concern with the handler. We use usecase stubs generated in the [Usecase](Use Case) section.  

```go
uc := mock.NewMockUseCase(ctrl)
```

#### EXPECT()

We expect the use case to accept a context and a `model.Book` and also return `model.Book`, error tuple once. Thus,

```go
uc.EXPECT().Create(ctx, testBookRequest).Return(ucResp, e).Times(1)
```

Create these input and output

```go
// input
ctx := context.Background()
testBookRequest := &models.Book{
    Title:         "test01",
    PublishedDate: now.MustParse("2020-02-02"),
    ImageURL: null.String{
        String: "https://example.com/image.png",
        Valid:  true,
    },
    Description: "test01",
}

// output
ucResp := &models.Book{
    BookID:        1,
    Title:         testBookRequest.Title,
    PublishedDate: testBookRequest.PublishedDate,
    ImageURL:      testBookRequest.ImageURL,
    Description:   testBookRequest.Description,
}
var e error
```


#### Test Handler

One more thing left before we call our API is to create a handler which is going to be used to call our API endpoint. API registration is done from `RegisterHTTPEndPoints()`. 

```go
h := RegisterHTTPEndPoints(router, val, uc)
```

For the other dependencies, they can be simply new instances of router and validation library.

#### Test Request

All preparation is in place. What's left is to make a request to our endpoint `/api/v1/books`.  Our `Create()` method has the standard http handler which takes a `http.ResponseWriter` and `*http.Request`. We make use of `httptest` library. Form the request with the http method, path and any payload. Result will be written into `ww`.

```go
ww := httptest.NewRecorder()
rr := httptest.NewRequest(http.MethodPost, "/api/v1/books", bytes.NewBuffer(body))
```

Finally, we can call our `Create()` function.
```go
h.Create(ww, rr)
```

#### Assert Results

Decode the json response into a `book.Res` struct - because that is what the handler ultimate turn our response from use case thanks to `book.Resource()` method.

```go
var gotBook book.Res
err = json.NewDecoder(ww.Body).Decode(&gotBook)
```

Now we can assert http status and all values

```go
assert.Equal(t, http.StatusCreated, ww.Code)
assert.Equal(t, gotBook.Title, ucResp.Title)
assert.Equal(t, gotBook.Description.String, ucResp.Description)
assert.Equal(t, gotBook.PublishedDate.String(), ucResp.PublishedDate.String())
assert.Equal(t, gotBook.ImageURL.String, ucResp.ImageURL.String)
```

## End-to-End Test

Technically End-to-End test (e2e test) can be done separately in another program and language. Having e2e binary integrated in the project has the advantage of reusing structs and migration which will be explained down below. 

The idea here is to run our application isolated in a container (along with database) and the e2e program calls known API of this program and checks if the output is what is expected.  

It creates two docker containers, one for the API, and second for postgres. Once the containers have started. It runs the app, and then the e2e binary.

Remember the `Server` struct in `internal/server/server.go` file? The `New()` function is called both by our API and e2e binary. We can also call `Migrate()` function because the e2e test uses the same `Server` struct as our API.

In our actual e2e implementation use cases, we can perform various CRUD operations.

For example, in an empty database, we expect no books should be returned.

```go
func testEmptyBook(t *E2eTest) {
	// call our API endpoint
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s/api/v1/books", t.server.Config().Api.Port))
	
	// The return should be an empty array
    if !bytes.Equal(expected, got) {
        log.Printf("handler returned unexpected body: got %v want %v", string(got), expected)
    }
}
```

### Run e2e test

Start

    task dockertest

or

```shell
 cd docker-test && docker-compose down -v --build && docker-compose up -d
 docker exec -t go8_container_test "/home/appuser/app/e2e"
```

# TODO

 - [ ] Fix end to end test
 - [ ] Complete HTTP integration test
 - [x] Better return response
 - [x] LRU cache
 - [X] Redis Cache
 - [ ] Tracing
 - [ ] Metric

# Acknowledgements

 * https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
 * https://github.com/moemoe89/integration-test-golang
 * https://github.com/george-e-shaw-iv/integration-tests-example
 * https://gist.github.com/Oxyrus/b63f51929d687c1e20cda3447f834147
 * https://github.com/sno6/gosane
 * https://github.com/AleksK1NG/Go-Clean-Architecture-REST-API
 * https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
 * https://github.com/arielizuardi/golang-backend-blog
 
# Appendix

## Dev Environment Installation

For Ubuntu:

```shell
sudo apt update && sudo apt install git curl build-essential jq
wget https://golang.org/dl/go1.17.6.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
echo 'PATH=$PATH:/usr/local/go/bin' >> ~/.bash_aliases
source ~/.bashrc
go get -u golang.org/x/tools/...

curl -s https://get.docker.com | sudo bash
sudo usermod -aG docker ${USER}
newgrp docker
su - ${USER} # or logout and login

sudo curl -L "https://github.com/docker/compose/releases/download/v2.2.3/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
```