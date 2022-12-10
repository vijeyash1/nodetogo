package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/vijeyash1/backend/models"
)

func (config *Config) DatasourceHandler(c *gin.Context) {
	data := &[]models.Datasource{}
	config.Db.Raw("SELECT * FROM datasources").Scan(data)
	json, err := json.Marshal(data)
	if err != nil {
		c.JSON(400, gin.H{
			"unable to marshal content": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"datasouce": string(json),
	})
}

func (config *Config) CreateDatasourceHandler(c *gin.Context) {
	data := &models.Datasource{}
	c.BindJSON(data)
	if data.Name == "" {
		c.JSON(400, gin.H{
			"data empty error": "Datasource name is Mandatory",
		})
		return
	}
	err := config.Db.Select("Name", "Type", "Metadata").Create(data).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Created DataSource successfully",
	})
}

func (config *Config) DeleteDatasourceHandler(c *gin.Context) {
	data := &models.Datasource{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(400, gin.H{
			"data empty error": "Datasource id is Mandatory",
		})
		return
	}
	err := config.Db.Delete(&models.Datasource{}, data.ID).Error
	if err != nil {
		c.JSON(400, gin.H{
			"unable to delete datasource": err.Error(),
		})
	}
}

func (config *Config) UpdateDatasourceHandler(c *gin.Context) {
	data := &models.Datasource{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(400, gin.H{
			"data empty error": "Datasource id is Mandatory",
		})
		return
	}
	if data.Name == "" {
		c.JSON(400, gin.H{
			"data empty error": "Datasource name is Mandatory",
		})
		return
	}
	err := config.Db.Model(&models.Datasource{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		c.JSON(400, gin.H{
			"unable to update datasource": err.Error(),
		})
	}
}
