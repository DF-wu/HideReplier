#!/bin/bash
# 不要在container裡面用maven build ， 不然會很慢
# 不要在container裡面用maven build ， 不然會很慢
# 不要在container裡面用maven build ， 不然會很慢
# by stanley2058@yahoo.com.tw
# 所以這是先在外面build ，再把build好的檔案 copy到working dir 再tag
printf "\n\n\n"
# echo "########## BOT_VERSION: $BOT_VERSION"
echo "########## BOT_VERSION: 1.0.2"
echo "########## TENOR_API_KEY: $TENOR_API_KEY"
printf "\n\n\n"
echo "build-deploy.sh : start git pull" | lolcat
git pull
echo "build-deploy.sh : git pull done." | lolcat


# WARN: please remember to set TENOR_API_KEY & BOT_VERSION environment variables
echo "build-deploy.sh : start npm build" | lolcat
pushd web || exit 1
./build-in-docker.sh
rm -rf ../src/main/resources/static/*
cp -r ./dist/* ../src/main/resources/static
popd || exit 1
echo "build-deploy.sh : done npm build" | lolcat

echo "build-deploy.sh : start maven build" | lolcat
mvn -T 1C clean install -Dmaven.test.skip=true
echo "build-deploy.sh : maven build successfully" | lolcat

echo "build-deploy.sh : start to copy built files" | lolcat
cp ./target/hidereplier-0.0.1-SNAPSHOT.jar app.jar
echo "build-deploy.sh : cp apps from target directory to root directory" | lolcat

echo "build-deploy.sh : building docker image" | lolcat 
docker build . -t dfder/hidereplier
echo "build-deploy.sh : docker build successfully" | lolcat

echo "build-deploy.sh : pushing docker image to dockerhub" | lolcat
docker push dfder/hidereplier
echo "build-deploy.sh : docker image uploaded." | lolcat

echo "build-deploy.sh : deploying to fly.io." | lolcat
fly deploy
echo "build-deploy.sh : deploying finished." | lolcat
