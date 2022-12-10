package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/vijeyash1/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	log.Println("Initializing database and tables...")
	log.Println("Initializing database with sample data...")

	dbURI := "host=localhost user=core password=core dbname=core port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to open database connection: ", err)
	}
	db.Migrator().HasTable(&models.Datasource{})
	db.Migrator().DropTable(&models.Datasource{})
	db.Migrator().HasTable(&models.Dashboard{})
	db.Migrator().DropTable(&models.Dashboard{})
	db.Migrator().HasTable(&models.Node{})
	db.Migrator().DropTable(&models.Node{})
	err = db.AutoMigrate(&models.Datasource{}, &models.Dashboard{}, &models.Node{})
	if err != nil {
		fmt.Println(err)
		return
	}
	Datasource := models.Datasource{
		ID:   uuid.New(),
		Name: "Postgres",
		Type: "POSTGRES",
		Metadata: models.Metadata{
			Host:     "localhost",
			Database: "core",
			Port:     5432,
			User:     "core",
			Password: "core",
			Ssl:      false,
		},
	}
	DatasourceResult := db.Create(&Datasource)
	fmt.Printf("DataSource added %v \n", DatasourceResult.RowsAffected)
	Node := []models.Node{{
		ID:       uuid.New(),
		Name:     "Superman",
		Visitors: 324,
		Count:    69,
	},
		{
			ID:       uuid.New(),
			Name:     "Spiderman",
			Visitors: 435,
			Count:    78,
		},
		{
			ID:       uuid.New(),
			Name:     "Batman",
			Visitors: 234,
			Count:    43,
		},
	}
	data := []models.Panel{

		{
			ID:    1,
			Name:  "Node Packages",
			Type:  "DATA_TABLE",
			Query: "SELECT name, visitors, count FROM node_packages",
		},
		{
			ID:    2,
			Name:  "Bar Chart",
			Type:  "DATA_TABLE",
			Query: "SELECT name, visitors, count FROM node_packages",
		},
	}
	json, err := json.Marshal(data)
	if err != nil {
		log.Println("unable to marshal")
	}
	Dashboard := models.Dashboard{
		ID:           uuid.New(),
		Name:         "Demo",
		Panel:        json,
		DatasourceID: Datasource.ID,
	}
	//
	DashboardResult := db.Create(&Dashboard)
	fmt.Printf("Dashboard created %v \n", DashboardResult.RowsAffected)

	NodeResult := db.Create(&Node)
	fmt.Printf("Nodes created %v \n", NodeResult.RowsAffected)

}
