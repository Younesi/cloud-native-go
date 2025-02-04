# Learning Cloud Native Go
> 🌱 Cloud Native Application Development is one way of speeding up the building of web applications using microservices, containers, and orchestration tools.

This repository is for practice/learning purposes.

## 🚀 Features included

- An idiomatic structure based on resource-oriented design.
- Dockerized development with Docker, Docker Compose, Alpine images, and linters.
- Healthcheck and CRUD API implementations following OpenAPI specifications.
- Database migrations using [Goose](https://github.com/pressly/goose) and ORM management with [GORM](https://gorm.io/).
- Structured logging with [`Slog`](https://pkg.go.dev/log/slog).
- Data validation using [Validator.v10](https://github.com/go-playground/validator).
- JWT-based authentication with role-based access control.
- Kubernetes manifests and Helm charts for deployment.
- A CI/CD pipeline implemented with GitHub Actions.

## ⌛ Upcoming Features

- The usage of GitHub actions to run tests and linters, generate OpenAPI specifications, and build and push production images to the Docker registry.
- Add gRPC endpoints for high-performance communication
- Set up Prometheus and Grafana for monitoring and observability
- Add distributed tracing with Jaeger
- Add API rate limiting and throttling mechanisms
- Implement centralized logging with ELK stack
- Integrate a service mesh using Istio
- Set up an API gateway with Kong
- Add Redis for caching
- Use Terraform for infrastructure provisioning
- Set up database read replicas for scaling 
- Create a serverless function using AWS Lambda
- Implement a simple event sourcing system for a feature
- Set up a webhook system for notifications
- Implement basic chaos testing with Chaos Mesh
- Implement asynchronous processing using RabbitMQ
- Create a basic GraphQL endpoint
- Integrate OpenTelemetry for distributed tracing
- Implement WebSocket support for real-time updates
- Set up a message queue system with RabbitMQ
- Implement a search engine with Elasticsearch
- Set up database connection pooling for scalability
- Implement a backup and restore system for the database
- Set up a cron job system for periodic tasks
- Add OAuth2 Support for Third-Party Authentication