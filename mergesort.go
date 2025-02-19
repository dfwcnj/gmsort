package main

// slightly modified version of
// https://gist.github.com/julianshen/3940045

import (
	"cmp"
)

// func Merge(ldata []string, rdata []string) (result []string) {
func Merge[S ~[]E, E cmp.Ordered](ldata S, rdata S) (result S) {
	result = make(S, len(ldata)+len(rdata))
	lidx, ridx := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case lidx >= len(ldata):
			result[i] = rdata[ridx]
			ridx++
		case ridx >= len(rdata):
			result[i] = ldata[lidx]
			lidx++
		case ldata[lidx] < rdata[ridx]:
			result[i] = ldata[lidx]
			lidx++
		default:
			result[i] = rdata[ridx]
			ridx++
		}
	}

	return result
}

// func MergeSort(data []string, r chan []string) {
func MergeSort[S ~[]E, E cmp.Ordered](data S) []E {
	if len(data) == 1 {
		return data
	}

	middle := len(data) / 2

	ldata := MergeSort(data[:middle])
	rdata := MergeSort(data[middle:])

	return Merge(ldata, rdata)
}
