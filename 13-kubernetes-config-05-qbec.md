# Домашнее задание к занятию "13.5 поддержка нескольких окружений на примере Qbec"
Приложение обычно существует в нескольких окружениях. Для удобства работы следует использовать соответствующие инструменты, например, Qbec.

## Задание 1: подготовить приложение для работы через qbec
Приложение следует упаковать в qbec. Окружения должно быть 2: stage и production. 

Требования:
* stage окружение должно поднимать каждый компонент приложения в одном экземпляре;
* production окружение — каждый компонент в трёх экземплярах;
* для production окружения нужно добавить endpoint на внешний адрес.
```
Решение
```
Файлы лежат [здесь](https://github.com/AnantaHari/devops-netology/tree/main/13-kubernetes-config-05-qbec/10-demo)  
![Stage](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-18%20в%2018.19.30.png)
![Prod1](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-18%20в%2018.42.07.png)
![Prod2](https://github.com/AnantaHari/devops-netology/blob/main/screenshots/Снимок%20экрана%202022-07-18%20в%2018.45.03.png)

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---

Команды занятия
cd 45-jsonnet/examples
cd 46-qbec/10-demo
kubectl create ns qbec
идем по реадме из 46-qbec
Для инициализации каталога с конфигами - 46-qbec/30-quick-tour.md
qbec apply default
kubectl -n misc exec multitool --curl -s demo.default
