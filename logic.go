package main

import (
	"math/rand"
	"time"
)

func szkuvify(text string) string {
	var szkuvifiedPhrase string

	for i, letter := range text {
		specialLetter, isLetterSpecial := yRules[letter]
		if isLetterSpecial && i != len(text)-1 && text[i+1] == 'y' {
			szkuvifiedPhrase += string(specialLetter)
			continue
		}

		if szkuviLetter, isLetterInRules := szkuviRules[letter]; isLetterInRules {
			szkuvifiedPhrase += string(szkuviLetter)
		} else {
			szkuvifiedPhrase += string(letter)
		}
	}
	return szkuvifiedPhrase
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
