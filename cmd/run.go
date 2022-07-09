package cmd

import (
	"errors"
	cli "github.com/jawher/mow.cli"
	"github.com/khodemobin/db-dumper/archive"
	"github.com/khodemobin/db-dumper/backup"
	"github.com/khodemobin/db-dumper/config"
	"github.com/khodemobin/db-dumper/storage"
	"os"
)

func Run(cmd *cli.Cmd) {
	cmd.Action = func() {
		cfg := config.GetConfig()
		for _, task := range cfg.Tasks {
			db, err := findDB(cfg, task)
			if err != nil {
				panic(err)
			}

			if err := runTask(cfg, task, db); err != nil {
				panic(err)
			}
		}
	}
}

func runTask(config *config.Config, task config.Task, db *config.Database) error {
	filePath, fileName, err := backup.Backup(db)
	if err != nil {
		return err
	}

	archivePath, archiveFileName, err := archive.Archive(filePath, fileName, &task)
	if err != nil {
		return err
	}

	for _, s := range task.Storages {
		st, err := findStorage(config, s)

		if err != nil {
			return err
		}

		if err := storage.Upload(archivePath, archiveFileName, st); err != nil {
			return err
		}
	}

	if err := os.Remove(filePath); err != nil {
		return nil
	}

	if err := os.Remove(archivePath); err != nil {
		return nil
	}

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

func findStorage(config *config.Config, storage string) (*config.Storage, error) {
	for _, v := range config.Storages {
		if v.Name == storage {
			return &v, nil
		}
	}

	return nil, errors.New("invalid storage name")
}
