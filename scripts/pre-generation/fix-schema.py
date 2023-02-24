#!/usr/bin/env python3

import json
import argparse
import logging
import sys


def init_logger(name, verbose, logfile=None):
    """
    Init logger object, prepare display for CLI and log to file
    :param name: Name to pass to Formatter
    :param verbose: Verbosity level will be used for console
    :param logfile: Path of the logfile
    :return: logging.logger object fully initialized
    """
    level_mapper = dict()
    level_mapper = {0: logging.CRITICAL, 1: logging.INFO, 2: logging.DEBUG}
    console = True if verbose > 0 else False
    logger = logging.getLogger(name)
    # logger.setLevel(logging.DEBUG)
    logger.setLevel(level_mapper[verbose])
    f = logging.Formatter(
        '%(asctime)s | %(name)s | [%(levelname)s] - %(message)s')
    if console is True:
        ch = logging.StreamHandler()
        ch.setLevel(logging.DEBUG)
        ch.setFormatter(f)
        logger.addHandler(ch)
        logger.info('Logger writing on console')
    if logfile is not None:
        fh = logging.FileHandler(logfile)
        fh.setLevel(logging.INFO)
        fh.setFormatter(f)
        logger.addHandler(fh)
        logger.info('Logger writing to %s log file' % logfile)
    logger.info('Succesfully initialized logger')
    return logger


def cli_arguments():
    """
    Get cli arguments.
    :return: argparse.parser object
    """
    parser = argparse.ArgumentParser(
        description=('Modifications on openapi.json.'))
    parser.add_argument('-v',
                        '--verbose',
                        action='count',
                        default=0,
                        help='Increase output verbosity.')
    parser.add_argument('-o',
                        '--output',
                        default='api/openapi.json',
                        help='Output file')
    parser.add_argument('input',
                        nargs='?',
                        default='api/openapi.json',
                        help='Original input json file')

    return parser.parse_args()


def complete_data(input_file, log):
    """
    Complete json data loaded from input file
    :return: json structure
    """
    try:
        f = open(input_file)
    except FileNotFoundError as e:
        print('Error: {}'.format(e))
        sys.exit(1)

    data = json.load(f)
    f.close()

    # Iterating over paths to add a default in responses for each verb
    # Add a default response for each verb of each paths
    # to get the return of netbox in case of failure
    verbs = ('get', 'delete', 'patch', 'post', 'put')
    missing_default = dict()
    missing_default['default'] = {
        'description': '',
        'schema': {
            'additionalProperties': True,
            'type': 'object'
        }
    }

    for path in data['paths']:
        for verb in data['paths'][path]:
            if verb in verbs:
                if 'default' not in data['paths'][path][verb]['responses']:
                    log.info('Adding default in {} verb of {} path'.format(
                        verb, path))
                    data['paths'][path][verb]['responses'].update(
                        missing_default)
                else:
                    log.info('default defined in {} ({})'.format(verb, path))
            else:
                # If verb is not known stop here
                if verb != 'parameters':
                    log.crit('Verb {} unknown.'.format(verb))
                    sys.exit(1)

    # Adding the x-omitempty option to avoid sending empty parameters
    # as netbox will return an error
    modify_properties = dict()
    modify_properties['tags'] = {'x-omitempty': True}
    modify_properties['asns'] = {'x-omitempty': True}

    # Iterating over definitions to complete properties
    for definition in data['definitions']:
        def_properties = data['definitions'][definition]['properties']
        for def_property in def_properties:
            for prop in modify_properties:
                if def_property in prop:
                    for k in modify_properties[prop]:
                        if k not in def_properties[def_property]:
                            log.info('Adding {} in {} property'.format(
                                k, prop))
                            def_properties[def_property].update(
                                modify_properties[prop])

    return (data)


def write_results(data, output_file, log):
    """
    Json dump data in output file
    """
    try:
        f = open(output_file, 'w')
    except FileNotFoundError as e:
        log.error('Error opening output file {} ({})'.format(output_file, e))
        sys.exit(1)
    except PermissionError as e:
        log.error('{}'.format(e))
        sys.exit(1)

    json.dump(data, f, indent=4)
    f.write('\n')
    f.close()
    log.info('Data written in {}'.format(output_file))


def main():
    args = cli_arguments()
    # Get logger
    log = init_logger('openapi', args.verbose, logfile='/tmp/openapi.log')

    modified_data = complete_data(args.input, log)
    write_results(modified_data, args.output, log)


if __name__ == '__main__':
    main()
