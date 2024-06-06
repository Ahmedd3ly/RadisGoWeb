package models

import (
	"time"

	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type RequestDetail struct {
	Id                       string    `gorm:"column:pk_Request_ID" json:"id"`
	RequestStatus            int16     `json:"request_status"`
	RequestStatusDescription *string   `json:"request_status_description"`
	PriorityDescription      *string   `json:"priority_description"`
	Alerts                   *string   `json:"alerts,omitempty"`
	ClinicalComments         *string   `json:"clinical_comments,omitempty"`
	DateReceived             time.Time `json:"date_received"`
	BookedDate               *string   `json:"booked_date"`
	BookedTime               *string   `json:"booked_time"`
	Examination              *string   `json:"examination"`
	SiteDescription          *string   `json:"site_description,omitempty"`
	LocationCode             *string   `json:"location_code,omitempty"`
	LocationDescription      *string   `json:"location_description,omitempty"`
	TransportDescription     *string   `json:"transport_description,omitempty"`
	AppointmentDate          time.Time `json:"appointment_date"`
	AppointmentTime          *string   `json:"appointment_time"`
	PracticeAddress1         *string   `gorm:"column:Address1" json:"practice_address1,omitempty"`
	PracticeAddress2         *string   `gorm:"column:Address2" json:"practice_address2,omitempty"`
	PracticeAddress3         *string   `gorm:"column:Address3" json:"practice_address3,omitempty"`
	PracticeAddress4         *string   `gorm:"column:Address4" json:"practice_address4,omitempty"`
	PracticeAddress5         *string   `gorm:"column:Address5" json:"practice_address5,omitempty"`
	PracticePostcode         *string   `gorm:"column:Postcode" json:"practice_postcode,omitempty"`
	PracticeDescription      *string   `json:"practice_description,omitempty"`
	SpecialtyCode            *string   `json:"specialty_code,omitempty"`
	SpecialtyDescription     *string   `json:"specialty_description,omitempty"`
	ClinicianTitle           *string   `gorm:"column:Title" json:"clinician_title,omitempty"`
	ClinicianInitials        *string   `gorm:"column:Initials" json:"clinician_initials,omitempty"`
	ClinicianSurname         *string   `gorm:"column:Surname" json:"clinician_surname,omitempty"`
	ClinicianIsGP            *string   `gorm:"column:IsGP" json:"clinician_is_gp"`
	BookedByTitle            *string   `gorm:"column:BookedbyTitle" json:"booked_by_title,omitempty"`
	BookedByForenames        *string   `gorm:"column:BookedbyForenames" json:"booked_by_forenames,omitempty"`
	BookedBySurname          *string   `json:"booked_by_surname,omitempty"`
	CancelledByTitle         *string   `json:"cancelled_by_title,omitempty"`
	CancelledByForenames     *string   `json:"cancelled_by_forenames,omitempty"`
	CancelledBySurname       *string   `json:"cancelled_by_surname,omitempty"`
	CancellationDate         time.Time `json:"cancellation_date"`
	CancellationTime         *string   `json:"cancellation_time,omitempty"`
	CancellationDescription  *string   `json:"cancellation_description,omitempty"`
	CancellationCode         *string   `json:"cancellation_code,omitempty"`
}

func GetRequestDetailsById(requestId string) (*RequestDetail, error) {
	var requestDetail RequestDetail
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebRequestDetails @pk_Request_ID = ?, @userID = ?", requestId, "").Scan(&requestDetail).Error
	if err != nil {
		return nil, err
	}
	if requestDetail.Id == "" {
		return nil, nil
	}
	return &requestDetail, nil
}
