package models

import (
	"time"

	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type Procedure struct {
	ScanDate             *time.Time `json:"scan_date"`
	ExamSide             *string    `json:"exam_side"`
	StartTime            *time.Time `json:"start_time"`
	RoomCode             *string    `json:"room_code"`
	RoomDescription      *string    `json:"room_description"`
	ProcedureCode        *string    `gorm:"column:ProcedureCodeCode" json:"procedure_code"`
	ProcedureDescription *string    `gorm:"column:ProcedureCodeDescription" json:"procedure_description"`
}

func GetProceduresByRequestId(requestId string) ([]Procedure, error) {
	var procedures []Procedure
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebProcedureDetails @pk_Request_ID = ?, @userID = ?", requestId, "").Scan(&procedures).Error
	if err != nil {
		return nil, err
	}
	if len(procedures) == 0 {
		return nil, nil
	}
	return procedures, nil
}
