

<div align="center">
  <br>
  <h1>Exam Service</h1>
</div>

##  Table of contents
- [**Summary**](#summary)
- [**Getting started**](#getting-started)
- [**Prerequisites**](#prerequisites)
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

## Environment variables file

- Rename app.env.example file to .env
- Modify app.env file according to your needs.

## Running

- In your cloned directory.
- open your terminal and run:

```
go run main.go
```

The server will start at:

- Local: http://localhost:[SERVER_PORT]

## Postman collection

you will find it [here](postman%20collection/exam%20service.postman_collection.json).
