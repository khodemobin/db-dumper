package storage

import (
	"io/ioutil"
	"os"
)

func localUpload(path, fileName, designationPath string) error {
	if _, err := os.Stat(designationPath); os.IsNotExist(err) {
		if err := os.Mkdir(designationPath, os.ModePerm); err != nil {
			return err
		}
	}

	designationPath = designationPath + "/" + fileName
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(designationPath, input, 0600); err != nil {
		return err
	}

	return nil
}
