# Pet Todo - Микросервисное приложение на Go

Проект демонстрирует архитектуру микросервисов с использованием Go, PostgreSQL, Docker и Kafka.

##  Архитектура

- **API-сервис** (порт 8080) - принимает HTTP-запросы от клиентов
- **БД-сервис** (порт 8081) - работает с PostgreSQL
- **Kafka-сервис** - читает события из Kafka и логирует их
- **PostgreSQL** - основная база данных
- **Kafka + Zookeeper** - система событий

## Быстрый старт

### Через Docker Compose (рекомендуется)

```bash
# Запуск всех сервисов
make docker-up

# Просмотр логов
make docker-logs

# Остановка
make docker-down
```

### Локальный запуск

```bash
# Установка зависимостей
go mod download

# Запуск сервисов
make run-api
make run-db
make run-kafka
```

## API Endpoints

### Создание задачи
```bash
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Купить хлеб"}' \
  http://localhost:8080/tasks
```

### Получение списка задач
```bash
curl http://localhost:8080/tasks
```

### Отметка задачи как выполненной
```bash
curl -X PUT -H "Content-Type: application/json" \
  -d '{"id":"task-id"}' \
  http://localhost:8080/tasks
```

### Удаление задачи
```bash
curl -X DELETE -H "Content-Type: application/json" \
  -d '{"id":"task-id"}' \
  http://localhost:8080/tasks
```

## Тестирование

```bash
# Все тесты
make test

# Тесты конкретного сервиса
make test-api
make test-db
```

## Линтинг

```bash
# Линтер для всех сервисов
make lint

# Линтер конкретного сервиса
make lint-api
make lint-db
```

## Мониторинг

### Проверка статуса сервисов
```bash
docker-compose ps
```

### Логи сервисов
```bash
# Все логи
docker-compose logs

# Конкретный сервис
docker-compose logs api-service
docker-compose logs db-service
docker-compose logs kafka-service
```

### Проверка базы данных
```bash
docker-compose exec postgres psql -U postgres -d pet_todo -c "SELECT * FROM tasks;"
```

## Структура проекта

```
pet-todo/
├── api/                    # API-сервис
│   ├── cmd/main.go        # Точка входа
│   ├── internal/          # Внутренняя логика
│   └── Dockerfile
├── db/                    # БД-сервис
│   ├── cmd/main.go        # Точка входа
│   ├── internal/          # Внутренняя логика
│   └── Dockerfile
├── kafka/                 # Kafka-сервис
│   ├── cmd/main.go        # Точка входа
│   ├── internal/          # Внутренняя логика
│   └── Dockerfile
├── docker-compose.yaml    # Конфигурация Docker
├── Makefile              # Команды для разработки
├── .golangci.yml         # Конфигурация линтера
└── README.md             # Документация
```

## Разработка

### Добавление новых функций
1. Создай feature branch
2. Реализуй функциональность
3. Добавь тесты
4. Запусти линтер: `make lint`
5. Запусти тесты: `make test`
6. Создай pull request

### Локальная разработка
```bash
# Сборка
make build

# Запуск тестов
make test

# Линтинг
make lint

# Полная проверка
make check
```

## События

Система отправляет события в Kafka при следующих операциях:
- `task_created` - создание задачи
- `task_updated` - обновление задачи
- `task_deleted` - удаление задачи
- `task_completed` - отметка как выполненной

Kafka-сервис читает эти события и логирует их в файл.

## Технологии

- **Go 1.22+** - основной язык
- **Gin** - HTTP фреймворк
- **PostgreSQL** - база данных
- **Kafka** - система событий
- **Docker** - контейнеризация
- **Make** - автоматизация
- **golangci-lint** - линтер


