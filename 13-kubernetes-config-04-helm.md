# Домашнее задание к занятию "13.4 инструменты для упрощения написания конфигурационных файлов. Helm и Jsonnet"
В работе часто приходится применять системы автоматической генерации конфигураций. Для изучения нюансов использования разных инструментов нужно попробовать упаковать приложение каждым из них.

## Задание 1: подготовить helm чарт для приложения
Необходимо упаковать приложение в чарт для деплоя в разные окружения. Требования:
* каждый компонент приложения деплоится отдельным deployment’ом/statefulset’ом;
* в переменных чарта измените образ приложения для изменения версии.

## Задание 2: запустить 2 версии в разных неймспейсах
Подготовив чарт, необходимо его проверить. Попробуйте запустить несколько копий приложения:
* одну версию в namespace=app1;
* вторую версию в том же неймспейсе;
* третью версию в namespace=app2.

## Задание 3 (*): повторить упаковку на jsonnet
Для изучения другого инструмента стоит попробовать повторить опыт упаковки из задания 1, только теперь с помощью инструмента jsonnet.

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

Команды урока
cd 40-helm/01-templating/charts
идем по файлу README
diff <(helm template 01-simple) <(helm template --set image.tag=1.18,replicaCount=2 01-simple)
{{- toYaml .Values.resources | nindent 12 }} - nindent 12 добавляет 12 отступов
toYaml - преобразует к виду Yaml
cd 40-helm/02-package-manager
идем по README
cd charts
прошлись по файлам
cd 40-helm/03-deploy
смотрим в README
cd 40-helm/01-templating/
helm install demo-release charts/01-simple
kubectl get deploy demo -o jsonpath={.spec.template.spec.containers[0].image}
helm upgrade demo-release charts/01-simple
helm upgrade demo-release -f charts/01-simple/new-values2.yaml charts/01-simple
helm get values demo-release
