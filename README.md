## Zero

Zero real estate application, single repo Microservice

## Table of Contents

- [Installation](#installation)

## Installation

1) Set up env variables: Make a const.go inside config folder, fill in the needed info. User example.const as template
2) You'll need a Postgres and mongodb account
3) Set up Makefile. Use example.Makefile as template.
4) Generate proto files using protoc, by running make commands:

``` sh
   $ make gen_listings 
   $ make gen_users 
   $ make gen_application 
   $ make gen_file_service 
   $ make gen_frontend
```

7) Download GRPC follow: Please see [GrpcGo](https://grpc.io/docs/languages/go/quickstart/).
8) Download Go dependencies with:

``` sh
   $ go mod download
   $ go mod tidy
   $ npm i grpc-web
```
