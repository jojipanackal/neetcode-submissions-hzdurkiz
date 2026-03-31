func topKFrequent(nums []int, k int) []int {
	b := make([][]int, len(nums)+1)
	f := make(map[int]int)
	var res []int

	for _ , num := range nums {
		f[num]++
	}

	for key, v := range f {
		b[v] = append(b[v], key)
	}

	for i := len(b)-1; i >= 0; i-- {
		if len(b[i]) > 0 {
			res = append(res, b[i]...)
		}
		
		if len(res) >= k {
			break
		}

	}
	return res
}