package main

func twoSum(numbers []int, target int) []int {
	lower, upper := 0, len(numbers)-1
	for lower < upper {
		value := numbers[lower] + numbers[upper]
		if value == target {
			break
		}
		if value < target {
			lower++
		} else {
			upper--
		}
	}
	return []int{lower + 1, upper + 1}
}
