package main

func makeRange(min, max int) []int {
	rangeSlice := make([]int, max-min+1)
	for i := range rangeSlice {
		rangeSlice[i] = min + i
	}
	return rangeSlice
}
