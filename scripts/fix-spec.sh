#!/usr/bin/env bash

set -euo pipefail

# Remove "null" item from nullable enums
# (it could be more elegant, but it works for the current schema)
sed -i '/^\s*- null$/d' api/openapi.yaml
