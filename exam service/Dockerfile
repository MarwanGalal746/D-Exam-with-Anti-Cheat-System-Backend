# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
# Specify that we now need to execute any commands in this directory.
WORKDIR app
# Copy everything from this project into the filesystem of the container.
COPY . .

RUN go mod download github.com/sagikazarmark/crypt
RUN go mod download github.com/jackc/chunkreader
RUN go mod download github.com/jackc/pgproto3
RUN go mod download github.com/ugorji/go

# Compile the binary exe for our app.
RUN go build -o main main.go

EXPOSE 8888
# Start the application.
CMD [ "main" ]