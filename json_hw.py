
import json
import socket
import yaml
log_list = {}
host_list = ('drive.google.com', 'mail.google.com', 'google.com')
js_dict = {}
log_host = ''
with open('2.log', 'r') as log:
    data = log.read()
for line in data.splitlines():
    (key, val) = line.split(' - ')
    log_list[key] = val
open('2.log', 'w').close()
for host in host_list:
    host_ip = socket.gethostbyname(host)
    log_host = log_list.get('<' + host + '>')
    if str(log_host).find(host_ip) == -1:
      print(f'[ERROR] <{host}> IP mismatch: {log_host} <{host_ip}>')
    result = '<' + host + '>' + ' - ' + '<' + host_ip + '>'
    print(result)
    with open('2.log', 'a') as log:
        log.write(result + '\n')
    js_dict[host] = host_ip
with open('2.json', 'w') as js:
    js.write(json.dumps(js_dict, indent=2))
with open('2.yaml', 'w') as ym:
    ym.write(yaml.dump(js_dict, explicit_start=True, explicit_end=True))