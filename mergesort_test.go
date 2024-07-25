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

	var l int = 32
	var r bool = true

	ls := []int{1, 2, 1 << 4, 1 << 8, 1 << 16, 1 << 20, 1 << 24}

	for _, nl := range ls {

		log.Print("testing sort of ", nl)
		lns := randomstrings(nl, l, r)
		if len(lns) != int(nl) {
			log.Fatal("lns: wanted len ", nl, " got ", len(lns))
		}

		slns := MergeSort(lns)

		if len(slns) != int(nl) {
			log.Fatal("slns: wanted len ", nl, " got ", len(slns))
		}

		if !sort.StringsAreSorted(slns) {
			t.Error("rsort2a failed for size ", nl)
		} else {
			log.Print("sort test passed for ", nl)
		}

		timemymergesort(lns)
		timeslicessort(lns)
	}
}
