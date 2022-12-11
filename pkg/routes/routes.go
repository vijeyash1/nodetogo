package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"github.com/vijeyash1/backend/pkg/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	Host     string `envconfig:"host" default:"localhost" required:"false"`
	Database string `envconfig:"db" default:"core" required:"false"`
	Port     int    `envconfig:"dbport" default:"5432" required:"false"`
	User     string `envconfig:"dbuser" default:"core" required:"false"`
	Password string `envconfig:"dbpassword" default:"core" required:"false"`
	Ssl      string `envconfig:"sslmode" default:"disable" required:"false"`
}

func Start() {
	r := gin.Default()
	dbmetadata := &Db{}
	err := envconfig.Process("", dbmetadata)
	if err != nil {
		log.Fatal(err.Error())
	}
	format := "host=%s user=%s password=%s dbname=%s port=%d sslmode=%v"
	dbURI := fmt.Sprintf(format, dbmetadata.Host, dbmetadata.User, dbmetadata.Password, dbmetadata.Database, dbmetadata.Port, dbmetadata.Ssl)
	// dbURI := "host=localhost user=core password=core dbname=core port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to open database connection: ", err)
	}

	app := handlers.NewAPIHandler(db)

	r.GET("/api/datasource", app.DatasourceHandler)
	r.POST("/api/datasource", app.CreateDatasourceHandler)
	r.DELETE("/api/datasource", app.DeleteDatasourceHandler)
	r.PUT("/api/datasource", app.UpdateDatasourceHandler)
	r.POST("/api/panel", app.ConnectorHandler)
	r.GET("/api/dashboard", app.GetDashboardHandler)
	r.POST("/api/dashboard", app.CreateDashboardHandler)
	r.DELETE("/api/dashboard", app.DeleteDashboardHandler)
	r.PUT("/api/dashboard", app.UpdateDashboardHandler)

	r.Run(":8000")
}
