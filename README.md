# Clean Architecture Golang

## Basic Operations Supported
 - Execute `go run cmd/main.go`
 - Create User 
    - `curl -d '{"name":"kelvin", "email":"kel@gmail.com"}' -H "Content-Type: application/json" -X POST http://localhost:8080/users`
 - Get All Users 
    - `curl -H "Content-Type: application/json" -X GET http://localhost:8080/users`
 - Get One User via ID or Email
    - `curl -H "Content-Type: application/json" -X GET http://localhost:8080/users/1`
    - `curl -H "Content-Type: application/json" -X GET http://localhost:8080/users/kel@gmail.com`
