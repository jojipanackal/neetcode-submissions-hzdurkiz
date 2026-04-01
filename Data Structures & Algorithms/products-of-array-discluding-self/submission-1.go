func productExceptSelf(nums []int) []int {
	var result []int

	p := make([]int, len(nums))
	s := make([]int, len(nums))

	p[0] = 1
	temp := 1
	for i:=0; i<len(nums)-1; i++ {
		temp *= nums[i]
		p[i+1] = temp
	}

	s[len(nums)-1] = 1
	temp = 1
	for i:=len(nums)-1; i>0; i-- {
		temp *= nums[i]
		s[i-1] = temp
	}

	for i:=0; i<len(nums); i++ {
		result = append(result, p[i]*s[i])
	}

	return result
}
