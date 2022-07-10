<div align="center">
  <h1>Exam Service</h1>
</div>

##  Table of contents
- [**Summary**](#summary)
- [**Getting Started**](#getting-started)
- [**Prerequisites**](#prerequisites)
- [**Packages**](#packages)
- [**Environment variables file**](environment-variables-file)
- [**Running**](#running)
- [**Postman Collection**](#postman-collection)
---
## Summary

This service allows teachers to create, get, update and delete exams and questions.

## Getting Started

Clone the repository.<br />
Follow the instructions to complete the installation.

## Prerequisites

- [Golang](https://golang.org/dl/)
- [Redis](https://redis.io/download/)
- [Postgresql](https://www.postgresql.org/download/)

## Packages

- [Gin](https://github.com/gin-gonic/gin)
- [Validator](https://github.com/go-playground/validator)
- [Go-Redis](https://github.com/go-redis/redis)
- [Go-Rejson](https://github.com/nitishm/go-rejson)
- [Viper](https://github.com/spf13/viper)
- [Gorm](https://github.com/go-gorm/gorm)

## Environment variables file

- Rename app.env.example file to .env
- Modify app.env file according to your needs.

## Running

- In your cloned directory.
- open your terminal and run:

```bash
go run main.go
```

The server will start at:

- Local: http://localhost:[EXAM_SERVER_PORT]

## Postman Collection

you will find the postman collection [here](postman collection/exam service.postman_collection.json).

You can get what is the right structure of JSON file to send requests and recieving responses from the postman collection after importing it in the [Postman](https://www.postman.com/).