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
for definition, definition_spec in data["definitions"].items():
    for prop, prop_spec in definition_spec["properties"].items():
        # for config_context, change additionalProperties.type from string to object
        if prop == "config_context":
            if prop_spec["additionalProperties"]["type"] == "string":
                logging.debug(
                    f"{definition}.{prop}.additionalProperties.type == string"
                )
                logging.debug("before: " + str(prop_spec))
                prop_spec["additionalProperties"] = {"type": "object"}
                logging.debug("after: " + str(prop_spec))
                logging.info(f"fixed {definition}.{prop}")

        # for local_context_data, change type from string to object
        # and add additionalProperties with type object
        if prop == "local_context_data":
            if prop_spec["type"] == "string":
                logging.debug(f"{definition}.{prop}.type == string")
                logging.debug("before: " + str(prop_spec))
                prop_spec["type"] = "object"
                prop_spec["additionalProperties"] = {"type": "object"}
                logging.debug("after: " + str(prop_spec))
                logging.info(f"fixed {definition}.{prop}")


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

# Remove omitempty for site attribute on clusters
for prop, prop_spec in data["definitions"]["WritableCluster"]["properties"].items():
    if prop == "site":
        prop_spec["x-omitempty"] = False
        logging.info(f"set x-nullable = false on WritableCluster.{prop}")

# Write output file
with open("swagger.processed.json", "w") as writefile:
    json.dump(data, writefile, indent=2)


print("Swaggerfile preprocession complete")
print("Do not forget to apply the IP Address change at")
print(
    "* https://github.com/fbreckle/go-netbox/commit/1363e14cfc7bce4bd3d5ee93c09ca70543c51279#diff-9688c8954edccf890a0b6a8e1bea4d9f81083d0fe1214be0fea2514c0ed3f651 *"
)
