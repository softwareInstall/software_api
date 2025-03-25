package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"time"
)

type Software struct {
	Model

	Platform           int        `json:"platform" gorm:"type:integer"`
	Name               string     `json:"name"`
	LatestReleaseDate  *time.Time `json:"latest_release_date" gorm:"default:NULL"`
	LatestObjectName   string     `json:"latest_object_name"`
	LatestOriginalName string     `json:"latest_original_name"`
	LatestVersion      string     `json:"latest_version"`
	Versions           []*Version
}

type SoftwareDto struct {
	Platform int    `json:"platform"`
	Name     string `json:"name"`
}

// ExistSoftwareByID 检查软件是否存在（基于 ID）
func ExistSoftwareByID(id int) (bool, error) {
	var software Software
	err := db.Select("id").Where("id = ?", id).First(&software).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("ExistSoftwareByID error: %v", err)
		return false, err
	}
	return software.ID > 0, nil
}

// ExistSoftwareByName 检查软件是否存在（基于名称）
func ExistSoftwareByName(name string) (bool, error) {
	var count int64
	err := db.Model(&Software{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		log.Printf("ExistSoftwareByName error: %v", err)
		return false, err
	}
	return count > 0, nil
}

// ExistSoftwareByFilter 检查软件是否存在（基于名称\平台等等）
func ExistSoftwareByFilter(maps map[string]interface{}) (bool, error) {
	var count int64
	var err error

	// 构建查询条件
	query := db.Model(&Software{})
	query = query.Where(maps)

	// 执行查询
	err = query.Count(&count).Error
	if err != nil {
		log.Printf("ExistSoftwareByFilter error: %v", err)
		return false, err
	}

	// 返回结果
	return count > 0, nil
}

// GetSoftwareNameList 获取软件名列表（支持分页和条件查询）
func GetSoftwareNameList(offset int, limit int, maps map[string]interface{}, orderBy string, desc string, versionsOrderBy string, versionsDesc string) ([]*Software, error) {
	if orderBy == "" {
		orderBy = "id"
	}
	if desc == "" {
		desc = "desc"
	}
	order := orderBy + " " + desc

	var software []*Software
	query := db.Select("id", "name")
	if len(maps) > 0 {
		query = query.Where(maps)
	}
	err := query.Offset(offset).Limit(limit).Order(order).Find(&software).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("GetSoftwareNameList error: %v", err)
		return nil, err
	}
	return software, nil
}

// GetSoftwareList 获取软件列表（支持分页和条件查询）
func GetSoftwareList(offset int, limit int, maps map[string]interface{}, orderBy string, desc string, versionsOrderBy string, versionsDesc string) ([]*Software, error) {
	if orderBy == "" {
		orderBy = "id"
	}
	if desc == "" {
		desc = "desc"
	}
	order := orderBy + " " + desc
	versionOrder := versionsOrderBy + " " + (versionsDesc)

	var software []*Software
	query := db.Preload("Versions", func(db *gorm.DB) *gorm.DB {
		return db.Order(versionOrder)
	})
	if len(maps) > 0 {
		query = query.Where(maps)
	}
	err := query.Offset(offset).Limit(limit).Order(order).Find(&software).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("GetSoftwareList error: %v", err)
		return nil, err
	}
	return software, nil
}

// EditSoftware 修改软件信息
func EditSoftware(id int, data map[string]interface{}) error {
	if id <= 0 {
		return errors.New("invalid software ID")
	}
	err := db.Model(&Software{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		log.Printf("EditSoftware error: %v", err)
		return err
	}
	return nil
}

// AddSoftware 添加新软件
func AddSoftware(software Software) error {
	err := db.Create(&software).Error
	if err != nil {
		log.Printf("AddSoftware error: %v", err)
		return err
	}
	return nil
}

// GetSoftwareTotal 获取符合条件的软件总数
func GetSoftwareTotal(maps map[string]interface{}) (int64, error) {
	var count int64
	query := db.Model(&Software{})
	if len(maps) > 0 {
		query = query.Where(maps)
	}
	err := query.Count(&count).Error
	if err != nil {
		log.Printf("GetSoftwareTotal error: %v", err)
		return 0, err
	}
	return count, nil
}
