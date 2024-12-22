package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{1, 2, 1}, 3))
}

func subarraySum(nums []int, k int) int {
	hashmap := make(map[int]int)
	prefixSum := 0
	count := 0
	hashmap[0] = 1
	for _, num := range nums {
		prefixSum += num
		x, ok := hashmap[prefixSum-k]
		if ok {
			count += x
		}
		hashmap[prefixSum] += 1
	}
	return count
}
