#!/bin/sh

docker build -t christmas_comp . 2>&1 >/dev/null
docker run --memory=1G christmas_comp
