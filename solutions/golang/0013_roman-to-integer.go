func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	cur, next, count := 0, roman[s[0]], roman[s[len(s)-1]]
	for i := 1; i < len(s); i++ {
		cur, next = next, roman[s[i]]
		if cur < next {
			count -= cur
		} else {
			count += cur
		}
	}
	return count
}
