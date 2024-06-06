package routes

import (
	"fmt"
	"net/http"

	"RadisGoWeb/server/models"

	"github.com/gin-gonic/gin"
)

func getPatientHistoryById(context *gin.Context) {
	patientId := context.Query("patientId")
	fmt.Print("Patient Id:", patientId)
	events, err := models.GetPatientHistoryById(patientId)

	if patientId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}
