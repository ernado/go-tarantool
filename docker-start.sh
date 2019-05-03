#!/usr/bin/env bash

# Starts docker container with name tarantool-go-test so tests can run.
docker rm -f tarantool-go-test

docker run --name tarantool-go-test -p 3013:3013 -d \
  -v $(pwd)/config-docker.lua:/opt/tarantool/config.lua:ro tarantool/tarantool tarantool \
  /opt/tarantool/config.lua
