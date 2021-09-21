1 урок

1. Работа c HTTP через телнет.
Подключитесь утилитой телнет к сайту stackoverflow.com telnet stackoverflow.com 80
отправьте HTTP запрос
GET /questions HTTP/1.0
HOST: stackoverflow.com
[press enter]
[press enter]
В ответе укажите полученный HTTP код, что он означает?
```
HTTP/1.1 301 Moved Permanently
cache-control: no-cache, no-store, must-revalidate
location: https://stackoverflow.com/questions
x-request-guid: da5ad868-168f-4fee-b880-b2bb1bbed687
feature-policy: microphone 'none'; speaker 'none'
content-security-policy: upgrade-insecure-requests; frame-ancestors 'self' https://stackexchange.com
Accept-Ranges: bytes
Date: Wed, 08 Sep 2021 18:16:55 GMT
Via: 1.1 varnish
Connection: close
X-Served-By: cache-hhn4042-HHN
X-Cache: MISS
X-Cache-Hits: 0
X-Timer: S1631125015.356543,VS0,VE85
Vary: Fastly-SSL
X-DNS-Prefetch-Control: off
Set-Cookie: prov=7c7d8c9f-380e-084c-dd67-b180b03306ea; domain=.stackoverflow.com; expires=Fri, 01-Jan-2055 00:00:00 GMT; path=/; HttpOnly

Получаем ответ сервера. Основное здесь это код 301 - это код постоянного редиректа на другой адрес
```

2. Повторите задание 1 в браузере, используя консоль разработчика F12.
откройте вкладку Network
отправьте запрос http://stackoverflow.com
найдите первый ответ HTTP сервера, откройте вкладку Headers
укажите в ответе полученный HTTP код.
проверьте время загрузки страницы, какой запрос обрабатывался дольше всего?
приложите скриншот консоли браузера в ответ.
```
Status Code: 301 Moved Permanently

![image](https://user-images.githubusercontent.com/87232557/132569492-c01345e6-687f-4e2b-8f63-5d4da3228901.png)

```
3. Какой IP адрес у вас в интернете?
```
Мой IP: 178.176.134.35
![image](https://user-images.githubusercontent.com/87232557/132570078-25697999-f45f-4c4c-a5d4-af207c2f893c.png)
```
4. Какому провайдеру принадлежит ваш IP адрес? Какой автономной системе AS? Воспользуйтесь утилитой whois
```
Провайдер: FCURP of PJSC MegaFon
AS: AS31133
```
5. Через какие сети проходит пакет, отправленный с вашего компьютера на адрес 8.8.8.8? Через какие AS? Воспользуйтесь утилитой traceroute
```
anantahari@ubuntu:~$ traceroute -An 8.8.8.8
traceroute: invalid option -- 'A'
Try 'traceroute --help' or 'traceroute --usage' for more information.
anantahari@ubuntu:~$ traceroute 8.8.8.8
traceroute to 8.8.8.8 (8.8.8.8), 64 hops max
  1   192.168.108.1  2.322ms  5.464ms  2.639ms 
  2   10.128.0.1  4.036ms  2.648ms  6.645ms 
  3   78.25.64.29  7.654ms  4.644ms  5.325ms 
  4   10.222.33.13  10.669ms  10.290ms  10.068ms 
  5   10.222.77.66  11.472ms  10.765ms  11.246ms 
  6   83.169.204.113  9.312ms  9.479ms  9.146ms 
  7   178.176.152.61  9.343ms  9.175ms  8.587ms 
  8   108.170.250.66  10.861ms  10.214ms  11.765ms 
  9   142.251.49.158  22.198ms  22.336ms  24.236ms 
 10   172.253.65.159  24.703ms  25.791ms  24.597ms 
 11   142.250.209.25  26.450ms  26.020ms  29.898ms 
 12   *  *  * 
 13   *  *  * 
 14   *  *  * 
 15   *  *  * 
 16   *  *  * 
 17   *  *  * 
 18   *  *  * 
 19   *  *  * 
 20   *  *  * 
 21   8.8.8.8  28.630ms  25.776ms  25.439ms 

traceroute -An - в Ubuntu не работает, нет такого параметра
```
6. Повторите задание 5 в утилите mtr. На каком участке наибольшая задержка - delay?
```
![image](https://user-images.githubusercontent.com/87232557/132626011-5efc93ff-729b-424a-9c9d-4fc1ecfb7dee.png)

На AS15169 ip 142.251.49.158 в данный момент самая большая задержка
```
7. Какие DNS сервера отвечают за доменное имя dns.google? Какие A записи? воспользуйтесь утилитой dig
```
dns.google.		10800	IN	NS	ns1.zdns.google.
dns.google.		10800	IN	NS	ns4.zdns.google.
dns.google.		10800	IN	NS	ns2.zdns.google.
dns.google.		10800	IN	NS	ns3.zdns.google.
Эти DNS сервера отвечают за доменное имя dns.google

dns.google.		900	IN	A	8.8.8.8
dns.google.		900	IN	A	8.8.4.4
Это A записи
```
8. Проверьте PTR записи для IP адресов из задания 7. Какое доменное имя привязано к IP? воспользуйтесь утилитой dig
```
anantahari@ubuntu:~$ dig -x 8.8.8.8

; <<>> DiG 9.16.8-Ubuntu <<>> -x 8.8.8.8
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 15524
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;8.8.8.8.in-addr.arpa.		IN	PTR

;; ANSWER SECTION:
8.8.8.8.in-addr.arpa.	3408	IN	PTR	dns.google.

;; Query time: 16 msec
;; SERVER: 127.0.0.53#53(127.0.0.53)
;; WHEN: Wed Sep 08 22:35:40 PDT 2021
;; MSG SIZE  rcvd: 73

nantahari@ubuntu:~$ dig -x 8.8.4.4

; <<>> DiG 9.16.8-Ubuntu <<>> -x 8.8.4.4
;; global options: +cmd
;; Got answer:
;; ->>HEADER<<- opcode: QUERY, status: NOERROR, id: 26647
;; flags: qr rd ra; QUERY: 1, ANSWER: 1, AUTHORITY: 0, ADDITIONAL: 1

;; OPT PSEUDOSECTION:
; EDNS: version: 0, flags:; udp: 65494
;; QUESTION SECTION:
;4.4.8.8.in-addr.arpa.		IN	PTR

;; ANSWER SECTION:
4.4.8.8.in-addr.arpa.	5396	IN	PTR	dns.google.

;; Query time: 8 msec
;; SERVER: 127.0.0.53#53(127.0.0.53)
;; WHEN: Wed Sep 08 22:39:01 PDT 2021
;; MSG SIZE  rcvd: 73

Для 8.8.8.8 и 8.8.4.4 доменное имя dns.google.


```

2 урок
1. Проверьте список доступных сетевых интерфейсов на вашем компьютере. Какие команды есть для этого в Linux и в Windows?
```
Список интерфейсов и их адреса в Линукс:
anantahari@ubuntu:~$ ip -c -br link
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
Для сети из 40-50 хостов можно использовать сеть и маску указанную ниже:
anantahari@ubuntu:~$ ipcalc 100.64.1.1/26
Address:   100.64.1.1           01100100.01000000.00000001.00 000001
Netmask:   255.255.255.192 = 26 11111111.11111111.11111111.11 000000
Wildcard:  0.0.0.63             00000000.00000000.00000000.00 111111
=>
Network:   100.64.1.0/26        01100100.01000000.00000001.00 000000
HostMin:   100.64.1.1           01100100.01000000.00000001.00 000001
HostMax:   100.64.1.62          01100100.01000000.00000001.00 111110
Broadcast: 100.64.1.63          01100100.01000000.00000001.00 111111
Hosts/Net: 62                    Class A
```

7. Как проверить ARP таблицу в Linux, Windows? Как очистить ARP кеш полностью? Как из ARP таблицы удалить только один нужный IP?
```
В Windows:
C:\Users\kveta>arp -a

Интерфейс: 10.12.124.125 --- 0x8
  адрес в Интернете      Физический адрес      Тип
  10.12.124.1           86-09-ca-de-65-27     динамический
  10.12.127.255         ff-ff-ff-ff-ff-ff     статический
  224.0.0.2             01-00-5e-00-00-02     статический
  224.0.0.22            01-00-5e-00-00-16     статический
  224.0.0.251           01-00-5e-00-00-fb     статический
  224.0.0.252           01-00-5e-00-00-fc     статический
  239.255.102.18        01-00-5e-7f-66-12     статический
  239.255.255.250       01-00-5e-7f-ff-fa     статический
  255.255.255.255       ff-ff-ff-ff-ff-ff     статический
  
  Удалить один ip - arp -d 192.168.1.1
  Очистить ARP кеш полностью: 
  Пуск – Выполнить – ввести нижеприведенную команду и нажать Ок
  netsh interface ip delete arpcache
  
В Linux:
anantahari@ubuntu:~$ ip neigh
192.168.23.2 dev ens33 lladdr 00:50:56:e8:60:9b STALE
192.168.23.254 dev ens33 lladdr 00:50:56:fd:39:d5 STALE

Удалить один ip - sudo ip neigh del 192.168.11.100 dev ens33
чистить ARP кеш полностью:
sudo ip neigh flush all
```
