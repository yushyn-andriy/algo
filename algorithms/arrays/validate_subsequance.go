package arrays

// ValidateSubsequent ...
// O(N) time | O(1) space
func ValidateSubsequent(sequent, subSequent []int) bool {
	seqI, subI := 0, 0

Loop:
	for {
		subEl, seqEl := subSequent[subI], sequent[seqI]
		if subEl == seqEl {
			subI++
		}
		seqI++

		if subI > len(subSequent)-1 || seqI > len(sequent)-1 {
			break Loop
		}
	}

	return subI == len(subSequent)
}
