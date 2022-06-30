# Домашнее задание к занятию "12.5 Сетевые решения CNI"
После работы с Flannel появилась необходимость обеспечить безопасность для приложения. Для этого лучше всего подойдет Calico.
## Задание 1: установить в кластер CNI плагин Calico
Для проверки других сетевых решений стоит поставить отличный от Flannel плагин — например, Calico. Требования:
* установка производится через ansible/kubespray;
* после применения следует настроить политику доступа к hello-world извне. Инструкции [kubernetes.io](https://kubernetes.io/docs/concepts/services-networking/network-policies/), [Calico](https://docs.projectcalico.org/about/about-network-policy)

```
20-network-policy git:(master) ✗ kubectl apply -f ./manifests/main/
deployment.apps/frontend created
service/frontend created
deployment.apps/backend created
service/backend created
deployment.apps/cache created

service/cache created
➜  20-network-policy git:(master) ✗ kubectl get po
NAME                       READY   STATUS              RESTARTS   AGE
backend-869fd89bdc-9fpjk   0/1     ContainerCreating   0          6s
cache-b7cbd9f8f-q7ljp      0/1     ContainerCreating   0          6s
frontend-c74c5646c-zj2bn   0/1     ContainerCreating   0          6s
➜  20-network-policy git:(master) ✗ kubectl get networkpolicies
No resources found in default namespace.

Настраиваем правила чтобы фронт мог ходить в бек, а бек в кеш
➜  20-network-policy git:(master) ✗ kubectl apply -f ./manifests/network-policy/10-frontend.yaml
networkpolicy.networking.k8s.io/frontend created
➜  20-network-policy git:(master) ✗ kubectl apply -f ./manifests/network-policy/20-backend.yaml
networkpolicy.networking.k8s.io/backend created
➜  20-network-policy git:(master) ✗ kubectl apply -f ./manifests/network-policy/30-cache.yaml
networkpolicy.networking.k8s.io/cache created

Проверяем что бек ходит только в кеш:
➜  20-network-policy git:(master) ✗ kubectl exec backend-869fd89bdc-9fpjk -- curl -s -m 1 cache
Praqma Network MultiTool (with NGINX) - cache-b7cbd9f8f-q7ljp - 10.233.90.2
➜  20-network-policy git:(master) ✗ kubectl exec backend-869fd89bdc-9fpjk -- curl -s -m 1 backend
command terminated with exit code 28
➜  20-network-policy git:(master) ✗ kubectl exec backend-869fd89bdc-9fpjk -- curl -s -m 1 frontend
command terminated with exit code 28

Фронт только в бек:
➜  20-network-policy git:(master) ✗ kubectl exec frontend-c74c5646c-zj2bn -- curl -s -m 1 backend
Praqma Network MultiTool (with NGINX) - backend-869fd89bdc-9fpjk - 10.233.96.2
➜  20-network-policy git:(master) ✗ kubectl exec frontend-c74c5646c-zj2bn -- curl -s -m 1 cache
command terminated with exit code 28
➜  20-network-policy git:(master) ✗ kubectl exec frontend-c74c5646c-zj2bn -- curl -s -m 1 frontend
command terminated with exit code 28

Кеш никуда не ходит
➜  20-network-policy git:(master) ✗ kubectl exec cache-b7cbd9f8f-q7ljp -- curl -s -m 1 backend
command terminated with exit code 28
➜  20-network-policy git:(master) ✗ kubectl exec cache-b7cbd9f8f-q7ljp -- curl -s -m 1 frontend
command terminated with exit code 28
➜  20-network-policy git:(master) ✗ kubectl exec cache-b7cbd9f8f-q7ljp -- curl -s -m 1 cache
command terminated with exit code 28
```

## Задание 2: изучить, что запущено по умолчанию
Самый простой способ — проверить командой calicoctl get <type>. Для проверки стоит получить список нод, ipPool и profile.
Требования:
* установить утилиту calicoctl;
* получить 3 вышеописанных типа в консоли.
```
➜  bin calicoctl get ipPool --allow-version-mismatch
NAME           CIDR             SELECTOR

default-pool   10.233.64.0/18   all()

➜  bin calicoctl get nodes --allow-version-mismatch
NAME
cp1
node1
node2

➜  bin calicoctl get profile --allow-version-mismatch
NAME
projectcalico-default-allow
kns.default
kns.kube-node-lease
kns.kube-public
kns.kube-system
ksa.default.default
ksa.kube-node-lease.default
ksa.kube-public.default
ksa.kube-system.attachdetach-controller
ksa.kube-system.bootstrap-signer
ksa.kube-system.calico-node
ksa.kube-system.certificate-controller
ksa.kube-system.clusterrole-aggregation-controller
ksa.kube-system.coredns
ksa.kube-system.cronjob-controller
ksa.kube-system.daemon-set-controller
ksa.kube-system.default
ksa.kube-system.deployment-controller
ksa.kube-system.disruption-controller
ksa.kube-system.dns-autoscaler
ksa.kube-system.endpoint-controller
ksa.kube-system.endpointslice-controller
ksa.kube-system.endpointslicemirroring-controller
ksa.kube-system.ephemeral-volume-controller
ksa.kube-system.expand-controller
ksa.kube-system.generic-garbage-collector
ksa.kube-system.horizontal-pod-autoscaler
ksa.kube-system.job-controller
ksa.kube-system.kube-proxy
ksa.kube-system.namespace-controller
ksa.kube-system.node-controller
ksa.kube-system.nodelocaldns
ksa.kube-system.persistent-volume-binder
ksa.kube-system.pod-garbage-collector
ksa.kube-system.pv-protection-controller
ksa.kube-system.pvc-protection-controller
ksa.kube-system.replicaset-controller
ksa.kube-system.replication-controller
ksa.kube-system.resourcequota-controller
ksa.kube-system.root-ca-cert-publisher
ksa.kube-system.service-account-controller
ksa.kube-system.service-controller
ksa.kube-system.statefulset-controller
ksa.kube-system.token-cleaner
ksa.kube-system.ttl-after-finished-controller
ksa.kube-system.ttl-controller
```

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

cd 16-networking/20-network-policy/
kubectl -n default  get po
kubectl apply -f ./manifests/main/
kubectl -n default  get po
выполняем ### Проверка доступности между подами
kubectl apply -f ./manifests/network-policy/00-default.yaml
kubectl get networkpolicies
kubectl apply -f ./manifests/network-policy/20-backend.yaml
