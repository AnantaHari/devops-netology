1. Подключитесь к публичному маршрутизатору в интернет. Найдите маршрут к вашему публичному IP

telnet route-views.routeviews.org
Username: rviews
show ip route x.x.x.x/32
show bgp x.x.x.x/32
```
route-views>show ip route 178.176.134.35
Routing entry for 178.176.134.0/23
  Known via "bgp 6447", distance 20, metric 0
  Tag 8283, type external
  Last update from 94.142.247.3 3w6d ago
  Routing Descriptor Blocks:
  * 94.142.247.3, from 94.142.247.3, 3w6d ago
      Route metric is 0, traffic share count is 1
      AS Hops 2
      Route tag 8283
      MPLS label: none

route-views>show bgp 178.176.134.35           
BGP routing table entry for 178.176.134.0/23, version 981072571
Paths: (24 available, best #18, table default)
  Not advertised to any peer
  Refresh Epoch 1
  53767 174 31133, (aggregated by 31133 10.222.253.97)
    162.251.163.2 from 162.251.163.2 (162.251.162.3)
      Origin IGP, localpref 100, valid, external, atomic-aggregate
      Community: 174:21101 174:22005 53767:5000
      path 7FE138FB7CF0 RPKI State not found
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 1
  3267 31133, (aggregated by 31133 10.222.253.197)
    194.85.40.15 from 194.85.40.15 (185.141.126.1)
      Origin IGP, metric 0, localpref 100, valid, external, atomic-aggregate
      path 7FE16F360730 RPKI State not found
      rx pathid: 0, tx pathid: 0
  Refresh Epoch 1
  20912 3257 1299 31133, (aggregated by 31133 10.222.253.97)
    212.66.96.126 from 212.66.96.126 (212.66.96.126)
      Origin IGP, localpref 100, valid, external, atomic-aggregate
      Community: 3257:8095 3257:30622 3257:50001 3257:53900 3257:53904 20912:65004
      path 7FE0AAC3E6C0 RPKI State not found
```

2. Создайте dummy0 интерфейс в Ubuntu. Добавьте несколько статических маршрутов. Проверьте таблицу маршрутизации.
```
auto eth0.1400
iface eth0.1400 inet static
        address 192.168.1.10
        netmask 255.255.255.0
        vlan_raw_device eth0
        post-up ip route add 192.168.1.0/24 via 192.168.1.1
        post-up ip route add 192.168.1.0/24 dev eth0.1400 src 192.168.1.10

auto dummy0
iface dummy0 inet static
        address 10.2.2.2/32
        pre-up ip link add dummy0 type dummy
        post-down ip link del dummy0
~                                                                               
~                                                                               
~                                                                               
~                                                                               
~                                                                               
~                                                                               
~                                                                               
~                                                                               
~                                                                               
~                                                                               
"/etc/network/interfaces" 13L, 379B                           7,31-38       All

oot@ubuntu:/home/anantahari# ip route show
default via 192.168.23.2 dev ens33 proto dhcp metric 100 
169.254.0.0/16 dev ens33 scope link metric 1000 
192.168.23.0/24 dev ens33 proto kernel scope link src 192.168.23.129 metric 100 
```

3. Проверьте открытые TCP порты в Ubuntu, какие протоколы и приложения используют эти порты? Приведите несколько примеров.
```
oot@ubuntu:/home/anantahari# ss -n
Netid State  Recv-Q Send-Q                                 Local Address:Port                       Peer Address:Port                   Process                 
u_str ESTAB  0      0                                 /run/user/1000/bus 82597                                 * 81880                                          
u_str ESTAB  0      0                        /run/dbus/system_bus_socket 66452                                 * 67189                                          
u_str ESTAB  0      0                                                  * 94278                                 * 94277                                          
u_str ESTAB  0      0                        /run/systemd/journal/stdout 67756                                 * 67755                                          
u_seq ESTAB  0      0                                                  * 101296                                * 101295                                         
u_str ESTAB  0      0                        /run/dbus/system_bus_socket 33827                                 * 33181                                          

root@ubuntu:/home/anantahari# ss -p
Netid  State   Recv-Q  Send-Q                                     Local Address:Port                                        Peer Address:Port                   Process                                                                         
u_str  ESTAB   0       0                                     /run/user/1000/bus 82597                                                  * 81880                   users:(("dbus-daemon",pid=5756,fd=88))                                         
u_str  ESTAB   0       0                            /run/dbus/system_bus_socket 66452                                                  * 67189                   users:(("dbus-daemon",pid=869,fd=31))                                          
u_str  ESTAB   0       0                                                      * 94278                                                  * 94277                   users:(("Web Content",pid=15683,fd=3))                                         
u_str  ESTAB   0       0                            /run/systemd/journal/stdout 67756                                                  * 67755                   users:(("systemd-journal",pid=8285,fd=78),("systemd",pid=1,fd=269))            
u_seq  ESTAB   0       0                                                      * 101296                                                 * 101295                  users:(("RDD Process",pid=16082,fd=22))                                        
u_str  ESTAB   0       0                            /run/dbus/system_bus_socket 33827                                                  * 33181                   users:(("dbus-daemon",pid=869,fd=17))                                          
u_str  ESTAB   0       0                            /run/systemd/journal/stdout 32507                                                  * 31712                   users:(("systemd-journal",pid=8285,fd=107),("systemd",pid=1,fd=299))           
u_str  ESTAB   0       0                                                      * 67802                                                  * 67803                   users:(("gvfsd-trash",pid=6011,fd=2),("gvfsd-trash",pid=6011,fd=1),("gvfsd-fuse",pid=5778,fd=2),("gvfsd-fuse",pid=5778,fd=1),("gvfsd",pid=5772,fd=2),("gvfsd",pid=5772,fd=1))
u_str  ESTAB   0       0                                                      * 94277                                                  * 94278                   users:(("firefox",pid=15592,fd=91))                                            
```

4. Проверьте используемые UDP сокеты в Ubuntu, какие протоколы и приложения используют эти порты?
```
root@ubuntu:/home/anantahari# ss -ua
State          Recv-Q         Send-Q                        Local Address:Port                     Peer Address:Port          Process        
UNCONN         0              0                                   0.0.0.0:36814                         0.0.0.0:*                            
UNCONN         0              0                             127.0.0.53%lo:domain                        0.0.0.0:*                            
ESTAB          0              0                      192.168.23.129%ens33:bootpc                 192.168.23.254:bootps                       
UNCONN         0              0                                   0.0.0.0:631                           0.0.0.0:*                            
UNCONN         0              0                                   0.0.0.0:mdns                          0.0.0.0:*                            
UNCONN         0              0                                      [::]:mdns                             [::]:*                            
UNCONN         0              0                                      [::]:56866                            [::]:*                            

```

5. Используя diagrams.net, создайте L3 диаграмму вашей домашней сети или любой другой сети, с которой вы работали. 
```

```

 ---
## Задание для самостоятельной отработки (необязательно к выполнению)

6*. Установите Nginx, настройте в режиме балансировщика TCP или UDP.

7*. Установите bird2, настройте динамический протокол маршрутизации RIP.

8*. Установите Netbox, создайте несколько IP префиксов, используя curl проверьте работу API.
