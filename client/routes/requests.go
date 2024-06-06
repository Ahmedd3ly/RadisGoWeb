package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestDetail struct {
	Id                       string    `gorm:"column:pk_Request_ID" json:"id"`
	RequestStatus            int16     `json:"request_status"`
	RequestStatusDescription string    `json:"request_status_description"`
	PriorityDescription      *string   `json:"priority_description"`
	Alerts                   *string   `json:"alerts,omitempty"`
	ClinicalComments         *string   `json:"clinical_comments,omitempty"`
	DateReceived             time.Time `json:"date_received"`
	BookedDate               time.Time `json:"booked_date"`
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

type Procedure struct {
	ScanDate             *time.Time `json:"scan_date"`
	ExamSide             *string    `json:"exam_side"`
	StartTime            *time.Time `json:"start_time"`
	RoomCode             *string    `json:"room_code"`
	RoomDescription      *string    `json:"room_description"`
	ProcedureCode        *string    `gorm:"column:ProcedureCodeCode" json:"procedure_code"`
	ProcedureDescription *string    `gorm:"column:ProcedureCodeDescription" json:"procedure_description"`
}

func requestProcedureDetailByIdHandler(context *gin.Context) {
	requestId := context.Param("requestId")

	if requestId == "" {
		context.HTML(http.StatusBadRequest, "appointments.html", gin.H{"Error": "A request id is required"})
		return
	}

	// Construct the API URLs
	requestDetailURL := fmt.Sprintf("http://localhost:8081/radisweb/requests/request-details?requestId=%s", requestId)
	procedureDetailURL := fmt.Sprintf("http://localhost:8081/radisweb/procedures?requestId=%s", requestId)

	// Make the API request for request details
	requestResp, err := http.Get(requestDetailURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "appointments.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}
	defer requestResp.Body.Close()

	// Read the response body for request details
	requestBody, err := ioutil.ReadAll(requestResp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "appointments.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Parse the response body into the RequestDetail struct
	var requestDetail RequestDetail
	err = json.Unmarshal(requestBody, &requestDetail)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "appointments.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	// Make the API request for procedure details
	procedureResp, err := http.Get(procedureDetailURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "appointments.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}
	defer procedureResp.Body.Close()

	// Read the response body for procedure details
	procedureBody, err := ioutil.ReadAll(procedureResp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "appointments.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Parse the response body into the ProcedureDetail struct
	var procedureDetails []Procedure
	err = json.Unmarshal(procedureBody, &procedureDetails)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "appointments.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	// Render the template with the combined data
	context.HTML(http.StatusOK, "appointments.html", gin.H{
		"Request":    requestDetail,
		"Procedures": procedureDetails,
	})
}
