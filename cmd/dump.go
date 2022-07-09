package cmd

import (
	"fmt"
	cli "github.com/jawher/mow.cli"
)

func Dump(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Printf("database dump")
	}
}
