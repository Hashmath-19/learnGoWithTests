package main

func Sum(nums []int) int {
	var sum int
	for i := range nums{
		sum += nums[i]
	}
	return sum
}

func SumAll(nums ...[]int) []int{
	result := make([]int, 0)
	for _, num := range nums{
		result = append(result, Sum(num))
	}
	return result
}