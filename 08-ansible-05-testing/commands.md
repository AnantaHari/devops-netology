molecule init role new_role
molecule init scenario open --driver-name openstack
molecule matrix converge
molecule matrix converge -s open - тестирование только указанного сценария
molecule login --host имя машины - для подключения к машине
molecule test
molecule converge - так он не удаляет машину
pip3 install molecule
pip3 install molecule_docker
pip3 install molecule_podman

docker run -v $(pwd):/opt/elasticsearch_role -it tox_docker:latest /bin/bash

tox -l
tox -e py39-ansible28
ansible-galaxy collection install "community.docker:>=1.9.1" -c
