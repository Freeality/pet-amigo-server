package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/freeality/pet-amigo-server/models"
)

var postgresDb *gorm.DB

// PostgresDb Retorna um gorm.DB válido
func PostgresDb() (*gorm.DB, error) {
	if postgresDb != nil {
		return postgresDb, nil
	}

	// db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	dbParams := fmt.Sprintf("%s %s", os.Getenv("DATABASE_URL"), "sslmode=disable")
	postgresDb, err := gorm.Open("postgres", dbParams)

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	return postgresDb, err
}

// Migrate cria, modifica ou apaga tabelas
func Migrate() {
	postgresDb, err := PostgresDb()
	if err != nil { return }
	defer postgresDb.Close()

	postgresDb.HasTable(&models.Usuario{})
}
