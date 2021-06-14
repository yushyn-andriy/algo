package recursion

func Factors(fn []int) [][][]int {
	r := [][][]int{}
	factors(fn, len(fn), [][]int{}, &r)
	return r
}

func factors(fn []int, fnLen int, currCombination [][]int, r *[][][]int) {
	if len(fn) == 0 {
		toAppend := make([][]int, len(currCombination))
		copy(toAppend, currCombination)
		*r = append(*r, toAppend)
		return
	}

	for i := 1; i <= fnLen; i++ {
		tmp := []int{}
		if len(fn) < i {
			break
		}
		for j := 0; j < i; j++ {
			if len(fn) == 0 {
				break
			}
			el := fn[len(fn)-1]
			fn = fn[:len(fn)-1]
			tmp = append(tmp, el)
		}
		appTmp := make([]int, len(tmp))
		copy(appTmp, tmp)

		currCombination = append(currCombination, appTmp)
		factors(fn, fnLen, currCombination, r)
		currCombination = currCombination[:len(currCombination)-1]
		for j := 0; j < i; j++ {
			if len(tmp) == 0 {
				break
			}
			el := tmp[len(tmp)-1]
			tmp = tmp[:len(tmp)-1]
			fn = append(fn, el)
		}
	}
}
