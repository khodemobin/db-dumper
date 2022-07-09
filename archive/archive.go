package archive

import (
	"errors"

	"github.com/khodemobin/db-dumper/config"
)

func Archive(path string, name string, task *config.Task) (filePath string, fileName string, err error) {
	if task.ArchiveDriver != "zip" {
		return "", "", errors.New("invalid archive driver")
	}

	return createZip(&ZipConfig{
		Path:     path,
		Name:     name,
		Password: task.ArchivePassword,
	})
}
