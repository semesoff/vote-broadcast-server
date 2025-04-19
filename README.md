# VOTE-BROADCAST-SERVER

Масштабируемое микросервисное приложение для создания, управления и проведения опросов с обновлением результатов голосования в реальном времени. Построено на современных технологиях для обеспечения безопасности, производительности и простоты развертывания.

## Обзор

**VOTE-BROADCAST-SERVER** упрощает процесс создания и управления опросами, предоставляя интерактивный опыт голосования. Система поддерживает безопасную аутентификацию пользователей, мгновенные обновления результатов и эффективное взаимодействие между сервисами, что делает ее идеальной для онлайн-мероприятий, публичных голосований или внутренних опросов.

## Основные функции

- Создание опросов через API
- Аутентификация на основе JWT
- Обновление данных (отображение опросов и голосов) в реальном времени на основе WebSockets
- Эффективное взаимодействие между сервисами с помощью gRPC
- Контейнеризация и развертывание с помощью Docker

## Стек

- Go
- JWT
- PostgreSQL
- gRPC
- WebSockets
- Docket

## Начало работы

### Требования

- Go
- Docker и Docker Compose
- Git

### Установка

```shell
git clone https://github.com/semesoff/vote-broadcast-server.git
cd vote-broadcast-server
```

### Сборка и запуск 

```shell
docker-compose up --build
```

### API-эндпоинты

Ниже приведены основные эндпоинты REST API и WebSocket с примерами запросов.

### REST API

**Регистрация пользователя**
- `POST /api/register` - Регистрация нового пользователя
- Тело запроса
```json
{
  "username": "string",
  "password": "string"
}
```

**Авторизация пользователя**
- `POST /api/login` - Авторизация пользователя
- Тело запроса
```json
{
  "username": "string",
  "password": "string"
}
```

**Получение списка опросов**
- `GET /api/polls` - Получение списка всех опросов

**Получение информации об опросе**
- `GET /api/polls/{poll_id}` - Получение информации об опросе по ID

**Создание опроса**
- `POST /api/polls` - Создание нового опроса
- Тело запроса
```json
{
  "title": "Poll Title",
  "options": [
    {
      "id": 0,
      "text": "Option 1"
    },
    {
      "id": 0,
      "text": "Option 2"
    }
  ],
  "type": 1
}
```
- Заголовки
```json
Authorization: Bearer <your_jwt_token>
```
- Примечание
- `id` в опциях не должен быть указан, он будет сгенерирован автоматически.
- `type` - тип опроса (1 - возможен выбор одного варианта ответа, 2 - возможен выбор всех)
- 
**Получение голосов**
- `GET /api/votes/{poll_id}` - Получение голосов по ID опроса

**Голосование**
- `POST /api/votes` - Голосование по опросу
- Тело запроса
```json
{
  "poll_id": 24,
  "options_id": [
    40
  ]
}
```
- Примечание
- Если опрос с типом 1, то `options_id` может содержать только один элемент.
- Если опрос с типом 2, то `options_id` может содержать несколько элементов.


### WebSocket-эндпоинты

**Дефолтный адрес** `ws://localhost:5005/`

**Получение списка опросов**
- `/getPolls`

**Получение голосов конкретного опроса**
- `/getVotes/{poll_id}`

## Структура проекта

```text
vote-broadcast-server/
├── init-scripts/ - скрипты инициализации базы данных
├── proto/ - proto-файлы для контрактов
├── services/ - сервисы
│   ├── auth/ - сервис  аутентификации
│   ├── gateway/ - сервис для работы с http-запросами
│   ├── poll/ - сервис опросов
│   ├── vote/ - сервис голосований
│   └── websocket/ - сервис websockets
├── .env - общий конфигурационный файл для сервисов
```