#!/bin/bash
git pull && \
sh build.sh && \
flyctl deploy