package routes

import (
	"net/http"

	"RadisGoWeb/server/models"

	"github.com/gin-gonic/gin"
)

func getConsultants(context *gin.Context) {
	consultants, err := models.GetConsultants()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get patient search data"})
		return
	}
	context.JSON(http.StatusOK, consultants)
}

func getLocations(context *gin.Context) {
	locations, err := models.GetLocations()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get patient search data"})
		return
	}
	context.JSON(http.StatusOK, locations)
}
