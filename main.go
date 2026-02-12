package main

import (
	"go-web/config"
	"go-web/router"
)

func main() {
	config.InitConfig()
	r := router.Router()
	r.Run(config.Cfg.Server.Port)
}
