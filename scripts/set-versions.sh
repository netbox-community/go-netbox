#!/usr/bin/env bash

NETBOX_VERSION='latest'
NETBOX_DOCKER_VERSION='auto'

set -euo pipefail

# Parse Netbox version
if [ "$#" -gt 0 ]; then
  NETBOX_VERSION="$1"
fi

if [ "${NETBOX_VERSION}" == 'latest' ]; then
  NETBOX_VERSION=$(curl --silent https://api.github.com/repos/netbox-community/netbox/releases/latest |
    jq '.tag_name' |
    sed -E 's/"v([^"]+)"/\1/')
fi

# Parse Netbox Docker version
if [ "$#" -gt 1 ]; then
  NETBOX_DOCKER_VERSION="$2"
fi

if [ "${NETBOX_DOCKER_VERSION}" == 'auto' ]; then
  NEXT='https://registry.hub.docker.com/v2/repositories/netboxcommunity/netbox/tags/?page=1&page_size=1000'
  NETBOX_DOCKER_VERSION=''

  while [ "${NEXT}" != "null" ] && [ "${NETBOX_DOCKER_VERSION}" == "" ]; do
    RESPONSE=$(curl --silent "${NEXT}")
    NEXT=$(echo -E "${RESPONSE}" | jq -r '.next')

    NETBOX_DOCKER_VERSION=$(echo -E "${RESPONSE}" |
      jq -r ".results[] | select((.name | startswith(\"v${NETBOX_VERSION}-\")) and (.name | index(\"ldap\") | not)) | .name")
    NETBOX_DOCKER_VERSION=${NETBOX_DOCKER_VERSION##*-}
  done

  unset NEXT
  unset RESPONSE
fi

# Print variables
echo ${NETBOX_VERSION} > api/netbox_version
echo ${NETBOX_DOCKER_VERSION} > api/netbox_docker_version
