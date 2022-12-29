package recursion

var mnemonics = map[rune]string{
	'1': "1",
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
	'0': "0",
}

func getMenonicsByNumber(n rune) string {
	v := mnemonics[n]
	return v
}

func PhoneNumberMnemonics(phoneNumber string) []string {
	sequance := []string{}
	res := []string{}
	for _, n := range phoneNumber {
		sequance = append(sequance, getMenonicsByNumber(n))
	}
	phoneNumberMnemonics(sequance, "", 0, &res)
	return res
}

func phoneNumberMnemonics(array []string, currMnem string, k int, res *[]string) {
	if len(array) == k {
		*res = append(*res, currMnem)
		return
	}

	for _, v := range array[k] {
		currMnem += string(v)
		phoneNumberMnemonics(array, currMnem, k+1, res)
		currMnem = currMnem[:len(currMnem)-1]
	}
}
