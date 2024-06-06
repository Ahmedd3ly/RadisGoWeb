package routes

import (
	"net/http"

	"RadisGoWeb/server/models"

	"github.com/gin-gonic/gin"
)

func getLibraryTransactionByPatientId(context *gin.Context) {
	patientId := context.Query("patientId")
	transactions, err := models.GetLibraryTransactionByPatientId(patientId)

	if patientId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if transactions == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "No library transactions found for this patient"})
		return
	}
	context.JSON(http.StatusOK, transactions)
}

func getLibraryDetailBySiteLibraryId(context *gin.Context) {
	siteLibraryId := context.Query("siteLibraryId")
	transactions, err := models.GetLibraryDetails(siteLibraryId)

	if siteLibraryId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if transactions == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "No library details found"})
		return
	}
	context.JSON(http.StatusOK, transactions)
}
