func productExceptSelf(nums []int) []int {
	total := 1
	zeroCount := 0
	zeroIndex := -1

	for i, num := range nums {
		if num == 0 {
			zeroCount++
			zeroIndex = i
		} else {
			total *= num
		}
	}

	output := make([]int, len(nums))

	if zeroCount >= 2 {
		return output
	}

	if zeroCount == 1 {
		output[zeroIndex] = total
		return output
	}

	for i, num := range nums {
		output[i] = total/num
	}

	return output
}

