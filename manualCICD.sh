#!/bin/bash
set -euo pipefail

git pull --ff-only
./build\&deploy.sh
