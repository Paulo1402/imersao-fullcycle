package main

import (
	"database/sql"
	"fmt"
	"imersaofc/internal/converter"
	"log/slog"
	"os"
	_ "github.com/lib/pq"
)


func connectPostgres() (*sql.DB, error) {
	user := getEnvOrDefault("POSTGRES_USER", "postgres")
	password := getEnvOrDefault("POSTGRES_PASSWORD", "postgres")
	dbName := getEnvOrDefault("POSTGRES_DB", "postgres")
	host := getEnvOrDefault("POSTGRES_HOST", "localhost")
	port := getEnvOrDefault("POSTGRES_PORT", "5433")
	sslMode := getEnvOrDefault("POSTGRES_SSLMODE", "disable")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", user, password, dbName, host, port, sslMode)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		slog.Error("Error connecting to database", slog.String("connStr", connStr))
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		slog.Error("Error pinging database", slog.String("connStr", connStr))
		return nil, err
	}

	slog.Info("Connected to Postgres successfully")

	return db, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv((key)); exists {
		return value
	}

	return defaultValue
}

func main() {
	db, err := connectPostgres()

	if err != nil {
		panic(err)
	}

	vc := converter.NewVideoConverter(db)
	vc.Handle([]byte(`{"video_id": 1, "path": "mediatest/media/uploads/1"}`))
}

