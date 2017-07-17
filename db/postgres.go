package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var postgresDb *gorm.DB

// PostgresDb Retorna um gorm.DB válido
func PostgresDb() *gorm.DB {
	if postgresDb != nil {
		return postgresDb
	}

	// db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	// db, err := gorm.Open("postgres", "host=myhost user=gorm dbname=gorm sslmode=disable password=mypassword")
	dbParams := fmt.Sprintf("host=%s", os.Getenv("DATABASE_URL"))
	postgresDb, err := gorm.Open("postgres", dbParams)

	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	return postgresDb
}

// Migrate cria, modifica ou apaga tabelas
func Migrate() {

	// self.PostgresDb.AutoMigrate(&models.Usuario{})
}
