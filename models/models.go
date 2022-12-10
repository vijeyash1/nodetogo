package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// type Base struct {
// 	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
// }

// func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.ID = uuid.New()
// 	return
// }

type Panel struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Query string `json:"quey"`
}
type JSON json.RawMessage

func (j *JSON) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*j = JSON(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (j JSON) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	return json.RawMessage(j).MarshalJSON()
}

type Metadata struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Ssl      bool   `json:"ssl"`
}
type Connector struct {
	Host     string `json:"host"`
	Database string `json:"database"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Ssl      bool   `json:"ssl"`
	Query 	string    `json:"query"`
}
type Datasource struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Metadata   `json:"metadata"`
	Dashboards []Dashboard `gorm:"foreignKey:DatasourceID;references:ID"`
}
type Dashboard struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;"`
	Name         string         `json:"name"`
	Panel        datatypes.JSON `json:"panel"`
	DatasourceID uuid.UUID      `gorm:"type:uuid;foreign_key;"`
}

type Node struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name     string    `json:"name"`
	Visitors int       `json:"visitors"`
	Count    int       `json:"count"`
}
