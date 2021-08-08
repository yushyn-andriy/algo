package strings

import (
	"strconv"
	"strings"
)

func ValidIPAddresses(str string) []string {
	ipAdressesFound := make([]string, 0)

	for i := 1; i < min(len(str), 4); i++ {
		currIPAddressParts := []string{"", "", "", ""}

		currIPAddressParts[0] = str[:i]
		if !isValidPart(currIPAddressParts[0]) {
			continue
		}

		for j := i + 1; j < i+min(len(str)-i, 4); j++ {
			currIPAddressParts[1] = str[i:j]
			if !isValidPart(currIPAddressParts[1]) {
				continue
			}

			for k := j + 1; k < j+min(len(str)-j, 4); k++ {
				currIPAddressParts[2] = str[j:k]
				currIPAddressParts[3] = str[k:]

				if isValidPart(currIPAddressParts[2]) && isValidPart(currIPAddressParts[3]) {
					ipAdressesFound = append(ipAdressesFound, strings.Join(currIPAddressParts, "."))
				}
			}
		}
	}

	return ipAdressesFound
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValidPart(str string) bool {
	i, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	if i > 255 {
		return false
	}

	return len(str) == len(strconv.Itoa(i))
}
