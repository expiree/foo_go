package main

import "fmt"

func main() {
	//nums := []int{3, 2, 4}
	//fmt.Println(twoSum(nums, 6))
	arr1 := []int{0, 1, 2, 4}
	arr2 := []int{3, 5, 7}

	fmt.Println(arrayCombine(arr1, arr2, 2))
}

/**
 * 输入一个整数，输出该数二进制表示中 1 的个数。例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
 */
func hammingWeight(num uint32) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num = num / 2
	}
	return res
}

func intToBin(num uint32) uint64 {
	if num <= 1 {
		return uint64(num)
	}
	var (
		res   uint64
		times uint64 = 1
	)
	for num > 0 {
		res += uint64(num%2) * times
		num /= 2
		times *= 10
	}
	return res
}

/**
 * 两个有序数组从小到大排列，取两个数组合并后第K大的元素，要求O(1)空间复杂度
 * 如 a = {0, 1, 2, 4} b = {3, 5, 7} 第4大元素，返回3
 */
func arrayCombine(a, b []int, k int) int {
	if len(a) > len(b) {
		return arrayCombine(b, a, k)
	}

	if len(a) == 0 {
		return b[k-1]
	}

	if k == 1 {
		if a[0] < b[0] {
			return a[0]
		} else {
			return b[0]
		}
	}

	x, y := len(a), len(b)
	if len(a) > k/2 {
		x = k / 2
	}
	if len(b) > k/2 {
		y = k / 2
	}

	if a[x-1] < b[y-1] {
		return arrayCombine(a[x:], b, k-x)
	} else if a[x-1] > b[y-1] {
		return arrayCombine(a, b[y:], k-y)
	} else {
		return a[x-1]
	}
}
