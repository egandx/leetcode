package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(x)
	fmt.Println(res)
}

// 双指针，三数之和， 时间O(n2), 空间O(1)，不重复的三元组
// a+b+c=0, b+c=-a -> -a即为target
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	ans := [][]int{}
	// a只需要到倒数第三个数为止
	for a := 0; a < length-2; a++ {
		// 剪枝
		// 因为已经排序，第一个数大于0，后面的数一定大于0，不可能组成三元组
		if nums[a] > 0 {
			break
		}
		// 避免重复
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		// 双指针实现
		c := length - 1
		target := -nums[a]
		// b最多到倒数第二个元素
		for b := a + 1; b < length-1; b++ {
			// 避免重复
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			// 相加大于target, c左移
			for b < c && nums[b]+nums[c] > target {
				c--
			}
			// b==c，说明该轮无合适值
			if b == c {
				break
			}
			// 找到合适的加入结果集中
			if nums[b]+nums[c] == target {
				ans = append(ans, []int{nums[a], nums[b], nums[c]})
			}
			// 继续在a的基础上寻找，b++, c--
		}
	}
	return ans
}
