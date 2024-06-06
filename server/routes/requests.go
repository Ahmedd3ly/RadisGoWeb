package routes

import (
	"net/http"

	"RadisGoWeb/server/models"

	"github.com/gin-gonic/gin"
)

func getRequestDetailsById(context *gin.Context) {
	requestId := context.Query("requestId")
	request, err := models.GetRequestDetailsById(requestId)

	if requestId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if request == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Request not found"})
		return
	}
	context.JSON(http.StatusOK, request)
}
