# syntax=docker/dockerfile:1
FROM golang:1.18-alpine
# Specify that we now need to execute any commands in this directory.
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . ./
RUN go mod tidy
RUN go build -o /main
# Compile the binary exe for our app.
# Start the application.
CMD ["/main"]
