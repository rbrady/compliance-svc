package api

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/rbrady/cappy/pkg/data"

	"github.com/gin-gonic/gin"
)

// PostFiles
// This function handles an API request to upload one or more files associated with a compliance report
func PostFiles(c *gin.Context) {
	reportId := c.Param("id")
	var parsedInt, err = strconv.ParseInt(reportId, 10, 32)

	// if we cannot find a matching reportId, no reason to continue
	var report data.Report
	db, _ := data.Connect()
	result := db.Find(&report, "Id = ?", reportId)
	if result.RowsAffected < 1 {
		// return a no report error
	}


	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		// TODO: (rbrady) get file storage location from config
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		reportFile := data.ReportFile{
			FileName: filename,
			ReportID: uint(parsedInt),
		}
		fileResult := db.Create(&reportFile)
		fmt.Println(fileResult)
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields reportId=%s.", len(files), reportId))
}
