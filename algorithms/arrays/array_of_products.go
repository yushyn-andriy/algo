package arrays

func ArrayOfProducts(array []int) []int {
	products := make([]int, len(array))

	n := 1
	zeros := 0
	for _, v := range array {
		if v != 0 {
			n *= v
		} else {
			zeros += 1
		}
	}

	for idx, v := range array {
		if v == 0 {
			if zeros > 1 {
				products[idx] = 0
			} else {
				products[idx] = n
			}
		} else {
			if zeros > 0 {
				products[idx] = 0
			} else {
				products[idx] = n / v
			}
		}
	}
	return products
}
