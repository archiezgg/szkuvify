package logic

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/IstvanN/szkuvify/rules"
)

var (
	triggerChanceString = os.Getenv("TRIGGER_CHANCE")
	triggerChance       int
)

func init() {
	var err error
	triggerChance, err = strconv.Atoi(triggerChanceString)
	if err != nil {
		log.Fatalln(err, "set TRIGGER_CHANCE env var")
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

// ShouldSzkuviReply returns with true if trigger chance hits, returns false if not
func ShouldSzkuviReply(chanceString string) bool {
	dice := GenRandomNumber(100 / triggerChance)
	return dice != 0
}

// GenRandomNumber generates a random number between 0 and max (max not included)
func GenRandomNumber(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

// GetRandomElementFromSlice returns a random element of a slice
func GetRandomElementFromSlice(slice []string) string {
	numberOfIndeces := len(slice) - 1
	randomIndex := GenRandomNumber(numberOfIndeces)
	return slice[randomIndex]
}

func isLetterFollowedByY(index int, text string) bool {
	return index != len(text)-1 && text[index+1] == 'y'
}
