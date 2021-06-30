# ASAPP Chat Backend Challenge v1
### Overview
This is a Go based boilerplate which runs an HTTP Server. It allow you to create an user, login using JWT.
It also allow the user to create a new message to another user, and also get all messages.

This project is made with:
  - Hexagonal Arhitecture (Infrastructure, appilacion, domain)
  - DDD
  - CQS Pattern
  - Repository pattern to handle database connections
  - SQLite
  - Chi Router framework
  - JWT for auth and authz

### Things to do
There are things that remains to be done:
  - Tests
  - Docker
### Instructions

They are located in the *docs/index.html* file

### Prerequisites

Installed Go version >= 1.12 since it uses Go Modules.

### How to run it
`
go run cmd/server.go
`
