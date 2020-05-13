package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IstvanN/szkuvify/handler"
	"github.com/bwmarrin/discordgo"
)

var token = os.Getenv("SZKUVI_TOKEN")

func main() {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("error creating Discord session:", err)
		return
	}
	defer discord.Close()

	discord.AddHandler(handler.SzkuviHandler)

	err = discord.Open()
	if err != nil {
		log.Fatalln("error opening Discord session:", err)
		return
	}

	log.Println("szkuvify is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
