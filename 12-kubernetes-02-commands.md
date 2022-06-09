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
```
apiVersion: apps/v1
kind: Deployment
metadata:
 labels:
   app.kubernetes.io/name: load-balancer-example
 name: hello-world
spec:
 replicas: 5
 selector:
   matchLabels:
     app.kubernetes.io/name: load-balancer-example
 template:
   metadata:
     labels:
       app.kubernetes.io/name: load-balancer-example
   spec:
     containers:
     - image: gcr.io/google-samples/node-hello:1.0
       name: hello-world
       ports:
       - containerPort: 8080

:~$ kubectl apply -f ./load-balancer-example.yaml
deployment.apps/hello-world created

:~$ kubectl get deployments hello-world
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
hello-world   0/5     5            0           7s

:~$ kubectl scale deploy hello-world --replicas=2
deployment.apps/hello-world scaled

:~$ kubectl get deployments hello-world
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
hello-world   2/2     2            2           103s

:~$ kubectl get deployment
NAME          READY   UP-TO-DATE   AVAILABLE   AGE
hello-world   2/2     2            2           4m41s

kubectl get pods
NAME                           READY   STATUS    RESTARTS   AGE
hello-world-6755976cfc-cx5hr   1/1     Running   0          5m6s
hello-world-6755976cfc-q2mm6   1/1     Running   0          5m6s
```

## Задание 2: Просмотр логов для разработки
Разработчикам крайне важно получать обратную связь от штатно работающего приложения и, еще важнее, об ошибках в его работе.
Требуется создать пользователя и выдать ему доступ на чтение конфигурации и логов подов в app-namespace.

Требования:
 * создан новый токен доступа для пользователя
 * пользователь прописан в локальный конфиг (~/.kube/config, блок users)
 * пользователь может просматривать логи подов и их конфигурацию (kubectl logs pod <pod_id>, kubectl describe pod <pod_id>)
```
Делал по этому примеру:
Create a ServiceAccount, say 'readonlyuser'.
kubectl create serviceaccount readonlyuser

Create cluster role, say 'readonlyuser'.
kubectl create clusterrole readonlyuser --verb=get --verb=list --verb=watch --resource=pods --resource=pods/log

Create cluster role binding, say 'readonlyuser'.
kubectl create clusterrolebinding readonlyuser --serviceaccount=default:readonlyuser --clusterrole=readonlyuser

Now get the token from secret of ServiceAccount we have created before. we will use this token to authenticate user.
TOKEN=$(kubectl describe secrets "$(kubectl describe serviceaccount readonlyuser | grep -i Tokens | awk '{print $2}')" | grep token: | awk '{print $2}')

Now set the credentials for the user in kube config file. I am using 'vikash' as username.
kubectl config set-credentials vikash --token=$TOKEN

Now Create a Context say podreader. I am using my clustername 'kubernetes' here.
kubectl config set-context podreader --cluster=kubernetes --user=vikash

Finally use the context .
kubectl config use-context podreader

And that's it. Now one can execute kubectl get pods --all-namespaces. One can also check the access by executing as given:
:~$ kubectl auth can-i get pods --all-namespaces
yes
:~$ kubectl auth can-i create pods
no
:~$ kubectl auth can-i delete pods
no
:~$ kubectl auth can-i get pod/logs
yes
:~$ kubectl logs hello-world-6755976cfc-2fzp7
:~$ kubectl describe pod hello-world-6755976cfc-2fzp7
Name:         hello-world-6755976cfc-2fzp7
Namespace:    default
Priority:     0
Node:         minikube/192.168.49.2
Start Time:   Thu, 09 Jun 2022 18:03:20 +0300
Labels:       app.kubernetes.io/name=load-balancer-example
              pod-template-hash=6755976cfc
Annotations:  <none>
Status:       Running
IP:           172.17.0.4
IPs:
  IP:           172.17.0.4
Controlled By:  ReplicaSet/hello-world-6755976cfc
Containers:
  hello-world:
    Container ID:   docker://5eff30739fc621e7581ccfa7e84de2b62e5a6696bb6c17aa93dc454d73742619
    Image:          gcr.io/google-samples/node-hello:1.0
    Image ID:       docker-pullable://gcr.io/google-samples/node-hello@sha256:d238d0ab54efb76ec0f7b1da666cefa9b40be59ef34346a761b8adc2dd45459b
    Port:           8080/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Thu, 09 Jun 2022 18:03:21 +0300
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-g2xsq (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  kube-api-access-g2xsq:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:                      <none>
```


## Задание 3: Изменение количества реплик
Поработав с приложением, вы получили запрос на увеличение количества реплик приложения для нагрузки. Необходимо изменить запущенный deployment, увеличив количество реплик до 5. Посмотрите статус запущенных подов после увеличения реплик.

Требования:
 * в deployment из задания 1 изменено количество реплик на 5
 * проверить что все поды перешли в статус running (kubectl get pods)
```
:~$ kubectl scale deploy hello-world --replicas=5
deployment.apps/hello-world scaled

:~$ kubectl get pods
NAME                           READY   STATUS    RESTARTS   AGE
hello-world-6755976cfc-2fzp7   1/1     Running   0          24s
hello-world-6755976cfc-cx5hr   1/1     Running   0          6m32s
hello-world-6755976cfc-l2nzd   1/1     Running   0          24s
hello-world-6755976cfc-nc4r6   1/1     Running   0          24s
hello-world-6755976cfc-q2mm6   1/1     Running   0          6m32s
```
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
