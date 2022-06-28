# Домашнее задание к занятию "12.4 Развертывание кластера на собственных серверах, лекция 2"
Новые проекты пошли стабильным потоком. Каждый проект требует себе несколько кластеров: под тесты и продуктив. Делать все руками — не вариант, поэтому стоит автоматизировать подготовку новых кластеров.

## Задание 1: Подготовить инвентарь kubespray
Новые тестовые кластеры требуют типичных простых настроек. Нужно подготовить инвентарь и проверить его работу. Требования к инвентарю:
* подготовка работы кластера из 5 нод: 1 мастер и 4 рабочие ноды;
* в качестве CRI — containerd;
* запуск etcd производить на мастере.

Решение
```
Сгенерировал hosts таким скриптом:
declare -a IPS=(51.250.47.243 51.250.38.183 51.250.33.181)
 2503  CONFIG_FILE=inventory/mycluster/hosts.yaml python3 contrib/inventory_builder/inventory.py ${IPS[@]}
```
 Поправил [hosts.yaml](https://github.com/AnantaHari/devops-netology/blob/main/12-kubernetes-04-install-part-2/hosts.yaml)
```
 Выполнил разворачивание облака командой: ansible-playbook -i inventory/mycluster/hosts.yaml  --become --become-user=root cluster.yml
 Потом выполнил:
 {
     mkdir -p $HOME/.kube
     sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
     sudo chown $(id -u):$(id -g) $HOME/.kube/config
 }
 После этого отработала команда:
 kubectl get nodes
 yc-user@cp1:~/.kube$ kubectl get nodes
NAME    STATUS   ROLES           AGE   VERSION
cp1     Ready    control-plane   16h   v1.24.2
node1   Ready    <none>          16h   v1.24.2
node2   Ready    <none>          16h   v1.24.2

Потом скопировал себе файл конфиг:
scp yc-user@51.250.47.243:/home/yc-user/.kube/config /Users/anantahari/.kube
Исправил в нем адрес, но подключиться не мог, т.к. для внешнего адреса небыло сертификата.
Выполнил следующее на control node:
yc-user@cp1:~/.kube$ sudo rm /etc/kubernetes/pki/apiserver.*
yc-user@cp1:~/.kube$ sudo kubeadm init phase certs all --apiserver-advertise-address=0.0.0.0 --apise
rver-cert-extra-sans=10.233.0.1,10.130.0.16,127.0.0.1,51.250.47.243
yc-user@cp1:~/.kube$ sudo systemctl restart kubelet
yc-user@cp1:~/.kube$ kubectl get nodes
NAME    STATUS   ROLES           AGE   VERSION
cp1     Ready    control-plane   16h   v1.24.2
node1   Ready    <none>          16h   v1.24.2
node2   Ready    <none>          16h   v1.24.2

В файле k8s-cluster.yml значение поумолчанию container_manager: containerd поэтому редактировать этот файл не пришлось.
```

## Задание 2 (*): подготовить и проверить инвентарь для кластера в AWS
Часть новых проектов хотят запускать на мощностях AWS. Требования похожи:
* разворачивать 5 нод: 1 мастер и 4 рабочие ноды;
* работать должны на минимально допустимых EC2 — t3.small.

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

cd kubernetes-for-beginners/99-misc
./list-vms.sh
./create-vms.sh
cp -rfp inventory/sample inventory/mycluster
В kubespray/README.md указываем внешние айпишники виртуалок в разделе # Update Ansible inventory file with inventory builder и выполняем эти строчки
Отредактировали hosts.yaml
ansible-playbook -i inventory/mycluster/hosts.yaml  --become --become-user=root cluster.yml
ssh yc-user@ip_address_cp1
kubectl get nodes
{
    mkdir -p $HOME/.kube
    sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
    sudo chown $(id -u):$(id -g) $HOME/.kube/config
}
kubectl get nodes
vim ~/.kube/config
копируем строку certificate-authority-data в такой же файл на нашем компе
меня ip
и также копируем строки пользователя(client-certificate-data и client-key-data)
в inventory/mycluster/group_vars/k8s_cluster/k8s-cluster.yml в параметре container_manager можно поменять менеджера
в этом же файле в параметр supplementary_addresses_in_ssl_keys добавляем внешний адрес машины
kubectl -context=efox -n kube-system get po -o wide
