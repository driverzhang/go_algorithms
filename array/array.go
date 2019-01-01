// 由于go的特性这里数组与切片当成同一操作
package array

/**************** reverse 数组反转函数 *********************/
// reverse 数组反转 逆序输出方法
// 开辟新内存空间
func OldReverse(s []int) []int {
	var newArr = make([]int, 0)

	for i := len(s) - 1; i >= 0; i-- {
		newArr = append(newArr, s[i])
	}

	return newArr
}

// 在原内存空间操作
// go slice 特性为指针类型可直接修改底层数组
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/**************** BinarySearch 二分查找 *********************/
/*
 * nums: 数组切片的元素
 * target: 给定要查找的元素值
 */
func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high { // 范围不能缩小到只包含一个元素
		mid := (low + high) / 2 // 当前比较的数组下标
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			mid = -1
		}
	}

	return -1
}
