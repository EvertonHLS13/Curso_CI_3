package database

import (
	"fmt"
	"log"
	"os"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func getEnv(keys []string, defaultValue string) string {
	for _, key := range keys {
		if value := os.Getenv(key); value != "" {
			return value
		}
	}

	return defaultValue
}

func ConectaComBancoDeDados() {

	host := getEnv(
		[]string{"DB_HOST", "DBHOST", "HOST"},
		"localhost",
	)

	user := getEnv(
		[]string{"DB_USER", "DBUSER", "USER"},
		"postgres",
	)

	password := getEnv(
		[]string{"DB_PASSWORD", "DBPASSWORD", "PASSWORD"},
		"postgres",
	)

	dbname := getEnv(
		[]string{"DB_NAME", "DBNAME"},
		"postgres",
	)

	port := getEnv(
		[]string{"DB_PORT", "DBPORT"},
		"5432",
	)


	// Define SSL automaticamente
	//
	// AWS RDS:
	// sslmode=require
	//
	// Local/GitHub Actions:
	// sslmode=disable

	sslmode := os.Getenv("DB_SSLMODE")

	if sslmode == "" {

		if host == "localhost" || host == "127.0.0.1" {
			sslmode = "disable"
		} else {
			sslmode = "require"
		}
	}


	log.Println("Conectando ao banco:")
	log.Printf(
		"host=%s user=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		dbname,
		port,
		sslmode,
	)


	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s connect_timeout=10",
		host,
		user,
		password,
		dbname,
		port,
		sslmode,
	)


	db, err := gorm.Open(
		postgres.Open(stringDeConexao),
		&gorm.Config{},
	)


	if err != nil {

		log.Println("Erro ao conectar com banco de dados:")
		log.Println(err)

		log.Panic("Falha na conexão com PostgreSQL")
	}


	DB = db


	log.Println("Banco conectado com sucesso")


	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {

		log.Println("Erro no AutoMigrate:")
		log.Println(err)
	}
}
