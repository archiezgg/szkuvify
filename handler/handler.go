package handler

import (
	"strings"

	"github.com/IstvanN/szkuvify/logic"
	"github.com/IstvanN/szkuvify/rules"
	"github.com/bwmarrin/discordgo"
)

// SzkuviHandler is the main handler to reply
func SzkuviHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}

	// szkuvi replies to a summoning with 100% chance
	if strings.Contains(message.Content, "szkufi") {
		discord.ChannelMessage(message.ChannelID, logic.GetRandomElementFromSlice(rules.Summonings))
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
