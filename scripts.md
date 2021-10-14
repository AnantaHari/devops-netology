1 урок - bash

    1. Есть скрипт:

    a=1
    b=2
    c=a+b
    d=$a+$b
    e=$(($a+$b))

        Какие значения переменным c,d,e будут присвоены?
        Почему?
        
```
c = a+b потому что переменным может быть присвоино целое число либо строка. Т.к. a+b не целое число, значит присвоилась строка.
d = 1+2 потому что переменная d не целочисленная, но т.к. перед a и b знак $, то с присвоилась строка значений a и b.
e = 3 потому что благодаря $(( )) происходит присвоение целочисленной суммы a и b в e.
```

2.    На нашем локальном сервере упал сервис и мы написали скрипт, который постоянно проверяет его доступность, записывая дату проверок до тех пор, пока сервис не станет доступным. В скрипте допущена ошибка, из-за которой выполнение не может завершиться, при этом место на Жёстком Диске постоянно уменьшается. Что необходимо сделать, чтобы его исправить:

    while ((1==1)
    do
    curl https://localhost:4757
    if (($? != 0))
    then
    date >> curl.log
    fi
    done
```
Т.к. в условии указано while ((1==1), то скрипт будет выполняться всегда. Нужно после date добавить break.
Оператор >> дописывает новые данные к уже существующим. Если нам этого не надо, то нужно использовать оператор > который удалит содержимое файла и запишет новым.
```

 3.   Необходимо написать скрипт, который проверяет доступность трёх IP: 192.168.0.1, 173.194.222.113, 87.250.250.242 по 80 порту и записывает результат в файл log. Проверять доступность необходимо пять раз для каждого узла.
```
#!/usr/bin/env bash
declare -i a
a=0
while (($a<5))
do 
curl 192.168.0.1:80
echo $? >> curl.log
curl 173.194.222.113:80
echo $? >> curl.log
curl 87.250.250.242:80
echo $? >> curl.log
let "a +=1"
done
```

4.    Необходимо дописать скрипт из предыдущего задания так, чтобы он выполнялся до тех пор, пока один из узлов не окажется недоступным. Если любой из узлов недоступен - IP этого узла пишется в файл error, скрипт прерывается
```
#!/usr/bin/env bash
ip_array=(192.168.0.1:80 173.194.222.113:80 87.250.250.242:80)
while ((1==1))
  do
  for i in ${ip_array[@]}
    do
    curl $i
    if (($? != 0))
      then
      echo $i > error
      break 2
    fi
  done
done

```
Дополнительное задание (со звездочкой*) - необязательно к выполнению

Мы хотим, чтобы у нас были красивые сообщения для коммитов в репозиторий. Для этого нужно написать локальный хук для git, который будет проверять, что сообщение в коммите содержит код текущего задания в квадратных скобках и количество символов в сообщении не превышает 30. Пример сообщения: [04-script-01-bash] сломал хук.
```
#!/usr/bin/env bash

while read line
    do
    # Skip comments
    if [ "${line:0:1}" == "#" ]
    then
        continue
    fi
    if [ ${#line} -ge 30 ]
    then
        echo "Комментарий коммита ограничен 30 символами."
        echo "Текущий комментарий коммита содержит ${#line} символов."
        echo "${line}"
        exit 1
    fi
    if [[ "${{line:0:1}" != *[* || "${{line:0:1}" != *]* ]]
    then
        echo "Коммит не содержит кода текущего задания"
        exit 1
    fi
done < "${1}"

exit 0

код другого студента:
#!/bin/bash
number=30
len=$(cat $1|wc -m)
if [[ $len -gt $number ]]
  then
  echo "Количество символов не должно превышать 30"
  exit 1
elif ! grep -qE '^(04-script-[0-9][0-9]-[a-z])' $1;
  then
  echo "Наименование коммита должно быть в формате [04-script-XX-XX]"
  exit 1
fi

```


2 урок - python
1.	Есть скрипт:
#!/usr/bin/env python3
a = 1
b = '2'
c = a + b
•	Какое значение будет присвоено переменной c? 
```
Будет ошибка
Traceback (most recent call last):
File "<stdin>", line 1, in <module>
TypeError: unsupported operand type(s) for +: 'int' and 'str'
```

•	Как получить для переменной c значение 12? 
```
        >>> c = str(a) + b
		>>> c
		'12'
```
•	Как получить для переменной c значение 3? 
```
        >>> c = a + int(b)
		>>> c
		3
```

2.	Мы устроились на работу в компанию, где раньше уже был DevOps Engineer. Он написал скрипт, позволяющий узнать, какие файлы модифицированы в репозитории, относительно локальных изменений. Этим скриптом недовольно начальство, потому что в его выводе есть не все изменённые файлы, а также непонятен полный путь к директории, где они находятся. Как можно доработать скрипт ниже, чтобы он исполнял требования вашего руководителя?
#!/usr/bin/env python3

import os

bash_command = ["cd ~/netology/sysadm-homeworks", "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
is_change = False
for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print(prepare_result)
        break
```
#!/usr/bin/env python3
import os
from os.path import expanduser
home = expanduser("~")
bash_command = ["cd ~/devops-netology/", "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
for result in result_os.split('\n'):
    if result.find('modified') != -1:
        prepare_result = result.replace('\tmodified:   ', '')
        print(f'{home}/devops-netology/{prepare_result}')
```

3.	Доработать скрипт выше так, чтобы он мог проверять не только локальный репозиторий в текущей директории, а также умел воспринимать путь к репозиторию, который мы передаём как входной параметр. Мы точно знаем, что начальство коварное и будет проверять работу этого скрипта в директориях, которые не являются локальными репозиториями.
```
#!/usr/bin/env python3
import os
import sys
import subprocess

git_dir = sys.argv[1]
bash_command = ["cd " + git_dir, "git status"]
os.chdir(git_dir)
with subprocess.Popen(['git', 'status'], stdout=subprocess.PIPE) as proc:
    result = proc.stdout.read().decode("utf-8")
if result.find('not') == -1:
    print('Данная директория не содержит репозитория!')
else:
  result_os =  os.popen(' && '.join(bash_command)).read()
  for result in result_os.split('\n'):
      if result.find('modified') != -1:
          prepare_result = result.replace('\tmodified:   ', '')
          print(f'{git_dir}{prepare_result}')
```
4.	Наша команда разрабатывает несколько веб-сервисов, доступных по http. Мы точно знаем, что на их стенде нет никакой балансировки, кластеризации, за DNS прячется конкретный IP сервера, где установлен сервис. Проблема в том, что отдел, занимающийся нашей инфраструктурой очень часто меняет нам сервера, поэтому IP меняются примерно раз в неделю, при этом сервисы сохраняют за собой DNS имена. Это бы совсем никого не беспокоило, если бы несколько раз сервера не уезжали в такой сегмент сети нашей компании, который недоступен для разработчиков. Мы хотим написать скрипт, который опрашивает веб-сервисы, получает их IP, выводит информацию в стандартный вывод в виде: <URL сервиса> - <его IP>. Также, должна быть реализована возможность проверки текущего IP сервиса c его IP из предыдущей проверки. Если проверка будет провалена - оповестить об этом в стандартный вывод сообщением: [ERROR] <URL сервиса> IP mismatch: <старый IP> <Новый IP>. Будем считать, что наша разработка реализовала сервисы: drive.google.com, mail.google.com, google.com.
```
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

```
```
log_list = {}

не очень хорошо в названии использовать list
в названии сказано что это список, а на деле словарь :)
if log_host.find(host_ip) == -1: - лучше сохранять ip и проверять на равенство
```

import subprocess  
import socket  
socket.gethostbyname('google.com')  
'64.233.164.101'  
socket.gethostbyname_ex('google.com')  
('google.com', [], ['64.233.164.100', '64.233.164.139', '64.233.164.138', '64.233.164.113', '64.233.164.102', '64.233.164.101'])  


