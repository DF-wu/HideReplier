#FROM openjdk:17
FROM azul/zulu-openjdk-alpine:17-latest AS javaenv
LABEL author="df" email="df@dfder.tw"
WORKDIR /usr/src/app
# copy files from outside to inside
COPY . . 
ENTRYPOINT ["sh","./run.sh"]
