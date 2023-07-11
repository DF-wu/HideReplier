#!/bin/bash
# 不要在container裡面用maven build ， 不然會很慢
# 不要在container裡面用maven build ， 不然會很慢
# 不要在container裡面用maven build ， 不然會很慢
# by stanley2058@yahoo.com.tw
# 所以這是先在外面build ，再把build好的檔案 copy到working dir 再tag
echo "build-deploy.sh : start git pull"
git pull
echo "build-deploy.sh : git pull done."


echo "build-deploy.sh : start maven build"
mvn -T 1C clean install -Dmaven.test.skip=true
echo "build-deploy.sh : maven build successfully"

echo "build-deploy.sh : start to copy built files"
cp ./target/hidereplier-0.0.1-SNAPSHOT.jar app.jar
echo "build-deploy.sh : cp apps from target directory to root directory"


echo "build-deploy.sh : building docker image"
docker build . -t dfder/hidereplier
echo "build-deploy.sh : docker build successfully"

echo "build-deploy.sh : pushing docker image to dockerhub"
docker push dfder/hidereplier
echo "build-deploy.sh : docker image uploaded."

echo "build-deploy.sh : deploying to fly.io."
flyctl deploy
echo "build-deploy.sh : deploying finished."