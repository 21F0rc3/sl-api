package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sl-api/domain"
)

var Database *gorm.DB

func Setup() {
	/* Connection string to Database */
	dsn := "host=localhost user=fred password=Fred2001 dbname=slapi port=5432 sslmode=disable TimeZone=Europe/Lisbon"

	var openErr error
	Database, openErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if openErr != nil {
		panic(openErr)
	}
	log.Println("Established connection to Database")

	/* Create tables that dont already exist */
	migErr := Database.AutoMigrate(&domain.OilBin{})
	if migErr != nil {
		panic(migErr)
	}
	log.Println("Automigration Complete!")
}

func seed() {
	bins := []domain.OilBin{
		{Address: "Rua das Oliveiras 114", CoordinateX: 12.32, CoordinateY: -141.99},
		{Address: "Rua das Jardineiras 32", CoordinateX: 41.11, CoordinateY: 92.13},
	}

	for _, b := range bins {
		Database.Create(&b)
	}
	Database.Commit()
}
