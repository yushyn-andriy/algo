package recursion

type ProductArr []interface{}

func ProductSum(arr ProductArr, m int) int {
	s := 0
	for _, el := range arr {
		if cast, ok := el.(ProductArr); ok {
			s += ProductSum(cast, m+1)
		} else if cast, ok := el.(int); ok {
			s += cast
		}
	}

	return s * m
}
