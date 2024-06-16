package main

import (
	"log"
	"sort"
	"testing"
)

func Test_MergeSort(t *testing.T) {

	var l uint = 32
	ls := []uint{1, 2, 2<<4, 2<<8, 2<<16, 1 << 20}

	for _, i := range ls {

		log.Print("testing sort of ", i)
		lns := randomstrings(i, l)
		if len(lns) != int(i) {
			log.Fatal("lns: wanted len ", i, " got ", len(lns))
		}
		result := make(chan []string)
		MergeSort(lns, result)
		rsl := <-result
		if len(rsl) != int(i) {
			log.Fatal("rsl: wanted len ", i, " got ", len(rsl))
		}

		if !sort.StringsAreSorted(rsl) {
			t.Error("rsort2a failed for size ", i)
		} else {
			log.Print("sort test passed for ", i)
		}
	}
}
