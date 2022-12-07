#!/bin/bash

IMAGE=csmith/aoc-2022-01

if ! docker image inspect $IMAGE >/dev/null 2>&1; then
  echo "One time setup: building docker image..."
  (cd .docker && docker build . -t $IMAGE)
fi

docker run --rm -it -v "$(pwd)":/code $IMAGE /entrypoint.sh "$@"
