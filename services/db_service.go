package services

import (
	"log"
	"os"
	"sl-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Setup() {
	/* Connection string to Database */
	//dsn := "host=ec2-52-4-104-184.compute-1.amazonaws.com user=zpgnxqyaesajeg password=f829f2c0d7e29fc770d47d57402252b0ee35b60a793445aca014a8c3f5cb4216 dbname=d1a6mks332cl55 port=5432 sslmode=disable TimeZone=Europe/Lisbon"

	var openErr error
	Database, openErr = gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	if openErr != nil {
		panic(openErr)
	}
	log.Println("Established connection to Database")

	/* Create tables that dont already exist */
	migErr := Database.AutoMigrate(&models.OilBin{}, &models.User{}, &models.Bottle{}, &models.Sensor{}, &models.SensorReading{}, &models.Deposit{})
	if migErr != nil {
		panic(migErr)
	}
	log.Println("Automigration Complete!")

}

func seed() {
	bins := []models.OilBin{
		{Address: "Rua das Oliveiras 114", CoordinateX: 12.32, CoordinateY: -141.99},
		{Address: "Rua das Jardineiras 32", CoordinateX: 41.11, CoordinateY: 92.13},
	}

	for _, b := range bins {
		Database.Create(&b)
	}
	Database.Commit()
}
