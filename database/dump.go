package database

import (
	"errors"
	"github.com/khodemobin/db-dumper/config"
	"github.com/khodemobin/db-dumper/database/mysql"
)

func Dump(db *config.Database) (filePath string, fileName string, err error) {
	if db.Driver != "mysql" {
		return "", "", errors.New("invalid database driver")
	}

	return mysql.Dump(&mysql.MysqlConfig{
		Host:     db.Host,
		Port:     db.Port,
		User:     db.User,
		Password: db.Password,
		Database: db.Database,
		TempPath: config.GetConfig().TempPath,
	})
}
