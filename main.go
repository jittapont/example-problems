package main

import (
	"fmt"
	"sort"
)

func Josephus(items []interface{}, k int) []interface{} {
  if len(items) == 0 {
    return []interface{}{}
  }
  if k > len(items) {
    k = k % len(items)
  }
	skip := k - 1
	i := skip
	items2 := make([]interface{}, 0)
	for {
		v := items[i]
		items = append(items[:i], items[i+1:]...)
		items2 = append(items2, v)
		if len(items) == 0 {
			break
		}
		i = (i + skip) % len(items)
	}
	return items2
}

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

func runningSum(nums []int) []int {
    sum := 0
    for i:=0; i < len(nums); i++ {
        if i == 0 {
            sum = nums[i]
        } else {
            sum += nums[i]
            nums[i] = sum   
        }
    }
    return nums   
}

func Solequa(n int) [][]int {
	s := make([][]int, 0)
	m := make(map[[2]int]bool)
	for i := n; i > 0; i-- {
		var max, min int
		if n%i == 0 {
			if i >= (n / i) {
				max = i
				min = (n / i)
			} else {
				min = i
				max = (n / i)
			}
			x := (max + min) / 2
			y := (max - min) / 4
			if (max-min)%4 == 0 && !m[[2]int{x, y}] {
				s = append(s, []int{x, y})
				m[[2]int{x, y}] = true
			}
		}
	}
	return s
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "floor", "florida"}))
	fmt.Println(longestCommonPrefix2([]string{"flower", "floor", "florida"}))
	fmt.Println(longestCommonPrefix2([]string{""}))
	fmt.Println(longestCommonPrefix2([]string{"test"}))
	fmt.Println(longestCommonPrefix2([]string{"flower", "floor", "d"}))
	fmt.Println(buddyStrings("aa", "aa"))
}
