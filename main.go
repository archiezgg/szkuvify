package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IstvanN/szkuvify/logic"
	"github.com/IstvanN/szkuvify/rules"
	"github.com/bwmarrin/discordgo"
)

var (
	token               = os.Getenv("SZKUVI_TOKEN")
	chanceToBeTriggered = 100
)

func main() {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("error creating Discord session:", err)
		return
	}
	defer discord.Close()

	discord.AddHandler(szkuviHandler)

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

func szkuviHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}

	// szkuvi replies with a 10% chance
	dice := logic.GenRandomNumber(100 / chanceToBeTriggered)
	if dice != 0 {
		return
	}

	// szkuvi compliments
	if message.Content == logic.Szkuvify(message.Content) {
		discord.ChannelMessageSend(message.ChannelID, logic.GetRandomElementFromSlice(rules.Compliments))
		return
	}

	// szkuvi corrects
	m := logic.GetRandomElementFromSlice(rules.Corrections) + " " + logic.Szkuvify(message.Content)
	discord.ChannelMessageSend(message.ChannelID, m)
}
