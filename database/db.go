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

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func ConectaComBancoDeDados() {

	host := getEnv("HOST", "localhost")
	user := getEnv("USER", "postgres")
	password := getEnv("PASSWORD", "postgres")
	dbname := getEnv("DBNAME", "postgres")
	port := getEnv("DBPORT", "5432")

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
		log.Println("Erro ao executar AutoMigrate:", err)
	}
}
