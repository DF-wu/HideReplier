#!/usr/bin/env bash

# WARN: please remember to set BOT_VERSION when building release assets.
docker run --rm \
  -v .:/workspace \
  -e VITE_BOT_VERSION="$BOT_VERSION" \
  --entrypoint /bin/sh \
  node:lts-alpine -c "cd /workspace && npm ci && npm run build"
