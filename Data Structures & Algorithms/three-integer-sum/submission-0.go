func threeSum(nums []int) [][]int {
	var result [][]int

	sort.Ints(nums)

	n := len(nums)
	for i:=0; i<n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i+1;
		r := n-1
		for l < r {
			if nums[i] + nums[l] + nums[r] > 0 {
				r--
				continue
			}
			if nums[i] + nums[l] + nums[r] < 0 {
				l++
				continue
			}
			result = append(result, []int{nums[i], nums[l], nums[r]})
			l++
			r--

			for l < r && nums[l] == nums[l-1] { l++ }
			for l < r && nums[r] == nums[r+1] { r-- }
		}
	}

	return result
}
