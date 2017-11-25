#!/bin/sh

export PRESENTS=50000000

docker build -t christmas_comp . 2>&1 >/dev/null
docker run --memory=1G -e PRESENTS christmas_comp
