type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var builder strings.Builder
	for _ , str := range strs {
		lengthStr := strconv.Itoa(len(str))

		builder.WriteString(lengthStr)
		builder.WriteString("#")
		builder.WriteString(str)
	}
	return builder.String()
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	
	len_encoded := len(encoded)
	i := 0
	for i < len(encoded) {
		length := 0

		for i < len_encoded && encoded[i] >= '0' && encoded[i] <= '9'{
			length = length * 10 + int(encoded[i] - '0')
			i++
		}

		if i < len_encoded && encoded[i] == '#' {
			i++
		}

		if i + length <= len_encoded {
			str := encoded[i : i+length]
			result = append(result, str)
			i += length
		} else {
			break
		}
	}
	return result 
}
