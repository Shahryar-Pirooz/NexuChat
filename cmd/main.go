package main

import (
	"flag"
	"nexu-chat/config"
	"os"
)

func main() {}

func loadConfig() config.Config {
	var configPath string
	const (
		defaultPath = "$HOME/nexu.config.yml"
		message     = "this flag passes config to app"
	)

	flag.StringVar(&configPath, "c", defaultPath, message)
	flag.StringVar(&configPath, "config", defaultPath, message)
	flag.Parse()

	if envConfig := os.Getenv("NEXU_CHAT_CONFIG_PATH"); len(envConfig) > 0 {
		configPath = envConfig
	}

	if len(configPath) < 0 {
		panic("configuration file not found")
	}
	return config.MustReadConfig(configPath)

}
