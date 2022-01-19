package tries

func check1(bigString, smallString string, start int) bool {
	bigStringStartIdx := start
	bigStringEndIdx := bigStringStartIdx + len(smallString) - 1
	smallStringStartIdx := 0
	smallStringEndIdx := len(smallString) - 1

	if bigStringEndIdx > len(bigString)-1 {
		return false
	}

	for bigStringStartIdx < bigStringEndIdx {
		if bigString[bigStringStartIdx] != smallString[smallStringStartIdx] ||
			bigString[bigStringEndIdx] != smallString[smallStringEndIdx] {
			return false
		}
		bigStringStartIdx++
		bigStringEndIdx--
		smallStringStartIdx++
		smallStringEndIdx--
	}

	return true
}

func isInBigString1(bigString, smallString string) bool {
	for startIdx := 0; startIdx < len(bigString)-(len(smallString)-1); startIdx++ {
		if in := check1(bigString, smallString, startIdx); in {
			return true
		}
	}

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
