package main

import "fmt"

func firstUniqChar(s string) int {
	a := [26]int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	b := [26]int{}
	c := 1<<31 - 1

	for i, v := range s {
		a[v-'a']++
		b[v-'a'] += (i + 1)
	}

	for i, v := range a {
		if v == 0 {
			if b[i] < c {
				c = b[i]
			}
		}
	}

	if c == 1<<31-1 {
		return -1
	}

	return int(c) - 1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))

	s = "loveleetcode"
	println(firstUniqChar(s))

	s = "aabb"
	println(firstUniqChar(s))

	s = "z"
	println(firstUniqChar(s))
}
