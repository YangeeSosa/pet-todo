package internal

import (
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@postgres:5432/pet_todo?sslmode=disable"
	}

	var err error
	for i := 0; i < 30; i++ {
		log.Printf("Попытка подключения к БД %d/30...", i+1)
		DB, err = sqlx.Connect("postgres", dsn)
		if err == nil {
			break
		}
		log.Printf("Ошибка подключения: %v", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("Не удалось подключиться к БД после 30 попыток: %v", err)
	}

	log.Println("Успешно подключились к PostgreSQL")
	createTable()
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS tasks (
        id UUID PRIMARY KEY,
        title TEXT NOT NULL,
        done BOOLEAN NOT NULL DEFAULT FALSE
    );`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Не удалось создать таблицу: %v", err)
	}
	log.Println("Таблица tasks готова")
}
