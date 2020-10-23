#!/usr/bin/env bash

docker build --tag challenge:1.0 .
docker run --publish 8080:8080 --detach --name challenge challenge:1.0