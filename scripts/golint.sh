#!/bin/bash

# Verify that all files are correctly golint'd.
set -e -o nounset -o pipefail
counter=0
while read -r line; do
	echo $line
	: $((counter++))
done < <(golint ./...)

if ((counter == 0)); then
	exit 0
fi
exit 1
