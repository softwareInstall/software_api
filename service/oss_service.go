package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"mime/multipart"
	"path/filepath"
	"software_api/pkg/oss"
	"strings"
	"time"
)

type OssService struct {
	File *multipart.FileHeader
	Ctx  context.Context
}

func (o *OssService) UploadFile() (map[string]string, error) {
	fileInfo := map[string]string{
		"object_name":   "",
		"original_name": "",
	}
	// 生成唯一 objectName
	objectName := o.GenerateUniqueObjectName(o.File.Filename)

	src, err := o.File.Open()
	if err != nil {
		return fileInfo, err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			log.Printf("%v", err)
		}
	}(src)

	// 上传文件到 oss
	err = oss.ManagerOssObj.UploadFile(o.Ctx, objectName, src, o.File.Size)
	if err != nil {
		log.Printf("%v", err)
		return fileInfo, err
	}
	fileInfo["object_name"] = objectName
	fileInfo["original_name"] = o.File.Filename
	return fileInfo, err

}

func (o *OssService) GenerateUniqueObjectName(originalName string) string {
	ext := filepath.Ext(originalName)                 // 获取文件扩展名
	baseName := strings.TrimSuffix(originalName, ext) // 去掉扩展名

	dateStr := time.Now().Format("20060102150405") // 获取当前日期，格式为 "YYYYMMDDHHMMSS"
	randomNum := rand.Intn(10000)                  // 生成随机数
	// 格式示例: originalName_YYYYMMDDHHMMSS_RRRR.ext
	uniqueName := fmt.Sprintf("%s_%s_%04d%s", baseName, dateStr, randomNum, ext)
	return uniqueName
}
