package models

import (
	"time"

	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type Event struct {
	SortOrder          string     `json:"sort_order"`
	EventDate          *time.Time `json:"event_date"`
	EventTime          *string    `json:"event_time"`
	DateReceived       *time.Time `json:"date_received"`
	Description        *string    `json:"doctor"`
	Doctor             *string    `json:"radis_number"`
	Site               *string    `json:"site"`
	Ward               *string    `json:"ward"`
	Pk_Request_Id      string     `gorm:"column:pk_Request_ID" json:"pk_Request_id"`
	Examination        *string    `json:"examination"`
	IsUltrasoundReport uint       `json:"is_ultrasound_report"`
	ProcedureId        string     `gorm:"column:ProcedureID" json:"procedure_id"`
}

func GetPatientHistoryById(patientId string) ([]Event, error) {
	var events []Event
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebPatientHistory @pk_Patient_ID = ?, @userID = ?", patientId, "").Scan(&events).Error
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, nil
	}
	return events, nil
}
