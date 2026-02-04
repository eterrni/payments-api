# Локальный запуск

1. Запустите PostgreSQL:
   ```bash
   docker compose -f local/docker-compose.yml up -d
   ```

2. Задайте переменные окружения (из корня проекта):
   ```bash
   export DB_DSN="host=localhost port=5432 user=postgres password=postgres dbname=payments sslmode=disable"
   ```
   Либо скопируйте `local/.env.example` в `local/.env`, отредактируйте при необходимости и выполните:
   ```bash
   set -a && . local/.env && set +a
   ```

3. Запустите приложение:
   ```bash
   go run ./cmd
   ```

API доступен на http://localhost:8080

Остановка PostgreSQL: `docker compose -f local/docker-compose.yml down`
