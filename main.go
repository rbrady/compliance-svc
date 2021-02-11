package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rbrady/cappy/pkg/api"
	"github.com/rbrady/cappy/pkg/db"
)

func main() {

	// database
	db.Migrate()
	// db.AddSampleReports()

	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// part of the example
	//router.Static("/", "./public")

	// TODO: (rbrady) Does it make sense to separate the service into two parts, so the only action offered to
	// compliance tools is to post a report and post files?

	// Compliance Tools Actions
	// Saves a given report and returns an identifier back to the caller
	// router.POST("/reports", api.PostReport)
	// provides path for uploading one or more files for a given report
	// router.POST("reports/:id/files", api.PostFiles)

	// Reports actions (Enterprise)
	// Get a list of reports
	router.GET("/reports", api.GetReports)
	// Get a specific report
	router.GET("/reports/:id", api.GetReport)
	// Update a report
	router.POST("/reports/:id", api.UpdateReport)
	// Delete a report
	// router.Delete("/reports/:id", api.DeleteReport)

	// Report files actions (Enterprise)
	// gets one or more files for a given report id
	// router.GET("reports/:reportId/files", api.GetFiles)

	router.Run(":8080")
}
