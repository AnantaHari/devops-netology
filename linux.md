1 урок
1. Ознакомьтесь с графическим интерфейсом VirtualBox, посмотрите как выглядит виртуальная машина, которую создал для вас Vagrant, какие аппаратные ресурсы ей выделены. Какие ресурсы выделены по-умолчанию?
```
Ей выделено 1Гб памяти, 2 ядра процессора, 64Гб HDD, video memory 4Гб
```

2. Ознакомьтесь с возможностями конфигурации VirtualBox через Vagrantfile: документация. Как добавить оперативной памяти или ресурсов процессора виртуальной машине?
```
config.vm.provider "virtualbox" do |v|
  v.memory = 2048
  v.cpus = 4
end
```


3. Ознакомиться с разделами man bash, почитать о настройках самого bash:
какой переменной можно задать длину журнала history, и на какой строчке manual это описывается?
```
HISTSIZE
Строчка 862
```
что делает директива ignoreboth в bash?
```
A  colon-separated  list of values controlling how commands are saved on the history list.  If the list of values includes ignorespace, lines which begin with a space character are not saved in  the  history list.  A value of ignoredups causes lines matching the previous history entry to not be saved.  A value of ignoreboth is shorthand for ignorespace and ignoredups.
Т.е. ignoreboth объединяет действие команд ignorespace и ignoredups (в историю не запишутся строки, которые начинаются с пробела и если в истории уже есть такая запись, то такая же записана не будет.
```

4. В каких сценариях использования применимы скобки {} и на какой строчке man bash это описано?
```
{ list; }
              list is simply executed in the current shell environment.  list must be terminated with  a  newline  or
              semicolon.  This is known as a group command.  The return status is the exit status of list.  Note that
              unlike the metacharacters ( and ), { and } are reserved words and must occur where a reserved  word  is
              permitted  to be recognized.  Since they do not cause a word break, they must be separated from list by
              whitespace or another shell metacharacter.
 Строка 257
{ список; }
Список просто выполняется в среде текущего командного интерпретатора.
Список должен завершаться переводом строки или точкой с запятой.
Эту команду называют командой группировки. Статусом возврата
является статус выхода списка.
```

5. Основываясь на предыдущем вопросе, как создать однократным вызовом touch 100000 файлов? А получилось ли создать 300000?
```
Для создания файлов - touch {000001..100000}.txt

Создать 300000 файлов не получилось - 
touch {000001..300000}.txt
-bash: /usr/bin/rm: Argument list too long
```

6. В man bash поищите по /\[\[. Что делает конструкция [[ -d /tmp ]]
```
Проверяет наличие директории tmp.
Конструкция [[ ... ]] более предпочтительна, нежели [ ... ], поскольку поможет избежать некоторых логических ошибок.
```

7. Основываясь на знаниях о просмотре текущих (например, PATH) и установке новых переменных; командах, которые мы рассматривали, добейтесь в выводе type -a bash в виртуальной машине наличия первым пунктом в списке:
bash is /tmp/new_path_directory/bash
bash is /usr/local/bin/bash
bash is /bin/bash
(прочие строки могут отличаться содержимым и порядком)

```
mkdir /tmp/new_path_directory/
cp /bin/bash /tmp/new_path_directory/
PATH=/tmp/new_path_directory/:$PATH
```

Чем отличается планирование команд с помощью batch и at?
 ```
  at      executes commands at a specified time. (выполняет команды в указанное время)
 batch   executes  commands  when  system load levels permit; in other words, when the load average drops below 1.5, or the
               value specified in the invocation of atd. (выполняет команды когда определенная загрузка системы, например ниже 1.5)
```

2 урок
1. Какого типа команда cd? Попробуйте объяснить, почему она именно такого типа; опишите ход своих мыслей, если считаете что она могла бы быть другого типа.
```
type -t cd
builtin - встроенная в оболочку программа. Она вызывется и выполняется в текущей оболочке.
Если сделать ее внешней, то она не будет менять каталог в текущей оболочке.
```

2. Какая альтернатива без pipe команде grep <some_string> <some_file> | wc -l? man grep поможет в ответе на этот вопрос. Ознакомьтесь с документом о других подобных некорректных вариантах использования pipe.
```
grep <some_string> <some_file> -c
```

3. Какой процесс с PID 1 является родителем для всех процессов в вашей виртуальной машине Ubuntu 20.04?
```
systemd(1)
```

4. Как будет выглядеть команда, которая перенаправит вывод stderr ls на другую сессию терминала?
```
Находясь в /dev/pts/0
cd qqq 2>/dev/pts/1
Тогда в /dev/pts/1 будет
-bash: cd: qqq: No such file or directory
```

5. Получится ли одновременно передать команде файл на stdin и вывести ее stdout в другой файл? Приведите работающий пример.
```
cat < b.txt > c.txt
```

6. Получится ли вывести находясь в графическом режиме данные из PTY в какой-либо из эмуляторов TTY? Сможете ли вы наблюдать выводимые данные?
```
Вывести сможем через перенаправление, но данные не будем наблюдать. Чтобы их увидеть надо перейти в эмулятор терминала с помощью Ctrl + Alt + F1..F6
```

7. Выполните команду bash 5>&1. К чему она приведет? Что будет, если вы выполните echo netology > /proc/$$/fd/5? Почему так происходит?
```
Первая команда создаст дескриптор c FD 5 и перенатправит его в stdout
вторая команда выведет netology в дескриптор "5", который был пернеаправлен в stdout и на экране увидим netology
```

8. Получится ли в качестве входного потока для pipe использовать только stderr команды, не потеряв при этом отображение stdout на pty? Напоминаем: по умолчанию через pipe передается только stdout команды слева от | на stdin команды справа. Это можно сделать, поменяв стандартные потоки местами через промежуточный новый дескриптор, который вы научились создавать в предыдущем вопросе.
```
ll \qqq 55>&2 2>&1 1>&55
```

9. Что выведет команда cat /proc/$$/environ? Как еще можно получить аналогичный по содержанию вывод?
```
Выведет переменные окружения.
Можно также вывести с помощью export или env.
```

10. Используя man, опишите что доступно по адресам /proc/<PID>/cmdline, /proc/<PID>/exe.
```
/proc/<PID>/cmdline
  Этот файл содержит полную командную строку запуска процесса, кроме тех процессов, что полностью ушли в своппинг, а также тех, что превратились в зомби. В этих двух случаях в файле ничего нет, то есть чтение этого файла вернет 0 символов. Аргументы командной строки в этом файле указаны как список строк, каждая из которых завешается нулевым символом, с добавочным нулевым байтом после последней строки.
/proc/<PID>/exe
  Под Linux 2.2 и 2.4 exe является символьной ссылкой, содержащей фактическое полное имя выполняемого файла. Символьная ссылка exe может использоваться обычным образом - при попытке открыть exe будет открыт исполняемый файл. Вы можете даже ввести /proc/[number]/exe чтобы запустить другую копию процесса такого же как и процесс с номером [число].
Под Linux 2.0 и в более ранних версиях exe является указателем на запущенный файл и является символьной ссылкой. Вызов readlink(2) на этот специальный файл exe под Linux 2.0 и более ранних версий возвращает строку формата:

[устройство]:индексный_дескриптор

Например, строка [0301]:1502 означает индексный дескриптор 1502 на устройстве со старшим номером устройства 03 (IDE, MFM и т. д.) и младшим номером устройства 01 (первый раздел на первом диске).

Для того, чтобы найти этот файл, может быть использована команда find(1) с опцией -inum.
```

11. Узнайте, какую наиболее старшую версию набора инструкций SSE поддерживает ваш процессор с помощью /proc/cpuinfo.
```
sse4_2
grep sse /proc/cpuinfo
```

12. При открытии нового окна терминала и vagrant ssh создается новая сессия и выделяется pty. Это можно подтвердить командой tty, которая упоминалась в лекции 3.2. Однако:

vagrant@netology1:~$ ssh localhost 'tty'
not a tty
Почитайте, почему так происходит, и как изменить поведение.
```
Пишут что при входе в shell удаленный хост предполагает, что соединение выполняется пользователем-человеком. Можно добавить флаг -t и тогда принудительно создается псевдотерминал.
```

13. Бывает, что есть необходимость переместить запущенный процесс из одной сессии в другую. Попробуйте сделать это, воспользовавшись reptyr. Например, так можно перенести в screen процесс, который вы запустили по ошибке в обычной SSH-сессии.
```
Делал по этой ссылке: https://github.com/nelhage/reptyr#readme
Только вместо указанной там команды выполнял sudo reptyr -T
```
  
14. sudo echo string > /root/new_file не даст выполнить перенаправление под обычным пользователем, так как перенаправлением занимается процесс shell'а, который запущен без sudo под вашим пользователем. Для решения данной проблемы можно использовать конструкцию echo string | sudo tee /root/new_file. Узнайте что делает команда tee и почему в отличие от sudo echo команда с sudo tee будет работать.
```
tee читает из input и пишет в output и в файл
Работает потому что команда запущена от sudo.
```

3 урок
1. Какой системный вызов делает команда cd? В прошлом ДЗ мы выяснили, что cd не является самостоятельной программой, это shell builtin, поэтому запустить strace непосредственно на cd не получится. Тем не менее, вы можете запустить strace на /bin/bash -c 'cd /tmp'. В этом случае вы увидите полный список системных вызовов, которые делает сам bash при старте. Вам нужно найти тот единственный, который относится именно к cd.
```
chdir("/tmp") 
```

2. Попробуйте использовать команду file на объекты разных типов на файловой системе. Например:
vagrant@netology1:~$ file /dev/tty
/dev/tty: character special (5/0)
vagrant@netology1:~$ file /dev/sda
/dev/sda: block special (8/0)
vagrant@netology1:~$ file /bin/bash
/bin/bash: ELF 64-bit LSB shared object, x86-64
Используя strace выясните, где находится база данных file на основании которой она делает свои догадки.
```
openat(AT_FDCWD, "/usr/share/misc/magic.mgc", O_RDONLY) = 3
```
3. Предположим, приложение пишет лог в текстовый файл. Этот файл оказался удален (deleted в lsof), однако возможности сигналом сказать приложению переоткрыть файлы или просто перезапустить приложение – нет. Так как приложение продолжает писать в удаленный файл, место на диске постепенно заканчивается. Основываясь на знаниях о перенаправлении потоков предложите способ обнуления открытого удаленного файла (чтобы освободить место на файловой системе).
```
lsof | grep do_not_delete_me 
python3   4259                      anantahari    3r      REG                8,3        5     656647 /home/anantahari/do_not_delete_me

echo '' >/proc/4259/fd/3
```

4. Занимают ли зомби-процессы какие-то ресурсы в ОС (CPU, RAM, IO)?
```
Зомби-процессы освобождает все свои ресурсы, но остается запись в таблице процессов
```

5. В iovisor BCC есть утилита opensnoop:
root@vagrant:~# dpkg -L bpfcc-tools | grep sbin/opensnoop
/usr/sbin/opensnoop-bpfcc
На какие файлы вы увидели вызовы группы open за первую секунду работы утилиты? Воспользуйтесь пакетом bpfcc-tools для Ubuntu 20.04. Дополнительные сведения по установке.
```
vagrant@vagrant:~$ sudo /usr/sbin/opensnoop-bpfcc
PID    COMM               FD ERR PATH
2172   tmux: server       10   0 /proc/2181/cmdline
577    irqbalance          6   0 /proc/interrupts
577    irqbalance          6   0 /proc/stat
577    irqbalance          6   0 /proc/irq/20/smp_affinity
577    irqbalance          6   0 /proc/irq/0/smp_affinity
577    irqbalance          6   0 /proc/irq/1/smp_affinity
577    irqbalance          6   0 /proc/irq/8/smp_affinity
577    irqbalance          6   0 /proc/irq/12/smp_affinity
577    irqbalance          6   0 /proc/irq/14/smp_affinity
577    irqbalance          6   0 /proc/irq/15/smp_affinity
2172   tmux: server       10   0 /proc/2181/cmdline
2172   tmux: server       10   0 /proc/2181/cmdline
2172   tmux: server       10   0 /proc/2181/cmdline
2172   tmux: server       10   0 /proc/2181/cmdline
1      systemd            22   0 /proc/5663/cgroup
2172   tmux: server       10   0 /proc/2181/cmdline
2172   tmux: server       10   0 /proc/2181/cmdline
2172   tmux: server       10   0 /proc/2181/cmdline
2172   tmux: server       10   0 /proc/2181/cmdline
2172   tmux: server       10   0 /proc/2181/cmdline
787    vminfo              5   0 /var/run/utmp
```

6. Какой системный вызов использует uname -a? Приведите цитату из man по этому системному вызову, где описывается альтернативное местоположение в /proc, где можно узнать версию ядра и релиз ОС.
```
uname()
Part of the utsname information is also accessible  via  /proc/sys/kernel/{ostype, hostname, osrelease, version, domainname}.
```

7. Чем отличается последовательность команд через ; и через && в bash? Например:
root@netology1:~# test -d /tmp/some_dir; echo Hi
Hi
root@netology1:~# test -d /tmp/some_dir && echo Hi
root@netology1:~#
Есть ли смысл использовать в bash &&, если применить set -e?
```
;  - позволяет запускать несколько команд за один раз, и выполнение команды происходит последовательно.
&& -  Оператор AND (&&) будет выполнять вторую команду только в том случае, если при выполнении первой команды SUCCEEDS, т.е. состояние выхода первой команды равно “0” — программа выполнена успешно. Эта команда очень полезна при проверке состояния выполнения последней команды.

Использовать set -e - не имеет смысла, т.к. при ошибке выполнение команд прекратиться. 
```

8. Из каких опций состоит режим bash set -euxo pipefail и почему его хорошо было бы использовать в сценариях?
```
-e - прекращает выполнение скрипта если команда завершилась ошибкой, выводит в stderr строку с ошибкой
-u - прекращает выполнение скрипта, если встретилась несуществующая переменная
-x - выводит выполняемые команды в stdout перед выполненинем
-o pipefail - прекращает выполнение скрипта, даже если одна из частей пайпа завершилась ошибкой

При таком запуске скрипт получается безопасным, происходит автоматическая обработка ошибок.
```

9. Используя -o stat для ps, определите, какой наиболее часто встречающийся статус у процессов в системе. В man ps ознакомьтесь (/PROCESS STATE CODES) что значат дополнительные к основной заглавной буквы статуса процессов. Его можно не учитывать при расчете (считать S, Ss или Ssl равнозначными).
```
Ss - процесс, ожидающий завершения, лидер сессии
R+ - процесс выполняется, фоновый процесс
```

4 урок
1. На лекции мы познакомились с node_exporter. В демонстрации его исполняемый файл запускался в background. Этого достаточно для демо, но не для настоящей production-системы, где процессы должны находиться под внешним управлением. Используя знания из лекции по systemd, создайте самостоятельно простой unit-файл для node_exporter:

поместите его в автозагрузку,
предусмотрите возможность добавления опций к запускаемому процессу через внешний файл (посмотрите, например, на systemctl cat cron),
удостоверьтесь, что с помощью systemctl процесс корректно стартует, завершается, а после перезагрузки автоматически поднимается.
```
vagrant@vagrant:/etc/systemd/system$ ps -e |grep node_exporter  
  12594 ?        00:00:00 node_exporter
vagrant@vagrant:/etc/systemd/system$ systemctl stop node_exporter
==== AUTHENTICATING FOR org.freedesktop.systemd1.manage-units ===
Authentication is required to stop 'node_exporter.service'.
Authenticating as: vagrant,,, (vagrant)
Password: 
==== AUTHENTICATION COMPLETE ===
vagrant@vagrant:/etc/systemd/system$ ps -e |grep node_exporter  
vagrant@vagrant:/etc/systemd/system$ systemctl start node_exporter
==== AUTHENTICATING FOR org.freedesktop.systemd1.manage-units ===
Authentication is required to start 'node_exporter.service'.
Authenticating as: vagrant,,, (vagrant)
Password: 
==== AUTHENTICATION COMPLETE ===
vagrant@vagrant:/etc/systemd/system$ ps -e |grep node_exporter  
  12707 ?        00:00:00 node_exporter

Сервис:
vagrant@vagrant:/etc/systemd/system$ cat /etc/systemd/system/node_exporter.service
[Unit]
Description=Node Exporter
Wants=network-online.target
After=network-online.target

[Service]
User=vagrant
Group=vagrant
Type=simple
ExecStart=/usr/local/bin/node_exporter
EnvironmentFile=/etc/default/node_exporter

[Install]
WantedBy=multi-user.target

Переменная окружения
vagrant@vagrant:/etc/default$ sudo cat /proc/12914/environ 
LANG=en_US.UTF-8LANGUAGE=en_US:PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/snap/binHOME=/home/vagrantLOGNAME=vagrantUSER=vagrantSHELL=/bin/bashINVOCATION_ID=2076e5eb3c64480985a09438e92520ebJOURNAL_STREAM=9:41650vagrant@vagrant:/etc/default$ 
```

2. Ознакомьтесь с опциями node_exporter и выводом /metrics по-умолчанию. Приведите несколько опций, которые вы бы выбрали для базового мониторинга хоста по CPU, памяти, диску и сети.
```
node_cpu_seconds_total{cpu="0",mode="idle"} 3251.81
node_cpu_seconds_total{cpu="0",mode="iowait"} 5.96
node_cpu_seconds_total{cpu="0",mode="irq"} 0
node_cpu_seconds_total{cpu="0",mode="nice"} 12.42
node_cpu_seconds_total{cpu="0",mode="softirq"} 7.5
node_cpu_seconds_total{cpu="0",mode="steal"} 0
node_cpu_seconds_total{cpu="0",mode="system"} 67.54
node_cpu_seconds_total{cpu="0",mode="user"} 6.39

# HELP node_disk_io_time_seconds_total Total seconds spent doing I/Os.
# TYPE node_disk_io_time_seconds_total counter
node_disk_io_time_seconds_total{device="dm-0"} 54.424
node_disk_io_time_seconds_total{device="dm-1"} 0.46
node_disk_io_time_seconds_total{device="sda"} 55.58
# HELP node_disk_read_bytes_total The total number of bytes read successfully.
# TYPE node_disk_read_bytes_total counter
node_disk_read_bytes_total{device="dm-0"} 3.5654144e+08
node_disk_read_bytes_total{device="dm-1"} 3.342336e+06
node_disk_read_bytes_total{device="sda"} 3.70340864e+08
# HELP node_disk_read_time_seconds_total The total number of seconds spent by all reads.
# TYPE node_disk_read_time_seconds_total counter
node_disk_read_time_seconds_total{device="dm-0"} 44.212
node_disk_read_time_seconds_total{device="dm-1"} 0.452
node_disk_read_time_seconds_total{device="sda"} 29.505

# HELP node_memory_MemAvailable_bytes Memory information field MemAvailable_bytes.
# TYPE node_memory_MemAvailable_bytes gauge
node_memory_MemAvailable_bytes 7.45787392e+08
# HELP node_memory_MemFree_bytes Memory information field MemFree_bytes.
# TYPE node_memory_MemFree_bytes gauge
node_memory_MemFree_bytes 3.44559616e+08
# HELP node_memory_MemTotal_bytes Memory information field MemTotal_bytes.
# TYPE node_memory_MemTotal_bytes gauge
node_memory_MemTotal_bytes 1.028694016e+09

# HELP node_network_receive_bytes_total Network device statistic receive_bytes.
# TYPE node_network_receive_bytes_total counter
node_network_receive_bytes_total{device="eth0"} 9.665855e+06
node_network_receive_bytes_total{device="lo"} 237886
# HELP node_network_receive_compressed_total Network device statistic receive_compressed.
# TYPE node_network_receive_compressed_total counter
node_network_receive_compressed_total{device="eth0"} 0
node_network_receive_compressed_total{device="lo"} 0
# HELP node_network_receive_drop_total Network device statistic receive_drop.
# TYPE node_network_receive_drop_total counter
node_network_receive_drop_total{device="eth0"} 0
node_network_receive_drop_total{device="lo"} 0
# HELP node_network_receive_errs_total Network device statistic receive_errs.
# TYPE node_network_receive_errs_total counter
node_network_receive_errs_total{device="eth0"} 0
node_network_receive_errs_total{device="lo"} 0
```

3. Установите в свою виртуальную машину Netdata. Воспользуйтесь готовыми пакетами для установки (sudo apt install -y netdata). После успешной установки:

в конфигурационном файле /etc/netdata/netdata.conf в секции [web] замените значение с localhost на bind to = 0.0.0.0,
добавьте в Vagrantfile проброс порта Netdata на свой локальный компьютер и сделайте vagrant reload:
config.vm.network "forwarded_port", guest: 19999, host: 19999
После успешной перезагрузки в браузере на своем ПК (не в виртуальной машине) вы должны суметь зайти на localhost:19999. Ознакомьтесь с метриками, которые по умолчанию собираются Netdata и с комментариями, которые даны к этим метрикам.
```
https://skr.sh/s9ibQghuCzv
```

4. Можно ли по выводу dmesg понять, осознает ли ОС, что загружена не на настоящем оборудовании, а на системе виртуализации?
```
vagrant@vagrant:~$ dmesg | grep -i virtual
[    0.000000] DMI: innotek GmbH VirtualBox/VirtualBox, BIOS VirtualBox 12/01/2006
[    0.015067] CPU MTRRs all blank - virtualized system.
[    0.121789] Booting paravirtualized kernel on KVM
[   23.419319] systemd[1]: Detected virtualization oracle.
```

5. Как настроен sysctl fs.nr_open на системе по-умолчанию? Узнайте, что означает этот параметр. Какой другой существующий лимит не позволит достичь такого числа (ulimit --help)?
```
vagrant@vagrant:~$ sysctl -n fs.nr_open
1048576
Лимит на кол-во открытых дескрипторов

ulimit -n	the maximum number of open file descriptors
```

6. Запустите любой долгоживущий процесс (не ls, который отработает мгновенно, а, например, sleep 1h) в отдельном неймспейсе процессов; покажите, что ваш процесс работает под PID 1 через nsenter. Для простоты работайте в данном задании под root (sudo -i). Под обычным пользователем требуются дополнительные опции (--map-root-user) и т.д.
```
В одном терминале выполнил:
root@vagrant:/# unshare -f --pid --mount-proc sleep 2m
А в другом:
root@vagrant:/# ps -e | grep sleep
  15478 pts/1    00:00:00 sleep
root@vagrant:/# nsenter --target 15478 --mount --uts --ipc --net --pid ps aux
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  0.0   8076   520 pts/1    S+   05:13   0:00 sleep 2m
root           2  0.0  0.3  11492  3324 pts/2    R+   05:13   0:00 ps aux
```

7. Найдите информацию о том, что такое :(){ :|:& };:. Запустите эту команду в своей виртуальной машине Vagrant с Ubuntu 20.04 (это важно, поведение в других ОС не проверялось). Некоторое время все будет "плохо", после чего (минуты) – ОС должна стабилизироваться. Вызов dmesg расскажет, какой механизм помог автоматической стабилизации. Как настроен этот механизм по-умолчанию, и как изменить число процессов, которое можно создать в сессии?
```
на bash'e запускает очень много шелов
это что-то типа fork-bomb

root@vagrant:/# dmesg | grep fork
[ 2170.377660] cgroup: fork rejected by pids controller in /user.slice/user-1000.slice/session-4.scope

limit -u	the maximum number of user processes
Например limit -u 10 - будет 10 процессов на пользователя
```
