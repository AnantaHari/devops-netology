# Домашнее задание к занятию "6.4. PostgreSQL"

## Задача 1

Используя docker поднимите инстанс PostgreSQL (версию 13). Данные БД сохраните в volume.

Подключитесь к БД PostgreSQL используя `psql`.

Воспользуйтесь командой `\?` для вывода подсказки по имеющимся в `psql` управляющим командам.

**Найдите и приведите** управляющие команды для:
- вывода списка БД
```
\l
```
- подключения к БД
```
\c dbname
```
- вывода списка таблиц
```
\dt
```
- вывода описания содержимого таблиц
```
\dS+ table_name
```
- выхода из psql
```
\q
```
```
psql -U pguser -d pgdb
```

## Задача 2

Используя `psql` создайте БД `test_database`.
```
create database test_database;
```
Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-04-postgresql/test_data).

Восстановите бэкап БД в `test_database`.
```
psql -U pguser -d test_database < test_dump.sql
```

Перейдите в управляющую консоль `psql` внутри контейнера.

Подключитесь к восстановленной БД и проведите операцию ANALYZE для сбора статистики по таблице.
```
test_database=# analyze VERBOSE orders;
INFO:  analyzing "public.orders"
INFO:  "orders": scanned 1 of 1 pages, containing 8 live rows and 0 dead rows; 8 rows in sample, 8 estimated total rows
ANALYZE
```

Используя таблицу [pg_stats](https://postgrespro.ru/docs/postgresql/12/view-pg-stats), найдите столбец таблицы `orders` 
с наибольшим средним значением размера элементов в байтах.


**Приведите в ответе** команду, которую вы использовали для вычисления и полученный результат.
```
test_database=# select avg_width from pg_stats where tablename='orders';
 avg_width 
-----------
         4
        16
         4
(3 rows)
```

## Задача 3

Архитектор и администратор БД выяснили, что ваша таблица orders разрослась до невиданных размеров и
поиск по ней занимает долгое время. Вам, как успешному выпускнику курсов DevOps в нетологии предложили
провести разбиение таблицы на 2 (шардировать на orders_1 - price>499 и orders_2 - price<=499).

Предложите SQL-транзакцию для проведения данной операции.
```
test_database=# create table orders_1 (check ( price > 499 )) inherits ( orders );
CREATE TABLE
test_database=# create table orders_2 (check ( price <= 499 )) inherits ( orders );
CREATE TABLE
test_database=# create rule order_price_more_499 as on insert to orders where ( price > 499 ) do instead insert into orders_1 values (new.*);
CREATE RULE
test_database=# create rule order_price_less_499 as on insert to orders where ( price <= 499 ) do instead insert into orders_2 values (new.*);
CREATE RULE
```

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?
```
Если при бизнес-анализе можно было бы понять что предполагается что заказов будет много, то можно было бы сразу сделать разбиение таблицы.
```

## Задача 4

Используя утилиту `pg_dump` создайте бекап БД `test_database`.

Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?
```
pg_dump -U pguser -W test_database > test_database.sql

Для уникальности добавил эти запросы в файл бекапа:
ALTER TABLE public.orders ADD CONSTRAINT unique_orders_title UNIQUE (title);
ALTER TABLE public.orders_1 ADD CONSTRAINT unique_orders_1_title UNIQUE (title);
ALTER TABLE public.orders_2 ADD CONSTRAINT unique_orders_2_title UNIQUE (title);
```
---
Комментарии преподавателя
```
Задание 3
Рассматривали декларативное партиционирование как еще один способ разбиения больших таблиц?

Задание 4
Если делать уникальным поле среди двух таблиц, то CONSTRAINT на одно поле не подходит, так как ограничение соблюдается только в рамках одной таблицы, необходимо добавить второе поле в CONSTRAINT, например, идентификатор записи. CONSTRAINT на поле в эфемерную таблицу orders не сработает, так как в таблице по сути нет данных
```
