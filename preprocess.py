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

for path, path_spec in data["paths"].items():
    logging.debug("checking path: " + path)
    for verb, verb_spec in path_spec.items():
        if "parameters" in verb_spec:
            tag_locations = [
                i for i, t in enumerate(verb_spec["parameters"]) if t["name"] == "tag"
            ]
            if tag_locations:
                for tag_location in tag_locations:
                    logging.info(
                        "Changing tag query to support multiple tags "
                        + path
                        + " "
                        + verb
                    )
                    verb_spec["parameters"][tag_location] = {
                        "name": "tag",
                        "in": "query",
                        "description": "",
                        "required": False,
                        "type": "array",
                        "items": {"type": "string"},
                        "collectionFormat": "multi",
                    }

# Second, fix the config contexts and local context data to our needs
# This implements https://github.com/fbreckle/go-netbox/commit/987669a91daf2d04a155feaf032a24ed684169b7 and https://github.com/fbreckle/go-netbox/commit/030637bab4bb25cec173035cd2d4a78fb3f47053
# for definition, definition_spec in data["definitions"].items():
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

for model in [
    "WritableRack",
    "WritablePowerPanel",
    "WritableModuleType",
    "WritableRearPort",
    "WritablePowerPort",
    "WritablePowerOutlet",
    "WritablePowerFeed",
    "WritableModule",
    "WritableModuleBay",
    "WritableFrontPort",
    "WritableConsoleServerPort",
    "WritableConsolePort",
    "WritableCable",
    "InventoryItemRole",
    "WritableInventoryItem",
]:
    for prop, prop_spec in data["definitions"][model]["properties"].items():
        if (
            "x-nullable" in prop_spec
            and prop_spec["x-nullable"] == True
            and prop_spec["type"] in ["integer", "number"]
        ):
            prop_spec["x-omitempty"] = False
            logging.info(f"set x-omitempty = false on {model}.{prop}")
        if prop in [
            "tags",
            "type",
            "outer_unit",
            "weight_unit",
            "asset_tag",
            "desc_units",
            "facility_id",
            "mark_connected",
            "feed_leg",
            "color",
            "component_type",
            "discovered",
        ]:
            prop_spec["x-omitempty"] = False
            logging.info(f"set x-omitempty = false on {model}.{prop}")

# fix generic objects that are incorrectly listed as strings
for model in data["definitions"]:
    for prop, prop_spec in data["definitions"][model]["properties"].items():
        if prop in ["link_peers", "connected_endpoints"]:
            prop_spec["items"]["type"] = "object"
            logging.info(f"set items.type = object on {model}.{prop}")

data["definitions"]["WritableCustomField"]["properties"]["required"][
    "x-omitempty"
] = False

data["definitions"]["WritableObjectPermission"]["properties"]["constraints"][
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

# Change model returned by paths /available-ips/ from AvailableIP to IPAddress.
data["paths"]["/ipam/ip-ranges/{id}/available-ips/"]["post"]["responses"]["201"][
    "schema"
]["items"] = {"$ref": "#/definitions/IPAddress"}
logging.info(f"Corrected reponse model when creating available-ips in an IP range")

data["paths"]["/ipam/prefixes/{id}/available-ips/"]["post"]["responses"]["201"][
    "schema"
]["items"] = {"$ref": "#/definitions/IPAddress"}
logging.info(f"Corrected reponse model when creating available-ips in a prefix")

# Add omitempty = false for site attribute on clusters
for prop, prop_spec in data["definitions"]["WritableCluster"]["properties"].items():
    if prop == "site":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableCluster.{prop}")

# Add omitempty = false for description attribute on tags
for prop, prop_spec in data["definitions"]["Tag"]["properties"].items():
    if prop == "description":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on Tag.{prop}")

# Add omitempty = false for prefix mark_utilized, is_pool, site_id, vlan_id, role_id, tenant_id and vrf_id
for prop, prop_spec in data["definitions"]["WritablePrefix"]["properties"].items():
    if prop in ["mark_utilized", "is_pool", "site", "vlan", "role", "tenant", "vrf"]:
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritablePrefix.{prop}")

# Add omitempty = false for ipaddress role
for prop, prop_spec in data["definitions"]["WritableIPAddress"]["properties"].items():
    if prop == "role":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableIPAddress.{prop}")

# Add omitempty = false for token write_enabled
for prop, prop_spec in data["definitions"]["WritableToken"]["properties"].items():
    if prop == "write_enabled":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableIPAddress.{prop}")

# Add omitempty = false for interface enabled
for prop, prop_spec in data["definitions"]["WritableInterface"]["properties"].items():
    if prop == "enabled":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableInterface.{prop}")

# Add omitempty = false for vminterface enabled
for prop, prop_spec in data["definitions"]["WritableVMInterface"]["properties"].items():
    if prop == "enabled":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-omitempty = false on WritableVMInterface.{prop}")

# Delete problematic maximums (might have to be replaced with a proper value)
for definition, definition_spec in data["definitions"].items():
    for prop, prop_spec in definition_spec["properties"].items():
        if (
            "maximum" in prop_spec.keys()
            and prop_spec["maximum"] == 9223372036854776000
        ):
            del prop_spec["maximum"]
            logging.info(f"deleted maximum of {definition}.{prop}")
        if (
            "minimum" in prop_spec.keys()
            and prop_spec["minimum"] == -9223372036854776000
        ):
            del prop_spec["minimum"]
            logging.info(f"deleted minimum of {definition}.{prop}")

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

# ASNs now have a full RIR object, not just a rir id
logging.info("Make ASN.rir a full RIR object")
data["definitions"]["ASN"]["properties"]["rir"] = {
  "$ref": "#/definitions/NestedRIR"
}

# Write output file
with open("swagger.processed.json", "w") as writefile:
    json.dump(data, writefile, indent=2)


print("Swaggerfile preprocession complete")
