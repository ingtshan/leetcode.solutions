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
	counter := 0
	for i := 1; i < len(s); i++ {
		if roman[s[i-1]] < roman[s[i]] {
			counter -= roman[s[i-1]]
		} else {
			counter += roman[s[i-1]]
		}
	}
	return counter + roman[s[len(s)-1]]
}
