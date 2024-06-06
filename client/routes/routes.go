package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// Templates
	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	server.GET("/search-results", func(c *gin.Context) {
		c.HTML(200, "search_results.html", gin.H{})
	})

	server.GET("/event-history", func(c *gin.Context) {
		c.HTML(200, "event_history.html", gin.H{})
	})

	server.GET("/appointments", func(c *gin.Context) {
		c.HTML(200, "appointments.html", gin.H{})
	})

	server.GET("/reports", func(c *gin.Context) {
		c.HTML(200, "reports.html", gin.H{})
	})

	//Patient
	server.GET("/radisweb/patients/id-search", basicSearchHandler)
	server.GET("/radisweb/patients/advanced-search", advancedSearchhandler)
	server.GET("/radisweb/patients/by-ward-last-week", searchByWardhandler)
	server.GET("/radisweb/patients/by-consultant-last-week", searchByConsultanthandler)
	server.GET("/radisweb/patients/by-ward", searchByWardAppointmentshandler)
	server.GET("/radisweb/patients/by-consultant", searchByConsultantAppointmentshandler)
	server.GET("/radisweb/patients/patient-details-history/:patientId", patientDetailsHistoryByIdHandler)
	server.GET("/radisweb/requests/request-details/:requestId", requestProcedureDetailByIdHandler)
	server.GET("/radisweb/reports/report-details/:requestId", requestReportDetailByIdHandler)
}
