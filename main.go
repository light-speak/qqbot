package main

import (
	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Logiase/MiraiGo-Template/config"
	"github.com/Logiase/MiraiGo-Template/utils"
	"os"
	"os/signal"
)

func init() {
	utils.WriteLogToFS()
	config.Init()
}

func main() {
	bot.Init()

	bot.StartService()

	bot.UseProtocol(bot.AndroidPhone)

	bot.Login()

	bot.RefreshList()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	bot.Stop()
}
