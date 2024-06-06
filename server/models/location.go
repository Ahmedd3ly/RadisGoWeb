package models

import (
	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type Location struct {
	Id          string `gorm:"column:pk_Location_ID" json:"id"`
	Description string `json:"description"`
}

func GetLocations() ([]Location, error) {
	var locations []Location
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC dbo.RadisWebListLocations").Scan(&locations).Error
	if err != nil {
		return nil, err
	}
	if len(locations) == 0 {
		return nil, nil
	}
	return locations, nil
}
