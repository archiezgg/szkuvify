package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var token = os.Getenv("SZKUVI_TOKEN")

func main() {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("error creating Discord session: ", err)
		return
	}
	defer discord.Close()

	discord.AddHandler(szkuviHandler)

	err = discord.Open()
	if err != nil {
		log.Println("error opening Discord session: ", err)
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
	dice := genRandomNumber(10)
	if dice != 5 {
		return
	}

	if message.Content == szkuvify(message.Content) {
		discord.ChannelMessageSend(message.ChannelID, compliment)
		return
	}
	m := baseCorrection + szkuvify(message.Content)
	discord.ChannelMessageSend(message.ChannelID, m)
}
