package compress

import (
	"errors"

	"github.com/khodemobin/db-dumper/config"
)

func Compress(path string, name string, task *config.Task) (filePath string, fileName string, err error) {
	if task.Compress.Driver != "zip" {
		return "", "", errors.New("invalid compress driver")
	}

	return createZip(&ZipConfig{
		Path:     path,
		Name:     name,
		Password: task.Compress.Password,
	})
}
