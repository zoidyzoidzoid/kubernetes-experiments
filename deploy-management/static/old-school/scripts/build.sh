#!/usr/bin/env bash

SCM_HASH="$(git rev-parse --short HEAD)"

PROJECT_ROOT="$(dirname "${BASH_SOURCE}")/.."
docker build -t "example-project:${SCM_HASH}" -f "${PROJECT_ROOT}/Dockerfile" "${PROJECT_ROOT}"
