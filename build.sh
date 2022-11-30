#!/bin/bash
# 不要在container裡面用maven build ， 不然會很慢
# 不要在container裡面用maven build ， 不然會很慢
# 不要在container裡面用maven build ， 不然會很慢
# by stanley2058@yahoo.com.tw
# 所以這是先在外面build ，再把build好的檔案 copy到working dir 再tag

echo "build.sh : start maven build"
mvn -T 1C clean install -Dmaven.test.skip=true
echo "build.sh : maven build successfully"

echo "build.sh : start to copy built files"
cp ./target/hidereplier-0.0.1-SNAPSHOT.jar app.jar
echo "build.sh : cp apps from target directory to root directory"


echo "build.sh : building docker image"
docker build . -t dfder/hidereplier
echo "build.sh : docker build successfully"

echo "build.sh : pushing to dockerhub"
docker push dfder/hidereplier
echo "build.sh : uploaded."