# ms-geolocation-restapi

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
