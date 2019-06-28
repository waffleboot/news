**Сборка**

Внутри Docker через Makefile

```
$ make all
```

**Запуск/Останов**

с помощью docker-compose (docker-compose.yaml и два Dockerfile)

```
$ make up
$ make down
```

**Очередь сообщений**

Redis

**База данных**

Redis

**Процессов**

4 процесса: rest-client, redis-pubsub, storage, redis-keystore

**ID новостей**

UUID

**Messaging**

Request/Reply на очередях "create-request-UUID", "create-reply-UUID", "find-request-UUID", "find-reply-UUID".
Каждой новости своя очередь в Redis.

**Protobuf**

news.proto

**Тесты**

```
curl -v -X POST http://localhost:8000/news -d '{"title":"hellow"}' -H "Content-Type: application/json"
curl -v http://localhost:8000/news/01ca3f73-cde0-4d35-bb8a-96128b00ebb0
```
