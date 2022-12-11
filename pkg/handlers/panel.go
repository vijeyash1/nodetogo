package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vijeyash1/backend/pkg/models"
)

func (config *APIHandler) ConnectorHandler(c *gin.Context) {
	request := &models.PanelQuery{}
	c.BindJSON(request)
	datasource := &models.Datasource{}
	err := config.Db.Model(&models.Datasource{}).Where("id = ?", request.ID).First(&datasource).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to find datasource",
		})
		return
	}
	switch datasource.Type {
	case "POSTGRES":
		response := config.PostgresConnectorHandler(c, datasource, request)
		c.IndentedJSON(http.StatusOK, *response)
	case "MYSQL":
		response := config.MysqlConnectorHandler(c, datasource, request)
		c.IndentedJSON(http.StatusOK, *response)
	case "CLICKHOUSE":
		response := config.ClickhouseConnectorHandler(c, datasource, request)
		c.IndentedJSON(http.StatusOK, *response)
	}
}
