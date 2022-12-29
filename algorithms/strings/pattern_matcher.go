package strings

func PatternMatcher(pattern, str string) []string {
	if len(pattern) > len(str) {
		return nil
	}

	newPattern := getNewPattern(pattern)
	didSwitch := newPattern[0] != string(pattern[0])
	counts := map[string]int{"x": 0, "y": 0}
	firstYPos := getCountAndFirstYPos(newPattern, counts)

	if counts["y"] != 0 {
		for lenOfX := 1; lenOfX < len(str); lenOfX++ {
			lenOfY := (len(str) - lenOfX*counts["x"]) / counts["y"]
			if lenOfY <= 0 || lenOfY%1 != 0 {
				continue
			}
			yIdx := firstYPos * lenOfX
			x := str[:lenOfX]
			y := str[yIdx : yIdx+lenOfY]
			potentialMatch := ""
			for _, ch := range newPattern {
				if ch == "x" {
					potentialMatch += x
				} else {
					potentialMatch += y
				}
			}
			if potentialMatch == str {
				if !didSwitch {
					return []string{x, y}
				} else {
					return []string{y, x}
				}
			}
		}
	} else {
		lenOfX := len(str) / counts["x"]
		if lenOfX%1 == 0 {
			x := str[:lenOfX]
			potentialMatch := ""
			for range newPattern {
				potentialMatch += x
			}
			if potentialMatch == str {
				if !didSwitch {
					return []string{x, ""}
				} else {
					return []string{"", x}
				}
			}
		}
	}

	return nil
}

func getNewPattern(pattern string) []string {
	newPattern := make([]string, len(pattern))
	for i := range pattern {
		newPattern[i] = string(pattern[i])
	}

	if pattern[0] == 'x' {
		return newPattern
	}
	for i := range newPattern {
		if newPattern[i] == "y" {
			newPattern[i] = "x"
		} else {
			newPattern[i] = "y"
		}
	}

	return newPattern
}

func getCountAndFirstYPos(pattern []string, counts map[string]int) int {
	firstYPos := -1
	for i, ch := range pattern {
		counts[ch] += 1
		if ch == "y" && firstYPos == -1 {
			firstYPos = i
		}
	}
	return firstYPos
}
