package romanization

func Romanize(word string) string {
	characters := []rune(word)
	result := ""
	rw := ""
	isThreeCh := false
	var vCh rune
	var lastCh rune

	for index, ch := range characters {
		// is Suffixable vowel? //ំ, ា, /េ, /ុ ,ោ, /ិ
		if IsSuffixableVowel(ch) && vCh == 0 {
			vCh = ch
			continue
		} else if vCh != 0 && (vCh+ch) == 12156 && isThreeCh == false {
			// if ាំ  special case
			isThreeCh = true
			continue
		}

		rw = Roman(ch)
		if isThreeCh {
			// if ch = ង
			if ch == 6020 {
				rw = SpecialVowel(12156, ch, IsKorGroup(lastCh))
			} else {
				rw = SpecialVowel(6070, 6086, IsKorGroup(lastCh)) + rw
			}

			isThreeCh = false
			vCh = 0
		} else if vCh != 0 {
			rw = SpecialVowel(vCh, ch, IsKorGroup(lastCh))
			vCh = 0
		} else if IsVowel(ch) {
			rw = VowelWord(ch, IsKorGroup(lastCh))
		} else if IsConsonant(lastCh) && IsConsonant(ch) {
			// ជ + ន = ជ + អr + ន
			// ជ + ណ = ជ + អ + ណ
			if index == 1 && lastCh == 6050 { //6050 = អ
				// if អ at the begin not cound
				rw = rw
			} else if IsKorGroup(lastCh) {
				rw = "o" + rw
			} else {
				rw = "a" + rw
			}
		}

		lastCh = ch
		result = result + rw
	}

	return result
}
