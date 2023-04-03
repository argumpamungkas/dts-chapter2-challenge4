package repo

import (
	"DTS/Chapter-2/chapter2-challenge-sesi-4/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "gume98"
	port     = "5432"
	dbname   = "api-gorm-book"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	db.Debug().AutoMigrate(models.Book{})
	// db.Debug().Migrator().DropTable(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
