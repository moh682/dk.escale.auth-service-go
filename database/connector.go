package database

import (
	"fmt"
	"log"

	"dk.escale.auth-service/database/models"

	"dk.escale.auth-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

// setup opens a connection to postgresql
func setup() *gorm.DB {
	env := config.GetEnvironments()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=5432", env.PostgresqlHost, env.PostgresqlUsername, env.PostgresqlPassword, env.PostgresqlDatabase)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Falied to connect to database err: " + err.Error())
	}
	config.Green("Connected to database have been established")

	err = db.AutoMigrate(
		&models.Location{},
		&models.Company{},
		// &models.Department{},
		&models.Employee{},
	)

	if err != nil {
		log.Fatalln(err.Error())
	}

	// db.Model(&models.Department{}).Association("Location")
	db.Model(&models.Employee{}).Association("Company")
	// db.Model(&models.Employee{}).Association("")

	if err != nil {
		log.Fatalln(err.Error())
	}

	return db
}

// GetConnection opens a connection to postgresql
func GetConnection() *gorm.DB {
	if connection != nil {
		return connection.Session(&gorm.Session{
			SkipHooks: false,
		})
	}
	connection = setup()
	return connection.Session(&gorm.Session{
		SkipHooks: false,
	})
}
