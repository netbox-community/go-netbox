#!/usr/bin/python3
import json
import logging
import os

# logging configuration
LOGLEVEL = os.environ.get("LOGLEVEL", default="INFO").upper()
logging.basicConfig(level=getattr(logging, LOGLEVEL))
logger = logging.getLogger(__name__)

source_file = "swagger.json"

# First, add a default response to all api calls. If these do not exists, we get no error messages from the API.
# This implements https://github.com/netbox-community/go-netbox/commit/674aa060f07457c7dccfc0597e35899448668ef5
with open(source_file, "r") as readfile:
    data = json.load(readfile)

default_response = {
    "description": "",
    "schema": {"type": "object", "additionalProperties": True},
}

for path, path_spec in data["paths"].items():
    logging.debug("checking path: " + path)
    for verb, verb_spec in path_spec.items():
        # ignore the parameters "verb"
        if verb == "parameters":
            continue
        if "responses" in verb_spec and "default" not in verb_spec["responses"]:
            logging.info("Adding default response to " + path + " " + verb)
            verb_spec["responses"]["default"] = default_response

# Second, fix the config contexts and local context data to our needs
# This implements https://github.com/fbreckle/go-netbox/commit/987669a91daf2d04a155feaf032a24ed684169b7 and https://github.com/fbreckle/go-netbox/commit/030637bab4bb25cec173035cd2d4a78fb3f47053
#for definition, definition_spec in data["definitions"].items():
#    for prop, prop_spec in definition_spec["properties"].items():
#        # for config_context, change additionalProperties.type from string to object
#        if prop == "config_context":
#            if prop_spec["additionalProperties"]["type"] == "string":
#                logging.debug(
#                    f"{definition}.{prop}.additionalProperties.type == string"
#                )
#                logging.debug("before: " + str(prop_spec))
#                prop_spec["additionalProperties"] = {"type": "object"}
#                logging.debug("after: " + str(prop_spec))
#                logging.info(f"fixed {definition}.{prop}")
#
#        # for local_context_data, change type from string to object
#        # and add additionalProperties with type object
#        if prop == "local_context_data":
#            if prop_spec["type"] == "string":
#                logging.debug(f"{definition}.{prop}.type == string")
#                logging.debug("before: " + str(prop_spec))
#                prop_spec["type"] = "object"
#                prop_spec["additionalProperties"] = {"type": "object"}
#                logging.debug("after: " + str(prop_spec))
#                logging.info(f"fixed {definition}.{prop}")


# Third, allow unsetting some optional attributes
# This implements https://github.com/fbreckle/go-netbox/commit/23f80e06cb3da6f0fd9974a2d986c09826e518b5
for prop, prop_spec in data["definitions"]["WritableIPAddress"]["properties"].items():
    if (
        "x-nullable" in prop_spec
        and prop_spec["x-nullable"] == True
        and prop_spec["type"] == "integer"
    ):
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableIPAddress.{prop}")
    if prop == "tags":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableIPAddress.{prop}")

data["definitions"]["WritableCustomField"]["properties"]["required"][
    "x-omitempty"
] = False

# This implements https://github.com/fbreckle/go-netbox/commit/1363e14cfc7bce4bd3d5ee93c09ca70543c51279
for prop, prop_spec in data["definitions"]["WritableVirtualMachineWithConfigContext"][
    "properties"
].items():
    if "x-nullable" in prop_spec and prop_spec["x-nullable"] == True:
        prop_spec["x-omitempty"] = False
        logging.info(
            f"set x-omitempty = false on WritableVirtualMachineWithConfigContext.{prop}"
        )
    if prop == "tags":
        prop_spec["x-omitempty"] = False
        logging.info(
            f"set x-omitempty = false on WritableVirtualMachineWithConfigContext.{prop}"
        )

# Fourth, add schema for the 200 response of the /status/ endpoint
data["paths"]["/status/"]["get"]["responses"]["200"]["schema"] = {
    "type": "object",
    "additionalProperties": True,
}
logging.info(f"Added schema for 200 response of /status/ get")

# change model IPAddress's assigned_object property from string to object otherwise it fails to unmarshall
data["definitions"]["IPAddress"]["properties"]["assigned_object"] = {
    "title": "Assigned object",
    "type": "object",
    "readOnly": True,
}
logging.info(
    f"Fix 'error: json: cannot unmarshal number into Go struct field IPAddress.assigned_object of type \
string' when creating available-ips"
)

# Change model returned by paths /available-ips/ from AvailableIP to IPAddress.
data["paths"]["/ipam/ip-ranges/{id}/available-ips/"]["post"]["responses"]["201"][
    "schema"
]["items"] = {"$ref": "#/definitions/IPAddress"}
logging.info(f"Corrected reponse model when creating available-ips in an IP range")

data["paths"]["/ipam/prefixes/{id}/available-ips/"]["post"]["responses"]["201"][
    "schema"
]["items"] = {"$ref": "#/definitions/IPAddress"}
logging.info(f"Corrected reponse model when creating available-ips in a prefix")

# Remove omitempty for site attribute on clusters
for prop, prop_spec in data["definitions"]["WritableCluster"]["properties"].items():
    if prop == "site":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableCluster.{prop}")

# Remove omitempty for description attribute on tags
for prop, prop_spec in data["definitions"]["Tag"]["properties"].items():
    if prop == "description":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on Tag.{prop}")

# Add omitempty = false for prefix mark_utilized and is_pool
for prop, prop_spec in data["definitions"]["WritablePrefix"]["properties"].items():
    if prop in ["mark_utilized", "is_pool"]:
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritablePrefix.{prop}")

# Add omitempty = false for prefix mark_utilized
for prop, prop_spec in data["definitions"]["WritablePrefix"]["properties"].items():
    if prop == "mark_utilized":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritablePrefix.{prop}")

# Add omitempty = false for ipaddress role
for prop, prop_spec in data["definitions"]["WritableIPAddress"]["properties"].items():
    if prop == "role":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableIPAddress.{prop}")

# Delete problematic maximums (might have to be replaced with a proper value)
for definition, definition_spec in data["definitions"].items():
    for prop, prop_spec in definition_spec["properties"].items():
        if (
            "maximum" in prop_spec.keys()
            and prop_spec["maximum"] == 9223372036854776000
        ):
            del prop_spec["maximum"]
            logging.info(f"deleted maximum of {definition}.{prop}")

# Add custom fields to PrefixLength (https://github.com/fbreckle/go-netbox/pull/11)
data["definitions"]["PrefixLength"]["properties"]["custom_fields"] = {
    "title": "Custom fields",
    "type": "object",
    "default": {},
}

# Restore old version of the available-ip endpoints
logging.info("Restore old version of the parameters of prefixes available ip endpoint")
data["paths"]["/ipam/prefixes/{id}/available-ips/"]["post"]["parameters"] = [
    {
        "name": "data",
        "in": "body",
        "required": True,
        "schema": {"type": "array", "items": {"$ref": "#/definitions/AvailableIP"}},
    }
]

logging.info(
    "Restore old version of the parameters of the ip-ranges available ip endpoint"
)
data["paths"]["/ipam/ip-ranges/{id}/available-ips/"]["post"]["parameters"] = [
    {
        "name": "data",
        "in": "body",
        "required": True,
        "schema": {"type": "array", "items": {"$ref": "#/definitions/AvailableIP"}},
    }
]

logging.info(
    "Restore old version of the parameters of the ip-ranges available ip endpoint"
)
data["paths"]["/ipam/ip-ranges/{id}/available-ips/"]["post"]["parameters"] = [
    {
        "name": "data",
        "in": "body",
        "required": True,
        "schema": {"type": "array", "items": {"$ref": "#/definitions/AvailableIP"}},
    }
]

logging.info("Restore old version of the response of the available ip endpoint")
data["paths"]["/ipam/prefixes/{id}/available-prefixes/"]["post"]["responses"]["201"] = {
    "description": "",
    "schema": {"$ref": "#/definitions/Prefix"},
}

# Write output file
with open("swagger.processed.json", "w") as writefile:
    json.dump(data, writefile, indent=2)


print("Swaggerfile preprocession complete")
