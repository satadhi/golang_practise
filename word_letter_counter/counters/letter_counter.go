package counters

import "unicode"

func CountLetters(letters string) map[rune]int {
	letterCount := make(map[rune]int)
	for _, ch := range letters {
		if unicode.IsLetter(ch) {
			ch = unicode.ToLower(ch)
			letterCount[ch]++
		}
	}
	return letterCount
}
