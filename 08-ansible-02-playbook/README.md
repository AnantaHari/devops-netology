# Домашнее задание к занятию "08.02 Работа с Playbook"

## Подготовка к выполнению
1. Создайте свой собственный (или используйте старый) публичный репозиторий на github с произвольным именем.
2. Скачайте [playbook](./playbook/) из репозитория с домашним заданием и перенесите его в свой репозиторий.
3. Подготовьте хосты в соотвтествии с группами из предподготовленного playbook. 
4. Скачайте дистрибутив [java](https://www.oracle.com/java/technologies/javase-jdk11-downloads.html) и положите его в директорию `playbook/files/`. 

## Основная часть
1. Приготовьте свой собственный inventory файл `prod.yml`.
```
---
elasticsearch:
  hosts:
    ubuntu:
      ansible_host: 62.84.124.71
      ansible_connection: ssh
      ansible_user: anantahari

```
2. Допишите playbook: нужно сделать ещё один play, который устанавливает и настраивает kibana.
```
- name: Install Kibana
  hosts: elasticsearch
  tasks:
    - name: Upload tar.gz Kibana from remote URL
      get_url:
        url: "https://artifacts.elastic.co/downloads/kibana/kibana-{{ kibana_version }}-linux-x86_64.tar.gz"
        dest: "/tmp/kibana-{{ kibana_version }}-linux-x86_64.tar.gz"
        mode: 0755
        timeout: 60
        force: true
        validate_certs: false
      register: get_kibana
      until: get_kibana is succeeded
      tags: kibana
    - name: Create directrory for Kibana
      become: true
      file:
        state: directory
        path: "{{ kibana_home }}"
      tags: kibana
    - name: Extract Kibana in the installation directory
      become: true
      unarchive:
        copy: false
        src: "/tmp/kibana-{{ kibana_version }}-linux-x86_64.tar.gz"
        dest: "{{ kibana_home }}"
        extra_opts: [--strip-components=1]
        creates: "{{ kibana_home }}/bin/kibana"
      tags:
        - kibana
    - name: Set environment Kibana
      become: true
      template:
        src: templates/kbn.sh.j2
        dest: /etc/profile.d/kbn.sh
      tags: kibana
```
3. При создании tasks рекомендую использовать модули: `get_url`, `template`, `unarchive`, `file`.
4. Tasks должны: скачать нужной версии дистрибутив, выполнить распаковку в выбранную директорию, сгенерировать конфигурацию с параметрами.
5. Запустите `ansible-lint site.yml` и исправьте ошибки, если они есть.
6. Попробуйте запустить playbook на этом окружении с флагом `--check`.
```
ansible-playbook -i inventory/prod.yml site.yml --check

PLAY [Install Java] *****************************************************************************************************************************************

TASK [Gathering Facts] **************************************************************************************************************************************
ok: [ubuntu]

TASK [Set facts for Java 11 vars] ***************************************************************************************************************************
ok: [ubuntu]

TASK [Upload .tar.gz file containing binaries from local storage] *******************************************************************************************
ok: [ubuntu]

TASK [Ensure installation dir exists] ***********************************************************************************************************************
ok: [ubuntu]

TASK [Extract java in the installation directory] ***********************************************************************************************************
skipping: [ubuntu]

TASK [Export environment variables] *************************************************************************************************************************
ok: [ubuntu]

PLAY [Install Elasticsearch] ********************************************************************************************************************************

TASK [Gathering Facts] **************************************************************************************************************************************
ok: [ubuntu]

TASK [Upload tar.gz Elasticsearch from remote URL] **********************************************************************************************************
changed: [ubuntu]

TASK [Create directrory for Elasticsearch] ******************************************************************************************************************
ok: [ubuntu]

TASK [Extract Elasticsearch in the installation directory] **************************************************************************************************
skipping: [ubuntu]

TASK [Set environment Elastic] ******************************************************************************************************************************
ok: [ubuntu]

PLAY [Install Kibana] ***************************************************************************************************************************************

TASK [Gathering Facts] **************************************************************************************************************************************
ok: [ubuntu]

TASK [Upload tar.gz Kibana from remote URL] *****************************************************************************************************************
changed: [ubuntu]

TASK [Create directrory for Kibana] *************************************************************************************************************************
ok: [ubuntu]

TASK [Extract Kibana in the installation directory] *********************************************************************************************************
skipping: [ubuntu]

TASK [Set environment Kibana] *******************************************************************************************************************************
ok: [ubuntu]

PLAY RECAP **************************************************************************************************************************************************
ubuntu                     : ok=13   changed=2    unreachable=0    failed=0    skipped=3    rescued=0    ignored=0   
```
7. Запустите playbook на `prod.yml` окружении с флагом `--diff`. Убедитесь, что изменения на системе произведены.
8. Повторно запустите playbook с флагом `--diff` и убедитесь, что playbook идемпотентен.
```
ansible-playbook -i inventory/prod.yml site.yml --diff 

PLAY [Install Java] *****************************************************************************************************************************************

TASK [Gathering Facts] **************************************************************************************************************************************
ok: [ubuntu]

TASK [Set facts for Java 11 vars] ***************************************************************************************************************************
ok: [ubuntu]

TASK [Upload .tar.gz file containing binaries from local storage] *******************************************************************************************
ok: [ubuntu]

TASK [Ensure installation dir exists] ***********************************************************************************************************************
ok: [ubuntu]

TASK [Extract java in the installation directory] ***********************************************************************************************************
skipping: [ubuntu]

TASK [Export environment variables] *************************************************************************************************************************
ok: [ubuntu]

PLAY [Install Elasticsearch] ********************************************************************************************************************************

TASK [Gathering Facts] **************************************************************************************************************************************
ok: [ubuntu]

TASK [Upload tar.gz Elasticsearch from remote URL] **********************************************************************************************************
ok: [ubuntu]

TASK [Create directrory for Elasticsearch] ******************************************************************************************************************
ok: [ubuntu]

TASK [Extract Elasticsearch in the installation directory] **************************************************************************************************
skipping: [ubuntu]

TASK [Set environment Elastic] ******************************************************************************************************************************
ok: [ubuntu]

PLAY [Install Kibana] ***************************************************************************************************************************************

TASK [Gathering Facts] **************************************************************************************************************************************
ok: [ubuntu]

TASK [Upload tar.gz Kibana from remote URL] *****************************************************************************************************************
ok: [ubuntu]

TASK [Create directrory for Kibana] *************************************************************************************************************************
ok: [ubuntu]

TASK [Extract Kibana in the installation directory] *********************************************************************************************************
skipping: [ubuntu]

TASK [Set environment Kibana] *******************************************************************************************************************************
ok: [ubuntu]

PLAY RECAP **************************************************************************************************************************************************
ubuntu                     : ok=13   changed=0    unreachable=0    failed=0    skipped=3    rescued=0    ignored=0   

```
9. Подготовьте README.md файл по своему playbook. В нём должно быть описано: что делает playbook, какие у него есть параметры и теги.
```
Playbook  создает fact для java home, загружает на клиента архив с java, создает директоию для java, разархивирует туда java и из template устанавливает переменные окружения. Аналогично все происходит для elasticsearch, Kibana. Logstash ставится на другую ноду, но действия происходят те же.
Теги которые есть: 
playbook: site.yml

  play #1 (all): Install Java   TAGS: []
      TASK TAGS: [java]

  play #2 (elasticsearch): Install Elasticsearch        TAGS: []
      TASK TAGS: [elastic]

  play #3 (elasticsearch): Install Kibana       TAGS: []
      TASK TAGS: [kibana]

  play #4 (logstash): Install Logstash  TAGS: []
      TASK TAGS: [logstash]
---
- name: Install Java
  hosts: all
  tasks:
    - name: Set facts for Java 11 vars
      set_fact:
        java_home: "/opt/jdk/{{ java_jdk_version }}"
      tags: java
    - name: Upload .tar.gz file containing binaries from local storage
      copy:
        src: "{{ java_oracle_jdk_package }}"
        dest: "/tmp/jdk-{{ java_jdk_version }}.tar.gz"
      register: download_java_binaries
      until: download_java_binaries is succeeded
      tags: java
    - name: Ensure installation dir exists
      become: true
      file:
        state: directory
        path: "{{ java_home }}"
      tags: java
    - name: Extract java in the installation directory
      become: true
      unarchive:
        copy: false
        src: "/tmp/jdk-{{ java_jdk_version }}.tar.gz"
        dest: "{{ java_home }}"
        extra_opts: [--strip-components=1]
        creates: "{{ java_home }}/bin/java"
      tags:
        - java
    - name: Export environment variables
      become: true
      template:
        src: jdk.sh.j2
        dest: /etc/profile.d/jdk.sh
      tags: java
- name: Install Elasticsearch
  hosts: elasticsearch
  tasks:
    - name: Upload tar.gz Elasticsearch from remote URL
      get_url:
        url: "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-{{ elastic_version }}-linux-x86_64.tar.gz"
        dest: "/tmp/elasticsearch-{{ elastic_version }}-linux-x86_64.tar.gz"
        mode: 0755
        timeout: 60
        force: true
        validate_certs: false
      register: get_elastic
      until: get_elastic is succeeded
      tags: elastic
    - name: Create directrory for Elasticsearch
      become: true
      file:
        state: directory
        path: "{{ elastic_home }}"
      tags: elastic
    - name: Extract Elasticsearch in the installation directory
      become: true
      unarchive:
        copy: false
        src: "/tmp/elasticsearch-{{ elastic_version }}-linux-x86_64.tar.gz"
        dest: "{{ elastic_home }}"
        extra_opts: [--strip-components=1]
        creates: "{{ elastic_home }}/bin/elasticsearch"
      tags:
        - elastic
    - name: Set environment Elastic
      become: true
      template:
        src: templates/elk.sh.j2
        dest: /etc/profile.d/elk.sh
      tags: elastic
- name: Install Kibana
  hosts: elasticsearch
  tasks:
    - name: Upload tar.gz Kibana from remote URL
      get_url:
        url: "https://artifacts.elastic.co/downloads/kibana/kibana-{{ kibana_version }}-linux-x86_64.tar.gz"
        dest: "/tmp/kibana-{{ kibana_version }}-linux-x86_64.tar.gz"
        mode: 0755
        timeout: 60
        force: true
        validate_certs: false
      register: get_kibana
      until: get_kibana is succeeded
      tags: kibana
    - name: Create directrory for Kibana
      become: true
      file:
        state: directory
        path: "{{ kibana_home }}"
      tags: kibana
    - name: Extract Kibana in the installation directory
      become: true
      unarchive:
        copy: false
        src: "/tmp/kibana-{{ kibana_version }}-linux-x86_64.tar.gz"
        dest: "{{ kibana_home }}"
        extra_opts: [--strip-components=1]
        creates: "{{ kibana_home }}/bin/kibana"
      tags:
        - kibana
    - name: Set environment Kibana
      become: true
      template:
        src: templates/kbn.sh.j2
        dest: /etc/profile.d/kbn.sh
      tags: kibana
- name: Install Logstash
  hosts: logstash
  tasks:
    - name: Upload tar.gz Logstash from remote URL
      get_url:
        url: "https://artifacts.elastic.co/downloads/logstash/logstash-{{ logstash_version }}-linux-x86_64.tar.gz"
        dest: "/tmp/logstash-{{ logstash_version }}-linux-x86_64.tar.gz"
        mode: 0755
        timeout: 60
        force: true
        validate_certs: false
      register: get_logstash
      until: get_logstash is succeeded
      tags: logstash
    - name: Create directrory for Logstash
      become: true
      file:
        state: directory
        path: "{{ logstash_home }}"
      tags: logstash
    - name: Extract Logstash in the installation directory
      become: true
      unarchive:
        copy: false
        src: "/tmp/logstash-{{ logstash_version }}-linux-x86_64.tar.gz"
        dest: "{{ logstash_home }}"
        extra_opts: [--strip-components=1]
        creates: "{{ logstash_home }}/bin/logstash"
      tags:
        - logstash
    - name: Set environment Logstash
      become: true
      template:
        src: templates/lgs.sh.j2
        dest: /etc/profile.d/lgs.sh
      tags: logstash

```
10. Готовый playbook выложите в свой репозиторий, в ответ предоставьте ссылку на него.

## Необязательная часть

1. Приготовьте дополнительный хост для установки logstash.
2. Пропишите данный хост в `prod.yml` в новую группу `logstash`.
```
logstash:
  hosts:
    ubuntu2:
      ansible_host: 178.154.202.39
      ansible_connection: ssh
      ansible_user: anantahari
```
3. Дополните playbook ещё одним play, который будет исполнять установку logstash только на выделенный для него хост.
4. Все переменные для нового play определите в отдельный файл `group_vars/logstash/vars.yml`.
5. Logstash конфиг должен конфигурироваться в части ссылки на elasticsearch (можно взять, например его IP из facts или определить через vars).
```
Не понял как выполнить этот шаг. Подскажите, пожалуйста, есть есть возможность.
```
6. Дополните README.md, протестируйте playbook, выложите новую версию в github. В ответ предоставьте ссылку на репозиторий.

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
