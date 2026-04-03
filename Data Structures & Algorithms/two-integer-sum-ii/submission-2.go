func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers)-1
	result := make([]int, 2)

	for i < j {
		if numbers[i] + numbers[j] == target {
			result = []int{i+1, j+1}
			break
		}
		if numbers[i] + numbers[j] < target {
			i++
		} else if numbers[i] + numbers[j] > target {
			j--
		}
	}
	return result
}
