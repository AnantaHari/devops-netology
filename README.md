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
