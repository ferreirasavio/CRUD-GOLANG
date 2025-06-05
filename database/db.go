package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: não foi possível carregar o arquivo .env")
	}

	dsn := os.Getenv("DATABASE_URL")
	log.Printf("DEBUG - DATABASE_URL: %q", dsn)
	if dsn == "" {
		log.Fatal("DATABASE_URL não definida")
	}

	DB, err = pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Erro ao verificar conexão com o banco: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida.")
}
