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

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func ConectaComBancoDeDados() {

	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "postgres")
	port := getEnv("DB_PORT", "5432")

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

	err = DB.AutoMigrate(&models.Aluno{})

	if err != nil {
		log.Println("Erro no AutoMigrate:", err)
	}
}
