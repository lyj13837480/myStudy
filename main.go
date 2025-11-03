package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, world!")

	//136. 只出现一次的数字
	fmt.Println("只出现一次的数字: ", singleNumber([]int{2, 2, 1}))
	//回文数
	fmt.Println("回文数: ", isPalindrome(101))
	//有效的括号
	fmt.Println("效的括号: ", isValid("([)]"))
	//最长公共前缀
	fmt.Println("最长公共前缀: ", longestCommonPrefix([]string{"abab", "aba", ""}))
}

// 136. 只出现一次的数字
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, i := range nums {
		if m[i] != 0 {
			m[i]++
		} else {
			m[i] = 1
		}
	}

	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

// 回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	a := strconv.Itoa(x)
	n := len(a)
	for i := 0; i < n/2; i++ {
		if a[i] != a[n-1-i] {
			return false
		}
	}
	return true
}

// 有效的括号
func isValid(s string) bool {
	a := map[rune]rune{'(': ')', '{': '}', '[': ']'}
	//b := []rune{'(', '{', '[', ')', '}', ']'}
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		if a[r[i]] == r[i+1] {
			i += 1
			continue
		}
		if a[r[i]] != r[len(r)-1-i] {
			return false
		}
	}
	return true
}

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs[0]) == 0 {
		return ""
	}
	r := []rune(strs[0])
	a := len(r) - 1
	for i := 1; i < len(strs); i++ {
		b := []rune(strs[i])
		if len(b) == 0 {
			return ""
		}
		m := 0
		for j := 0; j < len(r) && j < len(b); j++ {
			if r[j] != b[j] {
				if j == 0 {
					return ""
				}
				if j < m {
					m = j
				}
				break
			} else {
				m = j
			}
		}
		if m < a {
			a = m
		}
	}
	return string(r[:a+1])
}
