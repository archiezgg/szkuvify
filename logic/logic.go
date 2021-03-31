package logic

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/IstvanN/szkuvify/rules"
	"github.com/bwmarrin/discordgo"
)

var (
	triggerChanceString = os.Getenv("TRIGGER_CHANCE")
	summonChanceString  = os.Getenv("SUMMON_CHANCE")
)

func getTriggerChance() int {
	triggerChance, err := strconv.Atoi(triggerChanceString)
	if err != nil {
		log.Fatalln(err, "set TRIGGER_CHANCE env var")
	}
	return triggerChance
}

func getSummonChance() int {
	summonChance, err := strconv.Atoi(summonChanceString)
	if err != nil {
		log.Fatalln(err, "set SUMMON_CHANCE env var")
	}
	return summonChance
}

// Reply decides what and when does szkuvi replies
func Reply(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// someone posts an image/gif
	if message.Content == "" && szkuviGetsTriggered(getTriggerChance()) {
		reply := getRandomElementFromSlice(rules.ImageReplies)
		discord.ChannelMessageSend(message.ChannelID, reply)
		return
	}

	// someone thanks for something
	if messageContainsTrigger(message.Content, rules.ThankTriggers) && szkuviGetsTriggered(getSummonChance()) {
		reply := getRandomElementFromSlice(rules.ThankReplies)
		discord.ChannelMessageSend(message.ChannelID, reply)
		return
	}
	// szkuvi gets summoned
	if messageContainsTrigger(message.Content, rules.SummonTriggers) && szkuviGetsTriggered(getSummonChance()) {
		reply := getRandomElementFromSlice(rules.SummonReplies)
		discord.ChannelMessageSend(message.ChannelID, reply)
		return
	}

	if !szkuviGetsTriggered(getTriggerChance()) {
		return
	}

	// szkuvi compliments
	if strings.ToLower(message.Content) == szkuvify(message.Content) {
		reply := getRandomElementFromSlice(rules.Compliments)
		discord.ChannelMessageSend(message.ChannelID, reply)
		return
	}

	//szkuvi correkts
	reply := getRandomElementFromSlice(rules.Corrections) + " " + szkuvify(message.Content)
	discord.ChannelMessageSend(message.ChannelID, reply)

}

// szkuvify is the main logic function for forming the messages
func szkuvify(text string) string {
	text = strings.ToLower(text)
	var szkuvifiedPhrase string
	for i, letter := range text {
		if specialLetter, isLetterYRule := rules.YRules[letter]; isLetterYRule && isLetterFollowedByY(i, text) {
			szkuvifiedPhrase += string(specialLetter)
			continue
		}

		if szkuviLetter, isLetterInRules := rules.BaseRules[letter]; isLetterInRules {
			szkuvifiedPhrase += string(szkuviLetter)
		} else {
			szkuvifiedPhrase += string(letter)
		}
	}
	return szkuvifiedPhrase
}

func szkuviGetsTriggered(chance int) bool {
	dice := genRandomNumber(100 / chance)
	return dice == 0
}

func genRandomNumber(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

func getRandomElementFromSlice(slice []string) string {
	randomIndex := genRandomNumber(len(slice))
	return slice[randomIndex]
}

func isLetterFollowedByY(index int, text string) bool {
	return index != len(text)-1 && text[index+1] == 'y'
}

func messageContainsTrigger(message string, triggers []string) bool {
	for _, trigger := range triggers {
		if strings.Contains(strings.ToLower(message), trigger) {
			return true
		}
	}
	return false
}
