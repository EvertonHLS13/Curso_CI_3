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
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return defaultValue
}

func ConectaComBancoDeDados() {

	host := getEnv("DBHOST", "localhost")
	user := getEnv("DBUSER", "postgres")
	password := getEnv("DBPASSWORD", "postgres")
	dbname := getEnv("DBNAME", "postgres")
	port := getEnv("DBPORT", "5432")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	DB = db

	if err := DB.AutoMigrate(&models.Aluno{}); err != nil {
		log.Panic("Erro no AutoMigrate: ", err)
	}
}
