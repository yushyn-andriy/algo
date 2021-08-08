package strings

func ReverseWordsInString(str string) string {
	reversed := make([]byte, 0)

	i, j := len(str)-1, len(str)-1
	for ; i >= 0 && j >= 0; i-- {
		symbol := str[i]
		if isWhitespace(symbol) {
			reversed = append(reversed, symbol)
		} else {
			j = i
			word := []byte{}
			for ; j >= 0; j-- {
				currSumbol := str[j]
				if !isWhitespace(currSumbol) {
					word = append(word, currSumbol)
				} else {
					i = j + 1
					break
				}
			}
			reversed = append(reversed, reverseWord(word)...)
		}
	}

	return string(reversed)
}

func isWhitespace(b byte) bool {
	return b == 32
}

func reverseWord(word []byte) []byte {
	reversedWord := []byte{}
	for i := len(word) - 1; i >= 0; i-- {
		reversedWord = append(reversedWord, word[i])
	}
	return reversedWord
}
