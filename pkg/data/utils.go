package data

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, err
}

func Migrate()  {
	db, _ := Connect()
	db.AutoMigrate(&Report{}, &ReportFile{})
}

func AddSampleReports() {
	db, _ := Connect()
	report := Report{
		Source: "openscap",
		Result: "pass",
		Pod: "anchore-gub26sf",
		Namespace: "default",
		Tag: "anchore/enterprise-dev:latest",
		RepoDigest: "sha256:9fa0ec5177e761494ae5fb4d71a74437befc515f88cad03f31117996ab8f606f",
	}

	reportResult := db.Create(&report)
	if reportResult.Error != nil {
		fmt.Sprintf("An error occured: %s", reportResult.Error)
	}

	reportFiles := ReportFile {
			FileName: "2a493d4d-8ca5-4c94-bbcb-6cf9dd2c44f3.xml",
			FilePath: "/static/",
			ReportID: report.ID,
	}

	fileResult := db.Create(&reportFiles)
	if fileResult.Error != nil {
		fmt.Sprintf("An error occured: %s", fileResult.Error)
	}

	fmt.Sprintf("A record was created with id: %s", report)

}