package main

import (
	"log"
	"os"

	"github.com/gooneraki/blog-aggregator-go/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	appstate := state{
		cfg: &cfg,
	}

	appcommands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	appcommands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatalln("Not enough arguments were provided")
	}

	appcommand := command{
		name: args[1],
		args: args[2:],
	}

	err = appcommands.run(&appstate, appcommand)
	if err != nil {
		log.Fatalf("error running command with args '%v': %v", args, err)
	}
}
