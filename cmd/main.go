package main

import (
	"flag"
	"fmt"
	"nexu-chat/config"
	"os"
)

func main() {
	fmt.Println(loadConfig())
}

func loadConfig() config.Config {
	var path string
	var (
		defaultPath = "./sample-config.yml"
		message     = "pass the path to the config file"
	)
	flag.StringVar(&path, "config", defaultPath, message)
	flag.StringVar(&path, "c", defaultPath, message)
	flag.Parse()

	if envPath := os.Getenv("CONFIG_NEXU_CHAT"); len(envPath) > 0 {
		path = envPath
	}

	if len(path) < 1 {
		panic("config file path is required")
	}

	cnfg := config.MustRead(path)
	return cnfg
}
