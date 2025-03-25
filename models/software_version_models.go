package models

import (
	"log"
	"time"
)

// Version 表示软件的具体版本信息
type Version struct {
	Model

	SoftwareID   int        `json:"software_id"`
	Version      string     `json:"version"`
	ReleaseDate  *time.Time `json:"release_date"`
	Description  string     `json:"description"`
	ObjectName   string     `json:"object_name"`
	OriginalName string     `json:"original_name"`
}

// AddSoftwareVersion 添加新软件
func AddSoftwareVersion(data Version) error {
	err := db.Create(&data).Error
	if err != nil {
		log.Printf("AddSoftwareVersion error: %v", err)
		return err
	}
	softwareDate := Software{
		LatestReleaseDate:  data.ReleaseDate,
		LatestObjectName:   data.ObjectName,
		LatestOriginalName: data.OriginalName,
		LatestVersion:      data.Version,
	}
	err = db.Model(&Software{}).Where("id = ?", data.SoftwareID).Updates(softwareDate).Error
	if err != nil {
		log.Printf("EditSoftware error: %v", err)
		return err
	}
	return nil
}
