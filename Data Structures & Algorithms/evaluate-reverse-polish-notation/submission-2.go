func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _ , token := range tokens {
		tokenInt, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, tokenInt)
		} else {
			op1 := stack[len(stack)-1]
			op2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			if token == "+" {
				stack = append(stack, op1 + op2)
			}
			if token == "-" {
				stack = append(stack, op2 - op1)
			}
			if token == "*" {
				stack = append(stack, op1 * op2)
			}
			if token == "/" {
				stack = append(stack, op2 / op1)
			}
		}
	}

	return stack[len(stack)-1]
}
