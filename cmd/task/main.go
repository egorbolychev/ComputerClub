package main

import (
	"flag"
	"log"

	"github.com/egorbolychev/internal/app/controller"
)

var (
	configPath = "configs/test_1.txt"
)

func main() {
	flag.Parse()
	if flag.Arg(0) != "" {
		configPath = flag.Arg(0)
	}

	log.SetFlags(0)

	confStr, taskStr, err := controller.Parse(configPath)
	if err != nil {
		log.Fatal(err)
	}

	config := controller.NewConfig()
	if err := config.ConfigureConfig(confStr); err != nil {
		log.Fatal(err)
	}

	if err := controller.Start(config, taskStr); err != nil {
		log.Fatal(err)
	}
}
