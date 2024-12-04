package lib

func Reverse(str string) string {
	result := ""

	for _, v := range str {
		result = string(v) + result
	}

	return result
}

func AllIndexes(str string, sub string) []int {
	result := []int{}
	strLen := len(str)
	subLen := len(sub)

	for i := 0; i < strLen; i++ {
		if str[i] == sub[0] {
			if i+subLen <= strLen && str[i:i+subLen] == sub {
				result = append(result, i)
			}
		}
	}

	return result
}
