# Go Learning Project

This project serves as a learning journey through the Go programming language. Here's what it includes:

## Directories

### `/tour`

Contains lessons from [A Tour of Go](https://go.dev/tour). This is an interactive introduction to the Go language.

### `/auth`

This directory focuses on the initial steps in implementing authentication in Go. Key features include:

- **Redis Database**: Utilizes Redis as the datastore for login credentials.
- **Messaging with NATS**: Upon successful login, a message is dispatched using NATS.
- **JWT Authentication**: Successful logins respond with a cookie containing a JWT (JSON Web Token).
- **Architecture**: The design pattern followed in this section is the Hexagonal Architecture along with Domain-Driven Design (DDD).

## Getting Started

To get the project up and running, you need to have Docker installed. Once that's set up, navigate to the project's root directory and run:

```
docker-compose up
```

This will start all the necessary services and set up the environment for you.
