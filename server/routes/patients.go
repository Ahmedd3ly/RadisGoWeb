package routes

import (
	"net/http"

	"RadisGoWeb/server/models"

	"github.com/gin-gonic/gin"
)

func getPatientsBySurname(context *gin.Context) {
	surname := context.Query("surname")
	if surname == "" {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
		return
	}

	patients, err := models.GetPatientsBySurname(surname)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get patient search data"})
		return
	}

	if patients == nil {
		context.JSON(http.StatusNotFound, gin.H{"Error": "No patients found for this search"})
		return
	}
	context.JSON(http.StatusOK, patients)
}

func getPatientById(context *gin.Context) {
	id := context.Query("id")
	field := context.Query("field")

	if id == "" || field == "" {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
		return
	}

	patients, err := models.GetPatientsById(id, field)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get patient search data"})
		return
	}

	if patients == nil {
		context.JSON(http.StatusNotFound, gin.H{"Error": "No patients found for this search"})
		return
	}

	context.JSON(http.StatusOK, patients)
}

func getPatientsBySurnameForenameDobSex(context *gin.Context) {
	surname := context.Query("surname")
	forename := context.Query("forename")
	dob := context.Query("dob")
	sex := context.Query("sex")

	if surname == "" && dob == "" {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
		return
	}

	patients, err := models.GetPatientsBySurnameForenameDobSex(surname, forename, dob, sex)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get patient search data"})
		return
	}

	if patients == nil {
		context.JSON(http.StatusNotFound, gin.H{"Error": "No patients found for this search"})
		return
	}

	context.JSON(http.StatusOK, patients)
}

func getPatientsByWardWithScheduledExams(context *gin.Context) {
	locationId := context.Query("locationId")
	patients, err := models.GetPatientsByWardWithScheduledExams(locationId)

	if locationId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get patient search data"})
		return
	}

	if patients == nil {
		context.JSON(http.StatusNotFound, gin.H{"Error": "No patients found for this search"})
		return
	}
	context.JSON(http.StatusOK, patients)
}

func getPatientsForWardWithinLastWeek(context *gin.Context) {
	locationId := context.Query("locationId")
	patients, err := models.GetPatientsForWardWithinLastWeek(locationId)

	if locationId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get patient search data"})
		return
	}

	if patients == nil {
		context.JSON(http.StatusNotFound, gin.H{"Error": "No patients found for this search"})
		return
	}
	context.JSON(http.StatusOK, patients)
}

func getPatientsForConsultantWithScheduledExams(context *gin.Context) {
	clinicianPracticeId := context.Query("clinicianPracticeId")
	patients, err := models.GetPatientsForConsultantWithScheduledExams(clinicianPracticeId)

	if clinicianPracticeId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get patient search data"})
		return
	}

	if patients == nil {
		context.JSON(http.StatusNotFound, gin.H{"Error": "No patients found for this search"})
		return
	}
	context.JSON(http.StatusOK, patients)
}

func getPatientsForConsultantWithinTheLastWeek(context *gin.Context) {
	clinicianPracticeId := context.Query("clinicianPracticeId")
	patients, err := models.GetPatientsForConsultantWithinTheLastWeek(clinicianPracticeId)

	if clinicianPracticeId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get patient search data"})
		return
	}

	if patients == nil {
		context.JSON(http.StatusNotFound, gin.H{"Error": "No patients found for this search"})
		return
	}
	context.JSON(http.StatusOK, patients)
}

func getPatientDetailsById(context *gin.Context) {
	patientId := context.Query("patientId")
	patient, err := models.GetPatientDetails(patientId)

	if patientId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Error": "Could not get patient search data"})
		return
	}

	if patient == nil {
		context.JSON(http.StatusNotFound, gin.H{"Error": "No patients found for this search"})
		return
	}
	context.JSON(http.StatusOK, patient)
}
