package main

func subSlice(slice []int, begin int, length int, capacity int) []int {
	if slice == nil || begin < 0 || capacity < 0 {
		return nil
	}

	if capacity < length {
		capacity = length
	}
	newSlice := make([]int, length, capacity)

	if len(slice) < begin {
		return newSlice
	} else if len(slice) < begin+length {
		length = len(slice) - begin
	}
	copy(newSlice, slice[begin:begin+length])
	return newSlice
}
