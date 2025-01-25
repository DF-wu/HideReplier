#FROM openjdk:17
FROM azul/zulu-openjdk-alpine:17 AS javaenv
LABEL author="df" email="df@dfder.tw"
WORKDIR /usr/src/app
# copy files from outside to inside
COPY . . 

# this options for fly.io very low spec runtime spec.
ENV JAVA_OPTS="\
    -XX:+UseContainerSupport \
    -XX:MaxRAMPercentage=75.0 \
    -XX:InitialRAMPercentage=50.0 \
    -Xss256k \
    -XX:+UseSerialGC \
    -XX:+TieredCompilation"

ENTRYPOINT ["java $JAVA_OPTS","-Dspring.profiles.active=application","-jar","app.jar"]
