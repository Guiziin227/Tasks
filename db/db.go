package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"tasks/models"
)

var DB *gorm.DB

func Connect() error {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		host, user, password, dbName, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco: %v", err)
		return err
	}
	fmt.Println("Conexão com o PostgreSQL estabelecida com sucesso!")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Erro ao migrar: %v", err)
	}

	fmt.Println("Banco de dados conectado e migrado com sucesso!")

	return nil
}
