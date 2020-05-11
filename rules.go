package main

var (
	corrections = []string{
		"rozsul montot, hejesen:",
		"ozstopa tuk perekrin, asz hejesen:",
		"asz nem uty fna hanme: ",
		"falojaban asz: ",
		"hatyomanyos szkufiul asz: ",
	}
	compliments = []string{
		"azstakurfa esz iken prafo",
		"tabzs tabzs kecifei",
		"perfekt szkufinyelf",
		"szkufi apruvsz",
		"csicsi rektumphei",
	}

	szkuviRules = map[rune]rune{
		'v': 'f',
		'g': 'k',
		'b': 'p',
		'd': 't',
		'j': 'i',
	}

	yRules = map[rune]rune{
		'g': 't',
		'l': 'j',
	}
)
