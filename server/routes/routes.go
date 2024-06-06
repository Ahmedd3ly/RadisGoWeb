package routes

import (
	"RadisGoWeb/server/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//Patient
	server.GET("/radisweb/patients/surname", getPatientsBySurname)
	server.GET("/radisweb/patients/id-search", getPatientById)
	server.GET("/radisweb/patients/advanced-search", getPatientsBySurnameForenameDobSex)
	server.GET("/radisweb/patients/patient-history", getPatientHistoryById)
	server.GET("/radisweb/patients/patient-details", getPatientDetailsById)
	server.GET("/radisweb/patients/by-ward", getPatientsByWardWithScheduledExams)
	server.GET("/radisweb/patients/by-consultant", getPatientsForConsultantWithScheduledExams)
	server.GET("/radisweb/patients/by-ward-last-week", getPatientsForWardWithinLastWeek)
	server.GET("/radisweb/patients/by-consultant-last-week", getPatientsForConsultantWithinTheLastWeek)

	//Requests
	server.GET("/radisweb/requests/request-details", getRequestDetailsById)

	//Reports
	server.GET("/radisweb/reports/report-details", getReportDetailsById)

	//Procedures
	server.GET("/radisweb/procedures", getProceduresByRequestId)

	//Library
	server.GET("/radisweb/library-transactions", getLibraryTransactionByPatientId)
	server.GET("/radisweb/library-details", getLibraryDetailBySiteLibraryId)

	//Reference Data
	server.GET("/radisweb/data/consultants", getConsultants)
	server.GET("/radisweb/data/locations", getLocations)

	//protected routes
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	// authenticated.POST("/events", createEvent)
	// authenticated.PUT("/events/:id", updateEvent)
	// authenticated.DELETE("/events/:id", deleteEvent)

	// //signup
	// server.POST("/signup", signUp)
	// //login
	// server.POST("/login", login)
}
