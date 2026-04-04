func trap(height []int) int {
	n := len(height)
	prefix := make([]int, n)
	suffix := make([]int, n)

	temp := 0
	for i:=1; i<n; i++ {
		if height[i-1] > temp {
			temp = height[i-1]
		}
		prefix[i] = temp
	}

	temp = 0
	for i:=n-2; i>=0; i-- {
		if height[i+1] > temp {
			temp = height[i+1]
		}
		suffix[i] = temp
	}

	result := 0
	for i:=0; i<n; i++ {
		water := min(prefix[i], suffix[i]) - height[i]
		if water > 0 {
			result += water
		}
	}

	// fmt.Println(prefix)
	// fmt.Println(suffix)

	return result
}
