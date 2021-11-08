Задача 1
Опишите своими словами основные преимущества применения на практике IaaC паттернов.
Какой из принципов IaaC является основополагающим?
```
Быстрее можно развернуть инфраструктуру для разработки и тестирования и такая инфраструктура у всех, кто пользуется такими виртуалками будет иметь одинаковые настройки
и установленные приложения. Т.е. не получится такого что у кого-то функциональность разработанного приложения работает, а у другого сотрудника не работает.
Главный принцип это то, что каждый раз разворачивая инфрастуктуру с помощью кода, мы получим идетничный результат.
```
Задача 2
Чем Ansible выгодно отличается от других систем управление конфигурациями?
Какой, на ваш взгляд, метод работы систем конфигурации более надёжный push или pull?
```
В Ansible низкий порог вхождения, т.е. не нужно обладать большими навыками в разрабоке чтобы им пользоваться. Не требуется установка агента. И он очень популярен в компаниях.
По моему мнению pull, потому что сервер сам запрашивает конфигурацию после его разворачивания. В методе push может возникнуть ситуация когда сервер какое-то время будет сидеть 
и ждать когда же ему пришлют конфигурацию.
```
Задача 3
Установить на личный компьютер:

VirtualBox
Vagrant
Ansible
Приложить вывод команд установленных версий каждой из программ, оформленный в markdown.
```
anantahari@ubuntu:~/vagrant$ vagrant --version
Vagrant 2.2.18
anantahari@ubuntu:~/vagrant$ vboxmanage --version
6.1.26_Ubuntur145957
anantahari@ubuntu:~/vagrant$ ansible --version
ansible [core 2.11.6] 
  config file = /home/anantahari/vagrant/ansible.cfg
  configured module search path = ['/home/anantahari/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
  ansible python module location = /usr/lib/python3/dist-packages/ansible
  ansible collection location = /home/anantahari/.ansible/collections:/usr/share/ansible/collections
  executable location = /usr/bin/ansible
  python version = 3.9.7 (default, Sep 10 2021, 14:59:43) [GCC 11.2.0]
  jinja version = 2.11.3
  libyaml = True

```
Задача 4 (*)
Воспроизвести практическую часть лекции самостоятельно.

Создать виртуальную машину.
Зайти внутрь ВМ, убедиться, что Docker установлен с помощью команды
docker ps
```
root@server1:~# docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
root@server1:~# docker -v
Docker version 20.10.10, build b485636

```
