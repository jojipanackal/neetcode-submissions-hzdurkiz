func productExceptSelf(nums []int) []int {
	var result []int

	p := make([]int, len(nums))
	s := make([]int, len(nums))

	i := 1
	p[0] = 1
	pp := 1
	for i < len(nums) {
		pp *= nums[i-1]
		p[i] = pp
		i++ 
	}

	i = len(nums) - 2
	s[len(nums)-1] = 1
	sp := 1
	for i >= 0 {
		sp *= nums[i+1]
		s[i] = sp
		i--
	}

	i = 0
	for i < len(nums) {
		result = append(result, p[i]*s[i])
		i++
	}

	return result
}
