package cmd

import (
	"errors"
	cli "github.com/jawher/mow.cli"
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
	log.Println(task)
	log.Println(db)
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
