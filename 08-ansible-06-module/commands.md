ansible localhost -m ping -a "data=something"
ansible-doc ping

cd lib/ansible/modules
touch my_own_module.py
python -m ansible.modules.my_own_module payload.json

cd ansible/playbooks

Тесты находятся в ansible/test/units/modules
Для запуска теста: ansible-test units apt
Чтобы заработал тест надо ставить дополнительный модуль

ansible-test sanity my_own_module

В папке playbooks выполняем: ansible-galaxy collection init my_own_namespace.yandex_cloud_elk

deactivate

ansible-galaxy collection build
tar xvfz my_own_namespace-yandex_cloud_elk-1.0.0.tar.gz
Архив нужно распаковывать в ~/.ansible/collection, либо рядом с playbook создать ansible_collections и положить туда

cd user
ansible-galaxy collection install my_own_namespace-yandex_cloud_elk-1.0.0.tar.gz -p ./collections

cat ~/.ssh/id_rsa.pub

ansible-playbook -i inventory/prod site.yml