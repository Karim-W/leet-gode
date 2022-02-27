package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		} else {
			m[v] = i
		}
	}
	return []int{}
}
func main() {
	f := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(f)
}
