**Сборка**

Внутри Docker через Makefile

*make all*

**Запуск/Останов**

с помощью docker-compose

*make up*
*make down*

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
