# Домашнее задание к занятию "13.3 работа с kubectl"
## Задание 1: проверить работоспособность каждого компонента
Для проверки работы можно использовать 2 способа: port-forward и exec. Используя оба способа, проверьте каждый компонент:
* сделайте запросы к бекенду;
* сделайте запросы к фронту;
* подключитесь к базе данных.

```
Решение
```
[front.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-03-kubectl/front.yaml)  
[back.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-03-kubectl/back.yaml) 
[postgres-statefulset.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-03-kubectl/postgres-statefulset.yaml) 
![Скриншот 1 задания 1](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-13%20в%2010.04.51.png)
![Скриншот 2 задания 1](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-13%20в%2010.05.14.png)
![Скриншот 3 задания 1](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-13%20в%2010.05.49.png)
![Скриншот 4 задания 1](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-13%20в%2010.05.57.png)
![Скриншот 5 задания 1](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-13%20в%2016.00.51.png)
![Скриншот 6 задания 1](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-13%20в%2016.01.04.png)

## Задание 2: ручное масштабирование

При работе с приложением иногда может потребоваться вручную добавить пару копий. Используя команду kubectl scale, попробуйте увеличить количество бекенда и фронта до 3. Проверьте, на каких нодах оказались копии после каждого действия (kubectl describe, kubectl get pods -o wide). После уменьшите количество копий до 1.

```
Решение
```
[front.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-03-kubectl/front.yaml)  
[back.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-03-kubectl/back.yaml)  
![Скриншот задания 2](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-13%20в%2009.59.55.png)
```
describe тоже делал, просто вывод большой, не помещается на скриншот
```

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

kubectl completion --help
cd 30-kubectl
kubectl api-resources
kubectl explain pod
kubectl explain pod.spec
kubectl explain pod.spec.containers
kubectl describe nodes master
kubectl describe nodes
конфиг лежит в папке ~/.kube и относительного это пути указываются пути в конфиге
cd 00-kube-config
выполнили все из README
cd ..
kubectl delete -f . - удалить все
kubectl -n namespace_name get pods -A | grep что_искать
kubectl -n namespace_name logs -l app=имя_метки -c code --tail=5000 | grep что_искать | wc -l (wc -l считатет кол-во записей)
