package main

import (
	"fmt"
	"sort"
)


type Ans struct {
	From  string
	Var   string
	Count int
}

func count(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range s {
		c := strings.Count(s, string(w))
		if c > 1 && unicode.IsLetter(w) {
			m[string(w)] = c
		}

	}
	return m
}

func Mix(s1, s2 string) string {
	m1 := count(strings.ReplaceAll(s1, " ", ""))
	m2 := count(strings.ReplaceAll(s2, " ", ""))
	a := make([]Ans, 0)
	fmt.Println(m1, m2)
	for k, v := range m1 {
		_, ok := m2[k]
		if !ok {
			a = append(a, Ans{From: "1", Var: k, Count: v})
		}
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		var max int
		var i string
		if ok {
			if v2 > v1 {
				max = v2
				i = "2"
			} else if v1 == v2 {
				max = v1
				i = "="
			} else {
				max = v1
				i = "1"
			}
			a = append(a, Ans{From: i, Var: k, Count: max})
		}
	}
	for k, v := range m2 {
		_, ok := m1[k]
		if !ok {
			a = append(a, Ans{From: "2", Var: k, Count: v})
		}
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].Count != a[j].Count {
			return a[i].Count > a[j].Count
		}
		switch strings.Compare(a[i].Var, a[j].Var) {
		case -1:
			return true
		case 1:
			return false
		}
		return false
	})
	a2 := make([]string, 0)
	for _, i := range a {
		a2 = append(a2, i.From+":"+strings.Repeat(i.Var, i.Count))
	}
	return strings.Join(a2, "/")
}

func Solution(str string) []string {
	if len(str)%2 != 0 {
		str += strings.Repeat("_", ((len(str)/2)+1)*2-len(str))
	}
	st := []string{}
	for i := 0; i < len(str)-1; i+=2 {
		st = append(st, string([]rune(str)[i:i+2]))
	}
	return st
}

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

func FindOutlier(integers []int) int {
	if (integers[0]%2 == 0) != (integers[1]%2 == 0) && (integers[1]%2 == 0) == (integers[2]%2 == 0) {
		return integers[0]
	}
	for i, v := range integers[1:] {
		fmt.Println(integers[i], v)
		if (integers[i]%2 == 0) != (v%2 == 0) {
			return v
		}
	}
	return 0
}

func rotate(s string) string {
	if len(s) < 1 {
		return s
	}
	return string(s[1:]) + string(s[0])
}

func MaxRot(n int64) int64 {
	s := strconv.Itoa(int(n))
	max := int(n)
	for i := 0; i < len(s)-1; i++ {
		s = s[0:i] + rotate(s[i:])
		temp, _ := strconv.Atoi(s)
		if temp > max {
			max = temp
		}
	}
	return int64(max)
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
