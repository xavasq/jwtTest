package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

var DB *pgxpool.Pool

func ConnectDB() {
	dsn := "postgres://postgres:database153426@localhost:5432/sadboysdb"
	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("ошибка при подключения к базе")
	}

	DB = conn
	log.Println("база данных подключена успешно")
}
