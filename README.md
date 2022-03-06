# microservice-template-go2

go project template for microservice architecture structure.

# Description
This project is a set of ideas for implementing a microservice template with Go, I'll to ilustrate some Microservice patterns like:
- Messaging style patterns (REST)
- Reliable comunication patterns (circuit breaker)
- Security patterns 
- Cross-cutting concerns patterns

# Dependencies:
* [**Gorilla mux**](https://github.com/gorilla/mux)

# Go Directories   

<pre>
microservice-name
├── cmd
|   ├── service
|   |    └── main.go
├── internal/
|   ├── application
|   |    └── app.go
|   ├── configurations
|   |    └── setup.go
|   ├── controller
|   |    └── controller.go
|   ├── entity
|   |    └── entity.go
|   ├── routes
|   |    └── routes.go
|   ├── service
|   |    └── service.go
├── pkg/
|   ├── library01
|   |    ├── file01.go
|   |    └── file02.go
|   ├── library02
|   |    ├── file01.go
|   |    └── file02.go
├── go.mod
├── go.sum
└── README.md
</pre>

Bibliography
- https://microservices.io/patterns/
- https://semver.org/
