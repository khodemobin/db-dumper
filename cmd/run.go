package cmd

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	cli "github.com/jawher/mow.cli"
	"github.com/khodemobin/db-dumper/archive"
	"github.com/khodemobin/db-dumper/backup"
	"github.com/khodemobin/db-dumper/config"
	"log"
)

func Run(cmd *cli.Cmd) {
	cmd.Action = func() {
		cfg := config.GetConfig()
		for _, task := range cfg.Tasks {
			db, err := findDB(cfg, task)
			if err != nil {
				panic(err)
			}

			if err := runTask(task, db); err != nil {
				panic(err)
			}
		}
	}
}

func runTask(task config.Task, db *config.Database) error {
	filePath, fileName, err := backup.Backup(db)
	if err != nil {
		return err
	}

	archivePath, archiveFileName, err := archive.Archive(filePath, fileName, &task)
	if err != nil {
		return err
	}

	log.Println(archivePath)
	log.Println(archiveFileName)

	return nil
}

func findDB(config *config.Config, task config.Task) (*config.Database, error) {
	for _, v := range config.Databases {
		if v.Name == task.Database {
			return &v, nil
		}
	}

	return nil, errors.New("invalid database name")
}
