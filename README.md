# Segment Service

Сервис для управления сегментами пользователей и проведения экспериментов.

## Возможности
- Создание, удаление и редактирование сегментов
- Добавление и удаление пользователей в сегменты
- Случайное распределение сегмента на процент пользователей
- Получение сегментов пользователя
- Получение пользователей по сегменту

## Быстрый старт

1. **Запуск сервера:**
   ```sh
   go run cmd/service-segment/main.go
   ```

2. **Health-check:**
   ```sh
   curl http://localhost:8080/segment
   # Ответ: OK
   ```

## Примеры запросов

### Создать пользователя
```sh
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"user_id": 1}'
```

### Создать сегмент
```sh
curl -X POST http://localhost:8080/segments -H "Content-Type: application/json" -d '{"name": "MAIL_GPT", "description": "Использование GPT в письмах", "distribution_ratio": 0.0}'
```

### Добавить пользователя в сегмент
```sh
curl -X POST http://localhost:8080/users/1/segments -H "Content-Type: application/json" -d '{"segment_name": "MAIL_GPT"}'
```

### Получить сегменты пользователя
```sh
curl http://localhost:8080/users/1/segments
```

### Получить пользователей сегмента
```sh
curl http://localhost:8080/segments/MAIL_GPT/users
```

### Случайно распределить сегмент на процент пользователей
```sh
curl -X POST http://localhost:8080/segments/MAIL_GPT/assign_random -H "Content-Type: application/json" -d '{"ratio": 0.3}'
```

### Редактировать сегмент
```sh
curl -X PATCH http://localhost:8080/segments/MAIL_GPT -H "Content-Type: application/json" -d '{"description": "Новое описание", "distribution_ratio": 0.5}'
```

### Удалить сегмент
```sh
curl -X DELETE http://localhost:8080/segments/MAIL_GPT
```

## Структура проекта
- `cmd/service-segment/main.go` — точка входа
- `internal/models/` — бизнес-логика и работа с БД
- `internal/handlers/` — HTTP-обработчики
- `internal/router/` — маршрутизация
- `internal/db/` — инициализация и схема БД

## Требования
- Go 1.18+
- SQLite (создаётся автоматически)