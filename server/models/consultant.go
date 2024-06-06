package models

import (
	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type Consultant struct {
	Id   string `gorm:"column:pk_ClinicianPractice_ID" json:"id"`
	Name string `json:"name"`
}

func GetConsultants() ([]Consultant, error) {
	var consultants []Consultant
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC dbo.RadisWebListConsultants").Scan(&consultants).Error
	if err != nil {
		return nil, err
	}
	if len(consultants) == 0 {
		return nil, nil
	}
	return consultants, nil
}
