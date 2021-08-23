package main

import "fmt"

func main()  {
	sour := []int {2, 7, 6, 1}
	i  := twoSum(sour, 3)
	fmt.Printf("%d",i)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}
