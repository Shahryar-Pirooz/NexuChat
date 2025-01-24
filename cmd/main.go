package main

import (
	"flag"
	"fmt"
	"nexu-chat/config"
	natsserver "nexu-chat/pkg/nats-server"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {
	fmt.Println("run server")
	nc := natsserver.NewNatsService(nats.DefaultURL).Connect()
	defer nc.Close()
	nc.Subscribe("channel1", func(msg *nats.Msg) {
		h := msg.Header
		user := h.Get("user")
		fmt.Printf("%s\t->\t%s\n", user, msg.Data)
	})
	select {}
}

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
