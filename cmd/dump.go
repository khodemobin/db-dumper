package cmd

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
)

func CmdDump(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("database dump")
	}
}
