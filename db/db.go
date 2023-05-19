package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Função que conecta ao banco de dados
func ConnectDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	conexao := fmt.Sprintf(`user=%s dbname=%s password=%s host=localhost sslmode=disable`, os.Getenv("DB_USR"), os.Getenv("DB_NAME"), os.Getenv("DB_PWD"))

	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
