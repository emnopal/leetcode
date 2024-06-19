package main

import "fmt"

func IntToRoman(num int) string {
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousand := []string{"", "M", "MM", "MMM"}
	bases := [4]int{1000, 100, 10, 1}

	int_base := map[int][]string{
		1:    ones,
		10:   tens,
		100:  hundreds,
		1000: thousand,
	}

	var roman string

	for _, base := range bases {
		_num := int((num / base) % 10)
		roman += int_base[base][_num]
	}

	return roman
}

func IntToRoman2(num int) string {
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousand := []string{"", "M", "MM", "MMM"}

	return thousand[int(num/1000)%10] + hundreds[int(num/100)%10] + tens[int(num/10)%10] + ones[int(num/1)%10]
}

func main() {
	a := IntToRoman(3490)
	b := IntToRoman2(3490)
	fmt.Println(a, b)
}
