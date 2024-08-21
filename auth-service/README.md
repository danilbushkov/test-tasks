# news-api

Тестовое задание, заключающееся в предоставлении api для генерации и обновления токенов.

Подробное описание находится [здесь](./go-test.md).


## Запуск

### Установка

Скачайте репозиторий и перейдите в папку проекта:

```console
$ git clone https://github.com/danilbushkov/test-tasks.git
```
```console
$ cd test-tasks/auth-service

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

- Предоставление токенов

`POST /api/auth/get`

Тело запроса:

```json
{
    "uuid": "string"
}
```

В ответе содержатся токены:

```json
{
    "accessToken": "string",
    "refreshToken": "string"
}
```

- Обновление токенов

`POST /api/auth/refresh`

Тело запроса:

```json
{
    "refreshToken": "string"
}
```

В ответе содержатся токены:

```json
{
    "accessToken": "string",
    "refreshToken": "string"
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


