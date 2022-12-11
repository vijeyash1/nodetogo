package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vijeyash1/backend/pkg/models"
	"gorm.io/datatypes"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (config *APIHandler) PostgresConnectorHandler(c *gin.Context, metadata *models.Datasource, query *models.PanelQuery) *datatypes.JSON {
	m := &models.Metadata{}
	err := json.Unmarshal([]byte(metadata.Metadata), m)
	if err != nil {
		log.Println("unable to unmarshal metadata: ", err)
		return nil
	}
	dbURI := fmt.Sprintf("host=%s user=%s password =%s dbname=%s port=%d sslmode=disable", m.Host, m.User, m.Password, m.Database, m.Port)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Println("unable to open database connection: ", err)
		return nil
	}
	var result datatypes.JSON
	db.Raw(query.Query).Scan(&result)
	return &result
}
func (config *APIHandler) MysqlConnectorHandler(c *gin.Context, metadata *models.Datasource, query *models.PanelQuery) *datatypes.JSON {
	m := &models.Metadata{}
	err := json.Unmarshal([]byte(metadata.Metadata), m)
	if err != nil {
		log.Println("unable to unmarshal metadata: ", err)
		return nil
	}
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Println("unable to open database connection: ", err)
		return nil
	}
	var result datatypes.JSON
	err = db.Raw(query.Query).Scan(&result).Error
	if err != nil {
		log.Println("unsuccessfull query to database: ", err)
		return nil
	}
	return &result
}
func (config *APIHandler) ClickhouseConnectorHandler(c *gin.Context, metadata *models.Datasource, query *models.PanelQuery) *datatypes.JSON {
	m := &models.Metadata{}
	err := json.Unmarshal([]byte(metadata.Metadata), m)
	if err != nil {
		log.Println("unable to unmarshal metadata: ", err)
		return nil
	}
	dbURI := fmt.Sprintf("tcp://%s:%d?username=%s&password=%s&database=%s", m.Host, m.Port, m.User, m.Password, m.Database)
	db, err := gorm.Open(clickhouse.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Println("unable to open database connection: ", err)
		return nil
	}
	var result datatypes.JSON
	err = db.Raw(query.Query).Scan(&result).Error
	if err != nil {
		log.Println("unsuccessfull query to database: ", err)
		return nil
	}
	return &result

}
