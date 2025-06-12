#!/bin/sh

mkdir -p ./sample.d/public.d

echo hw1 > ./sample.d/public.d/hw1.txt
echo hw2 > ./sample.d/public.d/hw2.txt

export ENV_ROOT_DIR=./sample.d/public.d

export ENV_ADDR_PORT=127.0.0.1:61280

./go-simple-webdav-server
