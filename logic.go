package main

import (
	"math/rand"
	"time"
)

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

func genRandomNumber(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}
