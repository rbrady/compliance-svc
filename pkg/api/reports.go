package api

import (
	"fmt"
	"net/http"

	// "path/filepath"

	"github.com/rbrady/cappy/pkg/db"
	"github.com/gin-gonic/gin"
)

// Reports actions (Enterprise)
// Get a list of reports
// router.GET("/reports", api.GetReports)
// super simple method here.  probably need the ability to page the records
func GetReports(c *gin.Context) {
	var reports []db.Report
	db, _ := db.Connect()
	var result = db.Find(&reports)
	if result.Error != nil {
		c.Error(result.Error)
	}
	c.JSON(http.StatusOK, reports)
}

// Get a specific report
// router.GET("/reports/:id", api.GetReport)
func GetReport(c *gin.Context) {
	reportId := c.Param("id")

	var report db.Report
	db, _ := db.Connect()
	result := db.Find(&report, "Id = ?", reportId).Limit(1)
	if result.Error != nil {
		// log the error
		// return a no report error
		c.Error(result.Error) // obviously change this to something user friendly
	} else if result.RowsAffected != 1 {
		c.String(http.StatusOK, fmt.Sprintf("No reports forund for %d", reportId))
	}
	c.JSON(http.StatusOK, report)
}

// Update a report
func UpdateReport(c *gin.Context) {
	reportId := c.Param("id")
	var json db.Report
	if c.BindJSON(&json) == nil {
		// binding worked
	}

	var report db.Report
	db, _ := db.Connect()
	result := db.Find(&report, "Id = ?", reportId).Limit(1)
	if result.Error != nil {
		// log the error
		// return a no report error
		c.Error(result.Error) // obviously change this to something user friendly
	} else if result.RowsAffected != 1 {
		c.String(http.StatusOK, fmt.Sprintf("No reports found for %d", reportId))
	}

	// update report with json values
	report.Source = json.Source
	report.Result = json.Result
	report.Pod = json.Pod
	report.Namespace = json.Namespace
	report.Tag = json.Tag
	report.RepoDigest = json.RepoDigest

	db.Save(&report)

	c.String(http.StatusOK, "Success")

}
// Delete a report
func DeleteReport(c *gin.Context) {
	c.String(http.StatusOK, "{}")
}

// Report files actions (Enterprise)
// gets one or more files for a given report id
// this may be taken care of in GetReport...depending on the query results
//func GetReportFiles(c *gin.Context) {
//	c.String(http.StatusOK, "{}")
//}