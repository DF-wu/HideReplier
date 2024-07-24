#!/bin/bash
# author: df
# date: 2024-07-24
# desc: entrypoint.sh to run multiple process in one container

java -Dspring.profiles.active=application -jar ../app.jar 
