package api

import (
	"fmt"
	"net/http"
	"github.com/rbrady/cappy/pkg/data"
	"github.com/gin-gonic/gin"
)

func getReport(id string) (*data.Report, error) {
	var report data.Report
	var reportFiles []data.ReportFile

	db, _ := data.Connect()
	result := db.Find(&report, "Id = ?", id).Limit(1)
	if result.Error != nil {
		// log the error
		// return a no report error
		return &report, result.Error // obviously change this to something user friendly
	}

	db.Model(&report).Association("Files").Find(&reportFiles)
	report.Files = reportFiles
	return &report, nil
}

// Reports actions (Enterprise)
func PostReport(c *gin.Context) {
	var json data.Report
	if c.BindJSON(&json) == nil {
		// binding worked
	}

	var report data.Report
	db, _ := data.Connect()

	// update report with json values
	report.Source = json.Source
	report.Result = json.Result
	report.Pod = json.Pod
	report.Namespace = json.Namespace
	report.Tag = json.Tag
	report.RepoDigest = json.RepoDigest

	db.Save(&report)

	c.JSON(http.StatusCreated, report)

}


// Get a list of reports
// router.GET("/reports", api.GetReports)
// super simple method here.  probably need the ability to page the records
func GetReports(c *gin.Context) {
	var reports []data.Report
	db, _ := data.Connect()
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
	report, err := getReport(reportId)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, report)
}

// Update a report
func UpdateReport(c *gin.Context) {
	reportId := c.Param("id")
	var json data.Report
	if c.BindJSON(&json) == nil {
		// binding worked
	}

	report, err := getReport(reportId)
	if err != nil {
		c.Error(err)
	}

	// update report with json values
	report.Source = json.Source
	report.Result = json.Result
	report.Pod = json.Pod
	report.Namespace = json.Namespace
	report.Tag = json.Tag
	report.RepoDigest = json.RepoDigest

	db, _ := data.Connect()
	db.Save(&report)

	c.JSON(http.StatusOK, "Success")

}

// Delete a report
func DeleteReport(c *gin.Context) {
	reportId := c.Param("id")
	db, _ := data.Connect()
	// Soft Delete, record isn't gone just hidden from queries.
	// see https://gorm.io/id_ID/docs/delete.html for more info
	result := db.Delete(&data.Report{}, reportId)
	if result.Error != nil {
		// log the error
		// return a no report error
		c.Error(result.Error) // obviously change this to something user friendly
	} else if result.RowsAffected != 1 {
		c.JSON(http.StatusNotFound, fmt.Sprintf("No reports found for %d", reportId))
	}
	c.JSON(http.StatusOK, "Success")
}

// Report files actions (Enterprise)
// gets one or more files for a given report id
// this may be taken care of in GetReport...depending on the query results
func GetReportFiles(c *gin.Context) {
	reportId := c.Param("id")

	var report data.Report

	db, _ := data.Connect()
	result := db.Find(&report, "Id = ?", reportId).Limit(1)
	if result.Error != nil {
		// log the error
		// return a no report error
		c.Error(result.Error) // obviously change this to something user friendly
	}

	var reportFiles []data.ReportFile
	db.Model(&report).Association("Languages").Find(&reportFiles)

	c.JSON(http.StatusOK, reportFiles)
}