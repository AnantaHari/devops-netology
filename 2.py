#!/usr/bin/env python3

import socket
log_list = {}
host_list = ('drive.google.com', 'mail.google.com', 'google.com')
with open('2.log', 'r') as log:
    data = log.read()
for line in data.splitlines():
    (key, val) = line.split(' - ')
    log_list[key] = val
open('2.log', 'w').close()
for host in host_list:
    host_ip = socket.gethostbyname(host)
    log_host = log_list.get('<' + host + '>')
    if log_host.find(host_ip) == -1:
        print(f'[ERROR] <{host}> IP mismatch: {log_host} <{host_ip}>')
    result = '<' + host + '>' + ' - ' + '<' + host_ip + '>'
    print(result)
    with open('2.log', 'a') as log:
        log.write(result + '\n')
