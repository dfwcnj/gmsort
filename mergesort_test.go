package main

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"testing"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func timemymergesort(lns []string) {
	a := fmt.Sprintf("msort     %d", len(lns))
	defer timer(a)()
	MergeSort(lns)
}

func timeslicessort(lns []string) {
	a := fmt.Sprintf("slices sort %d", len(lns))
	defer timer(a)()
	slices.Sort(lns)
}

func Test_MergeSort(t *testing.T) {

	var l uint = 32
	ls := []uint{1, 2, 1 << 4, 1 << 8, 1 << 16, 1 << 20, 1 << 24}

	for _, i := range ls {

		log.Print("testing sort of ", i)
		lns := randomstrings(i, l)
		if len(lns) != int(i) {
			log.Fatal("lns: wanted len ", i, " got ", len(lns))
		}

		slns := MergeSort(lns)

		if len(slns) != int(i) {
			log.Fatal("slns: wanted len ", i, " got ", len(slns))
		}

		if !sort.StringsAreSorted(slns) {
			t.Error("rsort2a failed for size ", i)
		} else {
			log.Print("sort test passed for ", i)
		}

		timemymergesort(lns)
		timeslicessort(lns)
	}
}
