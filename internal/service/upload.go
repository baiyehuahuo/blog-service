package service

import (
	"blog-service/global"
	"blog-service/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
	"path"
)

type FileInfo struct {
	Name      string
	AccessURL string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(fileHeader.Filename)
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file type not support")
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("file size too big")
	}
	uploadSavePath := upload.GetSavePath()
	if upload.CheckSavePath(uploadSavePath) {
		if err := upload.CreateSavePath(uploadSavePath, os.ModePerm); err != nil {
			return nil, errors.New("create upload save path failed: " + err.Error())
		}
	}
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permission")
	}
	dst := path.Join(uploadSavePath, fileName)
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, errors.New("upload save file failed: " + err.Error())
	}
	accessURL := path.Join(global.AppSetting.UploadServerURL, fileName)
	return &FileInfo{
		Name:      fileName,
		AccessURL: accessURL,
	}, nil
}
