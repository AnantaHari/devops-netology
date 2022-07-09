# Домашнее задание к занятию "13.2 разделы и монтирование"
Приложение запущено и работает, но время от времени появляется необходимость передавать между бекендами данные. А сам бекенд генерирует статику для фронта. Нужно оптимизировать это.
Для настройки NFS сервера можно воспользоваться следующей инструкцией (производить под пользователем на сервере, у которого есть доступ до kubectl):
* установить helm: curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash
* добавить репозиторий чартов: helm repo add stable https://charts.helm.sh/stable && helm repo update
* установить nfs-server через helm: helm install nfs-server stable/nfs-server-provisioner

В конце установки будет выдан пример создания PVC для этого сервера.

## Задание 1: подключить для тестового конфига общую папку
В stage окружении часто возникает необходимость отдавать статику бекенда сразу фронтом. Проще всего сделать это через общую папку. Требования:
* в поде подключена общая папка между контейнерами (например, /static);
* после записи чего-либо в контейнере с беком файлы можно получить из контейнера с фронтом.

```
Решение
```


## Задание 2: подключить общую папку для прода
Поработав на stage, доработки нужно отправить на прод. В продуктиве у нас контейнеры крутятся в разных подах, поэтому потребуется PV и связь через PVC. Сам PV должен быть связан с NFS сервером. Требования:
* все бекенды подключаются к одному PV в режиме ReadWriteMany;
* фронтенды тоже подключаются к этому же PV с таким же режимом;
* файлы, созданные бекендом, должны быть доступны фронту.

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

cd 20-concepts/30-storage/10-volume
kubectl apply -f manifests/10-pod-int-volumes.yaml
kubectl get pods
kubectl exec pod-int-volumes -c nginx -- sh -c "echo '42' > /static/42.txt"
kubectl exec pod-int-volumes -c nginx -- ls -la /static
kubectl exec pod-int-volumes -c busybox -- ls -la /tmp/cache
kubectl exec pod-int-volumes -c busybox -- touch /tmp/cache/43.56
kubectl exec pod-int-volumes -c nginx -- ls -la /static
kubectl exec pod-int-volumes -c nginx -- dd if=/dev/zero of=/static/10mb.txt bs=1M count=10
kubectl get po -o wide
kubectl get po pod-int-volumes -o yaml | grep uid

# Заходим на эту ноду
ssh на ноду
find /var/lib/kubelet -name 42.txt
ls -la found-folder-name
find /var/lib/kubelet/ -name my-volume | grep volumes
find /var/lib/kubelet/ -name my-volume | grep volumes | xargs ls -la

Найденная папка вида:
`/var/lib/kubelet/pods/d7ecc26f-1d5e-4a91-99a1-3cb14dadf71b/volumes/kubernetes.io~empty-dir/my-volume/`

Где `d7ecc26f-1d5e-4a91-99a1-3cb14dadf71b` - это uid пода.
`my-volume` - имя тома в поде.

#### Удаляем под и проверяем наличие файлов.
```shell script
kubectl delete pod pod-int-volumes

# Смотрим на содержимое найденной папки. Нужно скопировать имя папки
ls -la found-folder-name
```
cd 20-concepts/30-storage/20-persistent-volume/10-manual/
выполнить все команды по readme
cd 20-concepts/30-storage/20-persistent-volume/20-delete/
выполнить все команды по readme
cd 20-concepts/30-storage/20-persistent-volume/40-dynamic-provisioning/
# Посмотреть список StorageClass
kubectl get storageclasses.storage.k8s.io
# Посмотреть список StorageClass (короткое имя)
kubectl get sc
# Показывает ноды, к которым могут быть примонтированы тома
kubectl get csinodes
# CSIDriver
kubectl get csidrivers
```shell script
watch 'kubectl get po,pvc,pv'
```

```shell script
kubectl apply -f manifests/10-pod.yaml
kubectl apply -f manifests/20-pvc.yaml
```
После этого будет создан PersistentVolume.
Создана связка PersistentVolumeClaim-PersistentVolume.
И запущен Pod.

```shell script
kubectl exec pod -- ls -la /static
kubectl exec pod -- sh -c "echo 'dynamic' > /static/dynamic.txt"

# Определим в какой папке у нас хранятся данные
kubectl get pv -o yaml | grep '^\s*path:'

# На ноде ищем файл. Например
sudo ls -la /var/snap/microk8s/common/default-storage/default-pvc-pvc-7bd66d4c-189e-44d1-ad0f-bc091491525e
```

Все работает как ожидалось.

## Удаление ненужного тома
В случае динамического создания томов их имена будут содержать uid PersistentVolumeClaim, который явился причиной создания PersistentVolume.

```shell script
# uid PVC
kubectl get pvc pvc -o yaml | grep uid

# название PV, связанного с данным PVC
kubectl get pvc pvc -o yaml | grep '^\s*volumeName'

# Удаление pvc
kubectl delete pod pod
kubectl delete pvc pvc
```

После удаления пода и PersistentVolumeClaim будет удален и PersistentVolume.
Все данные тоже будут удалены. Так как они больше не нужны.

## StorageClass
[Документация](https://kubernetes.io/docs/concepts/storage/storage-classes/)

StorageClass - это объект, который определяет параметры подключенного тома.
StorageClass указывается явно или неявно в спецификации PersistentVolumeClaim

```shell script
# Посмотреть список StorageClass
kubectl get storageclasses.storage.k8s.io

# Посмотреть список StorageClass (короткое имя)
kubectl get sc
```

Это выдал helm
```
---
    kind: PersistentVolumeClaim
    apiVersion: v1
    metadata:
      name: test-dynamic-volume-claim
    spec:
      storageClassName: "nfs"
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: 100Mi

```
