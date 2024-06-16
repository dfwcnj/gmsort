package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	var file string
	var dtype string
	flag.StringVar(&file, "file", "", "file to sort")
	flag.StringVar(&dtype, "type", "string", "type [string, int, datetime] to sort")
	flag.Parse()

	f := os.Stdin
	var err error
	if file != "" {
		f, err = os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	}

	// slice of lines
	br := bufio.NewReader(f)
	var ss []string
	var us []uint64
	//	var ds []time.Time

	for {
		s, err := br.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		if dtype == "int" {
			if dtype == "int" {
				un, err := strconv.ParseUint(s, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				us = append(us, un)
			}
		} else if dtype == "string" {
			ss = append(ss, s)
		} else {
			log.Fatal("only string, int, or datetime")
		}
	}
	switch dtype {
	case "int":
		result := make(chan []uint64)
		defer close(result)
		go MergeSort(us, result)
		r := <-result
		for _, v := range r {
			fmt.Print(v)
		}
	case "string":
		result := make(chan []string)
		defer close(result)
		go MergeSort(ss, result)
		r := <-result
		for _, v := range r {
			fmt.Print(v)
		}
	}

}
