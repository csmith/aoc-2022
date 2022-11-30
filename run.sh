#!/bin/bash

docker run --rm -it -v "$(pwd)":/code postgres:15.1 /code/.docker/entrypoint.sh $@
