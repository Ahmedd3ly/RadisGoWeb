package models

import (
	"errors"
	"strings"
	"time"

	"RadisGoWeb/server/database"

	"gorm.io/gorm"
)

type Patient struct {
	Id       string     `gorm:"column:pk_Patient_ID" json:"id"`
	Surname  string     `json:"surname"`
	Forename string     `json:"forename"`
	Dob      *time.Time `gorm:"column:DateOfBirth" json:"dob"`
	Address1 *string    `json:"address1"`
	Address2 *string    `json:"address2"`
	Address3 *string    `json:"address3"`
	Address4 *string    `json:"address4"`
}

type PatientDemographics struct {
	Id                  string     `gorm:"column:pk_Patient_ID" json:"id"`
	RadisNumber         string     `gorm:"unique" json:"radis_number"`
	NHSNumber           *string    `json:"nhs_number"`
	HospitalNumber      *string    `json:"hospital_number"`
	Title               *string    `json:"title"`
	Surname             string     `json:"surname"`
	Forename            *string    `json:"forename"`
	Sex                 *string    `json:"sex"`
	Dob                 *time.Time `gorm:"column:DateOfBirth" json:"dob"`
	Address1            *string    `json:"address1"`
	Address2            *string    `json:"address2"`
	Address3            *string    `json:"address3"`
	Address4            *string    `json:"address4"`
	Address5            *string    `json:"address5"`
	Postcode            *string    `json:"postcode"`
	DeceasedPatientFlag bool       `json:"deceased_patient_flag"`
}

func (Patient) TableName() string {
	return "Patient"
}

func GetPatientsBySurname(surname string) ([]Patient, error) {
	var patients []Patient

	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Where("Surname = ?", surname).Order("Forename").Find(&patients).Error
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, nil
	}
	return patients, nil
}

func GetPatientsById(id string, field string) ([]Patient, error) {
	var patients []Patient
	var whereClause string

	field = strings.ToLower(field)
	switch field {
	case "radisnumber", "nhsnumber", "hospitalnumber":
		whereClause = field + " = ?"
	default:
		return nil, errors.New("unexpected patient identifier field")
	}

	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Where(whereClause, id).Order("Surname").Order("Forename").Order("DateOfBirth").Find(&patients).Error
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, nil
	}
	return patients, nil
}

func GetPatientsBySurnameForenameDobSex(surname string, forename string, dob string, sex string) ([]Patient, error) {
	var patients []Patient
	surname = strings.ToLower(surname)
	forename = strings.ToLower(forename)
	sex = strings.ToLower(sex)

	tx := database.DB.Where("Sex = ?", sex)

	if dob != "" {
		_, err := time.Parse("2006-Jan-02", dob)
		if err != nil {
			return nil, errors.New("invalid date of birth")
		}
		tx = tx.Where("DateOfBirth = ?", dob)
	}

	if forename != "" {
		tx = tx.Where("Forename = ?", forename)
	}

	if surname != "" {
		tx = tx.Where("Surname = ?", surname)
	}

	tx = tx.Order("Surname").Order("Forename").Order("DateOfBirth")
	err := tx.Find(&patients).Error
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, nil
	}
	return patients, nil
}

func GetPatientsByWardWithScheduledExams(locationId string) ([]Patient, error) {
	var patients []Patient
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebPatientsForWardWithScheduledExams @pk_Location_ID = ?, @userID = ?", locationId, "").Scan(&patients).Error
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, nil
	}
	return patients, nil
}

func GetPatientsForWardWithinLastWeek(locationId string) ([]Patient, error) {
	var patients []Patient
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebPatientsForWardWithinLastWeek @pk_Location_ID = ?, @userID = ?", locationId, "").Scan(&patients).Error
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, nil
	}
	return patients, nil
}

func GetPatientsForConsultantWithScheduledExams(clinicianPracticeId string) ([]Patient, error) {
	var patients []Patient
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebPatientsForConsultantWithScheduledExams @pk_ClinicianPractice_ID = ?, @userID = ?", clinicianPracticeId, "").Scan(&patients).Error
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, nil
	}
	return patients, nil
}

func GetPatientsForConsultantWithinTheLastWeek(clinicianPracticeId string) ([]Patient, error) {
	var patients []Patient
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebPatientsForConsultantWithinTheLastSevenDays @pk_ClinicianPractice_ID = ?, @userID = ?", clinicianPracticeId, "").Scan(&patients).Error
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, nil
	}
	return patients, nil
}

func GetPatientDetails(patientId string) (*PatientDemographics, error) {
	var patientDemographics PatientDemographics
	tx := database.DB.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Raw("EXEC RadisWebPatientDetails @pk_Patient_ID = ?, @userID = ?", patientId, "").Scan(&patientDemographics).Error
	if err != nil {
		return nil, err
	}
	if patientDemographics.Id == "" {
		return nil, nil
	}
	return &patientDemographics, nil
}
