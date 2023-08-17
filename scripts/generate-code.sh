#!/usr/bin/env bash

set -euo pipefail

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.6.0 generate \
    --config /local/.openapi-generator/config.yaml \
    --input-spec /local/api/openapi.yaml \
    --output /local \
    --http-user-agent go-netbox/$(cat api/netbox_version)
