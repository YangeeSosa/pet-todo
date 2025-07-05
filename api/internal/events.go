package internal

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
)

type TaskEvent struct {
	Type      string    `json:"type"`
	TaskID    string    `json:"task_id"`
	Title     string    `json:"title"`
	Timestamp time.Time `json:"timestamp"`
}

var producer sarama.SyncProducer

func InitKafkaProducer() {
	kafkaURL := os.Getenv("KAFKA_URL")
	if kafkaURL == "" {
		log.Println("KAFKA_URL не указан")
		return
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	var err error
	producer, err = sarama.NewSyncProducer([]string{kafkaURL}, config)
	if err != nil {
		log.Fatalf("Ошибка при создании producer: %v", err)
		return
	}

	log.Println("Kafka producer инициализирован")
}

func SendTaskEvent(eventType, taskID, title string) {
	if producer == nil {
		log.Println("Producer не инициализирован")
		return
	}

	event := TaskEvent{
		Type:      eventType,
		TaskID:    taskID,
		Title:     title,
		Timestamp: time.Now(),
	}

	eventJSON, err := json.Marshal(event)
	if err != nil {
		log.Printf("Ошибка при маршалинге события: %v", err)
		return
	}
	msg := &sarama.ProducerMessage{
		Topic: "task_events",
		Value: sarama.StringEncoder(eventJSON),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Ошибка при отправке события: %v", err)
		return
	}

	log.Printf("Событие отправлено: Partition=%d, Offset=%d, type=%s, task_id=%s", partition, offset, eventType, taskID)
}
