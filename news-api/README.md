# news-api

Тестовое задание, заключающееся в предоставлении api для новостей.

Подробное описание находится [здесь](./go-test.md).


## Запуск

### Установка

Скачайте репозиторий и перейдите в папку проекта:

```console
$ git clone https://github.com/danilbushkov/test-tasks.git
```
```console
$ cd test-tasks/news-api

```

### Запуск

```console
$ make run
```

Доступ по `localhost:3000`.

### Остановка

```console
$ make stop
```

## API

### Список API:

- Добавление новости

`POST /add`

Тело запроса:

```json
{
    "Title": "string",
    "Content": "string"
}
```

В ответе содержится id вновь созданной новости в формате:

```json
{
    "Id": "int"
}
```

- Удаление новости

`DELETE /<id>`

- Обновление новости

`POST /edit/<id>`

Тело запроса:

```json
{
    "Id": "int",
    "Title": "string",
    "Content": "string"
    "Categories": "Array<int>"
}
```

- Список новостей

`GET /list`

Пример тела ответа:

```json
{
    "Success": true,
    "News": [
      {
        "Id": 64,
        "Title": "Lorem ipsum",
        "Content": "Dolor sit amet <b>foo</b>",
        "Categories": [1,2,3]
      },
      {
        "Id": 1,
        "Title": "first",
        "Content": "tratata",
        "Categories": [1]
      }
    ]
}
```

### Общий формат ответов

Статус кода успеных ответов: 
- 200 OK
- 201 Created
- 204 No Content

В случае пользовательской ошибки код ответа - 422 Unprocessable Entity c телом ответа 
следующего формата:

```json
{
    "Message": "string"
}
```

При неправильно составленном заросе - 400 Bad Request.

В случае ошибки сервера - 500 Internal Server Error.


