package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/vijeyash1/backend/models"
)

func (config *Config) GetDashboardHandler(c *gin.Context) {
	data := []models.Dashboard{}
	config.Db.Raw("SELECT * FROM dashboards").Scan(&data)
	json, err := json.Marshal(data)
	if err != nil {
		c.JSON(400, gin.H{
			"unable to marshal content": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"dashboard": string(json),
	})
}

func (config *Config) CreateDashboardHandler(c *gin.Context) {
	data := &models.Dashboard{}
	c.BindJSON(data)
	if data.Name == "" {
		c.JSON(400, gin.H{
			"data empty error": "Dashboard name is Mandatory",
		})
		return
	}
	if data.DatasourceID.String() == "" {
		c.JSON(400, gin.H{
			"data empty error": "Datasource id is Mandatory",
		})
		return
	}
	err := config.Db.Select("Name", "DatasourceID", "Metadata").Create(data).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Unable to Create Dashboard": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"success": "Created Dashboard successfully",
	})
}

func (config *Config) DeleteDashboardHandler(c *gin.Context) {
	data := &models.Dashboard{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(400, gin.H{
			"data empty error": "Dashboard id is Mandatory",
		})
		return
	}
	err := config.Db.Delete(&models.Dashboard{}, data.ID).Error
	if err != nil {
		c.JSON(400, gin.H{
			"unable to delete dashboard": err.Error(),
		})
	}
}

func (config *Config) UpdateDashboardHandler(c *gin.Context) {
	data := &models.Dashboard{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(400, gin.H{
			"data empty error": "Dashboard id is Mandatory",
		})
		return
	}
	err := config.Db.Model(&models.Dashboard{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		c.JSON(400, gin.H{
			"unable to update dashboard": err.Error(),
		})
	}
}

func (config *Config) GetDashboardByIDHandler(c *gin.Context) {
	data := &models.Dashboard{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(400, gin.H{
			"data empty error": "Dashboard id is Mandatory",
		})
		return
	}
	err := config.Db.Where("id = ?", data.ID).First(&data).Error
	if err != nil {
		c.JSON(400, gin.H{
			"unable to get dashboard": err.Error(),
		})
	}
	json, err := json.Marshal(data)
	if err != nil {
		c.JSON(400, gin.H{
			"unable to marshal content": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"dashboard": string(json),
	})
}
