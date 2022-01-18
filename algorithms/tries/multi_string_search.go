package tries

func isInBigString1(bigString, smallString string) bool {
	return false
}

func MultiStringSearch1(bigString string, smallStrings []string) []bool {
	output := make([]bool, len(smallStrings))
	for i, smallString := range smallStrings {
		output[i] = isInBigString1(bigString, smallString)
	}

	return output
}

func MultiStringSearch2(bigString string, smallStrings []string) []bool {
	// Write your code here.
	return nil
}

func MultiStringSearch3(bigString string, smallStrings []string) []bool {
	// Write your code here.
	return nil
}
