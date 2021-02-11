package api

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/rbrady/cappy/pkg/db"

	"github.com/gin-gonic/gin"
)

// PostFiles
// This function handles an API request to upload one or more files associated with a compliance report
func PostFiles(c *gin.Context) {
	reportId := c.PostForm("reportId")

	// if we cannot find a matching reportId, no reason to continue
	var report db.Report
	db, _ := db.Connect()
	result := db.Find(&report, "Id = ?", reportId)
	if result.RowsAffected < 1 {
		// return a no report error
	}


	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		// TODO: (rbrady) get file storage location from config
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}
		// TODO: (rbrady) after successful file save, update the db
		// fileObject.create(reportId, filePath)
	}

	c.String(http.StatusOK, fmt.Sprintf("Uploaded successfully %d files with fields reportId=%s.", len(files), reportId))
}
