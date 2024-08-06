# Импорт товаров

Задание представлено [здесь](./test_request_laravel_php_v2.docx).

## Зависимости

- `docker`
- `docker compose`
- `make`

## Установка

Скачайте репозиторий и перейдите в папку проекта:

```console
$ git clone https://github.com/danilbushkov/test-tasks.git
$ cd test-tasks/import-products

```


Для сборки выполните:

```console
$ make build
```

Для выполнения команды может потребоваться права суперпользователя.

## Запуск

```console
$ make run
```

Если не выполнена миграция базы данных, то выполните:

```console
$ make migrate
```

Для просмотра перейдите по `localhost:8080`.

## Остановка

```console
$ make stop
```



