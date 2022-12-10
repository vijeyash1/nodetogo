package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vijeyash1/backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (config *Config) PostgresConnectorHandler(c *gin.Context) {
	receivedData := make(map[string]interface{})
	data := &models.Connector{}
	c.BindJSON(data)
	// dbURI := "host=localhost user=core password=core dbname=core port=5432 sslmode=disable"
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", data.Host, data.User, data.Password, data.Database, data.Port)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		c.JSON(500, gin.H{
			"message": "unable to open database connection",
		})
		return
	}
	err = db.Raw(data.Query).Scan(&receivedData).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "unable to execute query",
		})
		return
	}
	c.JSON(200, gin.H{
		"result": receivedData,
	})
}
