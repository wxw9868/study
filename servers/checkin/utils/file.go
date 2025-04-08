package utils

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

// SaveFile 保存上传的文件
func SaveFile(file *multipart.FileHeader, saveDir string) (string, error) {
	// 生成唯一文件名
	ext := filepath.Ext(file.Filename)
	fileName := time.Now().Format("20060102150405") + ext
	filePath := filepath.Join(saveDir, fileName)

	// 创建保存目录
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return "", err
	}

	// 保存文件
	if err := os.WriteFile(filePath, []byte{}, 0644); err != nil {
		return "", err
	}
	return filePath, nil
}
