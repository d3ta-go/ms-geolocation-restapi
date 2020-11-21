# ms-geolocation-restapi ( [![Go Report Card](https://goreportcard.com/badge/github.com/d3ta-go/ms-geolocation-restapi)](https://goreportcard.com/report/github.com/d3ta-go/ms-geolocation-restapi) )

MicroService Interface/Presentation App: GeoLocation RestAPI

As a part of `Simple Implementation of Modular DDD Technical Architecture Patterns in Go`.

## Diagram v 0.2.2-Modular

![DDD-Technical-Architecture-Patterns-Golang-0.2.2-MS GeoLocation RESTAPI](docs/img/DDD-Technical-Architecture-Patterns-Golang-0.2.2-MS_GeoLocation_RestAPI.png)

## Components

A. Interface Layer (MicroService)

1. Microservice: GeoLocation REST API - using Echo Framework [ [d3ta-go/ms-geolocation-restapi](https://github.com/d3ta-go/ms-geolocation-restapi) ]

B. DDD Modules:

1. GeoLocation - using DDD Layered Architecture (CURD, GORM) [ [d3ta-go/ddd-mod-geolocation](https://github.com/d3ta-go/ddd-mod-geolocation) ]

C. Common System Libraries [ [d3ta-go/system](https://github.com/d3ta-go/system) ]:

1. Configuration - using yaml
2. Identity & Securities - using JWT, Casbin (RBAC)
3. Initializer
4. Email Sender - using SMTP
5. Handler
6. Migrations
7. Utils

D. Databases

1. MySQL (tested)
2. PostgreSQL (untested)
3. SQLServer (untested)
4. SQLite3 (untested)

E. Persistent Caches

1. Session/Token/JWT Cache (Redis, File, DB, etc) [tested: Redis]
2. Indexer/Search Cache (ElasticSearch)

F. Messaging [to-do]

G. Logs [to-do]

### Development

1. Clone

```shell
$ git clone https://github.com/d3ta-go/ms-geolocation-restapi.git
```

2. Setup

```
a. copy `conf/config-sample.yaml` to `conf/config.yaml`
b. copy `conf/data/test-data-sample.yaml` to `conf/data/test-data.yaml`
c. setup your dependencies/requirements (e.g: database, redis, smtp, etc.)
```

3. Runing on Development Stage

```shell
$ cd ms-geolocation-restapi
$ go run main.go db migrate
$ go run main.go server restapi
```

4. Build

```shell
$ cd ms-geolocation-restapi
$ go build
$ ./ms-geolocation-restapi db migrate
$ ./ms-geolocation-restapi server restapi
```

5. Distribution (binary)

Binary distribution (OS-arch):

- darwin-amd64
- linux-amd64
- linux-386
- windows-amd64
- windows-386

```shell
$ cd ms-geolocation-restapi
$ sh build.dist.sh
$ cd dist/[OS-arch]/
$ ./ms-geolocation-restapi db migrate
$ ./ms-geolocation-restapi server restapi
```

**RESTAPI (console):**

![Microservice: GeoLocation REST API](docs/img/geolocation-sample-ms-rest-api.png)

**Swagger UI (openapis docs):**

URL: http://localhost:20203/openapis/docs/index.html

![Openapis: Email REST AIP](docs/img/geolocation-sample-openapis-docs.png)

**Related Domain/Repositories:**

1. DDD Module: GeoLocation (Supporting Subdomain) - [d3ta-go/ddd-mod-geolocation](https://github.com/d3ta-go/ddd-mod-geolocation)
2. Common System Libraries - [d3ta-go/system](https://github.com/d3ta-go/system)
3. Need JWT Authorization/Token from: Account (Generic Subdomain) Module/Account REST API - [d3ta-go/ms-account-restapi](https://github.com/d3ta-go/ms-account-restapi). `Please use shared redis server to store/retrieve JWT Token.`

**Online Demo:\***

> Due the limitation of our server spec (for ELK: Elasticsearch/Kibana). We cannot provide the online demo for this repo. Very sorry for the inconvenience.

Available Online Demo:

1. Account REST API Microservice - [d3ta-go/ms-account-restapi](https://github.com/d3ta-go/ms-account-restapi)
2. Email REST API Microservice - [d3ta-go/ms-email-restapi](https://github.com/d3ta-go/ms-email-restapi)
3. Covid19 REST API Microservice - [d3ta-go/ms-covid19-restapi](https://github.com/d3ta-go/ms-covid19-restapi)

**References:**

1. [Book] Domain-Driven Design: Tackling Complexity in the Heart of Software 1st Edition (Eric Evans, 2004)

2. [Book] Patterns, Principles, and Practices of Domain-Driven Design (Scott Millett & Nick Tune, 2015)

**Team & Maintainer:**

1. Muhammad Hari (https://www.linkedin.com/in/muharihar/)
