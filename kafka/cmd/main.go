package main

import (
	"log"
	"os"

	"github.com/YangeeSosa/pet-todo-kafka/internal"
)

func main() {
	logFile, err := os.OpenFile("kafka.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Не удалось открыть файл для логов: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(os.Stdout)

	log.Println("Запуск Kafka-сервиса...")

	consumer := internal.NewConsumer()
	consumer.StartConsuming()
}
