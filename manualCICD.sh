#!/bin/bash
git pull && \
sh build.sh && \
fly deploy