# Домашнее задание к занятию "13.1 контейнеры, поды, deployment, statefulset, services, endpoints"
Настроив кластер, подготовьте приложение к запуску в нём. Приложение стандартное: бекенд, фронтенд, база данных. Его можно найти в папке 13-kubernetes-config.

## Задание 1: подготовить тестовый конфиг для запуска приложения
Для начала следует подготовить запуск приложения в stage окружении с простыми настройками. Требования:
* под содержит в себе 2 контейнера — фронтенд, бекенд;
* регулируется с помощью deployment фронтенд и бекенд;
* база данных — через statefulset.

```
Решение
```
[stage.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-01/stage.yaml)  
[stage-postgres.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-01/stage-postgres.yaml)  
![Скриншот задания 1](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-07%20в%2021.12.28.png)


## Задание 2: подготовить конфиг для production окружения
Следующим шагом будет запуск приложения в production окружении. Требования сложнее:
* каждый компонент (база, бекенд, фронтенд) запускаются в своем поде, регулируются отдельными deployment’ами;
* для связи используются service (у каждого компонента свой);
* в окружении фронта прописан адрес сервиса бекенда;
* в окружении бекенда прописан адрес сервиса базы данных.
```
Решение
```
[front.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-01/front.yaml)  
[back.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-01/back.yaml)  
[prod-postgres.yaml](https://github.com/AnantaHari/devops-netology/blob/main/13-kubernetes-config-01/prod-postgres.yaml)  
![Скриншот задания 2](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-07%20в%2021.16.07.png)


## Задание 3 (*): добавить endpoint на внешний ресурс api
Приложению потребовалось внешнее api, и для его использования лучше добавить endpoint в кластер, направленный на это api. Требования:
* добавлен endpoint до внешнего api (например, геокодер).

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

В качестве решения прикрепите к ДЗ конфиг файлы для деплоя. Прикрепите скриншоты вывода команды kubectl со списком запущенных объектов каждого типа (pods, deployments, statefulset, service) или скриншот из самого Kubernetes, что сервисы подняты и работают.

---

cd 20-network-policy/manifests/main
kubectl delete -f .
kubectl get po
kubectl get svc
kubectl delete svc demo
cd 20-concepts/10-workload/10-pod
kubectl apply -f manifests/10-pod.yaml
kubectl get pods
kubectl get pods nginx -o yaml
kubectl get pods -o wide
kubectl describe pods nginx
kubectl describe pods nginx | grep "^IP:"

kubectl apply -f manifests/30-pod-with-error.yaml
kubectl get pods -o wide
kubectl get pods pod-with-error -o yaml
kubectl describe pods pod-with-error
kubectl logs pod-with-error -c multitool

kubectl apply -f manifests/40-pod-wo-errors.yaml
kubectl get pods -o wide
kubectl exec -c multitool pod-wo-errors -- curl localhost
kubectl exec -c nginx pod-wo-errors -- curl localhost:8080

kubectl delete -f manifests/

kubectl get deploy
cd ../20-deployment/
kubectl apply -f manifests/10-multitool.yaml
kubectl get po
kubectl edit deployments.apps multitool
изменили кол-во реплик
kubectl delete po имя_пода
kubectl apply -f manifests/20-multitool-nginx.yaml
kubectl get rs
kubectl get rs имя_реплика_сета -o yaml

kubectl -n kube-system get po
kubectl -n kube-system get ds
kubectl get all


Команды из ДЗ
```
Задание 1
cd 13-kubernetes-config
kubectl apply -f stage.yaml
kubectl apply -f stage-postgres.yaml
kubectl get po
kubectl get deploy
kubectl get svc
kubectl get statefulset

Задание 2
kubectl apply -f front.yaml -f back.yaml -f prod-postgres.yaml
kubectl get po
kubectl get deploy
kubectl get pv
kubectl get pvc
kubectl get svc
kubectl get statefulset
```
