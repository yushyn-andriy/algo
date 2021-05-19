package arrays

const (
	FIRST = iota
	SECOND
	THIRD
)

type ElMax struct {
	IsSet bool
	Value int
}

type ThreeMax []ElMax

func ThreeNumbersMax(arr []int) ThreeMax {
	threeMaxArr := ThreeMax{
		{false, 0},
		{false, 0},
		{false, 0},
	}
	for _, v := range arr {
		if !threeMaxArr[THIRD].IsSet {
			threeMaxArr[THIRD].IsSet = true
			threeMaxArr[THIRD].Value = v
		} else {
			shiftUpdate(threeMaxArr, v)
		}
	}
	return threeMaxArr
}

func shiftUpdate(arr ThreeMax, value int) {
	if value > arr[THIRD].Value {
		second := arr[SECOND]
		arr[SECOND] = arr[THIRD]
		arr[FIRST] = second
		arr[THIRD] = ElMax{true, value}
	} else if value > arr[SECOND].Value {
		if !arr[SECOND].IsSet {
			arr[SECOND].IsSet = true
			arr[SECOND].Value = value
		} else {
			arr[FIRST] = arr[SECOND]
			arr[SECOND] = ElMax{true, value}
		}
	} else if value > arr[FIRST].Value || !arr[FIRST].IsSet {
		arr[FIRST].Value = value
		arr[FIRST].IsSet = true
	}
}
