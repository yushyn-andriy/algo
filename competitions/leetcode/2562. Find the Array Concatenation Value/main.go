import "strconv"

func findTheArrayConcVal(nums []int) int64 {
    var res int64
    i, j := 0, len(nums) - 1
    for i<=j {
        x, y := nums[i], nums[j]
        if i == j {
            res += int64(x)
        } else {
            a := strconv.Itoa(x)
            b := strconv.Itoa(y)
            c, _ := strconv.Atoi(a + b)
            res += int64(c)
        }
        
        i+=1
        j-=1
    }
    return res
}