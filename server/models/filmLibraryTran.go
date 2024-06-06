package models

import (
	"time"

	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type LibraryTransaction struct {
	Id              string     `gorm:"column:pk_LibraryTransaction_ID" json:"id"`
	TransactionDate *time.Time `json:"transaction_date"`
	TransactionTime *string    `json:"transaction_time"`
	Location        *string    `json:"location"`
}

func (LibraryTransaction) TableName() string {
	return "LibraryTransaction"
}

func GetLibraryTransactionByPatientId(patientId string) ([]LibraryTransaction, error) {
	var transactions []LibraryTransaction

	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Where("fk_Patient_ID = ?", patientId).Order("TransactionDate Desc").Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, nil
	}
	return transactions, nil

}
