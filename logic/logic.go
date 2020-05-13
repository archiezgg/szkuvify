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

// Szkuvify is the main logic function for forming the messages
func Szkuvify(text string) string {
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

// ReplyToSummon replies with a chance to a special trigger word with a chance
func ReplyToSummon(discord *discordgo.Session, channelID string, originalMsg string) {
	if strings.Contains(originalMsg, "szkufi") && szkuviGetsTriggered(summonChance) {
		reply := getRandomElementFromSlice(rules.SummonReplies)
		discord.ChannelMessageSend(channelID, reply)
	}
}

// Compliment sends a reply randomly from compliments
func Compliment(discord *discordgo.Session, channelID string) {
	if szkuviGetsTriggered(triggerChance) {
		reply := getRandomElementFromSlice(rules.Compliments)
		discord.ChannelMessageSend(channelID, reply)
	}
}

// Correkt sends a reply randomly from corrections
func Correkt(discord *discordgo.Session, channelID string, originalMsg string) {
	if szkuviGetsTriggered(triggerChance) {
		reply := getRandomElementFromSlice(rules.Corrections) + " " + Szkuvify(originalMsg)
		discord.ChannelMessageSend(channelID, reply)
	}
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
