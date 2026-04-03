func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	ns := []rune{}
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			ns = append(ns, unicode.ToLower(rune(ch)))
		}
	}

	i := 0
	j := len(ns)-1
	for i < j {
		if ns[i] != ns[j] {
			return false
		}
		i++
		j--
	}
	
	return true
}
