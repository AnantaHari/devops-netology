# Домашнее задание к занятию "12.2 Команды для работы с Kubernetes"
Кластер — это сложная система, с которой крайне редко работает один человек. Квалифицированный devops умеет наладить работу всей команды, занимающейся каким-либо сервисом.
После знакомства с кластером вас попросили выдать доступ нескольким разработчикам. Помимо этого требуется служебный аккаунт для просмотра логов.

## Задание 1: Запуск пода из образа в деплойменте
Для начала следует разобраться с прямым запуском приложений из консоли. Такой подход поможет быстро развернуть инструменты отладки в кластере. Требуется запустить деплоймент на основе образа из hello world уже через deployment. Сразу стоит запустить 2 копии приложения (replicas=2).

Требования:
 * пример из hello world запущен в качестве deployment
 * количество реплик в deployment установлено в 2
 * наличие deployment можно проверить командой kubectl get deployment
 * наличие подов можно проверить командой kubectl get pods


## Задание 2: Просмотр логов для разработки
Разработчикам крайне важно получать обратную связь от штатно работающего приложения и, еще важнее, об ошибках в его работе.
Требуется создать пользователя и выдать ему доступ на чтение конфигурации и логов подов в app-namespace.

Требования:
 * создан новый токен доступа для пользователя
 * пользователь прописан в локальный конфиг (~/.kube/config, блок users)
 * пользователь может просматривать логи подов и их конфигурацию (kubectl logs pod <pod_id>, kubectl describe pod <pod_id>)


## Задание 3: Изменение количества реплик
Поработав с приложением, вы получили запрос на увеличение количества реплик приложения для нагрузки. Необходимо изменить запущенный deployment, увеличив количество реплик до 5. Посмотрите статус запущенных подов после увеличения реплик.

Требования:
 * в deployment из задания 1 изменено количество реплик на 5
 * проверить что все поды перешли в статус running (kubectl get pods)

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

kubectl get nodes
kubectl get pods
kubectl get pods --selector app=main
или kubectl get pods -l app=main
kubectl get pods -l app=main --show-labels
k9s - для визуализации
kubectl logs -f имя_пода
kubectl logs -f имя_пода -c имя_контейнера
kubectl logs -f имя_пода -c имя_контейнера | grep что_искать
watch 'kubectl get pods -l app=main' - чтобы наблюдать
kubectl delete po имя_пода
watch 'kubectl get pods -l app=main -o wide' - более подробная информация
kubectl describe po имя_пода - описание паода
kubectl get po имя_пода -o yaml - более подробное описание пода
kubectl logs -f -c имя_контейнера
kubectl logs -f -c имя_контейнера -p - логи из предыдущего запуска пода
watch kubectl get po -l app=имя_пода
kubectl describe node имя | less
watch kubectl top po
kubectl top nodes
watch kubectl get po -l app=main
kubectl edit deploy main (deployment следит чтобы было запущено нужное кол-во подов)
kubectl scale deploy main --replicas=2
kubectl -n misc port-forward имя_пода 8080:80 (неймспейс misc, 8080 - локальный порт, 80 - удаленный порт)
kubectl -n misc logs -f имя_пода (чтобы смотреть логи)
kubectl config get-contexts
kubectl get nodes
kubectl --context=prod get nodes
kubectl config use-context prod - указание кластера по-умолчанию
watch kubectl get po -n default
kubectl create deployment nginx --namespace default --image=nginx:latest --replicas=2
kubectl delete deployments.apps -n default nginx

RBAC, ServiceAccount, Role, RoleBinding
kubectl create clusterrole readonlyuser
