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
            # List to store dynamically found nullable properties
            dynamic_non_required_props = []

            # Remove "null" item from nullable enums
            for propName, prop in schema['properties'].items():
                if 'enum' in prop and None in prop['enum']:
                    prop['enum'].remove(None)
                if 'properties' in prop and 'value' in prop['properties'] and 'enum' in prop['properties']['value'] and None in prop['properties']['value']['enum']:
                    prop['properties']['value']['enum'].remove(None)

                # If property is nullable, add to non-required props
                if prop.get('nullable', False):
                    dynamic_non_required_props.append(propName)

            # Fix nullable types
            nullable_types = [
                'parent_device',
                'primary_ip',
            ]

            for ntype in nullable_types:
                if ntype in schema['properties']:
                    schema['properties'][ntype]['nullable'] = True
                    # Add to non-required if marked as nullable
                    if ntype not in dynamic_non_required_props:
                        dynamic_non_required_props.append(ntype)

            # Fix non-nullable types
            # See: https://github.com/OpenAPITools/openapi-generator/issues/18006
            non_nullable_types = [
                'front_image',
                'rear_image',
            ]

            for ntype in non_nullable_types:
                if ntype in schema['properties']:
                    if schema['properties'][ntype]['format'] == 'binary':
                        schema['properties'][ntype].pop('nullable', None)

        # Combine the original non-required props with dynamically found nullable properties
        non_required_props = [
            'devicetype_count',
            'device_count',
            'virtualmachine_count',
            'prefix_count',
            'vlan_count',
            'display_url',
            'user_count',
            'aggregate_count',
            'asn_count',
            'choices_count',
            'circuit_count',
            'cluster_count',
            'contact_count',
            'file_count',
            'interface_count',
            'inventoryitem_count',
            'ipaddress_count',
            'member_count',
            'platform_count',
            'powerfeed_count',
            'provider_count',
            'rack_count',
            'site_count',
            'tenant_count',
            'terminations_count',
            'tunnel_count',
            'vrf_count',
            'wirelesslan_count',
            'created',  # for sites
            'provider_count', # ASNs
        ] + dynamic_non_required_props

        if 'required' in schema:
            # Filter out nullable properties from the required list
            schema['required'] = [prop for prop in schema['required'] if prop not in non_required_props]

# Add required fields to the Device schema
data['components']['schemas']['Device']['required'] = [
    'id',
    'url',
    'display',
    'name',
    'description',
]

# Save the spec file
with open(SPEC_PATH, 'w') as file:
    yaml.dump(data, file, Dumper=yaml.CDumper, sort_keys=False)