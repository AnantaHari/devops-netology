# Домашнее задание к занятию "13.3 работа с kubectl"
## Задание 1: проверить работоспособность каждого компонента
Для проверки работы можно использовать 2 способа: port-forward и exec. Используя оба способа, проверьте каждый компонент:
* сделайте запросы к бекенду;
* сделайте запросы к фронту;
* подключитесь к базе данных.

## Задание 2: ручное масштабирование

При работе с приложением иногда может потребоваться вручную добавить пару копий. Используя команду kubectl scale, попробуйте увеличить количество бекенда и фронта до 3. Проверьте, на каких нодах оказались копии после каждого действия (kubectl describe, kubectl get pods -o wide). После уменьшите количество копий до 1.

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
