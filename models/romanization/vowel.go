package romanization

var vowelForK = map[rune]string{
	6070: "a",   //ា
	6071: "e",   //ិ
	6072: "ei",  //ី
	6073: "oe",  //ឹ
	6074: "eu",  //ឺ
	6075: "o",   //ុ
	6076: "ou",  //ូ
	6077: "uo",  //ួ
	6078: "aeu", //ើ
	6079: "oea", //ឿ
	6080: "ie",  //ៀ
	6081: "e",   //េ
	6082: "ae",  //ែ
	6083: "ai",  //ៃ
	6084: "ao",  //ោ
	6085: "au",  //ៅ
	6086: "am",  //ំ
	6087: "ah",  //ះ
	6088: "ak",  //◌ៈ
	6089: "",    //=៉
	6091: "",    //់
	6092: "r",   //៌
	6095: "a",   //៏
	6096: "oa",  // ័
	6098: "",    //=្
}

var vowelForKor = map[rune]string{
	6070: "ea",  //ា
	6071: "i",   //ិ
	6072: "i",   //ី
	6073: "ue",  //ឹ
	6074: "ueu", //ឺ
	6075: "u",   //ុ
	6076: "u",   //ូ
	6077: "uo",  //ួ
	6078: "au",  //ើ
	6079: "oea", //ឿ
	6080: "ie",  //ៀ
	6081: "e",   //េ
	6082: "eae", //ែ
	6083: "ey",  //ៃ
	6084: "ou",  //ោ
	6085: "ov",  //ៅ
	6086: "um",  //ំ
	6087: "eah", //ះ
	6088: "eak", //◌ៈ
	6089: "",    //=៉
	6091: "",    //់
	6092: "r",   //៌
	6095: "a",   //៏
	6096: "oa",  // ័
	6098: "",    //=្
}

func IsVowel(ch rune) bool {
	_, isVowel := vowelForK[ch]

	return isVowel
}

func IsSuffixableVowel(ch rune) bool {
	suffixableVowelList := map[rune]string{
		6070: "ea", //ា => ាំ
		6071: "i",  //ិ => ិះ
		6075: "u",  //ុ => ុះ
		6081: "e",  //េ => េះ
		6084: "ou", //ោ => ោះ
		6086: "am", //ំ => ំុ
		6096: "oa", // ័
	}

	_, isSpecial := suffixableVowelList[ch]
	return isSpecial
}

func SpecialVowel(first rune, second rune, isKor bool) string {
	var specialVowelList map[rune]string
	if isKor {
		specialVowelList = map[rune]string{
			12161: "um",   //ំ => ំុ, 6086 + 6075 = 12161
			12156: "oam",  //ា => ាំ, 6070 + 6086 = 12156
			18176: "eang", //ាំង, 6070 + 6086 + 6020= 18176
			12168: "eh",   //េ => េះ, 6081 + 6087 = 12168
			12162: "uh",   //ុ => ុះ, 6075 + 6087 = 12162
			12171: "uoh",  //ោ => ោះ, 6084 + 6087 = 12171
			12158: "is",   //ិ => ិះ, 6071 + 6087 = 12158
			12116: "eang", //័ង , 6096 + 6020 = 12116
			12112: "eak",  //័ក , 6096 + 6016 = 12112
			12137: "ey",   //័យ , 6096 + 6041 = 12137
		}
	} else {
		specialVowelList = map[rune]string{
			12161: "om",  //ំ => ំុ, 6086 + 6075 = 12161
			12156: "am",  //ា => ាំ, 6070 + 6086 = 12156
			18176: "ang", //ាំង, 6070 + 6086 + 6020= 18176
			12168: "eh",  //េ => េះ, 6081 + 6087 = 12168
			12162: "oh",  //ុ => ុះ, 6075 + 6087 = 12162
			12171: "aoh", //ោ => ោះ, 6084 + 6087 = 12171
			12158: "eh",  //ិ => ិះ, 6071 + 6087 = 12158
			12116: "ang", //័ង , 6096 + 6020 = 12116
			12112: "ak",  //័ក , 6096 + 6016 = 12112
			12137: "ai",  //័យ , 6096 + 6041 = 12137
		}
	}

	if r_ch, isMatch := specialVowelList[first+second]; isMatch {
		return r_ch
	}

	if IsConsonant(second) {
		return VowelWord(first, isKor) + Roman(second)
	}

	return VowelWord(first, isKor) + VowelWord(second, false)
}

func VowelWord(ch rune, isKor bool) string {
	var wordsList map[rune]string
	if isKor {
		wordsList = vowelForKor
	} else {
		wordsList = vowelForK
	}
	if r_ch, isMatch := wordsList[ch]; isMatch {
		return r_ch
	}

	return ""
}
