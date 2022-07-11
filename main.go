package main

import (
	_ "github.com/go-sql-driver/mysql"

	cli "github.com/jawher/mow.cli"
	"github.com/khodemobin/db-dumper/cmd"
	"github.com/khodemobin/db-dumper/config"
	"log"
	"os"
)

func main() {
	_, err := config.LoadConfig("config.yaml", ".env")
	if err != nil {
		log.Fatalln(err)
	}

	app := cli.App("db-dumper", "Database backup/sync tools")

	app.Command("run", "run db-dumper", cmd.Run)
	app.Command("dump", "dump database", cmd.Dump)
	//app.Command("config", "manage accounts", func(config *cli.Cmd) {
	//	config.Command("list", "list accounts", cmdList)
	//	config.Command("add", "add an account", cmdAdd)
	//	config.Command("remove", "remove an account(s)", cmdRemove)
	//})
	err = app.Run(os.Args)
	if err != nil {
		return
	}
}
