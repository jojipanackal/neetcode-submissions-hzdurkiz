func topKFrequent(nums []int, k int) []int {
	f := make(map[int]int) // Freq Map

	// Building f
	for _ , num := range nums {
		f[num]++
	}

	b := make([][]int, len(nums)+1)
	for key, v := range f {
		b[v] = append(b[v], key)
	}

	var res []int
	var c int
	for i := len(b)-1; i >= 0; i-- {
		if len(b[i]) == 0 {
			continue
		}
		for _ , j := range b[i] {
			if k == c {
				return res
			}
			res = append(res, j)
			c++
		}

	}
	return res
}