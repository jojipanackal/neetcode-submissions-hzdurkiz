func maxArea(heights []int) int {
	i, j := 0, len(heights)-1

	max_area := 0
	for i < j {
		area := (j-i) * min(heights[i], heights[j])
		if area > max_area {
			max_area = area
		}
		if heights[i] < heights[j] {
			i++
			continue
		}
		j--
	}

	return max_area
}
