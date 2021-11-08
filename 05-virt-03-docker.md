Задача 1  
Сценарий выполения задачи:  

создайте свой репозиторий на https://hub.docker.com;  
выберете любой образ, который содержит веб-сервер Nginx;  
создайте свой fork образа;  
реализуйте функциональность: запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже:  
```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на https://hub.docker.com/username_repo.  
```
https://hub.docker.com/r/kvetalex/nginx-new-index
docker pull kvetalex/nginx-new-index
http://172.17.0.2/
```

Задача 2  
Посмотрите на сценарий ниже и ответьте на вопрос: "Подходит ли в этом сценарии использование Docker контейнеров или лучше подойдет виртуальная машина, физическая машина? Может быть возможны разные варианты?"  

Детально опишите и обоснуйте свой выбор.  

--  

Сценарий:  

Высоконагруженное монолитное java веб-приложение;  
```
Физическая или виртуальная машина с выделенными ресурсами чтобы ничем другим не были заняты ресурсы.
```
Nodejs веб-приложение;  
```
Docker контейнер, т.к. не надо много ресурсов и удобно обновлять
```
Мобильное приложение c версиями для Android и iOS;  
```
Не понятно зачему тут Docker и маины. Это же отдельный файл, который публикуется в магазине приложений. В случае с Андроид - это apk-файл, в iOS свой файл.
```
Шина данных на базе Apache Kafka;  
```
Можно Docker - один раз настроил, дал доступы к портам и все работает.
```
Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana; 
```
Виртуальные машины - нужно много ресурсов и настройка их кластера
```
Мониторинг-стек на базе Prometheus и Grafana; 
```
Можно в Docker, особенно если ставить заказчику. Если себе, то можно и на виртуалку.
```
MongoDB, как основное хранилище данных для java-приложения;  
```
Виртуалка и физическая машина. Даже лучше виртуалка чтобы удобнее было добавлять ресурсы.
```
Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.  
```
Виртуалка, т.к. надо много пространства для хранения контейнеров.
```

Задача 3
Запустите первый контейнер из образа centos c любым тэгом в фоновом режиме, подключив папку /data из текущей рабочей директории на хостовой машине в /data контейнера;  
Запустите второй контейнер из образа debian в фоновом режиме, подключив папку /data из текущей рабочей директории на хостовой машине в /data контейнера;  
Подключитесь к первому контейнеру с помощью docker exec и создайте текстовый файл любого содержания в /data;  
Добавьте еще один файл в папку /data на хостовой машине;  
Подключитесь во второй контейнер и отобразите листинг и содержание файлов в /data контейнера.  
```
[root@20c1dfd9f897 home]# vi deb
[root@20c1dfd9f897 home]# ls
1  11  12  3  deb  kasm-user  pam.conf

root@aaa26d0b1c6c:/home# ls
1  11  12  3  deb  kasm-user  pam.conf

anantahari@ubuntu:~/data$ ll
total 20
drwxrwxr-x  3 anantahari anantahari 4096 ноя  8 16:53 ./
drwxr-xr-x 26 anantahari anantahari 4096 ноя  8 16:49 ../
-rw-rw-r--  1 anantahari anantahari    0 ноя  8 16:23 1
-rw-rw-r--  1 anantahari anantahari    0 ноя  8 16:49 11
-rw-r--r--  1 root       root          4 ноя  8 16:49 12
-rw-rw-r--  1 anantahari anantahari    0 ноя  8 16:27 3
-rw-r--r--  1 root       root          0 ноя  8 16:53 deb
drwxr-xr-x  2 root       root       4096 ноя  8 16:29 kasm-user/
-rw-r--r--  1 root       root        552 ноя  8 16:52 pam.conf
```

Задача 4 (*)
Воспроизвести практическую часть лекции самостоятельно.

Соберите Docker образ с Ansible, загрузите на Docker Hub и пришлите ссылку вместе с остальными ответами к задачам.
```
```