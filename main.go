package main

import (
	"github.com/khodemobin/db-dumper/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(cfg.Tasks)
}
