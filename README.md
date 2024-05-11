# car-service-api

Api для добавления машин по их номеру

### Stack
  * Golang 1.22
  * Fiber
  * Swagger
  * Postgres
  * Docker

## Возможные действия:
  * Добавление машин
  * Получение списка всех машин с пагинацией и фильтрацией
  * Редактирование машины по одному или нескольким полям
  * Удаление машины по идентификатору

## Запуск проекта
  Необходимо склонировать репозиторий:
  ```
     git clone github.com/sh1neqd/car-service-api
     git checkout main
```
  Далее запускаем проект с помощью docker-compose:
```
  docker-compose build
  docker-compose up
```
  Запускаем Postman и теституем)

## Документация

Посмотреть swagger-документацию можно по запросу
```
http://localhost:8000/swagger/
```

### docs

* http://localhost:8000/api/ [post] adding 1 or more cars:
```
{
    "reg_nums":[reg_nums]
}
```
* http://localhost:8000/api/ [get] getting list of all cars
* http://localhost:8000/api/{id} [get] getting car by id
* http://localhost:8000/api/{id} [patch] editing car info by 1 or more fields:
```
{
    "regNum":reg_num,
    "mark":mark,
    "model":model
}
```
* http://localhost:8000/api/{id} [delete] deleting car by id
