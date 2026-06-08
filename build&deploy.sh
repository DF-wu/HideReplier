#!/usr/bin/env bash
set -euo pipefail

# Go is the primary runtime on this branch. Java/Spring Boot remains in src/ as
# legacy code, but the production deploy path builds the Go binary and Vite UI.

BOT_VERSION="${BOT_VERSION:-1.0.4}"
IMAGE_NAME="${IMAGE_NAME:-dfder/hidereplier}"

echo "build-deploy: checking Go tests"
go test ./...

echo "build-deploy: checking Vite frontend build"
pushd web >/dev/null
VITE_BOT_VERSION="$BOT_VERSION" npm ci
VITE_BOT_VERSION="$BOT_VERSION" npm run build
popd >/dev/null

echo "build-deploy: building Docker image $IMAGE_NAME"
docker build \
  --build-arg "VITE_BOT_VERSION=$BOT_VERSION" \
  -t "$IMAGE_NAME" .

if [[ "${PUSH_IMAGE:-0}" == "1" ]]; then
  echo "build-deploy: pushing Docker image $IMAGE_NAME"
  docker push "$IMAGE_NAME"
fi

echo "build-deploy: deploying to Fly.io"
fly deploy \
  --build-arg "VITE_BOT_VERSION=$BOT_VERSION"

echo "build-deploy: done"
