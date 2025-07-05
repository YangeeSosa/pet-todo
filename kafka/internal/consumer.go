package internal

import (
	"log"
	"os"
	"time"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumer sarama.Consumer
}

func NewConsumer() *Consumer {
	kafkaURL := os.Getenv("KAFKA_URL")
	if kafkaURL == "" {
		kafkaURL = "localhost:9092"
	}

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	for i := 0; i < 30; i++ {
		consumer, err := sarama.NewConsumer([]string{kafkaURL}, config)
		if err == nil {
			log.Printf("Подключение к Kafka успешно: %s", kafkaURL)
			return &Consumer{consumer: consumer}
		}
		log.Printf("Попытка подключения к Kafka %d/30: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	log.Fatal("Не удалось подключиться к Kafka после 30 попыток")
	return nil
}

func (c *Consumer) StartConsuming() {
	topic := "task_events"

	partitionConsumer, err := c.consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Ошибка создания partition consumer: %v", err)
	}
	defer partitionConsumer.Close()

	log.Printf("Начинаем чтение событий из топика: %s", topic)

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Получено событие: Topic=%s, Partition=%d, Offset=%d, Key=%s, Value=%s",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Printf("Ошибка чтения события: %v", err)
		}
	}
}
