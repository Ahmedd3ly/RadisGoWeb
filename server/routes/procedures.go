package routes

import (
	"net/http"

	"RadisGoWeb/server/models"

	"github.com/gin-gonic/gin"
)

func getProceduresByRequestId(context *gin.Context) {
	requestId := context.Query("requestId")
	procedures, err := models.GetProceduresByRequestId(requestId)

	if requestId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if procedures == nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "No procedures found for this request"})
		return
	}
	context.JSON(http.StatusOK, procedures)
}
