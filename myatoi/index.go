package main

import "fmt"

const (
	MaxInt32 = 1<<31 - 1
	MinInt32 = -1 << 31
)

func strip(s string) string {
	for _, v := range s {
		if v == 32 {
			s = s[1:]
		} else {
			break
		}
	}
	return s
}

func myAtoi(s string) int {
	if len(s) < 1 {
		return 0
	}

	s = strip(s)
	mult := 1

	if len(s) < 1 {
		return 0
	}

	if s[0] == 43 {
		s = s[1:]
	} else if s[0] == 45 {
		s = s[1:]
		mult = -1
	}

	if len(s) < 1 {
		return 0
	}

	if s[0] < 48 || s[0] > 57 {
		return 0
	}

	r := 0

	for _, v := range s {
		if v < 48 || v > 57 {
			break
		}
		if r > MaxInt32+1 {
			break
		}
		r = r*10 + (int(v) - 48)
	}

	result := mult * r

	if result <= MinInt32 {
		return MinInt32
	}

	if result > MaxInt32 {
		return MaxInt32
	}

	return result
}

func main() {
	fmt.Println(myAtoi("42"))            // 42
	fmt.Println(myAtoi("-42"))           // -42
	fmt.Println(myAtoi("+-42"))          // 0
	fmt.Println(myAtoi("+42"))           // 42
	fmt.Println(myAtoi("-+42"))          // 0
	fmt.Println(myAtoi("   -42"))        // -42
	fmt.Println(myAtoi("   -+42"))       // 0
	fmt.Println(myAtoi("1337c0d3"))      // 1337
	fmt.Println(myAtoi("0-1"))           // 0
	fmt.Println(myAtoi("words and 987")) // 0
	fmt.Println(myAtoi("2147483688"))    // 2147483647
	fmt.Println(myAtoi("-2147483688"))   // -2147483648

}
