Настройки playbook:
    mode: present - задание прав на файлы

cat ~/.ssh/id_rsa.pub | xclip -sel clip

ansible-playbook site.yml -i inventory/prod
ssh -i ~/.ssh/id_rsa centos@178.154.221.25 
sudo systemctl status elasticsearch
curl http://127.0.0.1:9200
sudo ls /etc/kibana/
sudo cp /etc/kibana/kibana.yml ~/
sudo chown anantahari:anantahari kibana.yml
scp <ip_address>:/home/anantahari/kibana.yml templates/kibana.yml.j2
curl http://localhost:5601 -v

 Команды для Яндекса:
yc init
yc vpc network create --name net --labels my-label=netology --description "my first network via yc"
yc vpc subnet create --name my-subnet-a --zone ru-central1-a --range 10.1.2.0/24 --network-name net --description "my first subnet via yc"
yc vpc subnet list 
yc vpc subnet delete <SUBNET-NAME>|<SUBNET-ID>
yc vpc network delete <SUBNET-NAME>|<SUBNET-ID>

Аутентифицируйтесь от имени сервисного аккаунта (https://cloud.yandex.ru/docs/cli/operations/authentication/service-account)
подключиться к виртуалке: ssh -i ~/.ssh/id_rsa centos@62.84.119.249 