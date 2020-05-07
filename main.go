package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token          = os.Getenv("SZKUVI_TOKEN")
	baseCorrection = "rozsul montot te ozstopa kecifei. hejesen: "
	clapclap       = "azstakurfa esz iken prafo tabzs tabzs kecifei, perfekt szkufinyelf"
)

var szkuviRules = map[rune]rune{
	'v': 'f',
	'g': 'k',
	'b': 'p',
	'd': 't',
	'j': 'i',
}

var yRules = map[rune]rune{
	'g': 't',
	'l': 'j',
}

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
	if message.Content == szkuvify(message.Content) {
		discord.ChannelMessageSend(message.ChannelID, clapclap)
		return
	}
	m := baseCorrection + szkuvify(message.Content)
	discord.ChannelMessageSend(message.ChannelID, m)
}

func szkuvify(text string) string {
	var szkuvifiedPhrase string

	for i, letter := range text {
		specialLetter, isSpecial := yRules[letter]
		if isSpecial && i != len(text)-1 && text[i+1] == 'y' {
			szkuvifiedPhrase += string(specialLetter)
			continue
		}

		szkuviLetter, isRule := szkuviRules[letter]
		if isRule {
			szkuvifiedPhrase += string(szkuviLetter)
		} else {
			szkuvifiedPhrase += string(letter)
		}
	}
	return szkuvifiedPhrase
}
