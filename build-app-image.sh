#!/bin/bash

# Variables
image="mongo-app"

echo "Building new image with tag: $TAGNAME"
docker build -t $image .