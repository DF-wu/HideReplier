#!/usr/bin/env bash

# WARN: please remember to set TENOR_API_KEY & BOT_VERSION environment variables
docker run --rm \
  -v .:/workspace \
  -e VITE_TENOR_API_KEY="$TENOR_API_KEY" \
  -e VITE_BOT_VERSION="$BOT_VERSION" \
  --entrypoint /bin/sh \
  node:lts-alpine -c "cd /workspace && npm i && npm run build"
