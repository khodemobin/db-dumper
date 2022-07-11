package archive

import (
	"archive/tar"
	"io/ioutil"
	"os"
)

type TarConfig struct {
	Path     string
	Name     string
	Password string
}

func createTar(cfg *TarConfig) (filePath string, fileName string, err error) {
	tFileName := cfg.Name + ".tar"
	tFilePath := cfg.Path + ".tar"

	content, err := ioutil.ReadFile(cfg.Path)
	if err != nil {
		return tFilePath, tFileName, err
	}

	file, err := os.Create(tFilePath)
	if err != nil {
		return tFilePath, tFileName, err
	}

	tw := tar.NewWriter(file)

	hdr := &tar.Header{
		Name: cfg.Name,
		Mode: 0600,
		Size: int64(len(content)),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		return tFilePath, tFileName, err
	}

	if _, err := tw.Write(content); err != nil {
		return tFilePath, tFileName, err
	}

	if err := tw.Close(); err != nil {
		return tFilePath, tFileName, err
	}

	return tFilePath, tFileName, err
}
