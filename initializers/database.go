package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn := "host=localhost user=postgres password=secret dbname=crud-pgsql port=5442 sslmode=disable TimeZone=Asia/Tokyo"
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Fail to connect to DB")
	}
}