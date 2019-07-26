#!/usr/bin/python
""" Verify ifconfig state sync """

#
# This file is part of Ansible
#
# Ansible is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Ansible is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Ansible. If not, see <http://www.gnu.org/licenses/>.
#

import shlex
import calendar

from collections import OrderedDict
import time
from ansible.module_utils.basic import AnsibleModule

DOCUMENTATION = """
---
module: test_ifconfig_state_sync
author: Platina Systems
short_description: Module to test ifconfig state sync.
description:
    Module to test ifconfig state sync.
options:
    switch_name:
      description:
        - Name of the switch on which tests will be performed.
      required: False
      type: str
    eth:
      description:
        - Name of the interface on which test is to be performed.
      required: False
      type: str
    start_time: 
      description:
        - The start time recorded in the starting of playbook execution.
      required: False
      type: str
    log_file_name:
      description:
        - Comma separated names of log files in which core dumps are to be checked(daemon.log, kern.log, syslog, dmesg).
      required: False
      type: str
    error_msg:
      description:
        - The msg which is to be checked for in the log file.
      required: False
      type: str
    hash_name:
      description:
        - Name of the hash in which to store the result in redis.
      required: False
      type: str
    log_dir_path:
      description:
        - Path to log directory where logs will be stored.
      required: False
      type: str
"""

EXAMPLES = """
- name: Verify ifconfig state out of sync with goes and ip route
  test_ifconfig_state_sync:
    switch_name: "{{ inventory_hostname }}"
    hash_name: "{{ hostvars['server_emulator']['hash_name'] }}"
    log_dir_path: "{{ issues_log_dir }}"
"""

RETURN = """
hash_dict:
  description: Dictionary containing key value pairs to store in hash.
  returned: always
  type: dict
"""

RESULT_STATUS = True
HASH_DICT = OrderedDict()
failure_summary = ''


def run_cli(module, cli):
    """
    Method to execute the cli command on the target node(s) and
    returns the output.
    :param module: The Ansible module to fetch input parameters.
    :param cli: The complete cli string to be executed on the target node(s).
    :return: Output/Error or None depending upon the response from cli.
    """
    cli = shlex.split(cli)
    rc, out, err = module.run_command(cli)

    if out:
        return out.rstrip()
    elif err:
        return err.rstrip()
    else:
        return None


def execute_commands(module, cmd):
    """
    Method to execute given commands and return the output.
    :param module: The Ansible module to fetch input parameters.
    :param cmd: Command to execute.
    :return: Output of the commands.
    """
    global HASH_DICT

    out = run_cli(module, cmd)

    # Store command prefixed with exec time as key and
    # command output as value in the hash dictionary
    exec_time = run_cli(module, 'date +%Y%m%d%T')
    key = '{0} {1} {2}'.format(module.params['switch_name'], exec_time, cmd)

    if out:
        HASH_DICT[key] = out[:512] if len(out.encode('utf-8')) > 512 else out
    else:
        HASH_DICT[key] = out

    return out


def verify_log_dumps(module):
    """
    Method to verify dumps in log files on invader.
    :param module: The Ansible module to fetch input parameters.
    """
    global RESULT_STATUS, failure_summary

    switch_name = module.params['switch_name']
    start_time = module.params['start_time']
    log_file_names = module.params['log_file_names'].split(',')
    error_msg = module.params['error_msg']

    # Check for errors in log files
    for file_name in log_file_names:
        with open(file_name, 'r') as f:
            content = f.read()
        log_file = content.split('\n\n')

        # Getting the time stamp in log file's time stamp format
        month = start_time[4:6]
        month_abb = calendar.month_abbr[int(month)]
        date = start_time[6:8]
        if date[0] == '0':
            date = ' {}'.format(date[1])
        time = start_time[8::]
        time_stamp = month_abb + ' ' + date + ' ' + time

        core_dump = []
        error_found = False

        for line in log_file:
            if time_stamp in line:
                for i in range(log_file.index(line), len(log_file)):
                    if error_msg in log_file[i]:
                        core_dump.append(log_file[i])
                        error_found = True
                break

            errors = '\n'.join(core_dump)
            if error_found:
                RESULT_STATUS = False
                failure_summary += 'On switch {} '.format(switch_name)
                failure_summary += 'runtime error found in {} file.\n'.format(file_name)
                failure_summary += 'Errors:\n{}\n'.format(errors)

    # Get the GOES status info
    execute_commands(module, 'goes status')


def check_interface_status_up(module):
    """
    Method to verify ifconfig state out of sync with goes and ip route when interface is up.
    :param module: The Ansible module to fetch input parameters.
    """
    global RESULT_STATUS, failure_summary

    switch_name = module.params['switch_name']
    eth = module.params['eth']

    # Get the interface status
    cmd = 'goes hget platina-mk1 xeth{}.admin'.format(eth)
    status = execute_commands(module, cmd)
    if 'true' not in status:
        RESULT_STATUS = False
        failure_summary += 'On switch {} '.format(switch_name)
        failure_summary += 'xeth{} status is false '.format(eth)
        failure_summary += 'for command {}\n'.format(cmd)

    cmd = 'goes hget platina-mk1 xeth{}.link'.format(eth)
    status = execute_commands(module, cmd)
    if 'true' not in status:
        RESULT_STATUS = False
        failure_summary += 'On switch {} '.format(switch_name)
        failure_summary += 'xeth{} status is false '.format(eth)
        failure_summary += 'for command {}\n'.format(cmd)

    cmd = 'ifconfig xeth{}'.format(eth)
    status = execute_commands(module, cmd)
    if 'UP BROADCAST RUNNING' not in status:
        RESULT_STATUS = False
        failure_summary += 'On switch {} '.format(switch_name)
        failure_summary += 'xeth{} status is false '.format(eth)
        failure_summary += 'for command {}\n'.format(cmd)

    cmd = 'ip link show xeth{}'.format(eth)
    status = execute_commands(module, cmd)
    if 'state UP' not in status:
        RESULT_STATUS = False
        failure_summary += 'On switch {} '.format(switch_name)
        failure_summary += 'xeth{} status is false '.format(eth)
        failure_summary += 'for command {}\n'.format(cmd)


def check_interface_status_down(module):
    """
    Method to verify ifconfig state out of sync with goes and ip route when interface is up.
    :param module: The Ansible module to fetch input parameters.
    """
    global RESULT_STATUS, failure_summary

    switch_name = module.params['switch_name']
    eth = module.params['eth']

    # Get the interface status
    cmd = 'goes hget platina-mk1 xeth{}.admin'.format(eth)
    status = execute_commands(module, cmd)
    if 'false' not in status:
        RESULT_STATUS = False
        failure_summary += 'On switch {} '.format(switch_name)
        failure_summary += 'xeth{} status is true even after bringing down the interface '.format(eth)
        failure_summary += 'for command {}\n'.format(cmd)

    cmd = 'goes hget platina-mk1 xeth{}.link'.format(eth)
    status = execute_commands(module, cmd)
    if 'false' not in status:
        RESULT_STATUS = False
        failure_summary += 'On switch {} '.format(switch_name)
        failure_summary += 'xeth{} status is true even after bringing down the interface '.format(eth)
        failure_summary += 'for command {}\n'.format(cmd)

    cmd = 'ifconfig xeth{}'.format(eth)
    status = execute_commands(module, cmd)
    if 'UP BROADCAST RUNNING' in status:
        RESULT_STATUS = False
        failure_summary += 'On switch {} '.format(switch_name)
        failure_summary += 'xeth{} status is true even after bringing down the interface '.format(eth)
        failure_summary += 'for command {}\n'.format(cmd)

    cmd = 'ip link show xeth{}'.format(eth)
    status = execute_commands(module, cmd)
    if 'state UP' in status:
        RESULT_STATUS = False
        failure_summary += 'On switch {} '.format(switch_name)
        failure_summary += 'xeth{} status is true even after bringing down the interface '.format(eth)
        failure_summary += 'for command {}\n'.format(cmd)


def verify_ifconfig_state(module):
    """
    Method to verify ifconfig state out of sync with goes and ip route.
    :param module: The Ansible module to fetch input parameters.
    """
    global RESULT_STATUS, HASH_DICT, failure_summary

    eth = module.params['eth']

    # Restart goes
    execute_commands(module, 'goes restart')
    time.sleep(10)
    # Check interface state when interface is up
    check_interface_status_up(module)

    # Bring the interface down
    cmd = 'ifdown xeth{}'.format(eth)
    execute_commands(module, cmd)
    time.sleep(15)
    # Check interface state when interface is down
    check_interface_status_down(module)

    # Bring the interface up
    cmd = 'ifup xeth{}'.format(eth)
    execute_commands(module, cmd)
    
    # Check interface state when interface is up
    #check_interface_status_up(module)

    # Check for errors in log file
    verify_log_dumps(module)

    HASH_DICT['result.detail'] = failure_summary


def main():
    """ This section is for arguments parsing """
    module = AnsibleModule(
        argument_spec=dict(
            switch_name=dict(required=False, type='str'),
            eth=dict(required=False, type='str'),
            start_time=dict(required=False, type='str'),
            log_file_names=dict(required=False, type='str'),
            error_msg=dict(required=False, type='str'),
            hash_name=dict(required=False, type='str'),
            log_dir_path=dict(required=False, type='str'),
        )
    )

    global RESULT_STATUS, HASH_DICT

    verify_ifconfig_state(module)

    # Calculate the entire test result
    HASH_DICT['result.status'] = 'Passed' if RESULT_STATUS else 'Failed'

    # Create a log file
    log_file_path = module.params['log_dir_path']
    log_file_path += '/{}.log'.format(module.params['hash_name'])
    log_file = open(log_file_path, 'a')
    for key, value in HASH_DICT.iteritems():
        log_file.write(key)
        log_file.write('\n')
        log_file.write(str(value))
        log_file.write('\n')
        log_file.write('\n')

    log_file.close()

    # Exit the module and return the required JSON.
    module.exit_json(
        hash_dict=HASH_DICT,
        log_file_path=log_file_path
    )

if __name__ == '__main__':
    main()


