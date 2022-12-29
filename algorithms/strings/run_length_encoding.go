package strings

import (
	"strconv"
)

// s += "asdd"

func RunLengthEncoding(str string) string {
	encoded := []byte{}

	counter := 1
	for i := 1; i < len(str); i++ {
		currentCharacter := str[i]
		previousCharacter := str[i-1]

		if currentCharacter != previousCharacter || counter == 9 {
			encoded = append(encoded, strconv.Itoa(counter)[0])
			encoded = append(encoded, previousCharacter)
			counter = 0
		}

		counter += 1
	}

	encoded = append(encoded, strconv.Itoa(counter)[0])
	encoded = append(encoded, str[len(str)-1])
	return string(encoded)
}
