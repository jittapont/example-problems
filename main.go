package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	i := 0
	w := ""
	for {
		if i >= len(strs[0]) {
			return w
		}
		char := strs[0][i]
		count := 0
		for _, str := range strs {
			if i >= len(str) || string(char) != string(str[i]) {
				return w
			} else {
				count += 1
			}
		}
		if count == len(strs) {
			w += string(char)
		}
		i += 1
	}
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	w := ""
	i := 0
	for {
		if i >= len(strs[0]) || strs[0][i] != strs[len(strs)-1][i] {
			return w
		} else {
			w += string(strs[0][i])
		}
		i += 1
	}
}

func isAllSameCharacter(s string) bool {
	m := make(map[rune]bool)
	count := 0
	for _, r := range s {
		ok := m[r]
		if !ok {
			m[r] = true
			count += 1
		}
	}
	return count == 1
}

func buddyStrings(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	if a == b {
		if isAllSameCharacter(a) && isAllSameCharacter(b) {
			return true
		}
		return false
	}
	diffPos := make([]int, 0)
	for i := 0; i < len(a); i++ {
		if string(a[i]) != string(b[i]) {
			diffPos = append(diffPos, i)
		}
	}
	if len(diffPos) > 2 {
		return false
	}
	if a[diffPos[0]] == b[diffPos[1]] && a[diffPos[1]] == b[diffPos[0]] {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "floor", "florida"}))
	fmt.Println(longestCommonPrefix2([]string{"flower", "floor", "florida"}))
	fmt.Println(longestCommonPrefix2([]string{""}))
	fmt.Println(longestCommonPrefix2([]string{"test"}))
	fmt.Println(longestCommonPrefix2([]string{"flower", "floor", "d"}))
	fmt.Println(buddyStrings("aa", "aa"))
}