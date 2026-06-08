#!/bin/bash
# author: df
# date: 2024-07-24
# desc: legacy wrapper retained for platforms that expect an entrypoint script

set -euo pipefail

exec /app/hidereplier-go
