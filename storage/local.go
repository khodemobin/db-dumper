package storage

import (
	"io/ioutil"
)

func localUpload(path, fileName, designationPath string) error {
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
