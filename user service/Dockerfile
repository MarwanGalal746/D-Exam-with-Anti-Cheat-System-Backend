FROM openjdk:17
ADD target/user-service.jar user-service
EXPOSE 8080
ENTRYPOINT ["java","-jar","/user-service"]