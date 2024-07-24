#!/bin/bash
# author: df
# date: 2024-07-24
# desc: entrypoint.sh to run multiple process in one container
cd ./cdn
npm run start && \
java -Dspring.profiles.active=application -jar ../app.jar && \
fg
