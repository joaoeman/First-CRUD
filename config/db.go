package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// inicializa a conexão com o db
func SetupDataBase() *sql.DB { //se começa com a letra maiuscula, pode chamar em outros arquivos
	err := godotenv.Load() //carregando do arquivo .env

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST") //pegando do arquivo .env
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connectionsStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	dbConnection, err := sql.Open("postgres", connectionsStr) //fazendo conexao com db

	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Ping() //verifica se a conexão deu certo

	fmt.Println("db conectado")
	fmt.Println(connectionsStr)

	return dbConnection
}
