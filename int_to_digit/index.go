package main

// Difficulty: Hard
// Link: https://leetcode.com/problems/integer-to-english-words/description/
// Complexity: O(1) both in time and space

func OneTeen(num int) string {
	ones := [20]string{"",
		"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen",
	}
	return ones[num]
}

func Ten(num int) string {
	if num < 20 {
		return OneTeen(num)
	}

	_Ten := [8]string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}

	var digit string

	if OneTeen(num%10) != "" {
		digit = " " + OneTeen(num%10)
	}

	return _Ten[int(num/10)-2] + digit
}

func Hundred(num int) string {
	if num < 100 {
		return Ten(num)
	}

	var digit string

	if num%100 > 0 {
		digit = " " + Ten(num%100)
	}

	return OneTeen(int(num/100)) + " Hundred" + digit
}

func Thousand(num int) string {
	if num < 1_000 {
		return Hundred(num)
	}

	var digit string

	if num%1_000 > 0 {
		digit = " " + Hundred(num%1_000)
	}

	return Hundred(int(num/1_000)) + " Thousand" + digit
}

func Million(num int) string {
	if num < 1_000_000 {
		return Thousand(num)
	}

	var digit string

	if num%1_000_000 > 0 {
		digit = " " + Thousand(num%1_000_000)
	}

	return Hundred(int(num/1_000_000)) + " Million" + digit
}

func Billion(num int) string {
	if num < 1_000_000_000 {
		return Million(num)
	}

	var digit string

	if num%1_000_000_000 > 0 {
		digit = " " + Million(num%1_000_000_000)
	}

	return Hundred(int(num/1_000_000_000)) + " Billion" + digit
}

func NumberToWords(num int) string {
	if num < 1 {
		return "Zero"
	}

	if num < 100 {
		return Ten(num)
	}

	if num >= 100 && num < 1_000 {
		return Hundred(num)
	}

	if num >= 1_000 && num < 1_000_000 {
		return Thousand(num)
	}

	if num >= 1_000_000 && num < 1_000_000_000 {
		return Million(num)
	}

	if num >= 1_000_000_000 && num < 2_147_483_648 {
		return Billion(num)
	}

	return ""
}

func main() {
	println(NumberToWords(0))              // Zero
	println(NumberToWords(1))              // One
	println(NumberToWords(12))             // Twelve
	println(NumberToWords(123))            // One Hundred Twenty Three
	println(NumberToWords(1_234))          // One Thousand Two Hundred Thirty Four
	println(NumberToWords(12_345))         // Twelve Thousand Three Hundred Forty Five
	println(NumberToWords(123_456))        // One Hundred Twenty Three Thousand Four Hundred Fifty Six
	println(NumberToWords(1_234_567))      // One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven
	println(NumberToWords(12_345_678))     // Twelve Million Three Hundred Forty Five Thousand Six Hundred Seventy Eight
	println(NumberToWords(123_456_789))    // One Hundred Twenty Three Million Four Hundred Fifty Six Thousand Seven Hundred Eighty Nine
	println(NumberToWords(1_234_567_890))  // One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety
	println(NumberToWords(12_345_678_901)) // ""
}
