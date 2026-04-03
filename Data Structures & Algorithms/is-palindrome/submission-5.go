func isPalindrome(s string) bool {
	i := 0
	j := len(s)-1

	for i <= j {
		for !unicode.IsLetter(rune(s[i])) && !unicode.IsDigit(rune(s[i])) && i < j {
			i++
		}
		for !unicode.IsLetter(rune(s[j])) && !unicode.IsDigit(rune(s[j])) && i < j {
			j--
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])){
			return false
		}
		i++
		j--
	}
	return true
}
