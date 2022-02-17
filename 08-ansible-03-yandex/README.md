# Домашнее задание к занятию "08.03 Использование Yandex Cloud"

## Подготовка к выполнению
1. Создайте свой собственный (или используйте старый) публичный репозиторий на github с произвольным именем.
2. Скачайте [playbook](./playbook/) из репозитория с домашним заданием и перенесите его в свой репозиторий.

## Основная часть
1. Допишите playbook: нужно сделать ещё один play, который устанавливает и настраивает kibana.
```
---
- name: Install Elasticsearch
  hosts: elasticsearch
  handlers:
    - name: restart Elasticsearch
      become: true
      service:
        name: elasticsearch
        state: restarted
  tasks:
    - name: "Download Elasticsearch's rpm"
      get_url:
        url: "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-{{ elk_stack_version }}-x86_64.rpm"
        dest: "/tmp/elasticsearch-{{ elk_stack_version }}-x86_64.rpm"
      register: download_elastic
      until: download_elastic is succeeded
    - name: Install Elasticsearch
      become: true
      yum:
        name: "/tmp/elasticsearch-{{ elk_stack_version }}-x86_64.rpm"
        state: present
    - name: Configure Elasticsearch
      become: true
      template:
        src: elasticsearch.yml.j2
        dest: /etc/elasticsearch/elasticsearch.yml
      notify: restart Elasticsearch
- name: Install Kibana
  hosts: kibana
  handlers:
    - name: restart Kibana
      become: true
      service:
        name: kibana
        state: restarted
  tasks:
    - name: "Download Kibana's rpm"
      get_url:
        url: "https://artifacts.elastic.co/downloads/kibana/kibana-{{ elk_stack_version }}-x86_64.rpm"
        dest: "/tmp/kibana-{{ elk_stack_version }}-x86_64.rpm"
      register: download_kibana
      until: download_kibana is succeeded
    - name: Install Kibana
      become: true
      yum:
        name: "/tmp/kibana-{{ elk_stack_version }}-x86_64.rpm"
        state: present
    - name: Configure Kibana
      become: true
      template:
        src: kibana.yml.j2
        dest: /etc/kibana/kibana.yml
      notify: restart Kibana
```
2. При создании tasks рекомендую использовать модули: `get_url`, `template`, `yum`, `apt`.
3. Tasks должны: скачать нужной версии дистрибутив, выполнить распаковку в выбранную директорию, сгенерировать конфигурацию с параметрами.
4. Приготовьте свой собственный inventory файл `prod.yml`.
5. Запустите `ansible-lint site.yml` и исправьте ошибки, если они есть.
6. Попробуйте запустить playbook на этом окружении с флагом `--check`.
```
ansible-playbook site.yml -i inventory/prod --check

PLAY [Install Elasticsearch] ******************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************
ok: [el-instance]

TASK [Download Elasticsearch's rpm] ***********************************************************************************************************************
ok: [el-instance]

TASK [Install Elasticsearch] ******************************************************************************************************************************
ok: [el-instance]

TASK [Configure Elasticsearch] ****************************************************************************************************************************
ok: [el-instance]

PLAY [Install Kibana] *************************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************
ok: [k-instance]

TASK [Download Kibana's rpm] ******************************************************************************************************************************
ok: [k-instance]

TASK [Install Kibana] *************************************************************************************************************************************
ok: [k-instance]

TASK [Configure Kibana] ***********************************************************************************************************************************
ok: [k-instance]

PLAY RECAP ************************************************************************************************************************************************
el-instance                : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
k-instance                 : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
```
7. Запустите playbook на `prod.yml` окружении с флагом `--diff`. Убедитесь, что изменения на системе произведены.
8. Повторно запустите playbook с флагом `--diff` и убедитесь, что playbook идемпотентен.
```
ansible-playbook site.yml -i inventory/prod --diff 

PLAY [Install Elasticsearch] ******************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************
ok: [el-instance]

TASK [Download Elasticsearch's rpm] ***********************************************************************************************************************
ok: [el-instance]

TASK [Install Elasticsearch] ******************************************************************************************************************************
ok: [el-instance]

TASK [Configure Elasticsearch] ****************************************************************************************************************************
ok: [el-instance]

PLAY [Install Kibana] *************************************************************************************************************************************

TASK [Gathering Facts] ************************************************************************************************************************************
ok: [k-instance]

TASK [Download Kibana's rpm] ******************************************************************************************************************************
ok: [k-instance]

TASK [Install Kibana] *************************************************************************************************************************************
ok: [k-instance]

TASK [Configure Kibana] ***********************************************************************************************************************************
ok: [k-instance]

PLAY RECAP ************************************************************************************************************************************************
el-instance                : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
k-instance                 : ok=4    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
```
9. Проделайте шаги с 1 до 8 для создания ещё одного play, который устанавливает и настраивает filebeat.
```
- name: Install filebeat
  hosts: app
  handlers:
    - name: restart filebeat
      become: true
      systemd:
        name: filebeat
        state: restarted
        enabled: true
  tasks:
    - name: "Download filebeat's rpm"
      get_url:
        url: "https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-{{ elk_stack_version }}-x86_64.rpm"
        dest: "/tmp/filebeat-{{ elk_stack_version }}-x86_64.rpm"
      register: download_filebeat
      until: download_filebeat is succeeded
    - name: Install filebeat
      become: true
      yum:
        name: "/tmp/filebeat-{{ elk_stack_version }}-x86_64.rpm"
        state: present
      notify: restart filebeat
    - name: Configure filebeat
      become: true
      template:
        src: filebeat.yml.j2
        dest: /etc/filebeat/filebeat.yml
      notify: restart filebeat
    - name: Set filebeat systemwork
      become: true
      command:
        cmd: filebeat modules enable system
        chdir: /usr/share/filebeat/bin
      register: filebeat_modules
      changed_when: filebeat_modules.stdout != 'Module system is already enabled'
    - name: Load Kibana dashboard
      become: true
      command:
        cmd: filebeat setup
        chdir: /usr/share/filebeat/bin
      register: filebeat_setup
      changed_when: false
      until: filebeat_setup is succeeded
```
10. Подготовьте README.md файл по своему playbook. В нём должно быть описано: что делает playbook, какие у него есть параметры и теги.
```
Playbook производит установку и настройку elasticsearch, kibana, filebeat чтобы filebeat собирал метрики с виртуалки, передавал их в elasticsearch и чтобы их можно было посмотреть в kibana.
```
11. Готовый playbook выложите в свой репозиторий, в ответ предоставьте ссылку на него.
```
https://github.com/AnantaHari/devops-netology/blob/main/08-ansible-03-yandex/playbook/site.yml
```

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
