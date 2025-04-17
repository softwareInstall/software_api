package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"software_api/dto"
	"software_api/models"
	"software_api/pkg/oss"
	"software_api/pkg/util"
	"strconv"
)

type SoftwareService struct {
	ID       int
	Name     string
	Platform int
	Offset   int
	Limit    int
	Filter   map[string]interface{}
}

func (s *SoftwareService) Add() error {
	Software := models.Software{
		Name:     s.Name,
		Platform: s.Platform,
	}

	if err := models.AddSoftware(Software); err != nil {
		return err
	}

	return nil
}

func (s *SoftwareService) Edit() error {
	return models.EditSoftware(s.ID, map[string]interface{}{
		"name": s.Name,
	})
}

func (s *SoftwareService) GetAll() ([]*models.Software, error) {
	var (
		Software []*models.Software
	)

	Software, err := models.GetSoftwareList(s.Offset, s.Limit, s.getMaps(), "id", "desc", "id", "desc")
	if err != nil {
		return nil, err
	}

	return Software, nil
}

func (s *SoftwareService) GetSimpleInfo() ([]*models.Software, error) {
	var (
		Software []*models.Software
	)

	Software, err := models.GetSoftwareNameList(s.Offset, s.Limit, s.getMaps(), "id", "desc", "id", "desc")
	if err != nil {
		return nil, err
	}

	return Software, nil
}

func (s *SoftwareService) ExistByID() (bool, error) {
	return models.ExistSoftwareByID(s.ID)
}

func (s *SoftwareService) ExistByName() (bool, error) {
	return models.ExistSoftwareByName(s.Name)
}

func (s *SoftwareService) ExistByFilter() (bool, error) {
	filter := map[string]interface{}{
		"name":     s.Name,
		"platform": s.Platform,
	}
	return models.ExistSoftwareByFilter(filter)
}

func (s *SoftwareService) Count() (int64, error) {
	return models.GetSoftwareTotal(s.getMaps())
}

func (s *SoftwareService) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	for key, val := range s.Filter {
		maps[key] = val
	}

	return maps
}

func (s *SoftwareService) ConvertSoftwareModelsToNameDts(ctx *gin.Context, models []*models.Software) []*dto.SoftwareNameIdDto {
	data := make([]*dto.SoftwareNameIdDto, 0)
	for _, model := range models {
		if model == nil {
			continue
		}
		item := &dto.SoftwareNameIdDto{
			Value: strconv.Itoa(model.ID),
			Label: model.Name,
		}
		data = append(data, item)
	}

	return data
}

func (s *SoftwareService) ConvertSoftwareModelsToDts(ctx *gin.Context, models []*models.Software) []*dto.SoftwareDto {
	services := make([]*dto.SoftwareDto, 0)
	for _, model := range models {
		if model == nil {
			continue
		}

		software := convertSoftwareModelToDto(ctx, model)

		if len(model.Versions) > 0 {
			software.Versions = convertVersionsToDto(ctx, model.Versions)
		}

		services = append(services, software)
	}

	return services
}

// 将单个 Software 模型转换为 DTO
func convertSoftwareModelToDto(ctx *gin.Context, model *models.Software) *dto.SoftwareDto {
	latestDownloadURL, err := oss.ManagerOssObj.GetFileUrl(ctx, model.LatestObjectName)
	if err != nil {
		log.Printf("获取文件下载地址失败：%v", err)
	}
	return &dto.SoftwareDto{
		ID:                model.ID,
		Name:              model.Name,
		LatestReleaseDate: util.FormatTimePtrCustom(model.LatestReleaseDate),
		LatestDownloadURL: latestDownloadURL,
		LatestVersion:     model.LatestVersion,
		Versions:          []dto.VersionDto{},
	}
}

// 将多个 Version 模型转换为 DTO
func convertVersionsToDto(ctx *gin.Context, versions []*models.Version) []dto.VersionDto {
	dtoVersions := make([]dto.VersionDto, 0)

	for _, version := range versions {
		if version == nil {
			continue
		}
		downloadURL, err := oss.ManagerOssObj.GetFileUrl(ctx, version.ObjectName)
		if err != nil {
			log.Printf("获取文件下载地址失败：%v", err)
		}
		dtoVersions = append(dtoVersions, dto.VersionDto{
			ID:          version.ID,
			SoftwareID:  version.SoftwareID,
			Version:     version.Version,
			ReleaseDate: util.FormatTimePtrCustom(version.ReleaseDate),
			Description: version.Description,
			DownloadURL: downloadURL,
		})
	}

	return dtoVersions
}
