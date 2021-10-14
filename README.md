1. Мы выгрузили JSON, который получили через API запрос к нашему сервису:
{ "info" : "Sample JSON output from our service\t",
    "elements" :[
        { "name" : "first",
        "type" : "server",
        "ip" : 7175 
        },
        { "name" : "second",
        "type" : "proxy",
        "ip : 71.78.22.43
        }
    ]
}
Нужно найти и исправить все ошибки, которые допускает наш сервис
```
{ "info" : "Sample JSON output from our service\\t",
    "elements" :[
        { "name" : "first",
        "type" : "server",
        "ip" : 7175
        },
        { "name" : "second",
        "type" : "proxy",
        "ip" : "71.78.22.43"
        }
    ]
}

```

2.В прошлый рабочий день мы создавали скрипт, позволяющий опрашивать веб-сервисы и получать их IP. К уже реализованному функционалу нам нужно добавить возможность записи JSON и YAML файлов, описывающих наши сервисы. Формат записи JSON по одному сервису: { "имя сервиса" : "его IP"}. Формат записи YAML по одному сервису: - имя сервиса: его IP. Если в момент исполнения скрипта меняется IP у сервиса - он должен так же поменяться в yml и json файле.
```
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
```
