package compress

import (
	"errors"
	"github.com/khodemobin/db-dumper/config"
)

func Compress(path string, name string, task *config.Task) (filePath string, fileName string, err error) {
	if task.Compress == (config.Compress{}) {
		return path, name, nil
	}

	if task.Compress.Driver == "zip" {
		return createZip(&ZipConfig{
			Path:     path,
			Name:     name,
			Password: task.Compress.Password,
		})
	}

	if task.Compress.Driver == "tar" {
		return createTar(&TarConfig{
			Path:     path,
			Name:     name,
			Password: task.Compress.Password,
		})
	}

	return "", "", errors.New("invalid compress driver")

}
