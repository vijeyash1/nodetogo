package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vijeyash1/backend/pkg/models"
)

func (config *APIHandler) GetDashboardHandler(c *gin.Context) {
	data := []models.Dashboard{}
	config.Db.Raw("SELECT * FROM dashboards").Scan(&data)
	c.IndentedJSON(http.StatusOK, data)
}

func (config *APIHandler) CreateDashboardHandler(c *gin.Context) {
	
	data := &models.Dashboard{}
	c.BindJSON(data)
	if data.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Dashboard name is Mandatory",
		})
		return
	}
	if data.DatasourceID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Datasource id is Mandatory",
		})
		return
	}
	err := config.Db.Select("Name", "DatasourceID", "Metadata").Create(data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Unable to Create Dashboard": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "Created Dashboard successfully",
	})
}

func (config *APIHandler) DeleteDashboardHandler(c *gin.Context) {
	
	data := &models.Dashboard{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Dashboard id is Mandatory",
		})
		return
	}
	err := config.Db.Delete(&models.Dashboard{}, data.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"unable to delete dashboard": err.Error(),
		})
	}
}

func (config *APIHandler) UpdateDashboardHandler(c *gin.Context) {
	
	data := &models.Dashboard{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Dashboard id is Mandatory",
		})
		return
	}
	if data.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Dashboard name is Mandatory",
		})
		return
	}
	if data.DatasourceID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Datasource id is Mandatory",
		})
		return
	}
	err := config.Db.Model(&models.Dashboard{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"unable to update dashboard": err.Error(),
		})
	}
}

func (config *APIHandler) GetDashboardByIDHandler(c *gin.Context) {
	data := &models.Dashboard{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Dashboard id is Mandatory",
		})
		return
	}
	err := config.Db.Where("id = ?", data.ID).First(&data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"unable to get dashboard": err.Error(),
		})
	}
	json, err := json.Marshal(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"unable to marshal content": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"dashboard": string(json),
	})
}
