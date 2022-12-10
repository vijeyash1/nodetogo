package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vijeyash1/backend/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start() {
	r := gin.Default()
	dbURI := "host=localhost user=core password=core dbname=core port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to open database connection: ", err)
	}

	app := handlers.NewConfig(db)

	r.GET("/api/datasource", app.DatasourceHandler)
	r.POST("/api/datasource", app.CreateDatasourceHandler)
	r.DELETE("/api/datasource", app.DeleteDatasourceHandler)
	r.PUT("/api/datasource", app.UpdateDatasourceHandler)
	r.POST("/api/connector", app.PostgresConnectorHandler)
	r.GET("/api/dashboard", app.GetDashboardHandler)
	r.POST("/api/dashboard", app.CreateDashboardHandler)
	r.DELETE("/api/dashboard", app.DeleteDashboardHandler)
	r.PUT("/api/dashboard", app.UpdateDashboardHandler)
	
	r.Run(":8000")
}
