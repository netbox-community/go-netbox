#!/usr/bin/env bash

set -euo pipefail

export SRC_DIR='./netbox'

# Prepare environment
rm -rf netbox/client netbox/models
go mod download

# Run pre-generation hooks
for SCRIPT in scripts/pre-generation/*; do
  "${SCRIPT}"
done

# Generate code
swagger generate client -f api/openapi.json --target="${SRC_DIR}" --copyright-file=assets/copyright_header.txt
go mod tidy

# Run post-generation hooks
for SCRIPT in scripts/post-generation/*; do
  "${SCRIPT}"
done

# Format generated code and fix imports
goimports -w "${SRC_DIR}"
