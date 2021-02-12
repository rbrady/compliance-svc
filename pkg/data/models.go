package data

import (
	"gorm.io/gorm"
)

/*
Reports

source - what tool created this report
created - the timestamp this record was created at
result - pass, fail, inconclusive
files - zero or more files associated with a compliance report
container_info - we need to know about the container image
*/

type Report struct {
	gorm.Model
	Source  		string			`form:"source" json:"source" binding:"required"`
	Result 		  	string			`form:"result" json:"result" binding:"required"`
	Files 		  	[]ReportFile	`json:"files"`
	Pod				string			`form:"pod" json:"pod" binding:"required"`
	Namespace		string			`form:"namespace" json:"namespace" binding:"required"`
	Tag				string			`form:"tag" json:"tag" binding:"required"`
	RepoDigest		string			`form:"repoDigest" json:"repoDigest" binding:"required"`
}

type ReportFile struct {
	gorm.Model
	ReportID	uint
	FileName 	string	`json:"file_name"`
	FilePath 	string	`json:"file_path"`
}

