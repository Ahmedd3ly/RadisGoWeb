package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func basicSearchHandler(context *gin.Context) {
	id := context.Query("id")
	field := context.Query("field")

	if id == "" || field == "" {
		context.HTML(http.StatusBadRequest, "search_results.html", gin.H{"Error": "A patient's number is required"})
		return
	}

	apiURL := fmt.Sprintf("http://localhost:8081/radisweb/patients/id-search?field=%s&id=%s", field, id)
	resp, err := http.Get(apiURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	var patients []Patient
	if err := json.Unmarshal(body, &patients); err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	context.HTML(http.StatusOK, "search_results.html", gin.H{"Patients": patients})
}

func advancedSearchhandler(context *gin.Context) {
	surname := context.Query("surname")
	forename := context.Query("forename")
	dob := context.Query("dob")
	sex := context.Query("sex")

	if surname == "" && dob == "" {
		context.HTML(http.StatusBadRequest, "search_results.html", gin.H{"Error": "A patient's surname or date of birth is required"})
		return
	}

	// Construct the API URL
	apiURL := fmt.Sprintf(
		"http://localhost:8081/radisweb/patients/advanced-search?surname=%s&forename=%s&dob=%s&sex=%s",
		surname, forename, dob, sex,
	)

	// Create a new HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to make request to external API"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Unmarshal the response body into a slice of patients
	var patients []Patient
	if err := json.Unmarshal(body, &patients); err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	context.HTML(http.StatusOK, "search_results.html", gin.H{"Patients": patients})
}

func searchByWardhandler(context *gin.Context) {
	locationId := context.Query("locationId")

	if locationId == "" {
		context.HTML(http.StatusBadRequest, "search_results.html", gin.H{"Error": "A location Id is required"})
		return
	}

	// Construct the API URL
	apiURL := fmt.Sprintf("http://localhost:8081/radisweb/patients/by-ward-last-week?locationId=%s", locationId)

	// Create a new HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to make request to external API"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Unmarshal the response body into a slice of patients
	var patients []Patient
	if err := json.Unmarshal(body, &patients); err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	context.HTML(http.StatusOK, "search_results.html", gin.H{"Patients": patients})
}

func searchByConsultanthandler(context *gin.Context) {
	clinicianPracticeId := context.Query("clinicianPracticeId")

	if clinicianPracticeId == "" {
		context.HTML(http.StatusBadRequest, "search_results.html", gin.H{"Error": "A clinician Id is required"})
		return
	}

	// Construct the API URL
	apiURL := fmt.Sprintf("http://localhost:8081/radisweb/patients/by-consultant-last-week?clinicianPracticeId=%s", clinicianPracticeId)

	// Create a new HTTP requestd
	resp, err := http.Get(apiURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to make request to external API"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Unmarshal the response body into a slice of patients
	var patients []Patient
	if err := json.Unmarshal(body, &patients); err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	context.HTML(http.StatusOK, "search_results.html", gin.H{"Patients": patients})
}

func searchByWardAppointmentshandler(context *gin.Context) {
	locationId := context.Query("locationId")

	if locationId == "" {
		context.HTML(http.StatusBadRequest, "search_results.html", gin.H{"Error": "A location Id is required"})
		return
	}

	// Construct the API URL
	apiURL := fmt.Sprintf("http://localhost:8081/radisweb/patients/by-ward?locationId=%s", locationId)

	// Create a new HTTP request
	resp, err := http.Get(apiURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to make request to external API"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Unmarshal the response body into a slice of patients
	var patients []Patient
	if err := json.Unmarshal(body, &patients); err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	context.HTML(http.StatusOK, "search_results.html", gin.H{"Patients": patients})
}

func searchByConsultantAppointmentshandler(context *gin.Context) {
	clinicianPracticeId := context.Query("clinicianPracticeId")

	if clinicianPracticeId == "" {
		context.HTML(http.StatusBadRequest, "search_results.html", gin.H{"Error": "A clinician Id is required"})
		return
	}

	// Construct the API URL
	apiURL := fmt.Sprintf("http://localhost:8081/radisweb/patients/by-consultant?clinicianPracticeId=%s", clinicianPracticeId)

	// Create a new HTTP requestd
	resp, err := http.Get(apiURL)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to make request to external API"})
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to read response from external API"})
		return
	}

	// Unmarshal the response body into a slice of patients
	var patients []Patient
	if err := json.Unmarshal(body, &patients); err != nil {
		context.HTML(http.StatusInternalServerError, "search_results.html", gin.H{"Error": "Failed to parse response from external API"})
		return
	}

	context.HTML(http.StatusOK, "search_results.html", gin.H{"Patients": patients})
}
