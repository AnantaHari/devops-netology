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
<pre>
root@vagrant:/# nsenter --target 1581 --mount --uts --ipc --net --pid ps aux
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.7  1.1 101692 11244 ?        Ss   14:52   0:10 /sbin/init
root           2  0.0  0.0      0     0 ?        S    14:52   0:00 [kthreadd]
root           3  0.0  0.0      0     0 ?        I<   14:52   0:00 [rcu_gp]
root           4  0.0  0.0      0     0 ?        I<   14:52   0:00 [rcu_par_gp]
root           5  0.0  0.0      0     0 ?        I    14:52   0:00 [kworker/0:0-events]
root           6  0.0  0.0      0     0 ?        I<   14:52   0:00 [kworker/0:0H-kblockd]
root           7  0.0  0.0      0     0 ?        I    14:52   0:00 [kworker/u4:0-events_power_efficient]
root           8  0.0  0.0      0     0 ?        I<   14:52   0:00 [mm_percpu_wq]
root           9  0.0  0.0      0     0 ?        S    14:52   0:00 [ksoftirqd/0]
root          10  0.1  0.0      0     0 ?        I    14:52   0:01 [rcu_sched]
root          11  0.0  0.0      0     0 ?        S    14:52   0:00 [migration/0]
root          12  0.0  0.0      0     0 ?        S    14:52   0:00 [idle_inject/0]
root          14  0.0  0.0      0     0 ?        S    14:52   0:00 [cpuhp/0]
root          15  0.0  0.0      0     0 ?        S    14:52   0:00 [cpuhp/1]
root          16  0.0  0.0      0     0 ?        S    14:52   0:00 [idle_inject/1]
root          17  0.0  0.0      0     0 ?        S    14:52   0:00 [migration/1]
root          18  0.0  0.0      0     0 ?        S    14:52   0:01 [ksoftirqd/1]
root          20  0.0  0.0      0     0 ?        I<   14:52   0:00 [kworker/1:0H-kblockd]
root          21  0.0  0.0      0     0 ?        S    14:52   0:00 [kdevtmpfs]
root          22  0.0  0.0      0     0 ?        I<   14:52   0:00 [netns]
root          23  0.0  0.0      0     0 ?        S    14:52   0:00 [rcu_tasks_kthre]
root          24  0.0  0.0      0     0 ?        S    14:52   0:00 [kauditd]
root          25  0.0  0.0      0     0 ?        S    14:52   0:00 [khungtaskd]
root          26  0.0  0.0      0     0 ?        S    14:52   0:00 [oom_reaper]
root          27  0.0  0.0      0     0 ?        I<   14:52   0:00 [writeback]
root          28  0.0  0.0      0     0 ?        S    14:52   0:00 [kcompactd0]
root          29  0.0  0.0      0     0 ?        SN   14:52   0:00 [ksmd]
root          30  0.0  0.0      0     0 ?        SN   14:52   0:00 [khugepaged]
root          77  0.0  0.0      0     0 ?        I<   14:52   0:00 [kintegrityd]
root          78  0.0  0.0      0     0 ?        I<   14:52   0:00 [kblockd]
root          79  0.0  0.0      0     0 ?        I<   14:52   0:00 [blkcg_punt_bio]
root          80  0.0  0.0      0     0 ?        I<   14:52   0:00 [tpm_dev_wq]
root          81  0.0  0.0      0     0 ?        I<   14:52   0:00 [ata_sff]
root          82  0.0  0.0      0     0 ?        I<   14:52   0:00 [md]
root          83  0.0  0.0      0     0 ?        I<   14:52   0:00 [edac-poller]
root          84  0.0  0.0      0     0 ?        I<   14:52   0:00 [devfreq_wq]
root          85  0.0  0.0      0     0 ?        S    14:52   0:00 [watchdogd]
root          88  0.0  0.0      0     0 ?        S    14:52   0:00 [kswapd0]
root          89  0.0  0.0      0     0 ?        S    14:52   0:00 [ecryptfs-kthrea]
root          91  0.0  0.0      0     0 ?        I<   14:52   0:00 [kthrotld]
root          92  0.0  0.0      0     0 ?        I<   14:52   0:00 [acpi_thermal_pm]
root          93  0.0  0.0      0     0 ?        S    14:52   0:00 [scsi_eh_0]
root          94  0.0  0.0      0     0 ?        I<   14:52   0:00 [scsi_tmf_0]
root          95  0.0  0.0      0     0 ?        S    14:52   0:00 [scsi_eh_1]
root          96  0.0  0.0      0     0 ?        I<   14:52   0:00 [scsi_tmf_1]
root          98  0.0  0.0      0     0 ?        I<   14:52   0:00 [vfio-irqfd-clea]
root          99  0.0  0.0      0     0 ?        I<   14:52   0:00 [ipv6_addrconf]
root         109  0.0  0.0      0     0 ?        I<   14:52   0:00 [kstrp]
root         112  0.0  0.0      0     0 ?        I<   14:52   0:00 [kworker/u5:0]
root         125  0.0  0.0      0     0 ?        I<   14:52   0:00 [charger_manager]
root         159  0.0  0.0      0     0 ?        I    14:52   0:01 [kworker/0:2-events]
root         168  0.3  0.0      0     0 ?        I    14:52   0:04 [kworker/1:3-events]
root         170  0.0  0.0      0     0 ?        S    14:52   0:00 [scsi_eh_2]
root         172  0.0  0.0      0     0 ?        I<   14:52   0:00 [scsi_tmf_2]
root         180  0.0  0.0      0     0 ?        I<   14:52   0:00 [ttm_swap]
root         181  0.0  0.0      0     0 ?        I<   14:52   0:00 [kworker/0:1H-kblockd]
root         192  0.0  0.0      0     0 ?        I<   14:52   0:00 [kdmflush]
root         194  0.0  0.0      0     0 ?        I<   14:52   0:00 [kdmflush]
root         225  0.1  0.0      0     0 ?        I<   14:52   0:01 [kworker/1:1H-kblockd]
root         227  0.0  0.0      0     0 ?        I<   14:52   0:00 [raid5wq]
root         282  0.0  0.0      0     0 ?        S    14:52   0:00 [jbd2/dm-0-8]
root         283  0.0  0.0      0     0 ?        I<   14:52   0:00 [ext4-rsv-conver]
root         348  0.0  1.9  51468 19392 ?        S<s  14:52   0:01 /lib/systemd/systemd-journald
root         365  0.0  0.0      0     0 ?        I<   14:52   0:00 [rpciod]
root         366  0.0  0.0      0     0 ?        I<   14:52   0:00 [xprtiod]
root         379  0.0  0.5  21252  5480 ?        Ss   14:52   0:01 /lib/systemd/systemd-udevd
systemd+     391  0.0  0.7  26604  7584 ?        Ss   14:52   0:00 /lib/systemd/systemd-networkd
root         393  0.0  0.0      0     0 ?        I<   14:52   0:00 [iprt-VBoxWQueue]
root         426  0.0  0.0      0     0 ?        I<   14:52   0:00 [nfit]
root         513  0.0  0.0      0     0 ?        I<   14:53   0:00 [kaluad]
root         514  0.0  0.0      0     0 ?        I<   14:53   0:00 [kmpath_rdacd]
root         516  0.0  0.0      0     0 ?        I<   14:53   0:00 [kmpathd]
root         517  0.0  0.0      0     0 ?        I<   14:53   0:00 [kmpath_handlerd]
root         518  0.2  1.7 280136 17940 ?        SLsl 14:53   0:03 /sbin/multipathd -d -s
_rpc         545  0.0  0.3   7104  3932 ?        Ss   14:53   0:00 /sbin/rpcbind -f -w
systemd+     546  0.0  1.2  23892 12192 ?        Ss   14:53   0:00 /lib/systemd/systemd-resolved
root         568  0.0  0.7 238192  7496 ?        Ssl  14:53   0:00 /usr/lib/accountsservice/accounts-daemon
message+     569  0.0  0.4   7616  4700 ?        Ss   14:53   0:00 /usr/bin/dbus-daemon --system --address=systemd: --nofork -
root         578  0.0  0.3  81828  3876 ?        Ssl  14:53   0:00 /usr/sbin/irqbalance --foreground
root         581  0.0  1.8  31636 18084 ?        Ss   14:53   0:00 /usr/bin/python3 /usr/bin/networkd-dispatcher --run-startup
syslog       584  0.0  0.4 224348  4356 ?        Ssl  14:53   0:00 /usr/sbin/rsyslogd -n -iNONE
root         591  0.0  0.7  16856  7848 ?        Ss   14:53   0:00 /lib/systemd/systemd-logind
netdata      618  4.7  5.3 308736 54108 ?        Ssl  14:53   1:02 /usr/sbin/netdata -D
vagrant      619  0.0  1.2 715024 12912 ?        Ssl  14:53   0:00 /usr/local/bin/node_exporter
root         627  0.0  0.2   9412  2940 ?        Ss   14:53   0:00 /usr/sbin/cron -f
daemon       630  0.0  0.2   3792  2340 ?        Ss   14:53   0:00 /usr/sbin/atd -f
root         645  0.0  0.1   8428  1832 tty1     Ss+  14:53   0:00 /sbin/agetty -o -p -- \u --noclear tty1 linux
root         647  0.0  0.7  12176  7516 ?        Ss   14:53   0:00 sshd: /usr/sbin/sshd -D [listener] 0 of 10-100 startups
root         674  0.0  0.6 232716  7004 ?        Ssl  14:53   0:00 /usr/lib/policykit-1/polkitd --no-debug
netdata      684  0.0  0.8  32176  8096 ?        Sl   14:53   0:00 /usr/sbin/netdata --special-spawn-server
vagrant      870  0.0  0.9  18376  9760 ?        Ss   14:53   0:00 /lib/systemd/systemd --user
vagrant      889  0.0  0.3 103044  3272 ?        S    14:53   0:00 (sd-pam)
root         910  0.1  0.3 295412  3016 ?        Sl   14:53   0:01 /usr/sbin/VBoxService --pidfile /var/run/vboxadd-service.sh
netdata     1000  0.2  0.3   4028  3244 ?        S    14:53   0:02 bash /usr/libexec/netdata/plugins.d/tc-qos-helper.sh 1
root        1003  0.0  0.8  13924  8928 ?        Ss   14:53   0:00 sshd: vagrant [priv]
root        1015  0.9  0.3 519404  3132 ?        Sl   14:53   0:13 /usr/libexec/netdata/plugins.d/ebpf.plugin 1
netdata     1018  1.4  2.1 723656 21424 ?        Sl   14:53   0:19 /usr/libexec/netdata/plugins.d/go.d.plugin 1
netdata     1021  1.0  0.3  53820  3852 ?        S    14:53   0:13 /usr/libexec/netdata/plugins.d/apps.plugin 1
root        1022  0.0  0.0   5120   912 ?        S    14:53   0:00 /usr/libexec/netdata/plugins.d/nfacct.plugin 1
root        1343  0.0  0.0      0     0 ?        I<   14:53   0:00 [cfg80211]
vagrant     1464  0.0  0.6  14056  6104 ?        S    14:53   0:00 sshd: vagrant@pts/0
vagrant     1465  0.0  0.4   9836  4288 pts/0    Ss   14:53   0:00 -bash
root        1544  0.0  0.0      0     0 ?        I    15:04   0:00 [kworker/u4:2-events_unbound]
root        1566  0.2  0.0      0     0 ?        I    15:05   0:01 [kworker/1:0-events]
root        1567  0.1  0.0      0     0 ?        I    15:05   0:00 [kworker/u4:3-events_power_efficient]
root        1568  0.0  0.4  11864  4700 pts/0    S    15:06   0:00 sudo su
root        1569  0.0  0.4  10972  4384 pts/0    S    15:06   0:00 su
root        1570  0.0  0.3   9836  3960 pts/0    S    15:06   0:00 bash
<b>root        1581  0.0  0.0   8076   592 pts/0    S+   15:07   0:00 sleep 1h</b>
root        1582  0.0  0.9  13928  9072 ?        Ss   15:07   0:00 sshd: vagrant [priv]
vagrant     1622  0.1  0.6  14060  6160 ?        S    15:07   0:00 sshd: vagrant@pts/1
vagrant     1623  0.0  0.4   9836  4192 pts/1    Ss   15:07   0:00 -bash
root        1643  0.0  0.4  11860  4564 pts/1    S    15:08   0:00 sudo nsenter --target 1581 --pid --mount
root        1644  0.0  0.0   8996   696 pts/1    S    15:08   0:00 nsenter --target 1581 --pid --mount
root        1645  0.0  0.4   9836  4140 pts/1    S    15:08   0:00 -bash
root        1654  0.0  0.4  12276  4612 pts/1    S    15:08   0:00 sudo -i
root        1655  0.0  0.4   9836  4180 pts/1    S    15:08   0:00 -bash
root        1670  0.0  0.0   8996   636 pts/1    S    15:09   0:00 nsenter --target 1581 --pid --mount
root        1671  0.0  0.4   9836  4232 pts/1    S    15:09   0:00 -bash
root        1690  0.0  0.0   8996   632 pts/1    S    15:13   0:00 nsenter --target 1581 --pid --mount
root        1691  0.0  0.4   9836  4092 pts/1    S    15:13   0:00 -bash
root        1708  0.0  0.0   8996   692 pts/1    S+   15:15   0:00 nsenter --target 1581 --mount --uts --ipc --net --pid ps au
root        1709  0.0  0.3  11492  3504 pts/1    R+   15:15   0:00 ps aux
</pre>
```

7. Найдите информацию о том, что такое :(){ :|:& };:. Запустите эту команду в своей виртуальной машине Vagrant с Ubuntu 20.04 (это важно, поведение в других ОС не проверялось). Некоторое время все будет "плохо", после чего (минуты) – ОС должна стабилизироваться. Вызов dmesg расскажет, какой механизм помог автоматической стабилизации. Как настроен этот механизм по-умолчанию, и как изменить число процессов, которое можно создать в сессии?
