package db

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
		Files: []ReportFile{
			ReportFile{
				FileName: "2a493d4d-8ca5-4c94-bbcb-6cf9dd2c44f3.xml",
				FilePath: "/static/",
			},
			ReportFile {
				FileName: "b3eb08ee-3706-45c5-b678-be12d45eaae2.json",
				FilePath: "/static/",
			},
		},
		Pod: "anchore-gub26sf",
		Namespace: "default",
		Tag: "anchore/enterprise-dev:latest",
		RepoDigest: "sha256:9fa0ec5177e761494ae5fb4d71a74437befc515f88cad03f31117996ab8f606f",
	}

	result := db.Create(&report)
	if result.Error != nil {
		fmt.Sprintf("An error occured: %s", result.Error)
	}

	fmt.Sprintf("A record was created with id: %s", report)

}