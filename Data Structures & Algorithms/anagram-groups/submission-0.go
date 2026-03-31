func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _ , str := range strs {
		var a [26]int
		for _ , ch := range str {
			a[ch - 'a']++

		}
		m[a] = append(m[a], str)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}