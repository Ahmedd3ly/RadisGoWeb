package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func requestReportDetailByIdHandler(context *gin.Context) {
	requestId := context.Param("requestId")

	if requestId == "" {
		context.HTML(http.StatusBadRequest, "reports.html", gin.H{"Error": "A request id is required"})
		return
	}

	// Construct the API URLs
	requestDetailURL := fmt.Sprintf("http://localhost:8081/radisweb/requests/request-details?requestId=%s", requestId)
	reportDetailURL := fmt.Sprintf("http://localhost:8081/radisweb/reports/report-details?requestId=%s", requestId)

	// Make the API request for request details
	requestResp, err := http.Get(requestDetailURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "reports.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}
	defer requestResp.Body.Close()

	// Read the response body for request details
	requestBody, err := ioutil.ReadAll(requestResp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "reports.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Parse the response body into the RequestDetail struct
	var requestDetail RequestDetail
	err = json.Unmarshal(requestBody, &requestDetail)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "reports.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	// Make the API request for report details
	reportResp, err := http.Get(reportDetailURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "reports.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}
	defer reportResp.Body.Close()

	// Read the response body for report details
	reportBody, err := ioutil.ReadAll(reportResp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "reports.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Parse the response body into the ProcedureDetail struct
	var reportDetails []ReportDetail
	err = json.Unmarshal(reportBody, &reportDetails)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "reports.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}
	fmt.Println("Data: ", reportDetails)
	// Render the template with the combined data
	context.HTML(http.StatusOK, "reports.html", gin.H{
		"Request": requestDetail,
		"Reports": reportDetails,
	})
}
