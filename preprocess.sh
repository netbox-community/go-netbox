#/usr/bin/env bash

### Set omitempty=false for the optional properties of the WritableVirtualMachineWithConfigContext object

set -x
# with_entries(foo) is a shorthand for to_entries | map(foo) | from_entries per https://stedolan.github.io/jq/manual/#to_entries,from_entries,with_entries
# jq '.definitions.WritableVirtualMachineWithConfigContext.properties = (.definitions.WritableVirtualMachineWithConfigContext.properties | to_entries | map(if .value."x-nullable" then .value."x-omitempty"=false else . end) | from_entries)' swagger.json > bla.json
jq '.definitions.WritableVirtualMachineWithConfigContext.properties = (.definitions.WritableVirtualMachineWithConfigContext.properties | with_entries(if .value."x-nullable" then .value."x-omitempty"=false else . end))' $1 | \
jq '.definitions.WritableIPAddress.properties = (.definitions.WritableIPAddress.properties | with_entries(if .value."x-nullable" then .value."x-omitempty"=false else . end))'
