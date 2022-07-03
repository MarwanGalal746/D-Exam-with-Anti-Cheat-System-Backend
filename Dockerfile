FROM openjdk:17-alpine
ADD target/user-service.jar user-service
EXPOSE 8080
ENTRYPOINT ["java","-jar","/user-service"]