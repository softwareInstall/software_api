package service

import (
	"software_api/models"
	"time"
)

type SoftwareVersionService struct {
	ID           int
	SoftwareID   int
	Version      string
	Description  string
	ObjectName   string
	OriginalName string
	Offset       int
	Limit        int
}

func (s *SoftwareVersionService) Add() error {
	currentTime := time.Now()
	SoftwareVersion := models.Version{
		SoftwareID:   s.SoftwareID,
		Version:      s.Version,
		Description:  s.Description,
		ObjectName:   s.ObjectName,
		OriginalName: s.OriginalName,
		ReleaseDate:  &currentTime,
	}

	if err := models.AddSoftwareVersion(SoftwareVersion); err != nil {
		return err
	}

	return nil
}

func (s *SoftwareVersionService) Edit() error {
	return models.EditSoftware(s.ID, map[string]interface{}{
		"Description": s.Description,
	})
}

func (s *SoftwareVersionService) ExistByID() (bool, error) {
	return models.ExistSoftwareByID(s.ID)
}

func (s *SoftwareVersionService) ExistByVersion() (bool, error) {
	return models.ExistSoftwareByName(s.Version)
}

func (s *SoftwareVersionService) Count() (int64, error) {
	return models.GetSoftwareTotal(s.getMaps())
}

func (s *SoftwareVersionService) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	return maps
}
