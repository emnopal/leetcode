package main

func lengthOfLastWord(s string) int {
	var b []string
	var c string
	for _, v := range s {
		if string(v) == " " {
			b = append(b, c)
			c = ""
		} else {
			c += string(v)
		}
	}

	b = append(b, c)
	var b_fix string

	for _, v := range b {
		if string(v) != "" {
			b_fix = v
		}
	}

	return len(b_fix)
}

func main() {
	println(lengthOfLastWord("Hello World"))                 // 5
	println(lengthOfLastWord("   fly me   to   the moon  ")) // 4
	println(lengthOfLastWord("luffy is still joyboy"))       // 6
}
