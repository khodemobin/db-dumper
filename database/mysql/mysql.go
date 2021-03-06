package mysql

import (
	"database/sql"
	"fmt"
	"os"
)

type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	TempPath string
}

func Dump(cfg *MysqlConfig) (filePath string, filename string, err error) {
	dumpFilenameFormat := fmt.Sprintf("%s-2006-01-02 15:04:05", cfg.Database)

	if _, err := os.Stat(cfg.TempPath); os.IsNotExist(err) {
		if err := os.Mkdir(cfg.TempPath, os.ModePerm); err != nil {
			return "", "", err
		}
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.User))
	if err != nil {
		return "", "", err
	}

	dmp, err := register(db, cfg.TempPath, dumpFilenameFormat)
	if err != nil {
		return "", "", err
	}

	fullPath, filename, err := dmp.dumper()
	if err != nil {
		return "", "", err
	}

	if err := dmp.Close(); err != nil {
		return "", "", err
	}

	filename += ".sql"

	return fullPath, filename, nil
}
