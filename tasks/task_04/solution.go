package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	if len(nums) < 2 {
		return Stats{}
	}

	var diff, sum, min, max int64
	count := len(nums) - 1
	min, max = nums[1]-nums[0], nums[1]-nums[0]

	for i := 1; i < len(nums); i++ {
		diff = nums[i] - nums[i-1]
		sum += diff
		if diff < min {
			min = diff
		}
		if diff > max {
			max = diff
		}
	}

	return Stats{
		Count: count,
		Sum:   sum,
		Min:   min,
		Max:   max,
	}
}
