package main

func romanToInt(s string) int {
	roman_base := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	r := 0
	for i, v := range s {
		if i != 0 {
			if s[i-1] == 73 && (v == 86 || v == 88) {
				r -= 2
			}
			if s[i-1] == 88 && (v == 76 || v == 67) {
				r -= 20
			}
			if s[i-1] == 67 && (v == 68 || v == 77) {
				r -= 200
			}
		}
		r += roman_base[string(v)]
	}

	return r
}

func main() {
	println(romanToInt("MCMXCIV"))   // 1994
	println(romanToInt("MMXCIV"))    // 2094
	println(romanToInt("MMXCIII"))   // 2093
	println(romanToInt("MCMXCIII"))  // 1993
	println(romanToInt("LVIII"))     // 58
	println(romanToInt("III"))       // 3
	println(romanToInt("MMMCMXCIX")) // 3999
	jj := []rune{'I', 'V', 'X', 'L', 'C', 'D', 'M'}
	for i, v := range jj {
		println(i, v, string(v))
	}
}
