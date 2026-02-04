# Payments API

REST API для управления платежами. Использует Go, Gorilla Mux и PostgreSQL (GORM).

## Требования

- Go 1.25+
- PostgreSQL

## Переменные окружения

| Переменная | Описание |
|------------|----------|
| `DB_DSN`   | Строка подключения к PostgreSQL (обязательно). Пример: `host=localhost user=postgres password=postgres dbname=payments sslmode=disable` |

## Запуск локально

Вариант с готовой конфигурацией (PostgreSQL в Docker и пример env) — см. **[local/README.md](local/README.md)**.

Минимальный запуск:

```bash
export DB_DSN="host=localhost user=postgres password=postgres dbname=payments sslmode=disable"
go run ./cmd
```

Сервер слушает порт **8080**.

## Запуск с Docker

Сборка образа:

```bash
docker build -t payments-api .
```

Запуск (база должна быть доступна по сети):

```bash
docker run -p 8080:8080 \
  -e DB_DSN="host=host.docker.internal user=postgres password=postgres dbname=payments sslmode=disable" \
  payments-api
```

Для Linux вместо `host.docker.internal` укажите IP хоста или имя сервиса БД в docker-сети.

## API

| Метод   | Путь            | Описание              |
|---------|-----------------|------------------------|
| POST    | `/payments`     | Создать платёж         |
| GET     | `/payments/{id}`| Получить платёж по ID  |
| PUT     | `/payments/{id}`| Обновить платёж        |
| DELETE  | `/payments/{id}`| Удалить платёж         |

### Примеры

**Создать платёж (POST /payments):**

```json
{
  "amount": 100.50,
  "currency": "USD"
}
```

**Обновить платёж (PUT /payments/{id}):** тело запроса — такой же JSON.

## Структура проекта

```
cmd/                 — точка входа
local/               — локальный запуск (docker-compose PostgreSQL, .env.example)
internal/
  handlers/          — HTTP-обработчики
  repository/        — работа с БД
  services/          — бизнес-логика
pkg/
  middleware/        — логирование, recovery
  utils/             — ответы JSON
```
