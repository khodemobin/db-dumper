package cmd

import (
	"errors"
	"fmt"
	"github.com/briandowns/spinner"
	cli "github.com/jawher/mow.cli"
	"github.com/khodemobin/db-dumper/compress"
	"github.com/khodemobin/db-dumper/config"
	"github.com/khodemobin/db-dumper/database"
	"github.com/khodemobin/db-dumper/storage"
	"os"
	"time"
)

var sp *spinner.Spinner

func Run(cmd *cli.Cmd) {
	sp = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
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
	/* Start taking dump  */
	fmt.Println("taking dump")
	sp.Start()
	filePath, fileName, err := database.Dump(db)
	if err != nil {
		return err
	}
	sp.Stop()

	/* Start running compress  */
	fmt.Println("creating compress")
	sp.Start()
	archivePath, archiveFileName, err := compress.Compress(filePath, fileName, &task)
	if err != nil {
		return err
	}
	sp.Stop()

	/* Start putting to storages  */
	fmt.Println("putting to storages")
	sp.Start()
	for _, s := range task.Storages {
		st, err := findStorage(config, s)

		if err != nil {
			return err
		}

		if err := storage.Upload(archivePath, archiveFileName, st); err != nil {
			return err
		}
	}
	sp.Stop()

	fmt.Println("cleanup temp files")
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
