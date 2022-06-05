# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
# Specify that we now need to execute any commands in this directory.
WORKDIR /go/src/exam-service
# Copy everything from this project into the filesystem of the container.
COPY . /go/src/exam-service
# Compile the binary exe for our app.
RUN go build -o /go/src/exam-servicemain .
# Start the application.
ENTRYPOINT ["/go/src/exam-service/main"]
