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

	host := getEnv([]string{"DB_HOST", "HOST"}, "localhost")
	user := getEnv([]string{"DB_USER"}, "postgres")
	password := getEnv([]string{"DB_PASSWORD", "PASSWORD"}, "postgres")
	dbname := getEnv([]string{"DB_NAME", "DBNAME"}, "postgres")
	port := getEnv([]string{"DB_PORT", "DBPORT"}, "5432")

	log.Println("Conectando ao banco:")
	log.Printf("host=%s user=%s dbname=%s port=%s", host, user, dbname, port)

	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		log.Println("Erro ao conectar com banco de dados")
		log.Panic(err)
	}

	DB = db

	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Println("Erro no AutoMigrate:", err)
	}
}
