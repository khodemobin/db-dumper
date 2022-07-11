package archive

import (
	"errors"
	"github.com/khodemobin/db-dumper/config"
)

func Archive(path string, name string, task *config.Task) (filePath string, fileName string, err error) {
	if task.Archive == (config.Archive{}) {
		return path, name, nil
	}

	if task.Archive.Driver == "zip" {
		return createZip(&ZipConfig{
			Path:     path,
			Name:     name,
			Password: task.Archive.Password,
		})
	}

	if task.Archive.Driver == "tar" {
		return createTar(&TarConfig{
			Path:     path,
			Name:     name,
			Password: task.Archive.Password,
		})
	}

	return "", "", errors.New("invalid archive driver")

}
