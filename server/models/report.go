package models

import (
	"time"

	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type ReportDetail struct {
	Pk_ReportChapter_ID                *string    `gorm:"column:pk_ReportChapter_ID" json:"id"`
	ReportChapterStatus                *string    `gorm:"column:fk_ReportChapterStatus_ID" json:"report_chapter_status"`
	Fk_Request_ID                      *string    `gorm:"column:fk_Request_ID" json:"request_id"`
	ReportText                         *string    `json:"report_text,omitempty"`
	UserDetailTypistID                 *string    `gorm:"column:fk_UserDetail_Typist_ID" json:"user_detail_typist_id,omitempty"`
	ChapterType                        *string    `json:"chapter_type,omitempty"`
	DateCreated                        *time.Time `json:"date_created"`
	TimeCreated                        *string    `json:"time_created"`
	UserDetailValidatedByID            *string    `gorm:"column:fk_UserDetail_ValidatedBy_ID" json:"user_detail_validated_by_id,omitempty"`
	UserDetailAuthorID                 *string    `gorm:"column:fk_UserDetail_Author_ID" json:"user_detail_author_id,omitempty"`
	DateValidated                      *time.Time `json:"date_validated,omitempty"`
	TimeValidated                      *string    `json:"time_validated,omitempty"`
	UserDetailLastPrintedByID          *string    `gorm:"column:fk_UserDetail_LastPrintedBy_ID" json:"user_detail_last_printed_by_id,omitempty"`
	DateLastPrinted                    *time.Time `json:"date_last_printed,omitempty"`
	TimeLastPrinted                    *string    `json:"time_last_printed,omitempty"`
	Publish                            *bool      `json:"publish"`
	UserDetailPublishedByID            *string    `gorm:"column:fk_UserDetail_PublishedBy_ID" json:"user_detail_published_by_id,omitempty"`
	DateDictated                       *time.Time `json:"date_dictated,omitempty"`
	TimeDictated                       *string    `json:"time_dictated,omitempty"`
	UserDetailDictationTranscribedByID *string    `gorm:"column:fk_UserDetail_DictationTranscribedBy_ID" json:"user_detail_dictation_transcribed_by_id,omitempty"`
	DateDictationTranscribed           *time.Time `json:"date_dictation_transcribed,omitempty"`
	TimeDictationTranscribed           *string    `json:"time_dictation_transcribed,omitempty"`
	Fk_Patient_ID                      *string    `gorm:"column:fk_Patient_ID" json:"patient_id,omitempty"`
	Fk_Procedure_ID                    *string    `gorm:"column:fk_Procedure_ID" json:"procedure_id,omitempty"`
	Fk_ReportBatch_ID                  *string    `gorm:"column:fk_ReportBatch_ID" json:"report_batch_id,omitempty"`
	Fk_UserDetail_Owner_ID             *string    `gorm:"column:fk_UserDetail_Owner_ID" json:"user_detail_owner_id,omitempty"`
	Fk_UserDetail_LastUpdatedBy_ID     *string    `gorm:"column:fk_UserDetail_LastUpdatedBy_ID" json:"user_detail_last_updated_by_id,omitempty"`
	ExpiryDate                         *time.Time `json:"expiry_date,omitempty"`
	StatusDescription                  *string    `json:"status_description,omitempty"`
}

func GetReportDetailsById(requestId string) ([]ReportDetail, error) {
	var reportDetails []ReportDetail
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebReportDetails @pk_Request_ID = ?, @userID = ?", requestId, "").Scan(&reportDetails).Error
	if err != nil {
		return nil, err
	}
	if len(reportDetails) == 0 {
		return nil, nil
	}
	return reportDetails, nil
}
