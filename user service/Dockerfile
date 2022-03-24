### BUILD a maven builder. This will contain all mvn dependencies and act as an abstraction for all mvn goals
FROM maven:3.5.4-jdk-8-alpine as builder

# create app folder for sources
RUN mkdir -p /build
RUN mkdir -p /build/logs

# The WORKDIR instruction sets the working directory for any RUN, CMD, ENTRYPOINT, COPY and ADD instructions that follow it in the Dockerfile.
WORKDIR /build
COPY pom.xml /build
#Download all required dependencies into one layer
RUN mvn dependency:resolve && mvn compile

ARG DB_HOSTNAME=postgresql
ARG POSTGRES_PORT=5432
ARG POSTGRES_DBNAME=docker
ARG POSTGRES_USERNAME=docker
ARG POSTGRES_PASS=docker
ARG VERIFICATION_MAIL_USERNAME=eyadyoussef69@gmail.com
ARG VERIFICATION_MAIL_PASSWORD=ifobbplpqanlsqjx


