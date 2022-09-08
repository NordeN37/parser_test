# parser_test

Тестовый сервер для запуска парсинга.

## Build

    cd cmd/server && go build -o parser_test

## Run

    ./parser_test

## docker

    $ docker-compose up -d

## Параметры окружения для запуска

| Параметр    | Описание            | По умолчанию |
|-------------|---------------------|--------------|
| PARSER_HOST | Хост запуска сервера | localhost    |
| PARSER_PORT | Порт запуска сервера | 8080         |
| MONGO_HOST | Хост запуска MONGO  | localhost    |
| MONGO_PORT | Порт запуска MONGO  | 27017        |
| MONGO_DB_NAME | Название базы mongo | parser_test  |
| MONGO_USERNAME | Имя пользователя    | admin        |
| MONGO_PASSWORD | Пароль пользователя | admin        |

## Пример ENV для старта сервера

```Bash
export PARSER_HOST=localhost
export PARSER_PORT=8080
```