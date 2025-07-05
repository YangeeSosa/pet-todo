.PHONY: help build test lint clean docker-build docker-up docker-down

# Переменные
API_DIR=./api
DB_DIR=./db
KAFKA_DIR=./kafka

# Команды по умолчанию
help: ## Показать справку
	@echo "Доступные команды:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Сборка
build: ## Собрать все сервисы
	@echo "Сборка API-сервиса..."
	cd $(API_DIR) && go build -o bin/api ./cmd
	@echo "Сборка БД-сервиса..."
	cd $(DB_DIR) && go build -o bin/db ./cmd
	@echo "Сборка Kafka-сервиса..."
	cd $(KAFKA_DIR) && go build -o bin/kafka ./cmd

# Тесты
test: ## Запустить все тесты
	@echo "Тесты API-сервиса..."
	cd $(API_DIR) && go test -v ./...
	@echo "Тесты БД-сервиса..."
	cd $(DB_DIR) && go test -v ./...
	@echo "Тесты Kafka-сервиса..."
	cd $(KAFKA_DIR) && go test -v ./...

test-api: ## Тесты только API-сервиса
	cd $(API_DIR) && go test -v ./...

test-db: ## Тесты только БД-сервиса
	cd $(DB_DIR) && go test -v ./...

# Линтер
lint: ## Запустить линтер для всех сервисов
	@echo "Линтер API-сервиса..."
	cd $(API_DIR) && golangci-lint run
	@echo "Линтер БД-сервиса..."
	cd $(DB_DIR) && golangci-lint run
	@echo "Линтер Kafka-сервиса..."
	cd $(KAFKA_DIR) && golangci-lint run

lint-api: ## Линтер только API-сервиса
	cd $(API_DIR) && golangci-lint run

lint-db: ## Линтер только БД-сервиса
	cd $(DB_DIR) && golangci-lint run

# Docker команды
docker-build: ## Собрать Docker образы
	docker-compose build

docker-up: ## Запустить все сервисы через Docker
	docker-compose up -d

docker-down: ## Остановить все сервисы
	docker-compose down

docker-logs: ## Показать логи всех сервисов
	docker-compose logs -f

# Очистка
clean: ## Очистить собранные файлы
	@echo "Очистка API-сервиса..."
	cd $(API_DIR) && rm -rf bin/
	@echo "Очистка БД-сервиса..."
	cd $(DB_DIR) && rm -rf bin/
	@echo "Очистка Kafka-сервиса..."
	cd $(KAFKA_DIR) && rm -rf bin/

# Полная проверка
check: lint test ## Запустить линтер и тесты
	@echo "Все проверки пройдены!"

# Запуск локально (без Docker)
run-api: ## Запустить API-сервис локально
	cd $(API_DIR) && go run ./cmd

run-db: ## Запустить БД-сервис локально
	cd $(DB_DIR) && go run ./cmd

run-kafka: ## Запустить Kafka-сервис локально
	cd $(KAFKA_DIR) && go run ./cmd