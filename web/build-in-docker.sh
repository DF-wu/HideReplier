#!/usr/bin/env bash

# WARN: please remember to set TENOR_API_KEY environment variable
docker run --rm \
  -v .:/workspace \
  -e VITE_TENOR_API_KEY="$TENOR_API_KEY" \
  --entrypoint /bin/sh \
  node:lts-alpine -c "cd /workspace && npm i && npm run build"
