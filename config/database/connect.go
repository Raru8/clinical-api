package database

import (
	"errors"
	"fmt"
	"os"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


var DB *gorm.DB

func Connect() (*gorm.DB, error){
	dsn := os.Getenv("DB_URL")

	if dsn == ""{
		return nil, errors.New("variável de ambiente DB_URL não foi definida")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir conexão com o banco: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter instância SQL do banco: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("Erro ao conectar ao banco: %w", err)
	}

	log.Printf("Conexão com o banco de dados estabelecida com sucesso")

	DB = db

	return DB, nil
}
