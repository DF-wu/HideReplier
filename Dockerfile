#FROM openjdk:17
FROM azul/zulu-openjdk-alpine:17
MAINTAINER df, df@dfder.tw
WORKDIR /usr/src/app
# copy files from outside to inside
COPY . .
ENTRYPOINT ["java","-Dspring.profiles.active=application","-jar","app.jar"]
