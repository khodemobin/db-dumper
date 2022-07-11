package storage

import (
	"errors"

	"github.com/khodemobin/db-dumper/config"
)

func Upload(filePath string, fileName string, storage *config.Storage) error {
	if storage.Driver == "ftp" {
		return ftpUpload(filePath, fileName, &FtpConfig{
			Host:       storage.Host,
			User:       storage.User,
			Port:       storage.Port,
			Password:   storage.Password,
			UploadPath: storage.Path,
		})
	}

	if storage.Driver == "local" {
		return localUpload(filePath, fileName, storage.Path)
	}

	return errors.New("invalid storage driver")
}
