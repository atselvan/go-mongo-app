#!/bin/bash

# Variables
image="mongo"
name="mongo-db"
network="isolated_nw"
port="27017:27017"

running=`docker ps | grep -c $name`
if [ $running -gt 0 ]
then
   echo "Stopping $name instance"
   docker stop $name
fi

existing=`docker ps -a | grep -c $name`
if [ $existing -gt 0 ]
then
   echo "Removing $name container"
   docker rm $name
fi

echo "Running a new instance with name $name"
echo "[INFO] IMAGE   : $image"
echo "[INFO] NAME    : $name"
echo "[INFO] NETWORK : $network"
echo "[INFO] PORT    : $port"

docker run --name $name -d -p $port --network $network $image
docker cp add-data.js $name:/tmp/add-data.js
docker exec -it $name mongo /tmp/add-data.js