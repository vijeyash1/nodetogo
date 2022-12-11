package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vijeyash1/backend/pkg/models"
)

func (config *APIHandler) DatasourceHandler(c *gin.Context) {
	data := []models.Datasource{}
	config.Db.Raw("SELECT * FROM datasources").Scan(&data)
	c.IndentedJSON(http.StatusOK, data)
}
func (config *APIHandler) CreateDatasourceHandler(c *gin.Context) {

	data := &models.Datasource{}
	c.BindJSON(data)
	if data.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Datasource name is Mandatory",
		})
		return
	}
	err := config.Db.Select("Name", "Type", "Metadata").Create(data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "Created DataSource successfully",
	})
}

func (config *APIHandler) DeleteDatasourceHandler(c *gin.Context) {

	data := &models.Datasource{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Datasource id is Mandatory",
		})
		return
	}
	err := config.Db.Delete(&models.Datasource{}, data.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"unable to delete datasource": err.Error(),
		})
	}
}

func (config *APIHandler) UpdateDatasourceHandler(c *gin.Context) {

	data := &models.Datasource{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Datasource id is Mandatory",
		})
		return
	}
	if data.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Datasource name is Mandatory",
		})
		return
	}
	err := config.Db.Model(&models.Datasource{}).Where("id = ?", data.ID).Updates(data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"unable to update datasource": err.Error(),
		})
	}
}
func (config *APIHandler) GetDatasourceByIdHandler(c *gin.Context) {

	data := &models.Datasource{}
	c.BindJSON(data)
	if data.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"data empty error": "Datasource id is Mandatory",
		})
		return
	}
	err := config.Db.Where("id = ?", data.ID).First(data).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"unable to get datasource": err.Error(),
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
		"datasouce": string(json),
	})
}
