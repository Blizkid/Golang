package main

import "sort"

/* LeetCode 136. 只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func singleNumber(nums []int) int {
	var siteMap map[int]int
	siteMap = make(map[int]int)
	for _, v := range nums {
		var count = siteMap[v]
		count += 1
		siteMap[v] = count
	}

	for v := range siteMap {
		if siteMap[v] == 1 {
			return v
		}
	}
	return -1
}

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
例如，121 是回文，而 123 不是。
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false // 负数不是回文数
	}
	var digits []int
	n := x
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	// 比较首尾
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}

/*
有效的括号
考察：字符串处理、栈的使用
题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/

func isValid(s string) bool {
	var stack []rune
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			//入栈
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			}
			//出栈
			top := stack[len(stack)-1]
			if (v == ')' && top != '(') ||
				(v == '}' && top != '{') ||
				(v == ']' && top != '[') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

/*最长公共前缀
考察：字符串处理、分治法
题目：编写一个函数来查找字符串数组中的最长公共前缀。
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = getCommonPrefix(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func getCommonPrefix(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

/*
题目：给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/

func plusOne(digits []int) []int {
	result := 0
	length := len(digits)
	for i, v := range digits {
		ten := 10
		for j := 0; j < length-1-i; j++ {
			v = v * ten
		}
		result += v
	}
	result += 1
	return intToIntSlice(result)
}

// 把整数 x 的每一位存入 int 切片（高位在前）
func intToIntSlice(x int) []int {
	if x == 0 {
		return []int{0}
	}
	var digits []int
	for x > 0 {
		digits = append([]int{x % 10}, digits...) // 高位在前
		x /= 10
	}
	return digits
}

/*
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
*/

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	// 按左端点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			// 有重叠，更新右端点
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			// 无重叠，保存当前区间
			result = append(result, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	// 最后一个区间
	result = append(result, []int{start, end})
	return result
}

/*
两数之和
*/
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok { // p 是 target-x 在 nums 中的下标(value)，ok 是是否存在
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
