func isValid(s string) bool {
    stack := make([]rune, 0, len(s))

	for _, ch := range s {

		n := len(stack)
		if n > 0 {
			val := stack[n-1]
			if (val == '(' && ch == ')') || (val == '[' && ch == ']') || (val == '{' && ch == '}') {
				stack = stack[:n-1]
				continue
			}
		}
		stack = append(stack, ch)
	}

	if len(stack) == 0 {
		return true
	}

	return false

}
