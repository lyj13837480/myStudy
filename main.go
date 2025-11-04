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
	//加一
	fmt.Println("加一: ", plusOne([]int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}))
	//删除有序数组中的重复项
	fmt.Println("删除有序数组中的重复项: ", removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	//合并区间
	fmt.Println("合并区间: ", merge([][]int{{4, 5}, {1, 4}, {0, 1}}))
	//两数之和
	fmt.Println("两数之和: ", twoSum([]int{2, 7, 11, 15}, 9))
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

// 加一
func plusOne(digits []int) []int {
	return addOne(digits, len(digits)-1)
}

func addOne(digits []int, index int) []int {
	if index < 0 {
		return append([]int{1}, digits...)
	} else if digits[index] < 9 {
		digits[index] += 1
		return digits
	} else {
		digits[index] = 0
		return addOne(digits, index-1)
	}
}

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				nums = append(nums[:j], nums[j+1:]...)
				j--
			}
		}
	}
	return len(nums)
}

// 合并区间
func merge(intervals [][]int) [][]int {
	r := [][]int{}
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]
		end := intervals[i][1]
		if len(r) == 0 {
			r = append(r, []int{start, end})
		} else {
			last := r[len(r)-1]
			if start <= last[1] {
				if end > last[1] {
					last[1] = end
				}
			} else {
				r = append(r, []int{start, end})
			}
		}
	}

	return r
}

// 两数之和
func twoSum(nums []int, target int) []int {
	a := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				a[0] = i
				a[1] = j
				return []int{a[0], a[1]}
			}
		}
	}
	return []int{}
}
