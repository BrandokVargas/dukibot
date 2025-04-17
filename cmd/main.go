package main

import (
	"github.com/BrandokVargas/dukibot/config"
	"github.com/BrandokVargas/dukibot/internal/bot"
	"github.com/BrandokVargas/dukibot/internal/constants"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.Load()

	discordBot, err := bot.New(cfg.Token)
	if err != nil {
		log.Fatalf("%s startup error: %v ", constants.NameBot, err)
	}

	router := bot.NewRouter(cfg.Prefix)
	router.Register(&bot.PingCommand{})
	router.Register(&bot.ResponseChannel{})
	router.Register(&bot.ResponseChannelCategory{})
	router.Register(&bot.ResponseChannelMultiplyCategory{})
	router.Register(&bot.ResponseChannelTextMultiply{})

	if err := discordBot.Start(router); err != nil {
		log.Fatalf("Error at startup: %v", err)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Off the", constants.NameBot)
	discordBot.Stop()
}
