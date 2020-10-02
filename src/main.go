package main

import (
	"fmt"
	"go-web-hook/src/config"
	"go-web-hook/src/server"
)

func main() {
	fmt.Println("hello web hook")
	config.LoadConfig()
	server.StartService(config.Conf.Host, config.Conf.Port)
}
