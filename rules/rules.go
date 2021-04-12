package rules

var (
	// Corrections stores what szkuvi appends if he corrects the sender
	Corrections = []string{
		"rozsul montot, hejesen:",
		"ozstopa tuk perekrin, asz hejesen:",
		"asz nem uty fna hanme:",
		"falojaban asz:",
		"hatyomanyos szkufiul asz:",
		"enket mek hoty kijafitsalka:",
		"ekészen pontosna:",
		"cska a pontosák ketféért:",
		"lyopp ha tutot:",
	}

	// Compliments stores what szkuvi says if the message is well formed
	Compliments = []string{
		"azstakurfa esz iken prafo",
		"tabzs tabzs kecifei",
		"perfekt szkufinyelf",
		"szkufi aprufsz",
		"csicsi rektumphei",
		"kratulálko",
		"asz työnyörüh",
		"prafo kecifei",
		"fékre falaki hejesen peszél",
		"szíl of aprufal krantit",
		"kut kém fel pléjt",
		"esz a peszté",
	}

	// SummonReplies stores what szkuvi says if he get's summoned
	SummonReplies = []string{
		"hattyál most nme",
		"montyat keci",
		"állantóan én kellekh hát peszarko",
		"mi fan mrá",
		"hattyatko pékén",
		"lezsarmo",
		"pisztos hoty nme",
		"keci",
		"pill, prp neför",
		"páj páj luszerek",
		"passzátko mek msto nme",
		"kus majt késöp",
	}

	// ThankReplies contains replies that szkuvi says if some thanks something
	ThankReplies = []string{
		"ninscmti kecife",
		"esz a lekefesepp",
		"szifesen szuka plijat",
		"ikaszán szifesen",
		"ninscmti örüllyé",
		"lele szifesen maskro se",
	}

	// SummonTriggers stores those words that summons szkuvi to discord
	SummonTriggers = []string{
		"szkuf",
		"szkuv",
		"szkúf",
		"szkúv",
		"keci",
	}

	// ThankTriggers contains the trigger words that triggers thank replies
	ThankTriggers = []string{
		"thenk",
		"thank",
		"tenk",
		"kösz",
		"közs",
		"köc",
		"thx",
	}

	// ImgeReplies contains the answers szkuvi replies for image posts
	ImageReplies = []string{
		"asz iken keci",
		"ftf",
		"tank shit",
		"lol asz otapasz",
		"geg",
		"ta fuk",
		"napaszek",
		"takarotty pasztmek",
		"opaszt",
		"ity ity",
		"regt",
	}

	// BaseRules stores the mapping for the single characters
	BaseRules = map[rune]rune{
		'v': 'f',
		'g': 'k',
		'b': 'p',
		'd': 't',
		'j': 'i',
		'w': 'f',
	}

	//YRules stores the mapping for the hungarian "y" characters (e.g "ly")
	YRules = map[rune]rune{
		'g': 't',
		'l': 'j',
	}
)
