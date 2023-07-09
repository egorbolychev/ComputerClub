package main

import (
	"flag"
	"log"

	"github.com/egorbolychev/internal/app/controller"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/test1.txt", "path to config txt file")
}

func main() {
	flag.Parse()

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
