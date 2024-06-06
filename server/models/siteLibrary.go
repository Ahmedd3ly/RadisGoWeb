package models

import (
	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type Library struct {
	Id          string  `gorm:"column:pk_SiteLibrary_ID" json:"id"`
	Code        *string `json:"code"`
	Description *string `json:"description"`
}

func GetLibraryDetails(siteLibraryId string) (*Library, error) {
	var library Library

	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("SELECT pk_SiteLibrary_ID, Code, Description FROM SiteLibrary JOIN dbo.Library ON pk_Library_ID = fk_Library_ID WHERE pk_SiteLibrary_ID = ?", siteLibraryId, "").Scan(&library).Error
	if err != nil {
		return nil, err
	}
	return &library, nil

}
