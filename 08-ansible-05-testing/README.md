# Домашнее задание к занятию "08.05 Тестирование Roles"

## Подготовка к выполнению
1. Установите molecule: `pip3 install "molecule==3.4.0"`
2. Соберите локальный образ на основе [Dockerfile](./Dockerfile)

## Основная часть

Наша основная цель - настроить тестирование наших ролей. Задача: сделать сценарии тестирования для kibana, filebeat. Ожидаемый результат: все сценарии успешно проходят тестирование ролей.

### Molecule

1. Запустите  `molecule test` внутри корневой директории elasticsearch-role, посмотрите на вывод команды.
```
Начинается тестирование
➜  elasticsearch_role git:(main) ✗ molecule test
INFO     default scenario test matrix: dependency, lint, cleanup, destroy, syntax, create, prepare, converge, idempotence, side_effect, verify, cleanup, destroy
INFO     Performing prerun...
INFO     Set ANSIBLE_LIBRARY=/Users/anantahari/.cache/ansible-compat/3c3d70/modules:/Users/anantahari/.ansible/plugins/modules:/usr/share/ansible/plugins/modules
INFO     Set ANSIBLE_COLLECTIONS_PATH=/Users/anantahari/.cache/ansible-compat/3c3d70/collections:/Users/anantahari/.ansible/collections:/usr/share/ansible/collections
INFO     Set ANSIBLE_ROLES_PATH=/Users/anantahari/.cache/ansible-compat/3c3d70/roles:/Users/anantahari/.ansible/roles:/usr/share/ansible/roles:/etc/ansible/roles
INFO     Running default > dependency
WARNING  Skipping, missing the requirements file.
WARNING  Skipping, missing the requirements file.
INFO     Running default > lint
INFO     Lint is disabled.
INFO     Running default > cleanup
WARNING  Skipping, cleanup playbook not configured.
INFO     Running default > destroy
INFO     Sanity checks: 'docker'

PLAY [Destroy] *****************************************************************

TASK [Destroy molecule instance(s)] ********************************************
changed: [localhost] => (item=centos7)
changed: [localhost] => (item=ubuntu)

TASK [Wait for instance(s) deletion to complete] *******************************
ok: [localhost] => (item=centos7)
ok: [localhost] => (item=ubuntu)

TASK [Delete docker networks(s)] ***********************************************

PLAY RECAP *********************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0

INFO     Running default > syntax

playbook: /Users/anantahari/Git/devops-netology/08-ansible-05-testing/playbook/roles/elasticsearch_role/molecule/default/converge.yml
INFO     Running default > create

PLAY [Create] ******************************************************************

TASK [Log into a Docker registry] **********************************************
skipping: [localhost] => (item=None) 
skipping: [localhost] => (item=None) 
skipping: [localhost]

TASK [Check presence of custom Dockerfiles] ************************************
ok: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True})
ok: [localhost] => (item={'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True})

TASK [Create Dockerfiles from image names] *************************************
skipping: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True}) 
skipping: [localhost] => (item={'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True}) 

TASK [Discover local Docker images] ********************************************
ok: [localhost] => (item={'changed': False, 'skipped': True, 'skip_reason': 'Conditional result was False', 'item': {'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True}, 'ansible_loop_var': 'item', 'i': 0, 'ansible_index_var': 'i'})
ok: [localhost] => (item={'changed': False, 'skipped': True, 'skip_reason': 'Conditional result was False', 'item': {'image': 'docker.io/pycontribs/ubuntu:latest', 'name': 'ubuntu', 'pre_build_image': True}, 'ansible_loop_var': 'item', 'i': 1, 'ansible_index_var': 'i'})

TASK [Build an Ansible compatible image (new)] *********************************
skipping: [localhost] => (item=molecule_local/docker.io/pycontribs/centos:7) 
skipping: [localhost] => (item=molecule_local/docker.io/pycontribs/ubuntu:latest) 
```
2. Перейдите в каталог с ролью kibana-role и создайте сценарий тестирования по умолчаню при помощи `molecule init scenario --driver-name docker`.
3. Добавьте несколько разных дистрибутивов (centos:8, ubuntu:latest) для инстансов и протестируйте роль, исправьте найденные ошибки, если они есть.
4. Добавьте несколько assert'ов в verify.yml файл, для  проверки работоспособности kibana-role (проверка, что web отвечает, проверка логов, etc). Запустите тестирование роли повторно и проверьте, что оно прошло успешно.
```
INFO     Running default > verify
INFO     Running Ansible Verifier

PLAY [Check centos service status] *********************************************

TASK [Set fact] ****************************************************************
changed: [centos7]

TASK [Print return information from the previous task] *************************
ok: [centos7] => {
    "kibana_service_status.stdout": "kibana is running"
}

TASK [Ensure status is correct] ************************************************
ok: [centos7] => {
    "changed": false,
    "msg": "All assertions passed"
}

PLAY [Check ubuntu service status] *********************************************

TASK [Set fact] ****************************************************************
changed: [ubuntu]

TASK [Ensure status is correct] ************************************************
ok: [ubuntu] => {
    "changed": false,
    "msg": "All assertions passed"
}

PLAY RECAP *********************************************************************
centos7                    : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
ubuntu                     : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

INFO     Verifier completed successfully.
INFO     Running default > cleanup
WARNING  Skipping, cleanup playbook not configured.
INFO     Running default > destroy

PLAY [Destroy] *****************************************************************

TASK [Destroy molecule instance(s)] ********************************************
changed: [localhost] => (item=centos7)
changed: [localhost] => (item=ubuntu)

TASK [Wait for instance(s) deletion to complete] *******************************
FAILED - RETRYING: [localhost]: Wait for instance(s) deletion to complete (300 retries left).
changed: [localhost] => (item=centos7)
changed: [localhost] => (item=ubuntu)

TASK [Delete docker networks(s)] ***********************************************

PLAY RECAP *********************************************************************
localhost                  : ok=2    changed=2    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0

INFO     Pruning extra files from scenario ephemeral directory
➜  kibana_role git:(main) ✗ 
```
5. Повторите шаги 2-4 для filebeat-role.
```
INFO     Running default > verify
INFO     Running Ansible Verifier

PLAY [Check centos service status] *********************************************

TASK [Set fact] ****************************************************************
changed: [ubuntu]
changed: [centos7]

TASK [Ensure status is correct] ************************************************
ok: [centos7] => {
    "changed": false,
    "msg": "All assertions passed"
}
ok: [ubuntu] => {
    "changed": false,
    "msg": "All assertions passed"
}

PLAY RECAP *********************************************************************
centos7                    : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
ubuntu                     : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

INFO     Verifier completed successfully.
INFO     Running default > cleanup
WARNING  Skipping, cleanup playbook not configured.
INFO     Running default > destroy

PLAY [Destroy] *****************************************************************

TASK [Destroy molecule instance(s)] ********************************************
changed: [localhost] => (item=centos7)
changed: [localhost] => (item=ubuntu)

TASK [Wait for instance(s) deletion to complete] *******************************
FAILED - RETRYING: [localhost]: Wait for instance(s) deletion to complete (300 retries left).
changed: [localhost] => (item=centos7)
changed: [localhost] => (item=ubuntu)

TASK [Delete docker networks(s)] ***********************************************

PLAY RECAP *********************************************************************
localhost                  : ok=2    changed=2    unreachable=0    failed=0    skipped=1    rescued=0    ignored=0

INFO     Pruning extra files from scenario ephemeral directory
➜  filebeat_role git:(main) ✗ 
```
6. Добавьте новый тег на коммит с рабочим сценарием в соответствии с семантическим версионированием.


### Tox

1. Запустите `docker run --privileged=True -v <path_to_repo>:/opt/elasticsearch-role -w /opt/elasticsearch-role -it <image_name> /bin/bash`, где path_to_repo - путь до корня репозитория с elasticsearch-role на вашей файловой системе.
2. Внутри контейнера выполните команду `tox`, посмотрите на вывод.
```
Начинает кешировать виртуальные среды для тестирования
[root@2024823d4ca1 elasticsearch-role]# tox
py36-ansible28 create: /opt/elasticsearch-role/.tox/py36-ansible28
py36-ansible28 installdeps: -rtest-requirements.txt, ansible<2.9
=============================================================== log start ===============================================================
Collecting ansible<2.9
  Using cached ansible-2.8.20.tar.gz (12.7 MB)
  Preparing metadata (setup.py): started
  Preparing metadata (setup.py): finished with status 'done'
Collecting molecule==3.5.2
  Using cached molecule-3.5.2-py3-none-any.whl (240 kB)
Collecting molecule_docker
  Using cached molecule_docker-1.1.0-py3-none-any.whl (16 kB)
Collecting molecule_podman
  Using cached molecule_podman-1.1.0-py3-none-any.whl (15 kB)
Collecting docker
  Using cached docker-5.0.3-py2.py3-none-any.whl (146 kB)
Collecting ansible-lint
  Using cached ansible_lint-5.4.0-py3-none-any.whl (119 kB)
Collecting yamllint
  Using cached yamllint-1.26.3.tar.gz (126 kB)
  Preparing metadata (setup.py): started
  Preparing metadata (setup.py): finished with status 'done'
Collecting importlib-metadata
  Using cached importlib_metadata-4.8.3-py3-none-any.whl (17 kB)
Collecting PyYAML<6,>=5.1
  Downloading PyYAML-5.4.1-cp36-cp36m-manylinux2014_aarch64.whl (726 kB)
Collecting click-help-colors>=0.9
  Using cached click_help_colors-0.9.1-py3-none-any.whl (5.5 kB)
Collecting pluggy<2.0,>=0.7.1
  Using cached pluggy-1.0.0-py2.py3-none-any.whl (13 kB)
Collecting enrich>=1.2.5
  Using cached enrich-1.2.7-py3-none-any.whl (8.7 kB)
Collecting paramiko<3,>=2.5.0
  Using cached paramiko-2.9.2-py2.py3-none-any.whl (210 kB)
Collecting subprocess-tee>=0.3.5
  Using cached subprocess_tee-0.3.5-py3-none-any.whl (8.0 kB)
Collecting cookiecutter>=1.7.3
  Using cached cookiecutter-1.7.3-py2.py3-none-any.whl (34 kB)
Collecting dataclasses
  Downloading dataclasses-0.8-py3-none-any.whl (19 kB)
Collecting packaging
  Using cached packaging-21.3-py3-none-any.whl (40 kB)
Collecting ansible-compat>=0.5.0
  Using cached ansible_compat-1.0.0-py3-none-any.whl (16 kB)
Collecting selinux
  Using cached selinux-0.2.1-py2.py3-none-any.whl (4.3 kB)
Collecting Jinja2>=2.11.3
  Using cached Jinja2-3.0.3-py3-none-any.whl (133 kB)
Collecting click<9,>=8.0
  Using cached click-8.0.4-py3-none-any.whl (97 kB)
Collecting cerberus!=1.3.3,!=1.3.4,>=1.3.1
  Using cached Cerberus-1.3.2.tar.gz (52 kB)
  Preparing metadata (setup.py): started
  Preparing metadata (setup.py): finished with status 'done'
  и т.д.
```
3. Добавьте файл `tox.ini` в корень репозитория каждой своей роли.
4. Создайте облегчённый сценарий для `molecule`. Проверьте его на исполнимость.
5. Пропишите правильную команду в `tox.ini` для того чтобы запускался облегчённый сценарий.
6. Запустите `docker` контейнер так, чтобы внутри оказались обе ваши роли.
```
➜  roles git:(main) ✗ docker run --privileged=True -v $(pwd)/elasticsearch_role:/opt/elasticsearch-role \
> -v $(pwd)/filebeat_role:/opt/filebeat-role \
> -v $(pwd)/kibana_role:/opt/kibana-role -w /opt/elasticsearch-role -it tox_docker:latest /bin/bash
[root@2024823d4ca1 elasticsearch-role]# cd ..
[root@2024823d4ca1 opt]# ls -la
total 8
drwxr-xr-x  1 root root 4096 Feb 26 11:22 .
drwxr-xr-x  1 root root 4096 Feb 26 11:22 ..
drwxr-xr-x 17 root root  544 Feb 26 08:18 elasticsearch-role
drwxr-xr-x 15 root root  480 Feb 26 11:05 filebeat-role
drwxr-xr-x 15 root root  480 Feb 26 11:05 kibana-role
```
7. Зайдти поочерёдно в каждую из них и запустите команду `tox`. Убедитесь, что всё отработало успешно.
```
[root@2024823d4ca1 elasticsearch-role]# tox -e py39-ansible28
py39-ansible28 installed: ansible==2.8.20,ansible-compat==1.0.0,ansible-lint==5.4.0,arrow==1.2.2,bcrypt==3.2.0,binaryornot==0.4.4,bracex==2.2.1,Cerberus==1.3.2,certifi==2021.10.8,cffi==1.15.0,chardet==4.0.0,charset-normalizer==2.0.12,click==8.0.4,click-help-colors==0.9.1,colorama==0.4.4,commonmark==0.9.1,cookiecutter==1.7.3,cryptography==36.0.1,distro==1.7.0,docker==5.0.3,enrich==1.2.7,idna==3.3,Jinja2==3.0.3,jinja2-time==0.2.0,MarkupSafe==2.1.0,molecule==3.5.2,molecule-docker==1.1.0,molecule-podman==1.1.0,packaging==21.3,paramiko==2.9.2,pathspec==0.9.0,pluggy==1.0.0,poyo==0.5.0,pycparser==2.21,Pygments==2.11.2,PyNaCl==1.5.0,pyparsing==3.0.7,python-dateutil==2.8.2,python-slugify==6.1.0,PyYAML==5.4.1,requests==2.27.1,rich==11.2.0,ruamel.yaml==0.17.21,ruamel.yaml.clib==0.2.6,selinux==0.2.1,six==1.16.0,subprocess-tee==0.3.5,tenacity==8.0.1,text-unidecode==1.3,urllib3==1.26.8,wcmatch==8.3,websocket-client==1.3.1,yamllint==1.26.3
py39-ansible28 run-test-pre: PYTHONHASHSEED='1311144589'
py39-ansible28 run-test: commands[0] | molecule test -s alternative
INFO     alternative scenario test matrix: destroy, create, converge, destroy
INFO     Performing prerun...
INFO     Set ANSIBLE_LIBRARY=/root/.cache/ansible-compat/2bda03/modules:/root/.ansible/plugins/modules:/usr/share/ansible/plugins/modules
INFO     Set ANSIBLE_COLLECTIONS_PATHS=/root/.cache/ansible-compat/2bda03/collections:/root/.ansible/collections:/usr/share/ansible/collections
INFO     Set ANSIBLE_ROLES_PATH=/root/.cache/ansible-compat/2bda03/roles:/root/.ansible/roles:/usr/share/ansible/roles:/etc/ansible/roles
INFO     Running alternative > destroy
INFO     Sanity checks: 'podman'

PLAY [Destroy] **************************************************************************************************************************

TASK [Destroy molecule instance(s)] *****************************************************************************************************
changed: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True})

TASK [Wait for instance(s) deletion to complete] ****************************************************************************************
changed: [localhost] => (item={'started': 1, 'finished': 0, 'ansible_job_id': '425786910408.52', 'results_file': '/root/.ansible_async/425786910408.52', 'changed': True, 'failed': False, 'item': {'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True}, 'ansible_loop_var': 'item'})

PLAY RECAP ******************************************************************************************************************************
localhost                  : ok=2    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

[root@2024823d4ca1 kibana-role]# tox -e py39-ansible28
py39-ansible28 create: /opt/kibana-role/.tox/py39-ansible28
py39-ansible28 installdeps: -rtest-requirements.txt, ansible<2.9
py39-ansible28 installed: ansible==2.8.20,ansible-compat==1.0.0,ansible-lint==5.4.0,arrow==1.2.2,bcrypt==3.2.0,binaryornot==0.4.4,bracex==2.2.1,Cerberus==1.3.2,certifi==2021.10.8,cffi==1.15.0,chardet==4.0.0,charset-normalizer==2.0.12,click==8.0.4,click-help-colors==0.9.1,colorama==0.4.4,commonmark==0.9.1,cookiecutter==1.7.3,cryptography==36.0.1,distro==1.7.0,docker==5.0.3,enrich==1.2.7,idna==3.3,Jinja2==3.0.3,jinja2-time==0.2.0,MarkupSafe==2.1.0,molecule==3.5.2,molecule-docker==1.1.0,molecule-podman==1.1.0,packaging==21.3,paramiko==2.9.2,pathspec==0.9.0,pluggy==1.0.0,poyo==0.5.0,pycparser==2.21,Pygments==2.11.2,PyNaCl==1.5.0,pyparsing==3.0.7,python-dateutil==2.8.2,python-slugify==6.1.0,PyYAML==5.4.1,requests==2.27.1,rich==11.2.0,ruamel.yaml==0.17.21,ruamel.yaml.clib==0.2.6,selinux==0.2.1,six==1.16.0,subprocess-tee==0.3.5,tenacity==8.0.1,text-unidecode==1.3,urllib3==1.26.8,wcmatch==8.3,websocket-client==1.3.1,yamllint==1.26.3
py39-ansible28 run-test-pre: PYTHONHASHSEED='4258917679'
py39-ansible28 run-test: commands[0] | molecule test -s alternative
INFO     alternative scenario test matrix: destroy, create, converge, destroy
INFO     Performing prerun...
INFO     Set ANSIBLE_LIBRARY=/root/.cache/ansible-compat/c5f5dd/modules:/root/.ansible/plugins/modules:/usr/share/ansible/plugins/modules
INFO     Set ANSIBLE_COLLECTIONS_PATHS=/root/.cache/ansible-compat/c5f5dd/collections:/root/.ansible/collections:/usr/share/ansible/collections
INFO     Set ANSIBLE_ROLES_PATH=/root/.cache/ansible-compat/c5f5dd/roles:/root/.ansible/roles:/usr/share/ansible/roles:/etc/ansible/roles
INFO     Running alternative > destroy
INFO     Sanity checks: 'podman'

PLAY [Destroy] **************************************************************************************************************************

TASK [Destroy molecule instance(s)] *****************************************************************************************************
changed: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True})

TASK [Wait for instance(s) deletion to complete] ****************************************************************************************
changed: [localhost] => (item={'started': 1, 'finished': 0, 'ansible_job_id': '151292425361.378', 'results_file': '/root/.ansible_async/151292425361.378', 'changed': True, 'failed': False, 'item': {'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True}, 'ansible_loop_var': 'item'})

PLAY RECAP ******************************************************************************************************************************
localhost                  : ok=2    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0

[root@2024823d4ca1 filebeat-role]# tox -e py39-ansible28
py39-ansible28 create: /opt/filebeat-role/.tox/py39-ansible28
py39-ansible28 installdeps: -rtest-requirements.txt, ansible<2.9
py39-ansible28 installed: ansible==2.8.20,ansible-compat==1.0.0,ansible-lint==5.4.0,arrow==1.2.2,bcrypt==3.2.0,binaryornot==0.4.4,bracex==2.2.1,Cerberus==1.3.2,certifi==2021.10.8,cffi==1.15.0,chardet==4.0.0,charset-normalizer==2.0.12,click==8.0.4,click-help-colors==0.9.1,colorama==0.4.4,commonmark==0.9.1,cookiecutter==1.7.3,cryptography==36.0.1,distro==1.7.0,docker==5.0.3,enrich==1.2.7,idna==3.3,Jinja2==3.0.3,jinja2-time==0.2.0,MarkupSafe==2.1.0,molecule==3.5.2,molecule-docker==1.1.0,molecule-podman==1.1.0,packaging==21.3,paramiko==2.9.2,pathspec==0.9.0,pluggy==1.0.0,poyo==0.5.0,pycparser==2.21,Pygments==2.11.2,PyNaCl==1.5.0,pyparsing==3.0.7,python-dateutil==2.8.2,python-slugify==6.1.0,PyYAML==5.4.1,requests==2.27.1,rich==11.2.0,ruamel.yaml==0.17.21,ruamel.yaml.clib==0.2.6,selinux==0.2.1,six==1.16.0,subprocess-tee==0.3.5,tenacity==8.0.1,text-unidecode==1.3,urllib3==1.26.8,wcmatch==8.3,websocket-client==1.3.1,yamllint==1.26.3
py39-ansible28 run-test-pre: PYTHONHASHSEED='2406244543'
py39-ansible28 run-test: commands[0] | molecule test -s alternative
INFO     alternative scenario test matrix: destroy, create, converge, destroy
INFO     Performing prerun...
INFO     Set ANSIBLE_LIBRARY=/root/.cache/ansible-compat/50ad40/modules:/root/.ansible/plugins/modules:/usr/share/ansible/plugins/modules
INFO     Set ANSIBLE_COLLECTIONS_PATHS=/root/.cache/ansible-compat/50ad40/collections:/root/.ansible/collections:/usr/share/ansible/collections
INFO     Set ANSIBLE_ROLES_PATH=/root/.cache/ansible-compat/50ad40/roles:/root/.ansible/roles:/usr/share/ansible/roles:/etc/ansible/roles
INFO     Running alternative > destroy
INFO     Sanity checks: 'podman'

PLAY [Destroy] **************************************************************************************************************************

TASK [Destroy molecule instance(s)] *****************************************************************************************************
changed: [localhost] => (item={'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True})

TASK [Wait for instance(s) deletion to complete] ****************************************************************************************
changed: [localhost] => (item={'started': 1, 'finished': 0, 'ansible_job_id': '890537725589.550', 'results_file': '/root/.ansible_async/890537725589.550', 'changed': True, 'failed': False, 'item': {'image': 'docker.io/pycontribs/centos:7', 'name': 'centos7', 'pre_build_image': True}, 'ansible_loop_var': 'item'})

PLAY RECAP ******************************************************************************************************************************
localhost                  : ok=2    changed=2    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
```
8. Добавьте новый тег на коммит с рабочим сценарием в соответствии с семантическим версионированием.

После выполнения у вас должно получится два сценария molecule и один tox.ini файл в каждом репозитории. Ссылки на репозитории являются ответами на домашнее задание. Не забудьте указать в ответе теги решений Tox и Molecule заданий.

## Необязательная часть

1. Проделайте схожие манипуляции для создания роли logstash.
2. Создайте дополнительный набор tasks, который позволяет обновлять стек ELK.
3. В ролях добавьте тестирование в раздел `verify.yml`. Данный раздел должен проверять, что logstash через команду `logstash -e 'input { stdin { } } output { stdout {} }'`  отвечате адекватно.
4. Создайте сценарий внутри любой из своих ролей, который умеет поднимать весь стек при помощи всех ролей.
5. Убедитесь в работоспособности своего стека. Создайте отдельный verify.yml, который будет проверять работоспособность интеграции всех инструментов между ними.
6. Выложите свои roles в репозитории. В ответ приведите ссылки.

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
