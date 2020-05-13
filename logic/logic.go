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
	triggerChance       int
	summonChance        int
)

func init() {
	var err error
	triggerChance, err = strconv.Atoi(triggerChanceString)
	if err != nil {
		log.Fatalln(err, "set TRIGGER_CHANCE env var")
	}

	summonChance, err = strconv.Atoi(summonChanceString)
	if err != nil {
		log.Fatalln(err, "set SUMMON_CHANCE env var")
	}
}

// Reply decides what and when does szkuvi replies
func Reply(discord *discordgo.Session, message *discordgo.MessageCreate) {

	// szkuvi gets summoned

	if messageContainsSummonTrigger(message.Content) && szkuviGetsTriggered(summonChance) {
		reply := getRandomElementFromSlice(rules.SummonReplies)
		discord.ChannelMessageSend(message.ChannelID, reply)
		return
	}

	if !szkuviGetsTriggered(triggerChance) {
		return
	}

	// szkuvi compliments
	if message.Content == szkuvify(message.Content) {
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
	numberOfIndeces := len(slice) - 1
	randomIndex := genRandomNumber(numberOfIndeces)
	return slice[randomIndex]
}

func isLetterFollowedByY(index int, text string) bool {
	return index != len(text)-1 && text[index+1] == 'y'
}

func messageContainsSummonTrigger(message string) bool {
	for _, trigger := range rules.SummonTriggers {
		if strings.Contains(message, trigger) {
			return true
		}
	}
	return false
}
