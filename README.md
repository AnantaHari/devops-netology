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
