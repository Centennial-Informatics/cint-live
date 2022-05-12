#!/bin/bash

docker build . --tag live-judge-image

docker run \
  -e CF_ACCOUNTS="$CF_ACCOUNTS" \
  -e ADMIN_TOKEN="$ADMIN_TOKEN" \
  -e DATABASE_URL="$DATABASE_URL" \
  -e BACKEND_HOST="$BACKEND_HOST" \
  --name live-judge \
  -p 8000:8000 \
  -it \
  --rm \
  live-judge-image