package compress

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"

	"github.com/yeka/zip"
)

type ZipConfig struct {
	Path     string
	Name     string
	Password string
}

func createZip(cfg *ZipConfig) (filePath string, fileName string, err error) {
	zFileName := cfg.Name + ".zip"
	zFilePath := cfg.Path + ".zip"

	content, err := ioutil.ReadFile(cfg.Path)
	if err != nil {
		return zFilePath, zFileName, err
	}

	fzip, err := os.Create(zFilePath)
	if err != nil {
		return zFilePath, zFileName, err
	}

	zipw := zip.NewWriter(fzip)
	w, err := zipw.Encrypt(cfg.Name, cfg.Password, zip.AES256Encryption)
	if err != nil {
		return zFilePath, zFileName, err
	}

	if _, err := io.Copy(w, bytes.NewReader(content)); err != nil {
		return zFilePath, zFileName, err
	}

	if err := zipw.Flush(); err != nil {
		return "", "", err
	}

	if err := zipw.Close(); err != nil {
		return "", "", err
	}

	return zFilePath, zFileName, nil
}
