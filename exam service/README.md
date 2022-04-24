

<div align="center">
  <br>
  <h1>Exam Service</h1>
</div>

##  Table of contents
- [**Summary**](#summary)
- [**Getting started**](#getting-started)
- [**Prerequisites**](#prerequisites)
- [**Packages**](#packages)
- [**Environment variables file**](environment-variables-file)
- [**Running**](#running)
- [**Postman collection**](#postman-collection)
---
## Summary

This service allows teachers to create, get, update and delete exams and questions.

## Getting Started

Clone the repository.<br />
Follow the instructions to complete the installation.

## Prerequisites

- [Golang](https://golang.org/dl/)
- [Redis](https://redis.io/download/)

## Packages

- [Gin](https://github.com/gin-gonic/gin)
- [Validator](https://github.com/go-playground/validator)
- [Go-Redis](https://github.com/go-redis/redis)
- [Go-Rejson](https://github.com/nitishm/go-rejson)
- [Viper](https://github.com/spf13/viper)

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

- Local: http://localhost:[SERVER_PORT]

## Postman collection

you will find the postman collection [here](postman%20collection/exam%20service.postman_collection.json).

You can get what is the right structure of JSON file to send requests and recieving responses from the postman collection after importing it in the [Postman](https://www.postman.com/).