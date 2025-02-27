package main

import (
	"screenshot-ai/config"
	"screenshot-ai/keyinput"
	"screenshot-ai/network"
	"screenshot-ai/pkg/log"
	"screenshot-ai/pkg/utils"
)

func main() {
	startKeyboardListener()
}

func startKeyboardListener() {
	keys := config.Cfg.KeyBoard
	keyborad := keyinput.NewBind(keys.Keys, keys.DoneKeys)
	go keyborad.Listen()
	keyborad.StartEvents()
}

func init() {
	config.InitConfig()
	log.InitLogger()
	network.NewHub()
	go network.HubServer.Start()
	engine := network.NewRouter()
	port := config.Cfg.Server.Port
	go engine.Run(port)
	go utils.CleanPeriodly()
}
