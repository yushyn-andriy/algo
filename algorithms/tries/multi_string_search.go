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

type SuffixTrie1 map[byte]SuffixTrie1

func NewSuffixTrie1() SuffixTrie1 {
	return SuffixTrie1{}
}

func (st SuffixTrie1) Populate(s string) {
	for i := range s {
		trie := st
		for j := i; j < len(s); j++ {
			ch := s[j]
			if _, found := trie[ch]; !found {
				trie[ch] = NewSuffixTrie1()
			}
			trie = trie[ch]
		}
		trie['*'] = nil
	}
}
func (st SuffixTrie1) Contains(s string) bool {
	trie := st
	for i := range s {
		ch := s[i]
		_, found := trie[ch]
		if !found {
			return false
		}
		trie = trie[ch]
	}

	// _, found := trie['*']
	return true // found
}

func MultiStringSearch2(bigString string, smallStrings []string) []bool {
	st := NewSuffixTrie1()
	st.Populate(bigString)
	result := make([]bool, len(smallStrings))
	for i := range smallStrings {
		result[i] = st.Contains(smallStrings[i])
	}
	return result
}

type Trie3 struct {
	children map[byte]Trie3

	word string
}

func (t Trie3) Add(word string) {
	current := t
	for i := range word {
		letter := word[i]
		if _, found := current.children[letter]; !found {
			current.children[letter] = Trie3{
				children: map[byte]Trie3{},
			}
		}
		current = current.children[letter]
	}
	current.children['*'] = Trie3{
		children: map[byte]Trie3{},
		word:     word,
	}
}

func MultiStringSearch3(bigString string, smallStrings []string) []bool {
	trie := Trie3{children: make(map[byte]Trie3)}
	for _, str := range smallStrings {
		trie.Add(str)
	}
	containedStrings := make(map[string]bool)
	for i := range bigString {
		findSmallStringsIn(bigString, i, trie, containedStrings)
	}

	output := make([]bool, len(smallStrings))
	for i, str := range smallStrings {
		output[i] = containedStrings[str]
	}

	return output
}

func findSmallStringsIn(str string, startIdx int, trie Trie3, containedStrings map[string]bool) {
	current := trie
	for i := startIdx; i < len(str); i++ {
		currentChar := str[i]
		if _, found := current.children[currentChar]; !found {
			break
		}
		current = current.children[currentChar]
		if end, found := current.children['*']; found {
			containedStrings[end.word] = true
		}
	}
}
