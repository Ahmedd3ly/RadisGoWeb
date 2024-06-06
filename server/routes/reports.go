package routes

import (
	"net/http"

	"RadisGoWeb/server/models"

	"github.com/gin-gonic/gin"
)

func getReportDetailsById(context *gin.Context) {
	requestId := context.Query("requestId")
	reports, err := models.GetReportDetailsById(requestId)

	if requestId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if reports == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "No reports found for this request"})
		return
	}
	context.JSON(http.StatusOK, reports)
}
