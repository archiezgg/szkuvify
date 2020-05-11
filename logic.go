package main

import (
	"math/rand"
	"time"
)

func szkuvify(text string) string {
	var szkuvifiedPhrase string

	for i, letter := range text {
		if specialLetter, isLetterYRule := yRules[letter]; isLetterYRule && isLetterFollowedByY(i, text) {
			szkuvifiedPhrase += string(specialLetter)
			continue
		}

		if szkuviLetter, isLetterInRules := baseRules[letter]; isLetterInRules {
			szkuvifiedPhrase += string(szkuviLetter)
		} else {
			szkuvifiedPhrase += string(letter)
		}
	}
	return szkuvifiedPhrase
}

func isLetterFollowedByY(index int, text string) bool {
	return index != len(text)-1 && text[index+1] == 'y'
}

func genRandomNumber(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}

func getElementRandomFromSlice(slice []string) string {
	numberOfIndeces := len(slice) - 1
	randomIndex := genRandomNumber(numberOfIndeces)
	return slice[randomIndex]
}
