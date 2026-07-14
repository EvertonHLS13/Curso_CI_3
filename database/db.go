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
	host := getEnv("DBHOST", "localhost")
	user := getEnv("DBUSER", "root")
	password := getEnv("DBPASSWORD", "postgres")
	dbname := getEnv("DBNAME", "postgres")
	port := getEnv("DBPORT", "5432")
	sslMode := getEnv("DBSSLMODE", "require")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=UTC",
		host,
		user,
		password,
		dbname,
		port,
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	DB = db

	log.Println("Conectado ao PostgreSQL com sucesso!")

	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Fatalf("Erro no AutoMigrate: %v", err)
	}

	log.Println("AutoMigrate executado com sucesso!")
}
