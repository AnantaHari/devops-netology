1. Проверьте список доступных сетевых интерфейсов на вашем компьютере. Какие команды есть для этого в Linux и в Windows?
```
Список интерфейсов и их адреса в Линукс:
nantahari@ubuntu:~$ ip -c -br link
lo               UNKNOWN        00:00:00:00:00:00 <LOOPBACK,UP,LOWER_UP> 
ens33            UP             00:0c:29:75:82:e0 <BROADCAST,MULTICAST,UP,LOWER_UP> 
anantahari@ubuntu:~$ ip -c -br address
lo               UNKNOWN        127.0.0.1/8 ::1/128 
ens33            UP             192.168.100.17/24 fe80::c261:a54a:914f:4a7e/64 

В Windows:
C:\Users\kveta>ipconfig

Настройка протокола IP для Windows


Адаптер беспроводной локальной сети Подключение по локальной сети* 1:

   Состояние среды. . . . . . . . : Среда передачи недоступна.
   DNS-суффикс подключения . . . . . :

Адаптер беспроводной локальной сети Подключение по локальной сети* 3:

   Состояние среды. . . . . . . . : Среда передачи недоступна.
   DNS-суффикс подключения . . . . . :

Адаптер Ethernet Ethernet 3:

   Состояние среды. . . . . . . . : Среда передачи недоступна.
   DNS-суффикс подключения . . . . . :

Адаптер Ethernet VMware Network Adapter VMnet1:

   DNS-суффикс подключения . . . . . :
   Локальный IPv6-адрес канала . . . : fe80::f1b2:b5fc:73d5:1214%16
   IPv4-адрес. . . . . . . . . . . . : 192.168.80.1
   Маска подсети . . . . . . . . . . : 255.255.255.0
   Основной шлюз. . . . . . . . . :

Адаптер Ethernet VMware Network Adapter VMnet8:

   DNS-суффикс подключения . . . . . :
   Локальный IPv6-адрес канала . . . : fe80::70b6:7aa6:7bf2:3e8e%17
   IPv4-адрес. . . . . . . . . . . . : 192.168.23.1
   Маска подсети . . . . . . . . . . : 255.255.255.0
   Основной шлюз. . . . . . . . . :

Адаптер беспроводной локальной сети Беспроводная сеть:

   DNS-суффикс подключения . . . . . :
   Локальный IPv6-адрес канала . . . : fe80::b57f:a3ec:8193:1dc5%8
   IPv4-адрес. . . . . . . . . . . . : 192.168.100.23
   Маска подсети . . . . . . . . . . : 255.255.255.0
   Основной шлюз. . . . . . . . . : 192.168.100.1
   
Еще можно посмотреть так:
anantahari@ubuntu:~$ ifconfig -a
ens33: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.100.17  netmask 255.255.255.0  broadcast 192.168.100.255
        inet6 fe80::c261:a54a:914f:4a7e  prefixlen 64  scopeid 0x20<link>
        ether 00:0c:29:75:82:e0  txqueuelen 1000  (Ethernet)
        RX packets 2806  bytes 1369323 (1.3 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 1899  bytes 549399 (549.3 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 761  bytes 77010 (77.0 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 761  bytes 77010 (77.0 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

```

2. Какой протокол используется для распознавания соседа по сетевому интерфейсу? Какой пакет и команды есть в Linux для этого?
```
Протокол LLDP
Используется пакет lldpd
Команда lldpctl
```

3. Какая технология используется для разделения L2 коммутатора на несколько виртуальных сетей? Какой пакет и команды есть в Linux для этого? Приведите пример конфига.
```
Технология VLAN позволяет разделить коммутатор на несколько виртуальных сетей.
Установка пакета производится командой apt install vlan
anantahari@ubuntu:~$ cat /etc/network/interfaces
auto eth0.1400
iface eth0.1400 inet static
        address 192.168.1.1
        netmask 255.255.255.0
        vlan_raw_device eth0
```

4. Какие типы агрегации интерфейсов есть в Linux? Какие опции есть для балансировки нагрузки? Приведите пример конфига.
```
Типы агрегации (LAG):
статический (на Cisco mode on);
динамический – LACP протокол (на Cisco mode active).
Сначала нужно установить пакет:
 apt install ifenslave
 
Потом выключить сетевые интерфейсы:
# ifdown eth0 (и другие)
# /etc/init.d/networking stop

Потом прописать например такие настройки для их объедщинения:
/etc/network/interfaces:

auto bond0

iface bond0 inet static
    address 10.31.1.5
    netmask 255.255.255.0
    network 10.31.1.0
    gateway 10.31.1.254
    bond-slaves eth0 eth1
    bond-mode active-backup
    bond-miimon 100
    bond-downdelay 200
    bond-updelay 200
```

5. Сколько IP адресов в сети с маской /29 ? Сколько /29 подсетей можно получить из сети с маской /24. Приведите несколько примеров /29 подсетей внутри сети 10.10.10.0/24.
```
В сети с маской 29 8 адресов.
Для хостов 6 адресов.
nantahari@ubuntu:~$ ipcalc 192.168.1.1/29
Address:   192.168.1.1          11000000.10101000.00000001.00000 001
Netmask:   255.255.255.248 = 29 11111111.11111111.11111111.11111 000
Wildcard:  0.0.0.7              00000000.00000000.00000000.00000 111
=>
Network:   192.168.1.0/29       11000000.10101000.00000001.00000 000
HostMin:   192.168.1.1          11000000.10101000.00000001.00000 001
HostMax:   192.168.1.6          11000000.10101000.00000001.00000 110
Broadcast: 192.168.1.7          11000000.10101000.00000001.00000 111
Hosts/Net: 6                     Class C, Private Internet

Можно получить 16 /29 подсетей из сети с маской /24.
Пример:
12. Requested size: 8 hosts
Netmask:   255.255.255.248 = 29 
Network:   192.168.1.176/29     
HostMin:   192.168.1.177        
HostMax:   192.168.1.182        
Broadcast: 192.168.1.183        
Hosts/Net: 6                     Class C, Private Internet

13. Requested size: 8 hosts
Netmask:   255.255.255.248 = 29 
Network:   192.168.1.192/29     
HostMin:   192.168.1.193        
HostMax:   192.168.1.198        
Broadcast: 192.168.1.199        
Hosts/Net: 6                     Class C, Private Internet

14. Requested size: 8 hosts
Netmask:   255.255.255.248 = 29 
Network:   192.168.1.208/29     
HostMin:   192.168.1.209        
HostMax:   192.168.1.214        
Broadcast: 192.168.1.215        
Hosts/Net: 6                     Class C, Private Internet
```

6. Задача: вас попросили организовать стык между 2-мя организациями. Диапазоны 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16 уже заняты. Из какой подсети допустимо взять частные IP адреса? Маску выберите из расчета максимум 40-50 хостов внутри подсети.
```
Можно взять из подсети 100.64.0.0 — 100.127.255.255 (маска подсети: 255.192.0.0 или /10) Carrier-Grade NAT.
```

7. Как проверить ARP таблицу в Linux, Windows? Как очистить ARP кеш полностью? Как из ARP таблицы удалить только один нужный IP?
```
```

Задание для самостоятельной отработки (необязательно к выполнению)
8*. Установите эмулятор EVE-ng.

Инструкция по установке - https://github.com/svmyasnikov/eve-ng

Выполните задания на lldp, vlan, bonding в эмуляторе EVE-ng.
