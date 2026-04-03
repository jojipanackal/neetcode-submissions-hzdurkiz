func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1

	for i < j {
		if numbers[i] + numbers[j] < target {
			i++
			continue
		}
		if numbers[i] + numbers[j] > target {
			j--
			continue
		}

		return []int{i+1, j+1}
	}
	return nil
}
