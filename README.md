# microservice-template-go2

go project template for microservice architecture structure ussing Gin-Gonic and Gorm (postgresql database).

# Description
This project is a set of ideas for implementing a microservice template with Go, I'll to ilustrate some Microservice patterns like:
- Messaging style patterns (REST)
- Reliable comunication patterns (circuit breaker)
- Security patterns 
- Cross-cutting concerns patterns

# Dependencies:
* [**Gin Gonic**](https://github.com/gin-gonic/gin)
* [**GORM**](https://gorm.io/)
* [**Postgresql**](https://hub.docker.com/_/postgres) 

# Go Directories   

<pre>
microservice-name
├── cmd
|   ├── service
|   |    └── main.go
├── config
|   └── config.yml
├── docs
|   └── people-project-insomnia.json
├── internal/
|   ├── adapters
|   |    └── db
|   |    |   └── db.go
|   |    └── web
|   |    |   └── routes.go
|   ├── application
|   |    └── app.go
|   ├── configurations
|   |    └── setup.go
|   ├── controller
|   |    └── controller.go
|   ├── entity
|   |    └── entity.go
|   ├── service
|   |    └── service.go
├── pkg/
|   ├── library01
|   |    ├── file01.go
|   |    └── file02.go
|   ├── library02
|   |    ├── file01.go
|   |    └── file02.go
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
</pre>

![Plantuml project](https://github.com/wilorios/microservice-template-go2/blob/main/docs/plantuml-go-project.png)

Bibliography
- https://microservices.io/patterns/
- https://semver.org/

Docker compose commands
- docker-compose build 
- docker-compose up
