package cmd

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
)

func CmdRun(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("db-dumper run")
	}
}
