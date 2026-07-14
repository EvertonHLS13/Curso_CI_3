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

func ConectaComBancoDeDados() {

	stringDeConexao := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
		os.Getenv("DBPORT"),
	)

	db, erro := gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})

	if erro != nil {
		log.Println("Erro ao conectar com banco de dados")
		log.Panic(erro)
	}

	DB = db

	DB.AutoMigrate(&models.Aluno{})
}
