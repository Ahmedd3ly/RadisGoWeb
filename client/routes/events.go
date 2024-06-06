package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PatientDemographics struct {
	Id                  string     `gorm:"column:pk_Patient_ID" json:"id"`
	RadisNumber         string     `gorm:"unique" json:"radis_number"`
	NHSNumber           *string    `json:"nhs_number"`
	HospitalNumber      *string    `json:"hospital_number"`
	Title               *string    `json:"title"`
	Surname             string     `json:"surname"`
	Forename            *string    `json:"forename"`
	Sex                 string     `json:"sex"`
	Dob                 *time.Time `gorm:"column:DateOfBirth" json:"dob"`
	Address1            *string    `json:"address1"`
	Address2            *string    `json:"address2"`
	Address3            *string    `json:"address3"`
	Address4            *string    `json:"address4"`
	Address5            *string    `json:"address5"`
	Postcode            *string    `json:"postcode"`
	DeceasedPatientFlag bool       `json:"deceased_patient_flag"`
}

type EventHistory struct {
	SortOrder          string     `json:"sort_order"`
	EventDate          *time.Time `json:"event_date"`
	EventTime          *string    `json:"event_time"`
	DateReceived       *time.Time `json:"date_received"`
	Description        string     `json:"doctor"`
	Doctor             *string    `json:"radis_number"`
	Site               *string    `json:"site"`
	Ward               *string    `json:"ward"`
	Pk_Request_Id      string     `gorm:"column:pk_Request_ID" json:"pk_Request_id"`
	Examination        *string    `json:"examination"`
	IsUltrasoundReport uint       `json:"is_ultrasound_report"`
	ProcedureId        string     `gorm:"column:ProcedureID" json:"procedure_id"`
}

func patientDetailsHistoryByIdHandler(context *gin.Context) {
	patientId := context.Param("patientId")

	if patientId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request: patientId is required"})
		return
	}

	// Construct the API URLs
	patientURL := fmt.Sprintf("http://localhost:8081/radisweb/patients/patient-details?patientId=%s", patientId)
	eventsURL := fmt.Sprintf("http://localhost:8081/radisweb/patients/patient-history?patientId=%s", patientId)

	// Get patient details
	patientResp, err := http.Get(patientURL)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get patient's details"})
		return
	}
	defer patientResp.Body.Close()

	// Read the details response body
	patientBody, err := ioutil.ReadAll(patientResp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	var patient PatientDemographics
	err = json.Unmarshal(patientBody, &patient)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to parse patient's details response"})
		return
	}

	// Get patient events history
	eventsResp, err := http.Get(eventsURL)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get patient  event history"})
		return
	}
	defer eventsResp.Body.Close()

	// Read the event history response body
	eventsBody, err := ioutil.ReadAll(eventsResp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	var events []EventHistory
	err = json.Unmarshal(eventsBody, &events)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to parse patient's history response"})
		return
	}

	context.HTML(http.StatusOK, "event_history.html", gin.H{
		"Patient": patient,
		"Events":  events,
	})
}
