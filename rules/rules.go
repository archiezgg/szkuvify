package rules

var (
	// Corrections stores what szkuvi appends if he corrects the sender
	Corrections = []string{
		"rozsul montot, hejesen:",
		"ozstopa tuk perekrin, asz hejesen:",
		"asz nem uty fna hanme: ",
		"falojaban asz: ",
		"hatyomanyos szkufiul asz: ",
	}

	// Compliments stores what szkuvi says if the message is well formed
	Compliments = []string{
		"azstakurfa esz iken prafo",
		"tabzs tabzs kecifei",
		"perfekt szkufinyelf",
		"szkufi apruvsz",
		"csicsi rektumphei",
	}

	// BaseRules stores the mapping for the single characters
	BaseRules = map[rune]rune{
		'v': 'f',
		'g': 'k',
		'b': 'p',
		'd': 't',
		'j': 'i',
	}

	//YRules stores the mapping for the hungarian "y" characters (e.g "ly")
	YRules = map[rune]rune{
		'g': 't',
		'l': 'j',
	}
)
