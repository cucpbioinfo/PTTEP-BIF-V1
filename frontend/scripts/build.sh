#!/bin/bash
# Build docker image
cd "$(dirname "$0")"
docker build -t edna-web ../
