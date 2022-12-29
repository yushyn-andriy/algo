package median


// O(n + m) time | O(n + m) space
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	merged := make([]int, len(nums1)+len(nums2))

	var current, i, j int
	for {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				merged[current] = nums1[i]
				i++
			} else {
				merged[current] = nums2[j]
				j++
			}
		} else if i < len(nums1) {
			merged[current] = nums1[i]
			i++
		} else if j < len(nums2) {
			merged[current] = nums2[j]
			j++
		} else {
			break
		}
		current++
	}

	var median float64
	if len(merged) == 0 {
		median = 0.0
	} else if len(merged)%2 == 0 {
		middle := len(merged) / 2
		median = float64(merged[middle-1]+merged[middle]) / 2
	} else {
		middle := len(merged) / 2
		median = float64(merged[middle])
	}
	return median
}
