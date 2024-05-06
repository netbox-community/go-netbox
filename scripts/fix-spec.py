#!/usr/bin/env python3

import yaml

SPEC_PATH = 'api/openapi.yaml'

# Load the spec file
with open(SPEC_PATH, 'r') as file:
    data = yaml.load(file, Loader=yaml.CLoader)

# Traverse schemas
if 'components' in data and 'schemas' in data['components']:
    for name, schema in data['components']['schemas'].items():
        if 'properties' in schema:
            # Remove "null" item from nullable enums
            for name, prop in schema['properties'].items():
                if 'enum' in prop and None in prop['enum']:
                    prop['enum'].remove(None)
                if 'properties' in prop and 'value' in prop['properties'] and 'enum' in prop['properties']['value'] and None in prop['properties']['value']['enum']:
                    prop['properties']['value']['enum'].remove(None)

            # Fix nullable types
            nullable_types = [
                'parent_device',
                'primary_ip',
            ]

            for ntype in nullable_types:
                if ntype in schema['properties']:
                    schema['properties'][ntype]['nullable'] = True

data["components"]["schemas"]["Site"]["properties"]["asns"]["items"].pop("type")
data["components"]["schemas"]["Interface"]["properties"]["tagged_vlans"]["items"].pop("type")
data["components"]["schemas"]["Site"]["properties"]["asns"]["items"]["$ref"] = "#/components/schemas/NestedASN"
data["components"]["schemas"]["Interface"]["properties"]["tagged_vlans"]["items"]["$ref"] = "#/components/schemas/NestedVLAN"
data["components"]["schemas"]["NestedASN"] = {
    "type": "object",
    "description": """Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
        dictionary of attributes which can be used to uniquely identify the related object. This class should be
        subclassed to return a full representation of the related object on read.""",
    "properties": {
        "id": {"type": "integer", "readOnly": True},
        "url": {"type": "string", "format": "uri", "readOnly": True},
        "display": {"type": "string", "readOnly": True},
        "asn": {
            "type": "integer",
            "maximum": 4294967295,
            "minimum": 1,
            "format": "int64",
            "description": "16- or 32-bit autonomous system number",
        },
    },
    "required": ["asn", "display", "id", "url"],
}

# Save the spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)
